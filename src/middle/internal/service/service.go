package service

import (
	"context"
	middleService "github.com/csa-f/go-macro-service-demo/common/proto/middle/service"
	"github.com/csa-f/go-macro-service-demo/middle/internal/repository"
	log "github.com/sirupsen/logrus"
)

type MiddleServer struct {
	middleService.UnimplementedMiddleServer
	repo repository.Repository
}

func NewMiddleServer() *MiddleServer {
	return &MiddleServer{}
}

func (s *MiddleServer) Ping(_ context.Context, in *middleService.PingReq) (*middleService.PongRes, error) {
	log.Println(in.Code)
	return &middleService.PongRes{Code: 2}, nil
}
