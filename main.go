//go:generate bash -c "cd ui && npm run build"
//go:generate astilectron-bundler -v

package main

import (
	"flag"
	"time"

	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
	"github.com/asticode/go-astilog"
	"github.com/pkg/errors"
)

// Vars
var (
	AppName string // astilectron-bundler will fill AppName
	BuiltAt string // astilectron-bundler will fill BuiltAt
	gWindow *astilectron.Window
	fDebug  = flag.Bool("d", false, "enables the debug mode")
)

func main() {
	// Init
	flag.Parse()
	astilog.FlagInit()

	// Run bootstrap
	astilog.Debugf("Running app built at %s", BuiltAt)
	if err := bootstrap.Run(bootstrap.Options{
		Asset:    Asset,
		AssetDir: AssetDir,
		AstilectronOptions: astilectron.Options{
			AppName:            AppName,
			AppIconDarwinPath:  "resources/icon.icns",
			AppIconDefaultPath: "resources/icon.png",
		},
		Debug:       *fDebug,
		MenuOptions: gMenuOptions,
		OnWait: func(_ *astilectron.Astilectron, ws []*astilectron.Window, _ *astilectron.Menu, _ *astilectron.Tray, _ *astilectron.Menu) error {
			gWindow = ws[0]
			go func() {
				time.Sleep(5 * time.Second)
				if err := SendMessage("go.check.out.menu", "Don't forget to check out the menu!"); err != nil {
					astilog.Error(errors.Wrap(err, "sending check.out.menu event failed"))
				}
			}()
			return nil
		},
		RestoreAssets: RestoreAssets,
		Windows: []*bootstrap.Window{{
			Homepage:       "index.html",
			MessageHandler: handleMessages,
			Options: &astilectron.WindowOptions{
				BackgroundColor: astilectron.PtrStr("#333"),
				Center:          astilectron.PtrBool(true),
				Height:          astilectron.PtrInt(700),
				Width:           astilectron.PtrInt(700),
			},
		}},
	}); err != nil {
		astilog.Fatal(errors.Wrap(err, "running bootstrap failed"))
	}
}

// SendMessage sends a message
func SendMessage(name string, payload interface{}, cs ...bootstrap.CallbackMessage) error {
	if gWindow == nil {
		return errors.New("global window not init")
	}
	return bootstrap.SendMessage(gWindow, name, payload, cs...)
}
