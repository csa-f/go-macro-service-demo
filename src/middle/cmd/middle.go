package main

import (
	middleService "github.com/csa-f/go-macro-service-demo/common/proto/middle/service"
	"github.com/csa-f/go-macro-service-demo/middle/internal/service"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":18080")
	if err != nil {
		panic(err)
	}

	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)
	middleService.RegisterMiddleServer(s, service.NewMiddleServer())
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
