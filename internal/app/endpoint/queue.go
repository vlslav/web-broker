package endpoint

import (
	"net/http"

	"github.com/gorilla/mux"
	web_broker "github.com/vlslav/web-broker/pkg/web-broker"
)

type queueSvc interface {
	Put(req *web_broker.PutValueReq) error
	Get(req *web_broker.GetValueReq) (*web_broker.GetValueResp, error)
}

type PublicHTTP struct {
	Router   *mux.Router
	queueSvc queueSvc
}

func RegisterPublicHTTP(queueSvc queueSvc) *mux.Router {
	r := mux.NewRouter()
	pHttp := &PublicHTTP{
		Router:   r,
		queueSvc: queueSvc,
	}

	r.HandleFunc("/{queue}", pHttp.putToQueue()).Methods(http.MethodPut)
	r.HandleFunc("/{queue}", pHttp.getFromQueue()).Methods(http.MethodGet)

	return r
}

func (pHttp *PublicHTTP) putToQueue() http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		// TODO: parse req and call queueSvc.Put(...)
	}
}

func (pHttp *PublicHTTP) getFromQueue() http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		// TODO: parse req and call queueSvc.Get(...)
	}
}
