package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	example()
	example2()
}

// 主动关闭连接
func example() {
	req, err := http.NewRequest("GET", "http://golang.org", nil)
	checkErr(err)

	req.Close = true
	//req.Header.Add("Connection", "close")    // 等效的关闭方式

	resp, err := http.DefaultClient.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	checkErr(err)

	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)

	fmt.Println(string(body))
}

func example2() {
	t := http.Transport{DisableKeepAlives: true}
	client := http.Client{Transport: &t}

	resp, err := client.Get("https://golang.org")
	if resp != nil {
		defer resp.Body.Close()
	}
	checkErr(err)

	fmt.Println(resp.StatusCode) //

	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)

	fmt.Println(string(body))
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
