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

func httpServer(host string, c chan int) {
	httpPort := "80"
	httpAddr := host + ":" + httpPort

	log.Println("Listening on " + httpAddr + "...")
	log.Fatal(http.ListenAndServe(httpAddr, nil))
}

func httpsServer(host string, c chan int) {
	httpsPort := "443"
	httpsAddr := host + ":" + httpsPort
	log.Println("Listening on " + httpsAddr + "...")

	// Debug Only
	// time.Sleep(5000 * time.Millisecond)
	// close(c)

	err := http.ListenAndServeTLS(httpsAddr, "./cert.pem", "./privkey.pem", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {

	host := ""

	rootHandler := func(w http.ResponseWriter, req *http.Request) {
		http.Redirect(w, req, "https://www.modernfidelity.us", http.StatusMovedPermanently)
	}

	http.HandleFunc("/", rootHandler)

	c := make(chan int)

	go httpServer(host, c)
	go httpsServer(host, c)

	_, ok := <-c

	log.Println(ok)

}
