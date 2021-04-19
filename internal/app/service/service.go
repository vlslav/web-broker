package service

import (
	"github.com/vlslav/web-broker/internal/pkg/model"
)

type (
	Repo interface {
		Get(key string) (string, error)
		Put(putReq *model.PutValue) error
	}

	Service struct {
		repo Repo
	}
)

func New(repo Repo) *Service {
	return &Service{repo: repo}
}
