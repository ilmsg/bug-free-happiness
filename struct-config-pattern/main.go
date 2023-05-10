package main

import "fmt"

type OptFunc func(*Opts)

type Opts struct {
	id      string
	maxConn int
	tls     bool
}

func defaultOpts() Opts {
	return Opts{
		id:      "default",
		maxConn: 10,
		tls:     false,
	}
}

func withTLS(opts *Opts) {
	opts.tls = true
}

func withID(id string) OptFunc {
	return func(opt *Opts) {
		opt.id = id
	}
}

func withMaxConn(n int) OptFunc {
	return func(opts *Opts) {
		opts.maxConn = n
	}
}

type Server struct {
	Opts
}

func NewServer(opts ...OptFunc) *Server {
	opt := defaultOpts()
	for _, fn := range opts {
		fn(&opt)
	}
	return &Server{
		Opts: opt,
	}
}

func main() {
	s := NewServer()
	fmt.Printf("%+v\n", s)

	s = NewServer(
		withTLS,
		withID("xyz"),
		withMaxConn(99),
	)
	fmt.Printf("%+v\n", s)
}
