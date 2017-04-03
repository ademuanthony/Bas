package controllers

import (
	"net/http"
)

type ContextAwareRequest struct {
	http.Request
}

func (r ContextAwareRequest) GetCurrentUserId() int64 {
	userId := r.Context().Value("UserId")

	return userId.(int64)
}

func (r ContextAwareRequest) GetPermissions() []int64 {
	permissions := r.Context().Value("Permissions");

	return permissions.([]int64);
}