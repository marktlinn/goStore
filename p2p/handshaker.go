package p2p

type HandshakeFunc func(any) error

func BypassHandshakeFunc(any) error { return nil }
