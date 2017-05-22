package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func main(){
	res, _ := http.Get("http://www.facebook.com")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Print("%s", page)
}
