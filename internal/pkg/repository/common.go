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
		if len(args) < 1 {
			return nil, fmt.Errorf("Not enough arguments!\n")
		}
		fileName, ok := args[0].(string)
		if ok != true {
			return nil, fmt.Errorf("Unsupported argument type!\n")
		}
		return NewFileRepo(fileName), nil
	case MemoryRepoType:
		return NewMemRepo(), nil
	case PostgresRepoType:
		return NewPgRepo(), nil
	default:
		return nil, fmt.Errorf("Unsupported repo type!\n")
	}
}
