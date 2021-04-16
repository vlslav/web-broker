package service

import (
	"log"

	"github.com/1r0npipe/web-broker/pkg/model"
	web_broker "github.com/1r0npipe/web-broker/pkg/web-broker"
)

func (s *Service) Put(req *web_broker.PutValueReq) error {
	if err := s.repo.Put(&model.PutValue{
		Key:   req.Key,
		Value: req.Value,
	}); err != nil {
		log.Printf("service/Put: put repo err: %v", err)
		return err
	}

	return nil
}
