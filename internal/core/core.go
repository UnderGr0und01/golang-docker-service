package core

import (
	"docker-service/internal/dcontainers"
)

type Controller interface {
	GetContainers() ([]dcontainers.DContainer, error)
	StartContainer(id string) error
	StopContainer(id string) error
	GetLogs(id string) (string, error)
}

func NewController() Controller {
	return &dcontainers.DContainer{}
}
