package db

import (
	"fmt"
	"gonote/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func New(dbCfg *config.DB) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		dbCfg.Host,
		dbCfg.User,
		dbCfg.Password,
		dbCfg.Name,
		dbCfg.Port,
	)
	log.Print(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
