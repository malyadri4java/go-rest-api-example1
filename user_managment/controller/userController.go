package controller

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"user_managment/model"
	userRepository "user_managment/repository"

	"github.com/gorilla/mux"
)

type Controller struct{}

var repo = userRepository.UserRepository{}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (c Controller) GetUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		users := []model.User{}
		users = repo.GetUsers(db)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}

func (c Controller) GetUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		user := model.User{}
		params := mux.Vars(req)
		id, _ := strconv.Atoi(params["id"])

		user = repo.GetUser(db, id)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}

func (c Controller) AddUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var userId int
		var user model.User
		json.NewDecoder(req.Body).Decode(&user)
		userId = repo.AddUser(db, user)
		json.NewEncoder(w).Encode(userId)
	}
}

func (c Controller) UpdateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		params := mux.Vars(req)
		userId, err := strconv.Atoi(params["id"])
		logFatal(err)

		var user model.User
		json.NewDecoder(req.Body).Decode(&user)
		if userId != 0 {
			user.UserId = userId
		}

		log.Println(" update request with user id ", user.UserId)
		userId = repo.UpdateUser(db, user)
		json.NewEncoder(w).Encode(userId)
	}
}

func (c Controller) DeleteUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var userId int
		params := mux.Vars(req)
		userId, err := strconv.Atoi(params["id"])
		logFatal(err)
		userId = repo.DeleteUser(db, userId)
		json.NewEncoder(w).Encode(userId)
	}
}
