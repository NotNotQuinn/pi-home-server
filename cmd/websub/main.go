package main

import "github.com/notnotquinn/pi-home-server/websub"

func main() {
	if err := websub.Run(); err != nil {
		panic(err)
	}
}
