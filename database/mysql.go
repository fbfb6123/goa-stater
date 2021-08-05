package database

import (
	"fmt"
	"os"
	"strconv"

	"goa_starter/gen/log"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DataSource string
	Host       string
	Port       int
	Database   string
	User       string
	Password   string
}

var logger = log.New("goa_starter", true)

func (c *Config) Open() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DriverName: "mysql",
		// DSN:        c.DataSource, // data source name, refer https://github.com/go-sql-driver/mysql#dsn-data-source-name
		DSN: "root:root@tcp(localhost:3306)/sf_offer?parseTime=true&loc=Local",
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	const numOfConns = 10
	gormDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// マイグレーションは一旦いいや

	// migrations := &migrate.FileMigrationSource{
	// 	Dir: "migrations",
	// }
	// n, err := migrate.Exec(gormDB, "mysql", migrations, migrate.Up)
	// if err != nil {
	// 	logger.Errorf("migration failed: %v", err)
	// 	return db, err
	// }
	// logger.Info(fmt.Sprintf("=======Applied %d migrations!\n========", n))

	gormDB.SetMaxIdleConns(numOfConns)
	gormDB.SetMaxOpenConns(numOfConns)

	return db, nil
}

func CloseConnection(db *gorm.DB) error {
	cdb, err := db.DB()
	if err != nil {
		return err
	}
	return cdb.Close()
}

func NewConfigsFromEnv(stage string) (*Config, error) {
	var host, database, user, pwd string
	var port int
	var err error
	var c Config

	user = os.Getenv("DB_USER")
	pwd = os.Getenv("DB_PASSWORD")
	if stage == "test" {
		port, err = strconv.Atoi(os.Getenv("DB_PORT_TEST"))
		if err != nil {
			return nil, err
		}
		host = os.Getenv("DB_HOST_TEST")
		database = os.Getenv("DATABASE_TEST")
	} else if stage == "local" {
		port, err = strconv.Atoi(os.Getenv("DB_PORT"))
		if err != nil {
			return nil, err
		}
		host = os.Getenv("DB_HOST")
		database = os.Getenv("DATABASE")
	}

	c = Config{
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			user, pwd, host, port, database),
		host,
		port,
		database,
		user,
		pwd,
	}

	return &c, nil
}
