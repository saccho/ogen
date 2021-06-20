package main

import (
	"context"
	"log"
	"net/http"

	"github.com/ernado/ogen/api"
	"github.com/go-chi/chi/v5"
)

type server struct{}

func (s server) PetGet(ctx context.Context, params *api.PetGetParameters) (*api.Pet, error) {
	return &api.Pet{
		ID:   params.Query.PetID,
		Name: "DOG",
	}, nil
}

func (s server) PetCreate(ctx context.Context, req *api.Pet) (*api.Pet, error) {
	req.ID = 1337
	return req, nil
}

func (s server) PetGetByName(ctx context.Context, params *api.PetGetByNameParameters) (*api.Pet, error) {
	return &api.Pet{
		ID:   1337,
		Name: params.Path.Name,
	}, nil
}

func main() {
	mux := chi.NewRouter()
	api.Register(mux, server{})
	log.Fatal(http.ListenAndServe(":8080", mux))
}