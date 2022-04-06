package main

import (
	"context"
	"fmt"

	"github.com/notnotquinn/pi-home-server/keylogger"
)

func main() {
	if err := keylogger.Run(context.Background()); err != nil {
		panic(fmt.Sprintf("%#v", err))
	}
}
