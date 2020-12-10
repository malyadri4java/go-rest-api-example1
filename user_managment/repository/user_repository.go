package userRepository

import (
	"database/sql"
	"fmt"
	"log"
	"user_managment/model"
)

type UserRepository struct{}

func logFatal(err error) {
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows found")
		} else {
			log.Fatal(err)
		}
	}
}

func (ur UserRepository) GetUsers(db *sql.DB) []model.User {
	users := []model.User{}
	rows, err := db.Query("SELECT * FROM users")
	logFatal(err)
	defer rows.Close()

	for rows.Next() {
		user := model.User{}
		err := rows.Scan(&user.UserId, &user.FirstName, &user.LastName, &user.Email, &user.MobileNumber, &user.Address, &user.City, &user.Role, &user.Status, &user.CreationDate)
		logFatal(err)
		users = append(users, user)
	}
	return users
}

func (ur UserRepository) GetUser(db *sql.DB, id int) model.User {
	rows := db.QueryRow("SELECT * FROM users where user_id=$1", id)
	user := model.User{}
	err := rows.Scan(&user.UserId, &user.FirstName, &user.LastName, &user.Email, &user.MobileNumber, &user.Address, &user.City, &user.Role, &user.Status, &user.CreationDate)
	logFatal(err)
	log.Println("Get Book with id", user.UserId)
	return user
}

func (ur UserRepository) AddUser(db *sql.DB, user model.User) int {
	err := db.QueryRow("insert into users (first_name, last_name, email, mobile_number, address, city, role) values ($1, $2, $3, $4, $5, $6, $7) RETURNING user_id;", user.FirstName, user.LastName, user.Email, user.MobileNumber, user.Address, user.City, user.Role).Scan(&user.UserId)
	logFatal(err)
	log.Println("User added successfully with id", user.UserId)
	return user.UserId
}

func (ur UserRepository) UpdateUser(db *sql.DB, user model.User) int {
	var UserId int
	row := db.QueryRow("update users set first_name=$1, last_name=$2, email=$3 where user_id=$4 RETURNING user_id;", &user.FirstName, &user.LastName, &user.Email, &user.UserId)
	err := row.Scan(&UserId)
	logFatal(err)
	log.Println("User updated successfully with id", UserId)
	return UserId
}

func (ur UserRepository) DeleteUser(db *sql.DB, id int) int {
	var UserId int
	row := db.QueryRow("delete from users where user_id=$1", id)
	err := row.Scan(&UserId)
	logFatal(err)
	log.Println("User deleted successfully with id", id)
	return id
}
