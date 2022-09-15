package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("id")
	if key == "42" {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `{"status":200,"resp":{"user":"42"}}`)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"status":500,"resp":{"err":"db_error"}}`)
	}

}

type TestCase struct {
	Id         string
	Response   string
	StatusCode int
}

func TestGetUser(t *testing.T) {
	cases := []TestCase{
		{
			Id:         "42",
			Response:   `{"status":200,"resp":{"user":"42"}}`,
			StatusCode: http.StatusOK,
		},
		{
			Id:         "4",
			Response:   `{"status":500,"resp":{"err":"db_error"}}`,
			StatusCode: http.StatusInternalServerError,
		},
	}
	for caseNum, item := range cases {
		url := "http://example.com/api/user?id=" + item.Id
		req := httptest.NewRequest(http.MethodGet, url, nil)
		w := httptest.NewRecorder()

		GetUser(w, req)

		if w.Code != item.StatusCode {
			t.Errorf("[%d] wrong StatusCode: got %d, expected %d", caseNum, w.Code, item.StatusCode)
		}

		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)

		bodyStr := string(body)
		if bodyStr != item.Response {
			t.Errorf("[%d] wrong Response: got %s, expected %s", caseNum, bodyStr, item.Response)
		}
	}

}
