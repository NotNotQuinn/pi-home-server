package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/notnotquinn/go-websub"
)

func main() {
	// initialize subscriber, hub, and publisher.
	baseUrl := "http://localhost:3033"
	mux := http.NewServeMux()

	s := websub.NewSubscriber(
		baseUrl+"/sub/",
		websub.SWithLeaseLength(time.Hour*24),
	)

	p := websub.NewPublisher(
		baseUrl+"/topic/",
		baseUrl+"/hub/",
		websub.PWithPostBodyAsContent(true),
		websub.PAdvertiseInvalidTopics(true),
	)

	// register handlers
	mux.Handle("/sub/", http.StripPrefix("/sub/", s))
	mux.Handle("/topic/", http.StripPrefix("/topic/", p))

	// listen for requests
	go http.ListenAndServe("127.0.0.1:3033", mux)
	fmt.Println("Listening on 127.0.0.1:3033")

	<-make(chan struct{})
}
