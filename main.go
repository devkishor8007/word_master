package main

import (
	"fmt"
	"github.com/devkishor8007/word_master/src/database"
	"github.com/devkishor8007/word_master/src/middleware"
	"github.com/devkishor8007/word_master/src/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"github.com/mvrilo/go-redoc"
)

func main() {
	doc := &redoc.Redoc{
		Title:       "Example API",
		Description: "Example API Description",
		SpecFile:    "./openapi.json",
		SpecPath:    "/docs/openapi.json",
	}

	database.InitDB()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	router := mux.NewRouter()

	router.PathPrefix("/docs").Handler(doc.Handler())

	router.Use(middleware.RateLimitMiddleware)

	routes.SetupRoutes(router)

	router.Use(mux.CORSMethodMiddleware(router))

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	fmt.Println("Server is running on :" + port)
	log.Fatal(server.ListenAndServe())
}
