package main

import (
	"encoding/json"

	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
)

// handleMessages handles messages
func handleMessages(_ *astilectron.Window, m bootstrap.MessageIn) (payload interface{}, err error) {
	switch m.Name {
	case "window.dosomething":
		return doSomething(m)
	}
	return
}

func doSomething(m bootstrap.MessageIn) (payload interface{}, err error) {
	// Unmarshal payload
	var msg string
	if len(m.Payload) > 0 {
		// Unmarshal payload
		if err = json.Unmarshal(m.Payload, &msg); err != nil {
			payload = err.Error()
			return
		}
	}

	// do something
	return
}
