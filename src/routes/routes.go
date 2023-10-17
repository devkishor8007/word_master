package routes

import (
	"github.com/devkishor8007/word_master/src/controllers"
	"github.com/devkishor8007/word_master/src/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/signup", controllers.Register).Methods("POST")
	r.HandleFunc("/signin", controllers.Login).Methods("POST")
	r.Handle("/protected", middleware.RequiredAuth(http.HandlerFunc(controllers.Home))).Methods("GET")
	r.Handle("/profile", middleware.RequiredAuth(http.HandlerFunc(controllers.ViewProfile))).Methods("GET")
	r.HandleFunc("/", controllers.Home).Methods("GET")

	// category endpoint
	r.HandleFunc("/category", controllers.GetCategories).Methods("GET")
	r.HandleFunc("/category", controllers.CreateCategory).Methods("POST")

	// articles endpoint
	r.HandleFunc("/article", controllers.GetArticles).Methods("GET")
	r.Handle("/article", middleware.RequiredAuth(http.HandlerFunc(controllers.CreateArticle))).Methods("POST")

	// comment endpoint
	r.HandleFunc("/comment", controllers.GetComments).Methods("GET")
	r.Handle("/comment", middleware.RequiredAuth(http.HandlerFunc(controllers.CreateComment))).Methods("POST")
}
