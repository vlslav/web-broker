package repository

import (
	"github.com/vlslav/web-broker/internal/app/service"
)

func New(newrepo string, repotype string) service.Repo {
	switch repotype {
	case "file":
		return NewFileRepo(newrepo)
	case "memory":
		return NewMemRepo()
	case "postgres":
		return NewPgRepo()
	default:
		return nil
	}
}
