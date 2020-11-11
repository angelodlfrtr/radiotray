// Package tray implement os systray for macos, linux and windows
package tray

import (
	"log"

	"github.com/angelodlfrtr/radiotray/cmd/config"
	"github.com/getlantern/systray"
)

var RadioSelectCH chan *config.Radio

func Init(cfg *config.Config, cbFunc func()) {
	RadioSelectCH = make(chan *config.Radio)
	systray.Run(onReady(cfg, cbFunc), onExit)
}

func onExit() {
	log.Println("Quitting.")
}
