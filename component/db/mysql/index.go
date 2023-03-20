package mysql

import (
	"log"
	"os"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/cast"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Mysql struct {
	DB                    *gorm.DB
	Host                  string
	Port                  uint
	IsLog                 bool
	User                  string
	Password              string
	DbName                string
	Charset               string
	Timeout               uint
	MaxIdleConn           int
	MaxOpenConn           int
	ConnMaxLifetimeSecond int
}

func (m Mysql) Run() error {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)
	config := &gorm.Config{}
	if m.IsLog {
		config = &gorm.Config{Logger: newLogger}
	}
	dsn := m.User + ":" + m.Password + "@(" + m.Host + ":" + cast.ToString(m.
		Port) + ")/" + m.DbName + "?charset=" + m.Charset + "&parseTime=True&loc" +
		"=Local" +
		"&timeout=" + cast.ToString(m.Timeout)

	db, err := gorm.Open(mysql.Open(dsn), config)

	if err != nil {
		return err
	}
	m.DB = db
	sqlDB, _ := m.DB.DB()

	sqlDB.SetMaxIdleConns(m.MaxIdleConn)
	sqlDB.SetMaxOpenConns(m.MaxOpenConn)
	sqlDB.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetimeSecond) * time.Second)

	return nil
}

func (m Mysql) GetHandel(tx *gorm.DB) *gorm.DB {
	var dbHandel *gorm.DB
	if tx != nil {
		dbHandel = tx
	} else {
		dbHandel = m.DB
	}
	return dbHandel
}
