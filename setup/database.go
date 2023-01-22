package setup

import (
	"fmt"
	"log"
	"os"
	"server/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Database() {
	dbTime := os.Getenv("APP_TIME")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	log.Println("Loaded database credentials")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		dbHost, dbUser, dbPassword, dbName, dbPort, dbTime)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Panicf("Failed to connect to database : %v\n", err)
	}
	log.Println("Connected to database")

	db.AutoMigrate(&models.HealthcheckLog{})
	db.AutoMigrate(&models.AccessRule{})
	db.AutoMigrate(&models.AccessLog{})
	db.AutoMigrate(&models.Personel{})
	db.AutoMigrate(&models.RSSILog{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Role{})
	db.AutoMigrate(&models.Lock{})
	db.AutoMigrate(&models.Key{})

	roles := []models.Role{
		{
			BaseModel: models.BaseModel{
				ID: 1,
			},
			Name: "Guest",
		},
		{
			BaseModel: models.BaseModel{
				ID: 2,
			},
			Name: "Doctor",
		},
		{
			BaseModel: models.BaseModel{
				ID: 3,
			},
			Name: "Nurse",
		},
		{
			BaseModel: models.BaseModel{
				ID: 4,
			},
			Name: "Staff",
		},
	}

	for _, r := range roles {
		db.Create(&r)
	}

	DB = db
}
