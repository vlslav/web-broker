package service

import (
	"github.com/vlslav/web-broker/internal/pkg/model"
)

type (
	Repo interface {
		Get(key string) (string, error)
		Put(putReq *model.PutValue) error
	}

	Client interface {
		GetInfo() string
	}

	Service struct {
		repo   Repo
		client Client
	}
)

func New(repo Repo, client Client) *Service {
	return &Service{repo: repo, client: client}
}
