package drivers

import (
	"database/sql"
	"economicus/config"
	"economicus/internal/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type DB struct {
	Config *config.DatabaseConfig
	SQL    *gorm.DB
}

func NewDatabase() *DB {
	dbConfig := config.NewDatabaseConfig()

	gormDB := getGormDB(dbConfig)
	newDB := DB{
		SQL:    gormDB,
		Config: dbConfig,
	}
	newDB.connectMySQL(dbConfig)
	testDBConnection(newDB.GetMysqlDB())
	newDB.autoMigrate()
	return &newDB
}

func (db *DB) GetMysqlDB() *sql.DB {
	mysqlDB, err := db.SQL.DB()
	if err != nil {
		log.Fatalf("error while getting database: %v", err)
	}
	return mysqlDB
}

func (db *DB) connectMySQL(dbConfig *config.DatabaseConfig) {
	mysqlDB := db.GetMysqlDB()
	mysqlDB.SetMaxOpenConns(dbConfig.MaxOpenDBConn)
	mysqlDB.SetMaxIdleConns(dbConfig.MaxIdleDBConn)
	mysqlDB.SetConnMaxLifetime(dbConfig.MaxDBLifeTime)
}

func (db *DB) autoMigrate() {
	if err := db.SQL.AutoMigrate(
		&models.User{},
		&models.Profile{},
		&models.Quant{},
		&models.QuantOption{},
		&models.MainSector{},
		&models.Reply{},
		&models.Comment{},
		&models.Reply{},
	); err != nil {
		log.Fatalf("error while auto migration: %v", err)
	}
}

func getGormDB(dbConfig *config.DatabaseConfig) *gorm.DB {
	gormDB, err := gorm.Open(mysql.Open(dbConfig.GetDSN()), &gorm.Config{})
	if err != nil {
		log.Fatalf("error while opening mysql with dsn '%s': %v", dbConfig.GetDSN(), err)
	}
	return gormDB
}

func testDBConnection(db *sql.DB) {
	maxIter := 20
	for curIter := 0; curIter < maxIter; curIter++ {
		if err := db.Ping(); err == nil {
			fmt.Println("===== Pinged database successfully! =====")
			return
		} else {
			fmt.Println("Database has not been prepared yet. sleeping for a second...")
			time.Sleep(1 * time.Second)
		}
	}
	log.Fatalf("error while testing connection: tried %d times but failed", maxIter)
}
