// Package config implement config struct and loading
package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Radios []*Radio `yaml:"radios"`

	// ChangedCH on config change
	ChangedCH chan bool `yaml:"-"`
}

func New() *Config {
	cfg := &Config{}
	cfg.SetDefaults()
	cfg.ChangedCH = make(chan bool)
	return cfg
}

// Load yaml config from fs
func (cfg *Config) Load() error {
	if fileExists(ConfigFilePath()) {
		yamlBytes, err := ioutil.ReadFile(ConfigFilePath())
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

func (cfg *Config) Write() error {
	yamlBytes, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(ConfigFilePath(), yamlBytes, 0644)
}

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

func ConfigFilePath() string {
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
