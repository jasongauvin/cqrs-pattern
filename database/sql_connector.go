package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jasongauvin/cqrs_pattern/models"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// InitializeDb connects to database
func InitializeDb(user string, password string, host string, name string) {
	dbURL := fmt.Sprintf("%s:%s@%s/%s?parseTime=True&charset=utf8", user, password, host, name)
	fmt.Println(" \n connexion ", dbURL)
	// var tmpDb, err = gorm.Open("mysql", dbURL)
	// if err != nil {
	// 	fmt.Printf("Cannot connect to database :")
	// 	fmt.Println(err)
	// 	// log.Fatal("error:", err)
	// 	// return
	// }

	// db := tmpDb.DB()

	var tmpDb *gorm.DB
	var err error

	for databaseConnectionAttemptLoop := 0; databaseConnectionAttemptLoop < 4; databaseConnectionAttemptLoop++ {
		err = nil

		tmpDb, err = gorm.Open("mysql", dbURL)

		if err != nil {
			fmt.Printf("\n error : %v", err)
		}

		if err == nil {
			break // Here, if there is no error, it simply breaks out and does not retry again.

		}

	}
	// var dbErr error

	// for i := 1; i <= 8; i++ {
	// 	dbErr = db.Ping()
	// 	if dbErr != nil {
	// 		if i < 8 {
	// 			log.Printf("db connection failed, %d retry : %v", i, dbErr)
	// 			time.Sleep(10 * time.Second)
	// 		}
	// 		continue
	// 	}

	// 	break
	// }

	// fmt.Printf("We are connected to database \n")
	DB = tmpDb
}

// MakeMigrations executes all migrations for our structs
func MakeMigrations() {
	DB.AutoMigrate(&models.User{})
}
