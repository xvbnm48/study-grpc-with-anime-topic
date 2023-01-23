package transport

import (
	gt "github.com/go-kit/kit/transport/grpc"
)

type Server interface {
	Serve(addr string)
	Stop()
}

type grpcServer struct {
	add gt.Handler
	get gt.Handler
}

// func NewGRPCServer(endpoints endopoint.End)
