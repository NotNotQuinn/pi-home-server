package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/NotNotQuinn/keylogger"
	"github.com/notnotquinn/pi-home-server/config"
	"github.com/notnotquinn/pi-home-server/keylogger/keys"
	"github.com/notnotquinn/wts"
)

// FindKeyboardDevice gets the keyboard device path to listen to
func FindKeyboardDevice() string {
	path := "/sys/class/input/event%d/device/name"
	resolved := "/dev/input/event%d"

	for i := 0; i < 255; i++ {
		buff, err := ioutil.ReadFile(fmt.Sprintf(path, i))
		if err != nil {
			continue
		}

		deviceName := strings.ToLower(string(buff))

		//fmt.Printf("Device: %q\n", deviceName)

		if deviceName == "mosart semi. 2.4g wireless keypad\n" {
			return fmt.Sprintf(resolved, i)
		}
	}

	return ""
}

func main() {
	// Load config
	c, err := config.Load("./config.yaml")
	if err != nil {
		panic(err)
	}

	// Get a key logger
	logger, err := keylogger.New(FindKeyboardDevice())
	if err != nil {
		panic(err)
	}

	// Create a node
	n := wts.NewNode(c.Keylogger.URL, c.Hub.URL)
	go http.ListenAndServe(fmt.Sprintf(":%d", c.Keylogger.Port), n)

	// Add an emitter to create events
	events := make(chan keys.Event)
	wts.AddEmitter(n, wts.NewBasicEmitter("wait a sec", events))

	// Read keys
	ch := logger.Read()
	for v := range ch {
		if keystr := v.KeyString(); v.Type == keylogger.EvKey &&
			keystr != "NUM_LOCK" && (v.Value == 1 || v.Value == 0) {
			key, err := keyFromString(keystr)
			if err != nil {
				continue
			}

			// Send key events
			go func(key keys.Key, v keylogger.InputEvent) {
				fmt.Printf("%s - %d", key, v.Value)

				events <- keys.Event{
					Key:  key,
					Mode: keys.KeyMode(v.Value),
				}
			}(key, v)
		}
	}

	close(ch)
}

// keyFromStr gets a key from the string output by the keylogger package
//
// If the key is not recognized (or not supported) an error is returned.
func keyFromString(keyStr string) (key keys.Key, err error) {
	switch keyStr {

	// Number keys
	case "INS":
		key = keys.NUMPAD_0
	case "END_1":
		key = keys.NUMPAD_1
	case "DOWN":
		key = keys.NUMPAD_2
	case "PGDN_3":
		key = keys.NUMPAD_3
	case "LEFT_4":
		key = keys.NUMPAD_4
	case "5":
		key = keys.NUMPAD_5
	case "RT_ARROW_6":
		key = keys.NUMPAD_6
	case "HOME":
		key = keys.NUMPAD_7
	case "UP_8":
		key = keys.NUMPAD_8
	case "PGUP_9":
		key = keys.NUMPAD_9

	// Math keys
	case "+":
		key = keys.PLUS
	case "-":
		key = keys.DASH
	case "*":
		key = keys.ASTERISK
	case "/":
		key = keys.SLASH
	case "=":
		key = keys.EQUALS

	// Arrow keys
	case "Up":
		key = keys.UP
	case "Down":
		key = keys.DOWN
	case "Left":
		key = keys.LEFT
	case "Right":
		key = keys.RIGHT

	// Navigation keys
	case "Home":
		key = keys.HOME
	case "End":
		key = keys.END
	case "PgUp":
		key = keys.PAGEUP
	case "PgDn":
		key = keys.PAGEDOWN
	case "Insert":
		key = keys.INSERT
	case "Del":
		key = keys.DELETE

	// Other keys
	case "R_ENTER":
		key = keys.ENTER
	case "BS":
		key = keys.BACKSPACE
	case "DEL":
		key = keys.DOT

	default:
		err = errors.New("key not recognized")
	}

	return
}
