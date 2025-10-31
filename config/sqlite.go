package config

import (
	"os"

	"github.com/lcaparada/gopportunities/schemas"
	"gorm.io/driver/sqlite"

	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")

	dbPath := "./db/main.db"

	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Info("database file not found, creating...")

		err = os.MkdirAll("./db", os.ModePerm)

		if err != nil {
			logger.Errorf("create folder error: %v", err)
			return nil, err
		}

		file, err := os.Create("main.db")

		if err != nil {
			logger.Errorf("create file error: %v", err)
		}

		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})

	if err != nil {
		logger.Errorf("sqlite automigration error: %v", err)
		return nil, err
	}

	return db, nil
}
