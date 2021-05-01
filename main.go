package main

import (
	"log"
	"net/http"
)

func main(){
	_, err := http.Get("https://google.com")
	if err != nil {
		panic(err)
		log.Println(err)
		return
	}
	panic("Hello")
	//log.Println("Success!")
}