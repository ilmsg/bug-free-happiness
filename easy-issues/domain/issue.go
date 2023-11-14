package domain

import "net/http"

type Issue struct {
	Id          int64
	Title       string
	Description string
	ProjectId   int64
	OwnerId     int64
	Status      Status
	Priority    Priority
}

type IssueController interface {
	List(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Show(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type IssueService interface {
	Issue(id int64) (*Issue, error)
	Issues() ([]*Issue, error)
	Create(issue *Issue) error
	Delete(id int64) error
}

type IssueRepository interface {
	GetById(id int64) (*Issue, error)
	All() ([]*Issue, error)
	Create(issue *Issue) error
	Delete(id int64) error
}
