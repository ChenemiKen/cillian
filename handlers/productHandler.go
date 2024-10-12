package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/chenemiken/cillian/models"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("post create product")
	var product models.Product
	json.NewDecoder(r.Body).Decode(&product)
	fmt.Println(product)

	json.NewEncoder(w).Encode(product)
}
