package main

//go:generate sqlboiler --wipe psql

import (
	"log"

	"github.com/ataleksand/gql/env"
	"github.com/ataleksand/gql/storage"

	"github.com/ataleksand/gql/server"
	_ "github.com/lib/pq"
)

func main() {
	environment, err := env.New()
	if err != nil {
		log.Fatalf("An error has occurred when initializing the Environment: %v", err)
	}

	strg, err := storage.New(environment.Storage)
	if err != nil {
		log.Fatalf("An error has occurred when initializing the Storage: %v", err)
	}

	api := server.NewApi(strg)

	srv := server.New(environment.Server, api)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("An error has occurred when listening: %v", err)
	}
}
