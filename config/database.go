package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=postgres user=postgres password=1111 dbname=gamestore port=5432 sslmode=disable"

	var database *gorm.DB
	var err error

	for i := 0; i < 10; i++ {
		database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("База данных подключена")
			DB = database
			return
		}

		log.Println("⏳ Ожидание подключения к базе данных... Попытка:", i+1)
		time.Sleep(3 * time.Second)
	}

	log.Fatal("❌ Ошибка подключения к базе данных после 10 попыток:", err)
}
