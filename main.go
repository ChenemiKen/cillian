package main

import (
	"os"
	"fmt"
	"log"
	"net/http"
	"github.com/joho/godotenv"
)

func home(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello world!")
}

func main(){
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
	
	port, exist := os.LookupEnv("PORT")
	if !exist {
		log.Fatal("PORT not set in .env")
	}

	http.HandleFunc("/", home)

	fmt.Println(fmt.Sprintf("Application running on port :%s", port))
	
	_= http.ListenAndServe(":"+port, nil)
}