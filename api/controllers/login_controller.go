package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/andreeabizu/xm-task/api/auth"
	models "github.com/andreeabizu/xm-task/api/model"
	"github.com/andreeabizu/xm-task/api/responses"
	"github.com/andreeabizu/xm-task/api/utils/formaterror"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := server.SignIn(user.Username, user.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

func (server *Server) SignIn(username, password string) (string, error) {

	var err error

	user := models.User{}

	err = server.DB.Debug().Model(models.User{}).Where("username = ?", username).Take(&user).Error
	if err != nil {
		return "", err
	}

	if user.Password == "admin" {
		return auth.CreateToken(username)
	}

	return "", errors.New("Unauthorized")

}
