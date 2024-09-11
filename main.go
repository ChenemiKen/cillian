package main

import (
	"os"
	"fmt"
	"log"
	"net/http"
	"github.com/joho/godotenv"
)

func main(){
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
	port, exist := os.LookupEnv("PORT")
	if !exist {
		log.Fatal("PORT not set in .env")
	}

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request){
		fmt.Println("Hello world!")
	})

	_= http.ListenAndServe(":"+port, nil)
}