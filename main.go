package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/chenemiken/cillian/handlers"
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

	http.HandleFunc("/", handlers.Home)

	fmt.Printf("Cillian running on port :%s \n", port)

	_ = http.ListenAndServe(":"+port, nil)
}
