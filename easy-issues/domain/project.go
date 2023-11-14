package domain

import "net/http"

type Project struct {
	Id          int64
	Name        string
	OwnerId     int64
	Description string
}

type ProjectController interface {
	List(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Show(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type ProjectService interface {
	Project(id int64) (*Project, error)
	Projects() ([]*Project, error)
	Create(p *Project) error
	Delete(id int64) error
}

type ProjectRepository interface {
	GetById(id int64) (*Project, error)
	All() ([]*Project, error)
	Create(p *Project) error
	Delete(id int64) error
}
