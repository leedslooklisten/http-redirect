package main

/*

Simple HTTP-redirect web server
Author: J. Benjamin Leeds
Created: February 8. 2019
Copyright 2019 Leeds Look Listen, Inc.

*/

import (
	"log"
	"net/http"
)

func main() {

	host := ""
	port := "443"
	addr := host + ":" + port

	rootHandler := func(w http.ResponseWriter, req *http.Request) {
		http.Redirect(w, req, "https://www.modernfidelity.us", http.StatusMovedPermanently)
	}

	http.HandleFunc("/", rootHandler)

	log.Println("Listening on " + addr + "...")

	err := http.ListenAndServeTLS(addr, "./cert.pem", "./privkey.pem", nil)
	if err != nil {
		log.Fatal(err)
	}

}
