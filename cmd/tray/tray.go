// Package tray implement os systray for macos, linux and windows
package tray

import (
	"log"

	"github.com/angelodlfrtr/radiotray/cmd/config"
	"github.com/getlantern/systray"
)

// RadioSelectCH chan to select radio
var RadioSelectCH chan *config.Radio

// StopCH to stop player
var StopCH chan bool

// SettingsCH to open settings interface
var SettingsCH chan bool

// Init tray menu
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

// SetDefaultIcon on tray menu
func SetDefaultIcon() {
	systray.SetIcon(Icon())
}

// SetPlayingIcon in tray menu
func SetPlayingIcon() {
	systray.SetIcon(IconRedBytes())
}

// SetTooltip in tray menu
func SetTooltip(tooltip string) {
	systray.SetTooltip(tooltip)
}

// SetTitle in tray menu
func SetTitle(title string) {
	systray.SetTitle(title)
}

// DisableStopItem disable stop menu item
func DisableStopItem() {
	stopItem.Disable()
}

// EnableStopItem enable stop menu item
func EnableStopItem() {
	stopItem.Enable()
}
