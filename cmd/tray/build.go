package tray

import (
	"fmt"
	"log"
	"os"

	"github.com/angelodlfrtr/radiotray/cmd/config"
	"github.com/getlantern/systray"
	autostart "github.com/protonmail/go-autostart"
)

var stopItem *systray.MenuItem

func Build(cfg *config.Config) {
	SetTooltip("Play radios your loving")
	SetDefaultIcon()

	// Stop control
	stopItem = systray.AddMenuItem("Stop", "Stop playing current radio")
	DisableStopItem()

	// Listen for stop click
	go func() {
		for {
			select {
			case <-stopItem.ClickedCh:
				select {
				case StopCH <- true:
				default:
				}
			}
		}
	}()

	// Separator
	systray.AddSeparator()

	handleRadioItemEvents := func(itm *systray.MenuItem, radio *config.Radio) {
		for {
			select {
			case <-itm.ClickedCh:
				UncheckRadioItems()
				itm.Check()
				stopItem.Enable()
				SetPlayingIcon()
				select {
				case RadioSelectCH <- radio:
				default:
				}
			}
		}
	}

	// Load radios in menu
	for _, r := range cfg.Radios {
		radioItem := systray.AddMenuItem(r.Name, fmt.Sprintf("Play %s", r.Name))
		radioMenuItems = append(radioMenuItems, radioItem)
		go handleRadioItemEvents(radioItem, r)
	}

	// Settings sep
	systray.AddSeparator()

	// Install / Uninstall launchd service
	lauchdServiceMenuItem := systray.AddMenuItem(
		"Run at login",
		"Enable / disable run at login",
	)

	// Check if lauchd service exist and is enable
	lnchApp := &autostart.App{
		Name:        "radiotray",
		DisplayName: "RadioTray",
		Exec:        []string{os.Args[0]},
	}

	if lnchApp.IsEnabled() {
		lauchdServiceMenuItem.Check()
	}

	go func() {
		for {
			select {
			case <-lauchdServiceMenuItem.ClickedCh:
				if lnchApp.IsEnabled() {
					if err := lnchApp.Disable(); err != nil {
						log.Println("ERR", err)
					}

					lauchdServiceMenuItem.Uncheck()
				} else {
					if err := lnchApp.Enable(); err != nil {
						log.Println("ERR", err)
					}

					lauchdServiceMenuItem.Check()
				}
			}
		}
	}()

	// Settings
	settingsMenuEntry := systray.AddMenuItem(
		"Settings",
		"Edit settings",
	)

	go func() {
		for {
			select {
			case <-settingsMenuEntry.ClickedCh:
				select {
				case SettingsCH <- true:
				default:
				}
			}
		}
	}()

	// Quit sep
	systray.AddSeparator()

	// Quit app
	mQuit := systray.AddMenuItem("Quit", "Quit radio tray")
	go func() {
		<-mQuit.ClickedCh
		os.Exit(0)
	}()
}