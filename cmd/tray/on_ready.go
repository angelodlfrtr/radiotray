package tray

import (
	"fmt"
	"os"

	"github.com/angelodlfrtr/radiotray/cmd/config"
	"github.com/angelodlfrtr/radiotray/cmd/player"
	"github.com/getlantern/systray"
)

func onReady(cfg *config.Config, cbFunc func()) func() {
	return func() {
		// systray.SetTitle("Radio tray")
		systray.SetTooltip("Play radios your loving")
		systray.SetTemplateIcon(Icon(), Icon())

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

		systray.AddSeparator()

		// Quit app
		mQuit := systray.AddMenuItem("Quit", "Quit radio tray")
		go func() {
			<-mQuit.ClickedCh
			os.Exit(0)
		}()

		cbFunc()
	}
}
