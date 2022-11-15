package controllers

import (
	"net/http"

	"github.com/andreeabizu/xm-task/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To Andreea's API")

}
