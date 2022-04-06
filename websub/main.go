package websub

import (
	"fmt"
	"io"
	"net/http"

	"github.com/notnotquinn/go-websub"
	"github.com/notnotquinn/pi-home-server/config"
)

var log = websub.Logger()

func Run() error {
	c, err := config.Load("./config.yaml")
	if err != nil {
		return err
	}

	h := websub.NewHub(c.Hub.URL,
		websub.HubAllowPostBodyAsContent(true),
		websub.HubExposeTopics(true),
		websub.HubWithUserAgent("pi-home-server-hub"),
		websub.HubWithHashFunction("sha512"),
	)

	count := 0
	h.AddSniffer("", func(topic, contentType string, body io.Reader) {
		count++
		content, err := io.ReadAll(body)
		if err != nil {
			log.Err(err).Msg("could not read publish body")
		}
		fmt.Printf("Publish (%d): %s (%s)\n  %s\n", count, topic, contentType, string(content))
	})

	return http.ListenAndServe(fmt.Sprintf(":%d", c.Hub.Port), h)
}
