package adapter

import (
	"bytes"
	"io"
	"testing"
)

func Test_NewStreamer(t *testing.T) {
	buf := bytes.NewBufferString("")
	streamer := NewStreamer(buf)

	var _ io.ReadCloser = streamer
}