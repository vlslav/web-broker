package repository

import (
	"github.com/1r0npipe/web-broker/pkg/model"
)

type FileRepo struct {
	fileName string
}

func NewFileRepo(fileName string) *FileRepo {
	return &FileRepo{fileName: fileName}
}

func (fr *FileRepo) Get(getReq string) (string, error) {
	// TODO: impl
	return "", nil
}

func (fr *FileRepo) Put(putReq *model.PutValue) error {
	// TODO: impl
	return nil
}
