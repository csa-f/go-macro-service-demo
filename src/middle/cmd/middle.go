package main

import (
	"net"

	middleService "github.com/csa-f/go-macro-service-demo/common/proto/middle/service"
	"github.com/csa-f/go-macro-service-demo/middle/config"
	"github.com/csa-f/go-macro-service-demo/middle/internal/repository"
	"github.com/csa-f/go-macro-service-demo/middle/internal/service"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", ":18080")
	if err != nil {
		panic(err)
	}

	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)
	server := service.NewMiddleServer(repository.NewRepository(nil, config.C))
	middleService.RegisterMiddleServer(s, server)
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
