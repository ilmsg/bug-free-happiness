package controller

import (
	"net/http"

	"github.com/bug-free-happiness/easy-issues/domain"
)

type IssueController struct {
	IssueService domain.IssueService
}

// Create implements domain.IssueController.
func (*IssueController) Create(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// Delete implements domain.IssueController.
func (*IssueController) Delete(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// List implements domain.IssueController.
func (*IssueController) List(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// Show implements domain.IssueController.
func (*IssueController) Show(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func NewIssueController(issueService domain.IssueService) domain.IssueController {
	return &IssueController{issueService}
}
