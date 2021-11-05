package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		type MuxRoute struct {
			Method   []string `json:"method"`
			Endpoint string   `json:"route"`
		}
		var routes []MuxRoute
		_ = router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
			methods, _ := route.GetMethods()
			endpoint, _ := route.GetPathRegexp()
			routes = append(routes, MuxRoute{
				Method:   methods,
				Endpoint: endpoint,
			})
			return nil
		})
		bytes, _ := json.Marshal(routes)
		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write(bytes)
	}).Methods("GET")
}
