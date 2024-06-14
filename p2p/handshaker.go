package p2p

type HandshakeFunc func(any) error

func BypassHandshake(any) error { return nil }
