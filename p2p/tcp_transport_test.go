package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransportInit(t *testing.T) {
	opts := TCPTransportOptions{
		ListenAddr:    ":8000",
		Decoder:       DefaultDecoder{},
		HandshakeFunc: BypassHandshakeFunc,
	}

	tp := NewTCPTransport(opts)
	assert.Equal(t, tp.ListenAddr, ":8000")

	assert.Nil(t, tp.ListenAccept())
}
