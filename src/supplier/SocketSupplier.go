package supplier

import (
	"net"
)

// SocketSupplier is a supplier of net.Conn instances.
type SocketSupplier interface {
	Supplier[[]byte]
	GetError() error
}

// NewSocketSupplier creates a new SocketSupplier instance.
func NewSocketSupplier(conn net.Conn, readFunc func(net.Conn) ([]byte, error)) SocketSupplier {
	var supp = &DefaultNetSocketSupplier{conn: conn, readFunc: readFunc}
	var nextRes, err = readFunc(conn)
	supp.actualThisResult = nextRes
	supp.Error = err
	return supp
}

// Default impl. Reads "1 ahead" to be able to provide HasNext.
type DefaultNetSocketSupplier struct {
	SocketSupplier
	Error            error
	conn             net.Conn
	readFunc         func(net.Conn) ([]byte, error)
	actualThisResult []byte
}

func (s *DefaultNetSocketSupplier) HasNext() bool {
	return s.Error == nil
}

func (s *DefaultNetSocketSupplier) Next() []byte {
	var res = s.actualThisResult
	var nextRes, err = s.readFunc(s.conn)
	s.actualThisResult = nextRes
	s.Error = err
	return res
}
