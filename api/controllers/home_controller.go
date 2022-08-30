package controllers

import (
	"github.com/selimaytac/TaskRegisterer/api/responses"
	"net/http"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To the API")
}