package databases

import (
	"fmt"
	"log"
	"thanapatjitmung/go-test-komgrip/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMariaDatabase(cfg *config.MariaDB) *gorm.DB {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	mariadbDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to MariaDB: %v", err)
	}
	log.Println("Connected to MariaDB!")

	return mariadbDB
}
