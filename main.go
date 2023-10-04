package main

import (
	"fmt"
    "net/http"
	"log"

	"github.com/joho/godotenv"
	"os"
    "github.com/gorilla/mux"
	"github.com/devkishor8007/word_master/src/database"
)

func main() {
	database.InitDB()

	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	router := mux.NewRouter()

	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	}
	
	router.HandleFunc("/", handler)

	server := &http.Server{
		Addr:    ":"+ port,
		Handler: router,
	}

	fmt.Println("Server is running on :"+port)
	log.Fatal(server.ListenAndServe())
}