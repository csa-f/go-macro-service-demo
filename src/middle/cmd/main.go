package main

import (
	"fmt"
	"net"

	middleService "github.com/csa-f/go-macro-service-demo/common/proto/middle/service"
	"github.com/csa-f/go-macro-service-demo/middle/config"
	"github.com/csa-f/go-macro-service-demo/middle/internal/repository"
	"github.com/csa-f/go-macro-service-demo/middle/internal/service"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	appConfig := config.Get()
	serverConfig := appConfig.Server
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", serverConfig.Ip, serverConfig.Port))
	if err != nil {
		panic(err)
	}

	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)
	server := service.NewMiddleServer(repository.NewRepository(nil, appConfig))
	middleService.RegisterMiddleServer(s, server)
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
