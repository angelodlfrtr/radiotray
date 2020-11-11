package tray

import (
	"fmt"
	"os"
	"time"

	"github.com/angelodlfrtr/radiotray/cmd/config"
	"github.com/angelodlfrtr/radiotray/cmd/launchd"
	"github.com/angelodlfrtr/radiotray/cmd/player"
	"github.com/getlantern/systray"
)

func onReady(cfg *config.Config, cbFunc func()) func() {
	return func() {
		// systray.SetTitle("Radio tray")
		systray.SetTooltip("Play radios your loving")

		setDefaultIcon := func() {
			systray.SetIcon(Icon())
		}

		setPlayingIcon := func() {
			systray.SetIcon(IconRedBytes())
		}

		setDefaultIcon()

		// Prepare radio menu items
		radioMenuItems := []*systray.MenuItem{}

		// Stop control
		stopItem := systray.AddMenuItem("Stop", "Stop playing current radio")
		stopItem.Disable()

		// Unckeck radio menu items
		uncheckRadioItems := func() {
			for _, mi := range radioMenuItems {
				mi.Uncheck()
			}
		}

		// Listen for stop click
		go func() {
			for {
				select {
				case <-stopItem.ClickedCh:
					setDefaultIcon()
					uncheckRadioItems()
					player.Stop()
					stopItem.Disable()
				}
			}
		}()

		// Separator
		systray.AddSeparator()

		handleRadioItemEvents := func(itm *systray.MenuItem, radio *config.Radio) {
			for {
				select {
				case <-itm.ClickedCh:
					uncheckRadioItems()
					itm.Check()
					stopItem.Enable()
					setPlayingIcon()
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

		// Settings / quit
		systray.AddSeparator()

		// Install / Uninstall launchd service
		lauchdServiceMenuItem := systray.AddMenuItem(
			"Run at startup",
			"Enable / disable run at starup",
		)

		// Check if lauchd service exist and is enable
		// @TODO: check working
		lnchdSrvc := launchd.NewService("radiotray")
		if lnchdSrvc.HasPlist() {
			lauchdServiceMenuItem.Check()
		}

		go func() {
			for {
				select {
				case <-lauchdServiceMenuItem.ClickedCh:
					lnchdSrvc := launchd.NewService("radiotray")
					if lnchdSrvc.HasPlist() {
						lnchdSrvc.Uninstall(5 * time.Second)
						lauchdServiceMenuItem.Uncheck()
					} else {
						// Create & install plist
						plist := launchd.NewPlist(
							"radiotray",
							os.Args[0],
							nil,
							nil,
							"",
							"",
						)

						lnchdSrvc.Install(plist, 5*time.Second)
						lauchdServiceMenuItem.Check()
					}
				}
			}
		}()

		// Quit app
		mQuit := systray.AddMenuItem("Quit", "Quit radio tray")
		go func() {
			<-mQuit.ClickedCh
			os.Exit(0)
		}()

		cbFunc()
	}
}
