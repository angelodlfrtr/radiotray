// Package tray implement os systray for macos, linux and windows
package tray

import (
	"log"

	"github.com/angelodlfrtr/radiotray/cmd/config"
	"github.com/getlantern/systray"
)

var RadioSelectCH chan *config.Radio
var StopCH chan bool
var SettingsCH chan bool

func Init(cfg *config.Config, cbFunc func()) {
	RadioSelectCH = make(chan *config.Radio)
	StopCH = make(chan bool)
	SettingsCH = make(chan bool)

	systray.Run(onReady(cfg, cbFunc), onExit)
}

func onReady(cfg *config.Config, cbFunc func()) func() {
	return func() {
		Build(cfg)
		cbFunc()
	}
}

func onExit() {
	log.Println("Quitting.")
}

func SetDefaultIcon() {
	systray.SetIcon(Icon())
}

func SetPlayingIcon() {
	systray.SetIcon(IconRedBytes())
}

func SetTooltip(tooltip string) {
	systray.SetTooltip(tooltip)
}

func SetTitle(title string) {
	systray.SetTitle(title)
}

func DisableStopItem() {
	stopItem.Disable()
}

func EnableStopItem() {
	stopItem.Enable()
}
