package datastore

import (
	"log"
	"os"

	"github.com/m-moo/k8s-plat/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// use only postgres
func Init() *gorm.DB {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pwd := os.Getenv("DB_PWD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dbURL := "host=" + host + " user=" + user +
		" password=" + pwd + " dbname=" + dbName + " port=" + port
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatalln(err)
	}

	if err = db.AutoMigrate(
		&models.Users{},
	); err != nil {
		log.Fatalln(err)
	}

	return db
}
