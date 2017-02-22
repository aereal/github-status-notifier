package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
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

	log.Println("---> start a server ...")
	err := http.ListenAndServe(":8000", nil)
	log.Fatal(err)
}
