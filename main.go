package main

/*
Simple HTTP-redirect web server
Author: J. Benjamin Leeds
Created: February 8, 2019
Modified: January 10, 2021
© 2021 Leeds Look Listen, Inc.
*/

import (
	"log"
	"net/http"
)

func httpServer(host string, c chan int) {
	httpPort := "80"
	httpAddr := host + ":" + httpPort

	/* http.ListenAndServe starts an HTTP server given an address and handler.
	* A nil handler means to use the DefaultServeMux. Use .Handle() and
	* .HandleFunc() to add handlers to the DefaultServeMux.
	*/
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
