package main

import (
	"encoding/json"
	"fmt"

	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
	"github.com/asticode/go-astilog"
	"github.com/pkg/errors"
)

var gMenuOptions = []*astilectron.MenuItemOptions{
	{
		Label: astilectron.PtrStr("File"),
		SubMenu: []*astilectron.MenuItemOptions{
			{
				Label:   astilectron.PtrStr("About"),
				OnClick: onAbout,
			},
			{Role: astilectron.MenuItemRoleClose},
		},
	},
}

func onAbout(e astilectron.Event) (deleteListener bool) {
	htmlAbout := fmt.Sprintf(`
Welcome on <b>%s</b>!<br>
This is using the go-astilectron as backend and vue.js as frontend.<br>
BuiltAt: %s`,
		AppName, BuiltAt)
	if err := SendMessage("menu.about", htmlAbout, func(m *bootstrap.MessageIn) {
		// Unmarshal payload
		var s string
		if err := json.Unmarshal(m.Payload, &s); err != nil {
			astilog.Error(errors.Wrap(err, "unmarshaling payload failed"))
			return
		}
		astilog.Infof("About modal has been displayed and payload is %s!", s)
	}); err != nil {
		astilog.Error(errors.Wrap(err, "sending about event failed"))
	}
	return
}
