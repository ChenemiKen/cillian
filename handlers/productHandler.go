package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/chenemiken/cillian/models"
	"github.com/chenemiken/cillian/utils"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("post create product")
	var product models.Product
	json.NewDecoder(r.Body).Decode(&product)
	fmt.Println(product)

	if err := utils.Validate(product); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ApiResponse{
			Timestamp: time.Now(),
			Success:   false,
			Message:   "Invalid input",
			Data:      err,
		})
		return
	}

	var response = models.NewApiResponse()
	response.Data = product

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
