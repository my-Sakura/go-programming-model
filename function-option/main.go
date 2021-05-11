package main

import (
	"crypto/tls"
	"fmt"
	"time"
)

type Server struct {
	Addr string
	Port string
	// option
	MaxConns int
	Timeout  time.Duration
	Protocol string
	TLS      *tls.Config
}

type Option func(*Server)

func New(addr, port string, options ...Option) (*Server, error) {
	srv := &Server{
		Addr:     addr,
		Port:     port,
		Protocol: "tcp",
		Timeout:  time.Second * 30,
		MaxConns: 100,
		TLS:      nil,
	}

	for _, option := range options {
		option(srv)
	}

	return srv, nil
}

func Protocol(protocol string) Option {
	return func(s *Server) {
		s.Protocol = protocol
	}
}

func MaxConns(maxConns int) Option {
	return func(s *Server) {
		s.MaxConns = maxConns
	}
}

func Timeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Timeout = timeout
	}
}

func TLS(tls *tls.Config) Option {
	return func(s *Server) {
		s.TLS = tls
	}
}

func main() {
	server, _ := New("localhost", ":8080", Protocol("TCP"), MaxConns(100), Timeout(time.Second*30), TLS(&tls.Config{}))
	fmt.Println(*server)
}
