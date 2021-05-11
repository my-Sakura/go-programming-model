package main

import (
	"crypto/tls"
	"time"
)

// It is convenient that use the following design model to init struct
// Server is necessary field and others is option field in example
type Server struct {
	Addr string
	Port string
}

type ServerBuilder struct {
	Server
	Timeout  time.Duration
	Protocol string
	MaxConns int
	TLS      *tls.Config
}

func New(addr, port string) *ServerBuilder {
	sb := ServerBuilder{}

	return sb.New(addr, port)
}

func (sb *ServerBuilder) New(addr, port string) *ServerBuilder {
	sb.Server.Addr = addr
	sb.Server.Port = port

	return sb
}

func (sb *ServerBuilder) WithProtocol(protocol string) *ServerBuilder {
	sb.Protocol = protocol

	return sb
}

func (sb *ServerBuilder) WithTimeout(timeout time.Duration) *ServerBuilder {
	sb.Timeout = timeout

	return sb
}

func (sb *ServerBuilder) WithMaxConns(maxConns int) *ServerBuilder {
	sb.MaxConns = maxConns

	return sb
}

func (sb *ServerBuilder) WithTLS(tls *tls.Config) *ServerBuilder {
	sb.TLS = tls

	return sb
}

func main() {
	serverBuilder := New("localhost", ":8080")

	serverBuilder.
		WithProtocol("HTTPS").
		WithTimeout(time.Second * 30).
		WithMaxConns(100).
		WithTLS(&tls.Config{})
}
