package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//logger "github.com/sirupsen/logrus"
)

// InitConnectionDb init database. For simplicity of the demo database is disabled
func InitConnectionDb(dbHost, dbUser, dbPassword, dbName string) *gorm.DB {
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True", dbUser, dbPassword, dbHost, dbName)
	db, _ := gorm.Open("mysql", dbDSN)
	/*if err != nil {
		logger.WithError(err).Fatalf("Cannot connect to database:")
	}*/
	return db
}
