package repository

import (
	"fmt"
	"github.com/vlslav/web-broker/internal/app/service"
)

const (
	FileRepoType     int = 1
	MemoryRepoType   int = 2
	PostgresRepoType int = 3
)

func New(repoType int, args ...interface{}) (service.Repo, error) {
	switch repoType {
	case FileRepoType:
		fileName := args[0].(string)
		return NewFileRepo(fileName), nil
	case MemoryRepoType:
		return NewMemRepo(), nil
	case PostgresRepoType:
		return NewPgRepo(), nil
	default:
		return nil, fmt.Errorf("Unsupported repo type!\n")
	}
}
