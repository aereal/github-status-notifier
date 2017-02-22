package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var listenPort = flag.Int("port", 8000, "port to listen")

func main() {
	flag.Parse()

	http.HandleFunc("/hook", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("method:%s\tpath:%s\thost:%s", r.Method, r.URL.Path, r.Host)

		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "bad body")
			return
		}

		githubEvent := r.Header.Get("X-GitHub-Event")
		githubDelivery := r.Header.Get("X-GitHub-Delivery")
		log.Printf("event:%s delivery:%s payload:%s", githubEvent, githubDelivery, string(bodyBytes))

		fmt.Fprintln(w, "ok")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("method:%s\tpath:%s\thost:%s", r.Method, r.URL.Path, r.Host)

		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "You can only get")
		} else {
			fmt.Fprintln(w, "ok")
		}
	})

	addr := fmt.Sprintf(":%d", *listenPort)
	log.Printf("---> start a server at %s ...", addr)
	err := http.ListenAndServe(addr, nil)
	log.Fatal(err)
}
