package dataBase

import (
	"beak/pkg/config"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB //база данных (todo еще один пиздый коммент))))

func init() {

	appConfig := config.GetConfig()

	username := appConfig.StorageDB.Username
	password := appConfig.StorageDB.Password
	dbName := appConfig.StorageDB.Database
	dbHost := appConfig.StorageDB.Host
	dbPort := appConfig.StorageDB.Port

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s",
		dbHost, username, dbName, password, dbPort) //Создать строку подключения
	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Println(err)
	}
	//TODO: ДОПИСАТЬ ИСТОРИЮ С ТОКЕНАМИ
	db = conn
	// db.Debug().AutoMigrate(&Account{}, &Contact{}) //Миграция базы данных
}

// возвращает дескриптор объекта DB (todo пиздатый коммент))) Я вахуе)
func GetDB() *gorm.DB {
	return db
}
