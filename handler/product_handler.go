package handler

import (
	"encoding/json"
	"local_my_api/model"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	var products = []model.Product{}

	json.NewEncoder(w).Encode(products)
}