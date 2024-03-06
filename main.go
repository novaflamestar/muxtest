package main

import (
	"fmt"
	handler "muxtest/handler"
	modles "muxtest/models"

	"net/http"
)

func main() {
	mux := http.NewServeMux()
	handler := handler.Handler{IdsModel: modles.NewIds()}
	mux.HandleFunc("GET /ids", handler.GetIds)
	mux.HandleFunc("POST /ids", handler.PostIds)

	fmt.Println("Starting server")
	http.ListenAndServe("localhost:8024", mux)
}
