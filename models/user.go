package models

import "time"

//"compress/gzip"

type User struct {
	ID          int    `json:id`
	Email       string `json:email`
	Password    string `json:string`
	DateCreated string `json:datecreated`
}

func CreateTestUsers() []User {
	return []User{
		User{ID: 0, Email: "example@where.com", Password: "jellyfish", DateCreated: time.Now().String()},
		User{ID: 1, Email: "person@human.com", Password: "chocolate", DateCreated: time.Now().String()},
		User{ID: 2, Email: "krunked@krunker.io", Password: "slim jim", DateCreated: time.Now().String()},
	}
}

func NewUser(id int) User {
	return User{ID: id}
}
