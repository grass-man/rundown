package controllers

import (
	"net/http"

	"github.com/grass-man/rundown/api/responses"
)

// Home base home controller...
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}
