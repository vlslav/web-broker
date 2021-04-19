package endpoint

import "net/http"

func (pHttp *PublicHTTP) getFromQueue() http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		// TODO: parse req and call queueSvc.Get(...)
	}
}
