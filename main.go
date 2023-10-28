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

	_ "github.com/devkishor8007/word_master/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title WordMaster API
// @version 1.0
// @description content management system for for bloggers and content creators
// @contact.email test@gmail.com
// @license.name Apache 2.0
// @host localhost:3002
// @BasePath /api/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	database.InitDB()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	router := mux.NewRouter()

	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), // Make sure the URL points to your generated Swagger JSON.
	))

	router.Use(middleware.RateLimitMiddleware)

	routes.SetupRoutes(router)

	router.Use(mux.CORSMethodMiddleware(router))

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	fmt.Println("Server is running on :" + port)
	baseURL := "http://localhost:" + port
	swaggerURL := baseURL + "/swagger/index.html"
	fmt.Println("Swagger Documentation URL:", swaggerURL)
	log.Fatal(server.ListenAndServe())
}
