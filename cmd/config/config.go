// Package config implement config struct and loading
package config

type Config struct {
	Radios []*Radio `yaml:"radios"`

	// ChangedCH on config change
	ChangedCH chan bool
}

func New() *Config {
	cfg := &Config{}
	cfg.SetDefaults()
	cfg.ChangedCH = make(chan bool)
	return cfg
}

// Load yaml config from fs
func (cfg *Config) Load() error {
	return nil
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
