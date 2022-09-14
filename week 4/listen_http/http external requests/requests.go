package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)


func getFullRequest(){
	req := &http.Request{
		Method: http.MethodGet,
		Header: http.Header{
			"User-Agent": {"coursera/golang"},
		}
	}

	//same as url:=http://127.0.0.1:8080/?id=42&user=yerassyl...

	req.URL,_=url.Parse("http://127.0.0.1:8080/?id=42")
	req.URL.Query().Set("user","yerassyl")

	resp, err:=http.DefaultClient.Do(req)
	if err != nil{
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Printf("testGetFullRequest %#v\n\n\n", string(respBody))
}

func runGet() {
	url := "http://localhost:8080/?param==123&param2=test"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("http.Get Body %#v\n\n\n", string(respBody))
}

func main() {

}
