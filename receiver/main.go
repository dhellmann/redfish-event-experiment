package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/dhellmann/redfish-event-experiment/config"
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
	appConfig, err := config.LoadFromFile("config.yaml")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", handler)
	fmt.Printf("listening on https://%s\n", appConfig.Receiver.Endpoint)
	log.Fatal(http.ListenAndServeTLS(
		appConfig.Receiver.Endpoint,
		appConfig.Receiver.CertFile,
		appConfig.Receiver.KeyFile,
		nil))
}
