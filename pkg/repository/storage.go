package repository

import (
	"github.com/1r0npipe/web-broker/pkg/model"
)

type (
	Storager interface {
		New() (interface{}, error)
	}

	Storage struct {
		stor Storager
	}
)

func New(stor Storager) (*Storage, error) {
	return &Storage{stor: stor}, nil
}

func (s *Storage) Get(str string) (string, error) {
	return str, nil
}

func (s *Storage) Put(putReq *model.PutValue) error {
	return nil
}
