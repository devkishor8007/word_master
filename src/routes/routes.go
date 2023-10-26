package routes

import (
	"github.com/devkishor8007/word_master/src/controllers"
	"github.com/devkishor8007/word_master/src/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

func SetupRoutes(router *mux.Router) {
	r := router.PathPrefix("/api/v1").Subrouter()
	r.HandleFunc("/signup", controllers.Register).Methods("POST")
	r.HandleFunc("/signin", controllers.Login).Methods("POST")
	r.Handle("/protected", middleware.RequiredAuth(http.HandlerFunc(controllers.Home))).Methods("GET")
	r.Handle("/profile", middleware.RequiredAuth(http.HandlerFunc(controllers.ViewProfile))).Methods("GET")
	r.HandleFunc("/", controllers.Home).Methods("GET")

	// category endpoint
	r.HandleFunc("/category", controllers.GetCategories).Methods("GET")
	r.HandleFunc("/category", controllers.CreateCategory).Methods("POST")

	// articles endpoint
	r.HandleFunc("/articles", controllers.GetArticles).Methods("GET")
	r.Handle("/articles", middleware.RequiredAuth(http.HandlerFunc(controllers.CreateArticle))).Methods("POST")
	r.Handle("/articles/contributors", middleware.RequiredAuth(http.HandlerFunc(controllers.GetOwnArticles))).Methods("GET")
	r.Handle("/article/{article_id:[0-9]}", middleware.RequiredAuth(http.HandlerFunc(controllers.DeleteArticle))).Methods("DELETE")

	// comment endpoint
	r.HandleFunc("/comment", controllers.GetComments).Methods("GET")
	r.Handle("/comment", middleware.RequiredAuth(http.HandlerFunc(controllers.CreateComment))).Methods("POST")
}
