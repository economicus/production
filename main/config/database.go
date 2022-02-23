package config

import (
	"fmt"
	"log"
	"os"
	"time"
)

const (
	maxOpenDBConn = 25
	maxIdleDBConn = 25
	maxDBLifeTime = 5 * time.Minute
)

type DatabaseConfig struct {
	dial     string
	user     string
	password string
	host     string
	port     string
	name     string

	MaxOpenDBConn int
	MaxIdleDBConn int
	MaxDBLifeTime time.Duration
}

func NewDatabaseConfig() *DatabaseConfig {
	dial := os.Getenv("DB_DIAL")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pwd := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")

	if dial == "" {
		log.Fatalln("error while configuring database: empty dial")
	}
	if host == "" {
		log.Fatalln("error while configuring database: empty host")
	}
	if port == "" {
		log.Fatalln("error while configuring database: empty port")
	}
	if user == "" {
		log.Fatalln("error while configuring database: empty user")
	}
	if pwd == "" {
		log.Fatalln("error while configuring database: empty pwd")
	}
	if name == "" {
		log.Fatalln("error while configuring database: empty name")
	}
	return &DatabaseConfig{
		dial:          dial,
		user:          user,
		password:      pwd,
		host:          host,
		port:          port,
		name:          name,
		MaxIdleDBConn: maxIdleDBConn,
		MaxOpenDBConn: maxOpenDBConn,
		MaxDBLifeTime: maxDBLifeTime,
	}
}

func (db *DatabaseConfig) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db.user, db.password, db.host, db.port, db.name)
}

func (db *DatabaseConfig) PrintInfo() {
	fmt.Println("========== DB ==========")
	fmt.Println("Dial: ", db.dial)
	fmt.Println("User: ", db.user)
	fmt.Println("Password: ", db.password)
	fmt.Println("Host: ", db.host)
	fmt.Println("Port: ", db.port)
	fmt.Println("Name: ", db.name)
	fmt.Println("MaxIdleDBConn: ", db.MaxIdleDBConn)
	fmt.Println("MaxOpenDBConn: ", db.MaxOpenDBConn)
	fmt.Println("MaxDBLifeTime: ", db.MaxDBLifeTime)
}
