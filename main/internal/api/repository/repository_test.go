package repository

import (
	"economicus/internal/drivers"
	"economicus/internal/models"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"testing"
	"time"
)

var (
	db          *gorm.DB
	userRepo    UserRepositoryFactory
	quantRepo   QuantRepositoryFactory
	commentRepo CommentRepositoryFactory
	replyRepo   ReplyRepositoryFactory
)

var (
	dbUser = os.Getenv("DB_USER")
	dbPass = os.Getenv("DB_PASSWORD")
	dbHost = "localhost"
	dbPort = os.Getenv("DB_PORT")
	dbName = os.Getenv("TEST_DB_NAME")
)

var mainLogger log.Logger

const NumberOfTestCase = 10

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	err := tearDown()
	if err != nil {
		panic(err)
	}
	os.Exit(code)
}

func setup() {
	var err error
	db, err = setupDB()
	if err != nil {
		_ = tearDown()
		log.Fatalf("error while initializing test db :%s", err.Error())
	}
	userRepo = NewUserRepository(db, drivers.NewAWS(), &mainLogger)
	quantRepo = NewQuantRepository(db, &mainLogger)
	commentRepo = NewCommentRepository(db, &mainLogger)
	replyRepo = NewReplyRepository(db, &mainLogger)
}

func setupDB() (*gorm.DB, error) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)
	newLogger := logger.New(
		log.New(),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("err in ConnectMySQL while connecting mysql for test: %w", err)
	}
	if err = autoMigrate(db); err != nil {
		return nil, fmt.Errorf("error in ConnectMySQL while auto migrate: %w", err)
	}
	return db, nil
}

func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.Profile{},
		&models.Quant{},
		&models.QuantOption{},
		&models.MainSector{},
		&models.Reply{},
		&models.Comment{},
	)
}

func tearDown() error {
	tables := []string{
		"comments",
		"quants",
		"replies",
		"main_sectors",
		"user_favorite_quants",
		"quant_options",
		"quant_results",
		"followings",
		"profiles",
		"users",
	}
	for _, table := range tables {
		if err := db.Migrator().DropTable(table); err != nil {
			return fmt.Errorf("error while drop table: %s", table)
		}
	}
	return nil
}
