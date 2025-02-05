package server

import (
	"fmt"
	"net"

	"github.com/redis/go-redis/v9"

	middleService "github.com/csa-f/go-macro-service-demo/common/proto/middle/service"
	"github.com/csa-f/go-macro-service-demo/middle/config"
	"github.com/csa-f/go-macro-service-demo/middle/internal/repository"
	"github.com/csa-f/go-macro-service-demo/middle/internal/service"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func NewRpcServer(db *gorm.DB, redis *redis.Client, appConf *config.Config) {
	serverConf := appConf.Server
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", serverConf.Ip, serverConf.Port))
	if err != nil {
		panic(err)
	}

	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)

	repo := repository.NewRepository(db, redis, appConf)
	server := service.NewMiddleServer(repo)

	middleService.RegisterMiddleServer(s, server)

	log.Warnf("server listening at %v", listen.Addr())
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
