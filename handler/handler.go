package handler

import (
	"encoding/json"
	"fmt"
	models "muxtest/models"
	"net/http"
)

type Handler struct {
	IdsModel models.Ids `json:"ids"`
}

func (handler *Handler) GetIds(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Printing Ids: ", handler.IdsModel)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(handler.IdsModel)
}

func (handler *Handler) PostIds(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	id := models.Id{}
	decoder.Decode(&id)
	fmt.Println("Received Id: ", id)
	handler.IdsModel.Ids = append(handler.IdsModel.Ids, id)
	handler.IdsModel.Count = len(handler.IdsModel.Ids)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(id)
}
