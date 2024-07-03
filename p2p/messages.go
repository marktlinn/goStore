package p2p

import "net"

// Messages represents a vehicle for delivering data over a network
// between two nodes.
type RPC struct {
	Payload []byte
	From    net.Addr
}
