package mysql

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Sqlite struct {
	DB                    *gorm.DB
	Path                  string
	IsLog                 bool
	MaxIdleConn           int
	MaxOpenConn           int
	ConnMaxLifetimeSecond int
}

func (s Sqlite) Run() error {
	if s.Path == "" {
		return errors.New("SqlitePath can not be empty")
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)
	config := &gorm.Config{}
	if s.IsLog {
		config = &gorm.Config{Logger: newLogger}
	}

	db, err := gorm.Open(sqlite.Open(s.Path), config)
	if err != nil {
		return err
	}

	sqlDB, _ := db.DB()

	sqlDB.SetMaxIdleConns(s.MaxIdleConn)
	sqlDB.SetMaxOpenConns(s.MaxOpenConn)
	sqlDB.SetConnMaxLifetime(time.Duration(s.ConnMaxLifetimeSecond) * time.Second)

	s.DB = db
	return nil
}
