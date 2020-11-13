package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	bodyRaw, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("[%s] %s: %q\n", r.Method, r.URL, bodyRaw)
	r.Header.Write(os.Stdout)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "got it\n")
	fmt.Printf("\n\n")
}

func main() {
	var (
		endpoint string
		certFile string
		keyFile  string
	)
	flag.StringVar(&endpoint, "endpoint", ":9090", "endpoint to listen on")
	flag.StringVar(&certFile, "certfile", "localhost.crt", "TLS certificate file")
	flag.StringVar(&keyFile, "keyfile", "localhost.key", "TLS key")
	flag.Parse()

	http.HandleFunc("/", handler)
	fmt.Printf("listening on https://%s\n", endpoint)
	log.Fatal(http.ListenAndServeTLS(endpoint, certFile, keyFile, nil))
}
