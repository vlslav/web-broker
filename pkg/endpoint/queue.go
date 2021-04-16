package endpoint

import (
	"net/http"

	web_broker "github.com/1r0npipe/web-broker/pkg/web-broker"
	"github.com/gorilla/mux"
)

type queueSvc interface {
	Put(req *web_broker.PutValueReq) error
	Get(req *web_broker.GetValueReq) (*web_broker.GetValueResp, error)
}

func RegisterPublicHTTP(queueSvc queueSvc) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/{queue}", putToQueue(queueSvc)).Methods(http.MethodPut)
	r.HandleFunc("/{queue}", getFromQueue(queueSvc)).Methods(http.MethodGet)

	return r
}

func putToQueue(queueSvc queueSvc) http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		// TODO: parse req and call queueSvc.Put(...)
	}
}

func getFromQueue(queueSvc queueSvc) http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		// TODO: parse req and call queueSvc.Get(...)
	}
}
