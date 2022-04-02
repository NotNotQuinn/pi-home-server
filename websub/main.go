package main

import (
	"net/http"

	"github.com/notnotquinn/go-websub"
)

func main() {
	http.ListenAndServe("127.0.0.1:3033", websub.NewHub("http://localhost:3033/",
		websub.HubAllowPostBodyAsContent(true),
		websub.HubExposeTopics(true),
	))

	<-make(chan struct{})
}
