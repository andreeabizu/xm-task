package modeltests

import (
	"log"
	"testing"

	models "github.com/andreeabizu/xm-task/api/model"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres driver
	"gopkg.in/go-playground/assert.v1"
)

func TestFindAllCompanyies(t *testing.T) {

	err := RefreshCompanyTable()
	if err != nil {
		log.Fatal(err)
	}

	err = seedCompanies()
	if err != nil {
		log.Fatal(err)
	}

	Companys, err := CompanyInstance.FindAllCompanys(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the Companys: %v\n", err)
		return
	}
	assert.Equal(t, len(*Companys), 2)
}

func TestSaveCompany(t *testing.T) {

	err := RefreshCompanyTable()
	if err != nil {
		log.Fatal(err)
	}
	newCompany := models.Company{
		ID:                1,
		Name:              "Company6",
		AmountOfEmployees: 540,
		Registered:        true,
		Type:              "Cooperative",
	}
	savedCompany, err := newCompany.SaveCompany(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the Companys: %v\n", err)
		return
	}
	assert.Equal(t, newCompany.ID, savedCompany.ID)
	assert.Equal(t, newCompany.Name, savedCompany.Name)
	assert.Equal(t, newCompany.AmountOfEmployees, savedCompany.AmountOfEmployees)
}

func TestGetCompanyByID(t *testing.T) {

	err := RefreshCompanyTable()
	if err != nil {
		log.Fatal(err)
	}

	Company, err := seedOneCompany()
	if err != nil {
		log.Fatalf("cannot seed Companys table: %v", err)
	}
	foundCompany, err := CompanyInstance.FindCompanyByID(server.DB, Company.ID)
	if err != nil {
		t.Errorf("this is the error getting one Company: %v\n", err)
		return
	}
	assert.Equal(t, foundCompany.ID, Company.ID)
	assert.Equal(t, foundCompany.AmountOfEmployees, Company.AmountOfEmployees)
	assert.Equal(t, foundCompany.Name, Company.Name)
}

func TestUpdateACompany(t *testing.T) {

	err := RefreshCompanyTable()
	if err != nil {
		log.Fatal(err)
	}

	Company, err := seedOneCompany()
	if err != nil {
		log.Fatalf("Cannot seed Company: %v\n", err)
	}

	CompanyUpdate := models.Company{
		ID:                1,
		Name:              "Modify",
		AmountOfEmployees: 546,
	}
	updatedCompany, err := CompanyUpdate.UpdateACompany(server.DB, Company.ID)
	if err != nil {
		t.Errorf("this is the error updating the Company: %v\n", err)
		return
	}
	assert.Equal(t, updatedCompany.ID, CompanyUpdate.ID)
	assert.Equal(t, updatedCompany.AmountOfEmployees, CompanyUpdate.AmountOfEmployees)
	assert.Equal(t, updatedCompany.Name, CompanyUpdate.Name)
}

func TestDeleteACompany(t *testing.T) {

	err := RefreshCompanyTable()
	if err != nil {
		log.Fatal(err)
	}

	Company, err := seedOneCompany()

	if err != nil {
		log.Fatalf("Cannot seed Company: %v\n", err)
	}

	isDeleted, err := CompanyInstance.DeleteACompany(server.DB, Company.ID)
	if err != nil {
		t.Errorf("this is the error updating the Company: %v\n", err)
		return
	}

	assert.Equal(t, isDeleted, int64(1))
}
