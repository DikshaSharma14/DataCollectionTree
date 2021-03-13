package main

import (
	"log"
	"net/http"

	openapi "github.com/DikshaSharma14/data-collection-tree/openapi"
)

func main() {
	log.Printf("Server started")

	InsertApiService := openapi.NewInsertApiService()
	InsertApiController := openapi.NewInsertApiController(InsertApiService)

	QueryApiService := openapi.NewQueryApiService()
	QueryApiController := openapi.NewQueryApiController(QueryApiService)

	router := openapi.NewRouter(InsertApiController, QueryApiController)

	log.Fatal(http.ListenAndServe(":9091", router))
}
