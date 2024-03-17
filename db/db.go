package db

import (
	"assessment/db/model"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Conn *gorm.DB
)

func StartDBConnection() {
	var err error
	for {
		Conn, err = gorm.Open(
			postgres.Open(os.Getenv("DB_CONN_STRING")),
			&gorm.Config{SkipDefaultTransaction: true},
		)
		if err != nil {
			time.Sleep(1 * time.Second)
			continue
		}
		break
	}

	err = initTables()
	if err != nil {
		log.Fatal(err)
	}
}

func initTables() error {
	err := Conn.AutoMigrate(&model.Movie{})
	if err != nil {
		return err
	}
	return nil
}

func CloseDBConnection() {
	sql, err := Conn.DB()
	if err != nil {
		log.Println(err)
		return
	}

	err = sql.Close()
	if err != nil {
		log.Println(err)
	}
}
