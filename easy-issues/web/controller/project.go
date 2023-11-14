package controller

import (
	"net/http"

	"github.com/bug-free-happiness/easy-issues/domain"
)

type ProjectController struct {
	ProjectService domain.ProjectService
}

// Create implements domain.ProjectController.
func (*ProjectController) Create(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// Delete implements domain.ProjectController.
func (*ProjectController) Delete(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// List implements domain.ProjectController.
func (*ProjectController) List(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// Show implements domain.ProjectController.
func (*ProjectController) Show(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func NewProjectController(projectService domain.ProjectService) domain.ProjectController {
	return &ProjectController{projectService}
}
