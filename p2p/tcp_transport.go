package p2p

import (
	"fmt"
	"log"
	"net"
	"sync"
)

type TCPTransport struct {
	listener   net.Listener
	listenAddr string
	peerLock   sync.RWMutex
	peers      map[net.Addr]Peer
}

func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		listenAddr: listenAddr,
	}
}

func (t *TCPTransport) ListenAccept() error {
	var err error

	t.listener, err = net.Listen("tcp", t.listenAddr)
	if err != nil {
		return fmt.Errorf("failed to Listen on TCP %s: %w", t.listenAddr, err)
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
	log.Printf("new connection established: %+v\n", conn)
}
