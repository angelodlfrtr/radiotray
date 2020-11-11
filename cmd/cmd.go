// Package cmd implement main cmd interface
package cmd

import (
	"log"

	"github.com/angelodlfrtr/radiotray/cmd/config"
	"github.com/angelodlfrtr/radiotray/cmd/player"
	"github.com/angelodlfrtr/radiotray/cmd/tray"
)

func Main() {
	// Load config
	cfg := config.New()
	if err := cfg.Load(); err != nil {
		log.Fatal(err)
	}

	tray.Init(cfg, func() {
		for {
			select {
			case r := <-tray.RadioSelectCH:
				player.Play(r)
			}
		}
	})
}
