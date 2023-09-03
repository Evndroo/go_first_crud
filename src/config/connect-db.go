package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	envs := GetAllEnvs()

	dsn := fmt.Sprintf("host=%s user=%s password=%s database=%s port=%s TimeZone=Brazil/East", envs.DB_HOST, envs.DB_USER, envs.DB_PASSWORD, envs.DB_NAME, envs.DB_PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return db, err

}
