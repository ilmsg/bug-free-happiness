package domain

import "net/http"

type User struct {
	Id      int64
	Name    string
	Surname string
	Email   string
}

type UserController interface {
	List(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Show(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type UserService interface {
	User(id int64) (*User, error)
	Users() ([]*User, error)
	Create(u *User) error
	Delete(id int64) error
}

type UserRepository interface {
	GetById(id int64) (*User, error)
	All() ([]*User, error)
	Create(u *User) error
	Delete(id int64) error
}
