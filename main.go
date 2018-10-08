package main

import (
	"io/ioutil"
	"log"
	"net/http"
)


// Getメソッドからレスポンスボディーを返す
func main() {
	response, err := http.Get("http://localhost:18888")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
}
