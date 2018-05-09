package dao

import (
	"fmt"

	"github.com/5sigma/vox"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	// Sqlite driver import
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	// Postgres driver import
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DAO - Data access layer
type DAO struct {
	DB *gorm.DB
}

// New - Initialize a new data access layer
func New() *DAO {
	connStr := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
		viper.GetString("DB_HOST"),
		viper.GetString("DB_USER"),
		viper.GetString("DB_NAME"),
		viper.GetString("DB_PASS"),
	)

	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		vox.Fatal(err.Error())
	}
	migrate(db)
	return &DAO{
		DB: db,
	}
}

// NewMemory - Creates a new data access layer that uses an internal memory
// store.
func NewMemory() *DAO {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		vox.Fatal(err.Error())
	}
	migrate(db)
	return &DAO{
		DB: db,
	}
}
