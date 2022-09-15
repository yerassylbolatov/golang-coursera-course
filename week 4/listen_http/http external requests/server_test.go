package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type TestCases struct {
	Id      string
	Result  *CheckoutResult
	IsError bool
}

type CheckoutResult struct {
	Status  int
	Balance int
	Err     string
}

func CheckoutDummy(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("id")
	switch key {
	case "1":
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `{"status": 200, "balance": 100500}`)
	case "2":
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `{"status": 400, "err": "bad_balance"}`)
	case "_broken_json_":
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `{"status": 400`)
	case "_internal_error_":
		fallthrough
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}

type Cart struct {
	PaymentApiURL string
}

func TestCartCheckout(t *testing.T) {
	cases := []TestCases{
		{
			Id: "1",
			Result: &CheckoutResult{
				Status:  200,
				Balance: 100500,
				Err:     "",
			},
			IsError: false,
		},
		{
			Id: "2",
			Result: &CheckoutResult{
				Status:  400,
				Balance: 0,
				Err:     "bad_balance",
			},
			IsError: false,
		},
		{
			Id:      "_broken_json_",
			Result:  nil,
			IsError: true,
		},
		{
			Id:      "_internal_error_",
			Result:  nil,
			IsError: true,
		},
	}
	ts := httptest.NewServer(http.HandlerFunc(CheckoutDummy))

	for caseNum, item := range cases {
		c := &Cart{
			PaymentApiURL: ts.URL,
		}
		result, err := c.Checkout(item.Id)
		if err != nil && !item.IsError {
			t.Errorf("[%d] unexpected error found: %#v", caseNum, err)
		}
		if err == nil && item.IsError {
			t.Errorf("[%d] expected error, got nil", caseNum)
		}
		if !reflect.DeepEqual(item.Result, result) {
			t.Errorf("[%d] results are not equal; got: %#v, expected: %#v", caseNum, result, item.Result)
		}
	}
	ts.Close()
}

func (c *Cart) Checkout(id string) (*CheckoutResult, error) {
	url := c.PaymentApiURL + "?id=" + id
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	result := &CheckoutResult{}
	err = json.Unmarshal(data, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
