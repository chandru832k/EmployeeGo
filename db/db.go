package db

import (
	config "employee/config"
	"fmt"
	"time"

	mysql "go.elastic.co/apm/module/apmgormv2/v2/driver/mysql"
	"gorm.io/gorm"
)

var DbClient *gorm.DB

func ConnectToDb() error {
	fmt.Println("%#v", config.EnvironmentData)

	ConnectionURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.EnvironmentData.DbUser, config.EnvironmentData.DbPass, config.EnvironmentData.DbUrl, config.EnvironmentData.DbPort, config.EnvironmentData.DbName)
	fmt.Println("%v", ConnectionURI)
	dbconn, err := gorm.Open(mysql.Open(ConnectionURI), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return err
	}
	sqlDB, err := dbconn.DB()
	if err != nil {
		return err
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Minute * 1)

	DbClient = dbconn
	return nil
}
