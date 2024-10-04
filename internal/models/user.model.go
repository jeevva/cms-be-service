package models

import "time"

type User struct {
	Id           string 		`db:"id" form:"id" json:"id"`
	First_name   string 		`db:"first_name" form:"first_name" json:"first_name"`
	Last_name    string 		`db:"last_name" form:"last_name" json:"last_name"`
	Email        string 		`db:"email" form:"email" json:"email"`
	Password     string 		`db:"password" form:"password" json:"password"`
	Role         string 		`db:"role" form:"role" json:"role"`
	Image        string 		`db:"image" form:"image" json:"image"`
	Phone_number string 		`db:"phone_number" form:"phone_number" json:"phone_number"`
	Created_at   *time.Time `db:"created_at" form:"created_at" json:"created_at"`
	Updated_at   *time.Time `db:"updated_at" form:"updated_at" json:"updated_at"`
}