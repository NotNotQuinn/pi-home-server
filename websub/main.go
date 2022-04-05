package main

import (
	"fmt"
	"net/http"

	"github.com/notnotquinn/go-websub"
	"github.com/notnotquinn/pi-home-server/config"
)

func main() {
	c, err := config.Load("./config.yaml")
	if err != nil {
		panic(err)
	}

	http.ListenAndServe(fmt.Sprintf(":%d", c.Hub.Port),
		websub.NewHub(c.Hub.URL,
			websub.HubAllowPostBodyAsContent(true),
			websub.HubExposeTopics(true),
			websub.HubWithUserAgent("pi-home-server-hub"),
			websub.HubWithHashFunction("sha512"),
		),
	)

	<-make(chan struct{})
}
