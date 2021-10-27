package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	url := "http://127.0.0.1:9000/index"
	doGet(url)
	doPost(url)
	doDiy(url)
}

func doPost(addr string) {
	postParams := url.Values{}

	postParams.Set("a", "1")
	postParams.Add("b", "2")

	resp, _ := http.PostForm(addr, postParams)
	defer resp.Body.Close()

	content, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(content))
}

func doGet(url string) {
	resp, _ := http.Get(url + "?a=1&b=2")
	defer resp.Body.Close()

	content, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(content))
}

func doDiy(url string) {
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet, url+"?a=1&b=2", nil)
	//req.Form.Set("a", "1")
	//req.Form.Set("b","2")
	resp, _ := client.Do(req)

	defer resp.Body.Close()

	content, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(content))
}
