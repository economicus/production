package db

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

const maxIter = 20

const (
	maxOpenDBConn = 25
	maxIdleDBConn = 25
	maxDBLifeTime = 5 * time.Minute
)

type MySql struct {
	conf
	DB *gorm.DB
}

func NewMySql() *MySql {
	ms := MySql{}
	ms.conf = newConf()
	ms.openGorm()
	ms.setup()
	ms.testDBConnection()
	return &ms
}

func (ms *MySql) GetSqlDB() *sql.DB {
	db, err := ms.DB.DB()
	if err != nil {
		log.Fatalf("error while getting sql db: %v", err)
	}
	return db
}

func (ms *MySql) openGorm() {
	var err error
	dsn := ms.conf.GetDSN()

	for i := 0; i < maxIter; i++ {
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			ms.DB = db
			return
		} else {
			log.Println("MySQL connection has failed. Waiting for MySQL. Sleeping for a second...")
			time.Sleep(1 * time.Second)
		}
	}
	log.Fatalf("error while opening mysql with dsn '%s': %v", dsn, err)
}

func (ms *MySql) setup() {
	mysqlDB := ms.GetSqlDB()
	mysqlDB.SetMaxOpenConns(maxOpenDBConn)
	mysqlDB.SetMaxIdleConns(maxIdleDBConn)
	mysqlDB.SetConnMaxLifetime(maxDBLifeTime)
}

func (ms *MySql) testDBConnection() {
	for i := 0; i < maxIter; i++ {
		if err := ms.GetSqlDB().Ping(); err == nil {
			fmt.Println("===== Pinged database successfully! =====")
			return
		} else {
			fmt.Println("Database has not been prepared yet. sleeping for a second...")
			time.Sleep(1 * time.Second)
		}
	}
	log.Fatalf("error while testing connection: tried %d times but failed", maxIter)
}

func (ms *MySql) Migrate(objs []interface{}) {
	for i := range objs {
		if err := ms.DB.AutoMigrate(objs[i]); err != nil {
			log.Fatalf("error while auto migration: %v", err)
		}
	}
}
