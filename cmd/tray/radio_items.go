package tray

import "github.com/getlantern/systray"

var radioMenuItems []*systray.MenuItem

// UncheckRadioItems radio menu items
func UncheckRadioItems() {
	for _, mi := range radioMenuItems {
		mi.Uncheck()
	}
}
