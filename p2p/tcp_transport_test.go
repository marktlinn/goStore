package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransportInit(t *testing.T) {
	listenAddr := ":8000"
	tp := NewTCPTransport(listenAddr)

	assert.Equal(t, tp.listenAddr, listenAddr)
}

func TestTCPTransportAccept(t *testing.T) {
	listenAddr := ":8000"
	tp := NewTCPTransport(listenAddr)

	assert.Nil(t, tp.ListenAccept())
}
