package datastore

import (
	"log"

	"github.com/m-moo/k8s-plat/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// use only postgres
func Init() *gorm.DB {

	dbURL := "host=cks0504.iptime.org user=postgres password=postgres#0524 dbname=kube-plat port=5432"
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
