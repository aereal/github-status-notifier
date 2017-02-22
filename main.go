package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var listenPort = flag.Int("port", 8000, "port to listen")

var config Config

type GitHubStatusEvent struct {
	Context     string           `json:"context"`
	State       string           `json:"state"`
	Description string           `json:"description"`
	TargetURL   string           `json:"target_url"`
	Repository  GithubRepository `json:"repository"`
}

type GithubRepository struct {
	Name     string `json:"name"`
	FullName string `json:"full_name"`
}

func notifyStatusEvent(statusEvent GitHubStatusEvent) {
	post := GitHubStatusEventAsPost(statusEvent)
	post.Channel = config.Notifications.Slack.Channel
	post.Username = config.Notifications.Slack.Username
	post.IconEmoji = config.Notifications.Slack.IconEmoji
	postJson, _ := json.Marshal(post)
	http.Post(config.Notifications.Slack.WebhookURL, "application/json", bytes.NewBuffer(postJson))
}

func handleHook(w http.ResponseWriter, r *http.Request) {
	log.Printf("method:%s\tpath:%s\thost:%s", r.Method, r.URL.Path, r.Host)

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "bad body")
		return
	}

	githubEvent := r.Header.Get("X-GitHub-Event")
	githubDelivery := r.Header.Get("X-GitHub-Delivery")

	if githubEvent == "status" {
		var statusEvent GitHubStatusEvent
		if err := json.Unmarshal(bodyBytes, &statusEvent); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error:%q", err)
			return
		}

		log.Printf("%#q", statusEvent)

		notifyStatusEvent(statusEvent)
	}

	log.Printf("event:%s delivery:%s", githubEvent, githubDelivery)

	fmt.Fprintln(w, "ok")
}

func main() {
	flag.Parse()

	config, err := ParseConfigFile("config.json")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%#v\n", config)

	http.HandleFunc("/hook", handleHook)

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
	err = http.ListenAndServe(addr, nil)
	log.Fatal(err)
}
