package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/chenemiken/cillian/db"
	"github.com/chenemiken/cillian/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}

	port, exist := os.LookupEnv("PORT")
	if !exist {
		log.Fatal("PORT not set in .env")
	}

	err := db.ConnectDb()
	if err != nil {
		log.Fatal("failed to connect to database: " + err.Error())
	}

	r := mux.NewRouter()

	r.HandleFunc("/", handlers.Home).Methods("GET")
	r.HandleFunc("/products", handlers.CreateProduct).Methods("POST")

	fmt.Printf("Cillian running on port :%s \n", port)

	_ = http.ListenAndServe(":"+port, r)
}
