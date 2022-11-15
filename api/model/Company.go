package models

import (
	"errors"
	"fmt"
	"html"
	"os"
	"strings"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jinzhu/gorm"
)

type Company struct {
	ID                uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Name              string `gorm:"size:15;not null;unique" json:"name"`
	Description       string `gorm:"size:100" json:"description"`
	AmountOfEmployees int32  `gorm:"not null" json:"amountOfEmployees"`
	Registered        bool   `gorm:"not null" json:"registered"`
	Type              string `gorm:"not null;" json:"type"`
}

var producer *kafka.Producer = CreateKafkaProducer()
var topic = "xm-task-kafka"

func CreateKafkaProducer() *kafka.Producer {

	producer, producerCreateErr := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": os.Getenv("KAFKA_SERVERS")})

	if producerCreateErr != nil {
		fmt.Println("Failed to create producer due to ", producerCreateErr)
		os.Exit(1)
	}
	return producer

}
func (c *Company) Prepare() {
	c.ID = 0
	c.Name = html.EscapeString(strings.TrimSpace(c.Name))
	c.Description = html.EscapeString(strings.TrimSpace(c.Description))
	c.Type = html.EscapeString(strings.TrimSpace(c.Type))

}

func (c *Company) Validate() error {
	if c.Name == "" {
		return errors.New("Required Name")
	}
	if c.Type == "" {
		return errors.New("Required Type")
	}
	return nil

}

func (c *Company) SaveCompany(db *gorm.DB) (*Company, error) {

	var err error
	err = db.Debug().Create(&c).Error
	if err != nil {
		return &Company{}, err
	}

	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte("Company " + c.Name + " created. Timestamp: " + time.Now().String()),
	}, nil)
	return c, nil
}

func (c *Company) FindAllCompanys(db *gorm.DB) (*[]Company, error) {
	var err error
	Companys := []Company{}
	err = db.Debug().Model(&Company{}).Limit(100).Find(&Companys).Error
	if err != nil {
		return &[]Company{}, err
	}

	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte("Return all companies " + " created. Timestamp: " + time.Now().String()),
	}, nil)

	return &Companys, err
}

func (c *Company) FindCompanyByID(db *gorm.DB, uid uint32) (*Company, error) {
	var err error
	err = db.Debug().Model(Company{}).Where("id = ?", uid).Take(&c).Error
	if err != nil {
		return &Company{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Company{}, errors.New("Company Not Found")
	}

	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte("Company " + c.Name + " returned. Timestamp: " + time.Now().String()),
	}, nil)

	return c, err
}

func (c *Company) UpdateACompany(db *gorm.DB, uid uint32) (*Company, error) {

	db = db.Debug().Model(&Company{}).Where("id = ?", uid).Take(&Company{}).UpdateColumns(
		map[string]interface{}{
			"name":              c.Name,
			"description":       c.Description,
			"amountOfEmployees": c.AmountOfEmployees,
			"registered":        c.Registered,
			"type":              c.Type,
		},
	)
	if db.Error != nil {
		return &Company{}, db.Error
	}

	err := db.Debug().Model(&Company{}).Where("id = ?", uid).Take(&c).Error
	if err != nil {
		return &Company{}, err
	}

	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte("Company " + " updated. Timestamp: " + time.Now().String()),
	}, nil)

	return c, nil
}

func (c *Company) DeleteACompany(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&Company{}).Where("id = ?", uid).Take(&Company{}).Delete(&Company{})

	if db.Error != nil {
		return 0, db.Error
	}
	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte("Company " + " deleted. Timestamp: " + time.Now().String()),
	}, nil)
	return db.RowsAffected, nil
}
