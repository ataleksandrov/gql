package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const gqlPath = "/gql"

func New(settings *Settings, api *Api) *http.Server {
	r := mux.NewRouter()
	r.Handle(gqlPath, GqlHandlerFunc(api.Gqlhandler))

	return &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("127.0.0.1:%d", settings.Port),
		WriteTimeout: settings.Timeout * time.Second,
		ReadTimeout:  settings.Timeout * time.Second,
	}
}
