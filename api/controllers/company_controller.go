package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	models "github.com/andreeabizu/xm-task/api/model"
	"github.com/andreeabizu/xm-task/api/responses"
	"github.com/andreeabizu/xm-task/api/utils/formaterror"
	"github.com/gorilla/mux"
)

func (server *Server) CreateCompany(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	Company := models.Company{}
	err = json.Unmarshal(body, &Company)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	Company.Prepare()
	err = Company.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	CompanyCreated, err := Company.SaveCompany(server.DB)

	if err != nil {

		formattedError := formaterror.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, CompanyCreated.ID))
	responses.JSON(w, http.StatusCreated, CompanyCreated)
}

func (server *Server) GetCompanies(w http.ResponseWriter, r *http.Request) {

	Company := models.Company{}

	Companys, err := Company.FindAllCompanys(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, Companys)
}

func (server *Server) GetCompany(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	Company := models.Company{}
	CompanyGotten, err := Company.FindCompanyByID(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, CompanyGotten)
}

func (server *Server) UpdateCompany(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	Company := models.Company{}
	err = json.Unmarshal(body, &Company)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	Company.Prepare()
	updatedCompany, err := Company.UpdateACompany(server.DB, uint32(uid))
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, updatedCompany)
}

func (server *Server) DeleteCompany(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	Company := models.Company{}

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	_, err = Company.DeleteACompany(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
	responses.JSON(w, http.StatusNoContent, "")
}
