package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDB connect to db
func ConnectDB() {

	var err error
	// p := config.Config("DB_PORT")
	// port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		panic("failed to parse database port")
	}

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=require",
		"ep-shiny-truth-89466193.ap-southeast-1.aws.neon.fl0.io",
		5432,
		"fl0user",
		"3qFsTVJQ9zNp",
		"database",
	)
	fmt.Println(dsn)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
	// DB.AutoMigrate(&model.User{}, &model.Leave{}, &model.Role{})
	// fmt.Println("Database Migrated")
}
