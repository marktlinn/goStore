package p2p

import (
	"fmt"
	"log"
	"net"
	"sync"
)

// TCPPeer represents the remote entity/node to which a connection is established via TCP.
type TCPPeer struct {
	conn net.Conn

	// outbound represents the direction of the connection
	// False = the connection is incoming
	// True = the connection is outgoing
	outbound bool
}

type TCPTransportOptions struct {
	ListenAddr    string
	Decoder       Decoder
	HandshakeFunc HandshakeFunc
}

type TCPTransport struct {
	TCPTransportOptions
	listener net.Listener

	peerLock sync.RWMutex
	peers    map[net.Addr]Peer
}

// NewTCPPeer creates a new TCPPeer and returns a reference to the newly created TCPPeer.
func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

func NewTCPTransport(opts TCPTransportOptions) *TCPTransport {
	return &TCPTransport{
		TCPTransportOptions: opts,
	}
}

func (t *TCPTransport) ListenAccept() error {
	var err error

	t.listener, err = net.Listen("tcp", t.ListenAddr)
	if err != nil {
		return fmt.Errorf("failed to Listen on TCP %s: %w", t.ListenAddr, err)
	}

	go t.handleTCPAccept()

	return nil
}

func (t *TCPTransport) handleTCPAccept() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			log.Printf("failed to accept TCP connection %s\n", err)
		}
		go t.handleTCPConnection(conn)
	}
}

func (t *TCPTransport) handleTCPConnection(conn net.Conn) {
	tcpPeer := NewTCPPeer(conn, true)

	if err := t.HandshakeFunc(tcpPeer); err != nil {
		log.Printf("failed to complete TCP handshake: %s\n", err)
		return
	}

	msg := &RPC{}
	for {
		if err := t.Decoder.Decode(conn, msg); err != nil {
			log.Printf("failed to decode from TCP %s\n", err)
			continue
		}

		msg.From = conn.RemoteAddr()
		log.Printf("new connection established: %+v\n", conn)

	}
}
