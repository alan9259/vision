package platform

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func New() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost password=WorkHappily123 user=postgres dbname=dbmessyitup sslmode=disable connect_timeout=30")
	//defer db.Close()
	if err != nil {
		log.Fatalln("Can't connect to the db: ", err)
		fmt.Println("storage err: ", err)
	}
	db.DB().SetMaxIdleConns(3)
	db.LogMode(true)
	return db
}

func AutoMigrate(db *gorm.DB) {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "miu." + defaultTableName
	}

}
