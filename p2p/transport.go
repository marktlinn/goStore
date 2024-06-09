package p2p

// Transport represents the interface defining the communication between entities over the network.
// This encompasses various methods of communication e.g UDP, TCP, sockets etc.
type Transport interface {
	ListenAccept() error
}

// Peer represents the interface of a remote entity/node.
type Peer interface{}
