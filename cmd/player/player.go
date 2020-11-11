// Package playe load audio streams and play to speaker
package player

import (
	"log"
	"net/http"
	"time"

	"github.com/angelodlfrtr/radiotray/cmd/config"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

var _currentStreamer *beep.StreamSeekCloser

// Play radio to speaker => blocking function
func Play(radio *config.Radio) error {
	// Stop any previous stream play
	Stop()

	resp, err := http.Get(radio.Source)
	if err != nil {
		log.Fatal(err)
	}
	reader := resp.Body

	streamer, format, err := mp3.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	_currentStreamer = &streamer

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(streamer)

	return nil
}

func Stop() {
	if _currentStreamer != nil {
		(*_currentStreamer).Close()
		_currentStreamer = nil
	}

	speaker.Clear()
}
