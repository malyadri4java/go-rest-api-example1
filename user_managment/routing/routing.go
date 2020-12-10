package routing

import (
	"database/sql"
	"user_managment/controller"
	"user_managment/driver"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var db *sql.DB

func init() {
	gotenv.Load()
}

func Router() *mux.Router {
	db = driver.CreateConnection()
	control := controller.Controller{}
	router := mux.NewRouter()

	router.HandleFunc("/users", control.GetUsers(db)).Methods("GET")
	router.HandleFunc("/users/{id}", control.GetUser(db)).Methods("GET")
	router.HandleFunc("/users", control.AddUser(db)).Methods("POST")
	router.HandleFunc("/users/{id}", control.UpdateUser(db)).Methods("PUT")
	router.HandleFunc("/users", control.UpdateUser(db)).Methods("PUT")
	router.HandleFunc("/users/{id}", control.DeleteUser(db)).Methods("DELETE")
	return router
}
