package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://keirin.jp/pc/dfw/portal/guest/data/winner_g3/winner_g3_hakodate.html"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(url)
	fmt.Println(string(byteArray))
}
