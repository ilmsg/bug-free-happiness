package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bug-free-happiness/easy-issues/domain"
)

type UserController struct {
	UserService domain.UserService
}

// Create implements domain.UserController.
func (c *UserController) Create(w http.ResponseWriter, r *http.Request) {

	// r.Form.Get("")

	message := struct{ msg string }{msg: "create successfuly"}
	createMessage, err := json.Marshal(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(createMessage)
}

// Delete implements domain.UserController.
func (c *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	// if id, ok := r.URL.Query(); !ok {
	// 	http.Error(w, "id exists", http.StatusInternalServerError)
	// 	return
	// }

	// queryValue := r.URL.Query()
	// id := queryValue.Get("id")

	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := c.UserService.Delete(userId); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	message := struct{ msg string }{msg: "delete successfuly"}
	deleteMessage, err := json.Marshal(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(deleteMessage)
}

// List implements domain.UserController.
func (c *UserController) List(w http.ResponseWriter, r *http.Request) {
	users, err := c.UserService.Users()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userJson, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(userJson)
}

// Show implements domain.UserController.
func (c *UserController) Show(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := c.UserService.User(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userJson, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(userJson)
}

func NewUserController(userService domain.UserService) domain.UserController {
	return &UserController{userService}
}
