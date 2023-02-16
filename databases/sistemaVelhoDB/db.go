package sistemaVelhoDB

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaDBAlugueis() {
	user := os.Getenv("DBUSER")
	pass := os.Getenv("DBPASS")
	host := os.Getenv("DBHOST")
	database := os.Getenv("DATABASEVELHO")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, database)

	DB, _ = gorm.Open(mysql.Open(dsn))

}
