package p2p

import (
	"encoding/gob"
	"io"
)

type Decoder interface {
	Decode(io.Reader, any) error
}

type GOBDecoder struct{}

func (d GOBDecoder) Decode(r io.Reader, val any) error {
	return gob.NewDecoder(r).Decode(val)
}
