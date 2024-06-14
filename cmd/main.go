package main

import (
	"fmt"
	"log"

	"github.com/marktlinn/goStore/p2p"
)

func main() {
	tcpOpts := p2p.TCPTransportOptions{
		ListenAddr:    ":8000",
		HandshakeFunc: p2p.BypassHandshake,
	}

	tr := p2p.NewTCPTransport(tcpOpts)

	if err := tr.ListenAccept(); err != nil {
		log.Fatalf("failed to listen to or accept connections %s\n", err)
	}

	fmt.Println("Working...")

	select {}
}
