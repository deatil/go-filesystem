package adapter

import (
	"io"
)

type Streamer struct {
	io.Reader
}

func NewStreamer(reader io.Reader) *Streamer {
	streamer := new(Streamer)
	streamer.Reader = reader

	return streamer
}

func (s *Streamer) Close() error {
	return nil
}
