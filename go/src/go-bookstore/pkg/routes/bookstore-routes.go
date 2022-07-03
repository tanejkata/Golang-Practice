package routes

import (
	"github.com/gorilla/mux"
	"github.com/Tanej98/go-bootstore/pkg/controllers"
)

var RegisterRoutes = func (router *mux.Router)  {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.GetBookByID).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.UpdateBook).Methods("PATCH")
	router.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("DELETE")

}