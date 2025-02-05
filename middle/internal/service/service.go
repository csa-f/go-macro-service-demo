package service

import (
	"context"

	middleService "github.com/csa-f/macro-service/common/proto/middle/service"
	"github.com/csa-f/macro-service/middle/internal/repository"
	log "github.com/sirupsen/logrus"
)

type MiddleServer struct {
	middleService.UnimplementedMiddleServer
	repo *repository.Repository
}

func NewMiddleServer(repo *repository.Repository) *MiddleServer {
	return &MiddleServer{
		repo: repo,
	}
}

func (s *MiddleServer) Ping(_ context.Context, req *middleService.PingReq) (*middleService.PongRes, error) {
	log.Println(req.Code)
	return &middleService.PongRes{Code: 2}, nil
}
