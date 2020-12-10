package model

type User struct {
	UserId       int    "json:userId"
	FirstName    string "json:firstName"
	LastName     string "json:lastName"
	Email        string "json:email"
	MobileNumber int    "json:mobileNumber"
	Address      string "json:address omitempty"
	City         string "json:city"
	Role         string "json:role"
	Status       string "json:status"
	CreationDate string "json:createdOn"
}
