package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
)

func main(){
	res, err := http.Get("https://www.facebook.com/")
	if err != nil {
		log.Fatal(err)
	}
	page, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", page)
}
