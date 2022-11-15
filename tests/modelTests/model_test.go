package modeltests

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/andreeabizu/xm-task/api/controllers"
	models "github.com/andreeabizu/xm-task/api/model"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}
var CompanyInstance = models.Company{}

func TestMain(m *testing.M) {
	var err error
	err = godotenv.Load(os.ExpandEnv("../../.env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}
	Database()

	os.Exit(m.Run())
}

func Database() {

	var err error

	TestDbDriver := os.Getenv("TEST_DB_DRIVER")
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("TEST_DB_HOST"), os.Getenv("TEST_DB_PORT"), os.Getenv("TEST_DB_USER"), os.Getenv("TEST_DB_NAME"), os.Getenv("TEST_DB_PASSWORD"))
	server.DB, err = gorm.Open(TestDbDriver, DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database\n", TestDbDriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database\n", TestDbDriver)
	}

}

func RefreshCompanyTable() error {
	err := server.DB.DropTableIfExists(&models.Company{}).Error
	if err != nil {
		return err
	}
	err = server.DB.AutoMigrate(&models.Company{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed table")
	return nil
}

func seedOneCompany() (models.Company, error) {

	RefreshCompanyTable()

	Company := models.Company{
		Name:              "CompanyTest1",
		AmountOfEmployees: 670,
		Registered:        true,
		Type:              "Cooperative",
	}

	err := server.DB.Model(&models.Company{}).Create(&Company).Error
	if err != nil {
		log.Fatalf("cannot seed Companies table: %v", err)
	}
	return Company, nil
}

func seedCompanies() error {

	Companies := []models.Company{
		models.Company{
			Name:              "CompanyTest3",
			AmountOfEmployees: 540,
			Registered:        true,
			Type:              "Cooperative",
		},
		models.Company{
			Name:              "CompanyTest4",
			AmountOfEmployees: 540,
			Registered:        true,
			Type:              "Cooperative",
		},
	}

	for i, _ := range Companies {
		err := server.DB.Model(&models.Company{}).Create(&Companies[i]).Error
		if err != nil {
			return err
		}
	}
	return nil
}
