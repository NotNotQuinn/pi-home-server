package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/notnotquinn/go-websub"
	"github.com/notnotquinn/pi-home-server/config"
)

var log = websub.Logger()

func main() {
	c, err := config.Load("./config.yaml")
	if err != nil {
		panic(err)
	}

	h := websub.NewHub(c.Hub.URL,
		websub.HubAllowPostBodyAsContent(true),
		websub.HubExposeTopics(true),
		websub.HubWithUserAgent("pi-home-server-hub"),
		websub.HubWithHashFunction("sha512"),
	)

	h.AddSniffer("", func(topic, contentType string, body io.Reader) {
		content, err := io.ReadAll(body)
		if err != nil {
			log.Err(err).Msg("could not read publish body")
		}
		fmt.Printf("Publish: %s (%s)\n  %s\n", topic, contentType, string(content))
	})

	http.ListenAndServe(fmt.Sprintf(":%d", c.Hub.Port), h)

	<-make(chan struct{})
}
