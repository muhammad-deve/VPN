package service

import (
	"github.com/muhammad-deve/vpn/internal/config"
)

type Service interface {
	Start() error
	Stop() error
}

type service struct {
	cfg *config.Config
}

func NewService(cfg *config.Config) Service {
	return &service{
		cfg: cfg,
	}
}

func (s *service) Start() error {
	return nil
}

func (s *service) Stop() error {
	return nil
}
