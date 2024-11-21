package migrator

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewDB new a gorm db instance
func NewDB(debug bool) (*gorm.DB, error) {
	fmt.Println("Connecting to MySQL database...")

	dsn := "root:123456@tcp(localhost:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"

	if debug {
		// return the gorm db instance which prints debug log
		return gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	}
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
