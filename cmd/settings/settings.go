// Package settings contain setting window specific code
package settings

import (
	"os/exec"
	"runtime"

	"github.com/angelodlfrtr/radiotray/cmd/config"
)

// Init settings
func Init() {
	// nothing
}

// Open settings
func Open() {
	openCommand := "xdg-open"

	if runtime.GOOS == "darwin" {
		openCommand = "open"
	}

	configFilePath := config.FilePath()
	cmd := exec.Command(openCommand, configFilePath)
	cmd.Start()
}
