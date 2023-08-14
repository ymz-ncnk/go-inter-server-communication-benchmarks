// Code generated by Kitex v0.6.1. DO NOT EDIT.

package kitexechoservice

import (
	server "github.com/cloudwego/kitex/server"
	echo "github.com/ymz-ncnk/go-inter-server-communication-benchmarks/kitex-mux_ttheader_protobuf/kitex_gen/echo"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler echo.KitexEchoService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
