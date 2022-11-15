package seed

import (
	"log"

	models "github.com/andreeabizu/xm-task/api/model"
	"github.com/jinzhu/gorm"
)

var Companys = []models.Company{
	models.Company{
		Name:              "Company1",
		AmountOfEmployees: 200,
		Registered:        true,
		Type:              "Cooperative",
	},
	models.Company{
		Name:              "Company2",
		AmountOfEmployees: 540,
		Registered:        true,
		Type:              "Cooperative",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Company{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.Company{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range Companys {
		err = db.Debug().Model(&models.Company{}).Create(&Companys[i]).Error
		if err != nil {
			log.Fatalf("cannot seed Companys table: %v", err)
		}
	}
}
