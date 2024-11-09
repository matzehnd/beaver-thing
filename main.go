/*

A blinker example using go-rpio library.
Requires administrator rights to run

Toggles a LED on physical pin 19 (mcu pin 10)
Connect a LED with resistor from pin 19 to ground.

*/

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/gorilla/websocket"
	"github.com/stianeikeland/go-rpio/v4"
)

var (
	pins = make(map[int]rpio.Pin)
)

type ActionTO struct {
	Action string  `json:"action"`
	Item   string  `json:"item"`
	Value  *string `json:"value,omitempty"`
}

func main() {
	u := url.URL{Scheme: "wss", Host: "xxx", Path: "/ws"}
	header := http.Header{}
	header.Add("Authorization", "Bearer xxx")

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), header)
	if err != nil {
		fmt.Println("Dial error:", err)
	}
	defer conn.Close()

	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Unmap gpio memory when done
	defer rpio.Close()

	// Set pin to output mode
	pin1.Output()
	pin1.Low()

	// Toggle pin 20 times
	for {
		// Antwort lesen
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("error on readMessage: %s", err)
		}
		var action ActionTO

		if err := json.Unmarshal(message, &action); err == nil {
			if action.Value != nil && *action.Value == "1" {
				pin1.High()
			} else {
				pin1.Low()
			}
		}

	}
}
