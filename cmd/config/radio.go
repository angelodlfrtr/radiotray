package config

// Radio implement radio config entry
type Radio struct {
	Name   string `yaml:"name"`   // Custom name
	Format string `yaml:"format"` // mp3, wav, ogg, flac
	Source string `yaml:"source"` // URI to file (https://.../...stream.mp3)
}
