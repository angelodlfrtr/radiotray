// Package config implement config struct and loading
package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// Config contains radios available
type Config struct {
	Radios []*Radio `yaml:"radios"`

	// ChangedCH on config change
	ChangedCH chan bool `yaml:"-"`
}

// New config with defaults
func New() *Config {
	cfg := &Config{}
	cfg.SetDefaults()
	cfg.ChangedCH = make(chan bool)
	return cfg
}

// Load yaml config from fs
func (cfg *Config) Load() error {
	if fileExists(FilePath()) {
		yamlBytes, err := ioutil.ReadFile(FilePath())
		if err != nil {
			return err
		}

		// Try to unmarshal config
		if err := yaml.Unmarshal(yamlBytes, cfg); err == nil {
			return nil
		}
	}

	return cfg.Write()
}

// Write config to disk
func (cfg *Config) Write() error {
	yamlBytes, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(FilePath(), yamlBytes, 0644)
}

// SetDefaults on config
func (cfg *Config) SetDefaults() {
	cfg.Radios = append(cfg.Radios, &Radio{
		Name:   "TSF Jazz",
		Source: "https://tsfjazz.ice.infomaniak.ch/tsfjazz-high.mp3",
		Format: "mp3",
	})

	cfg.Radios = append(cfg.Radios, &Radio{
		Name:   "France Inter",
		Source: "https://icecast.radiofrance.fr/franceinter-midfi.mp3",
		Format: "mp3",
	})
}

// FilePath returns default config file path
func FilePath() string {
	return fmt.Sprintf("%s/.radiotray.yaml", os.Getenv("HOME"))
}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
