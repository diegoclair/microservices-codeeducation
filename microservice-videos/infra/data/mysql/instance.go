package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/GuiaBolso/darwin"
	"github.com/diegoclair/go_utils-lib/logger"
	"github.com/diegoclair/microservices-codeeducation/microservice-videos/contract"
	"github.com/diegoclair/microservices-codeeducation/microservice-videos/infra/config"
	"github.com/diegoclair/microservices-codeeducation/microservice-videos/infra/data/migrations"

	mysqlDriver "github.com/go-sql-driver/mysql"
)

var (
	conn    *mysqlConn
	onceDB  sync.Once
	connErr error
)

// mysqlConn is the database connection manager
type mysqlConn struct {
	db *sql.DB
}

//Instance returns an instance of a MySQLRepo
func Instance() (contract.MySQLRepo, error) {
	onceDB.Do(func() {
		cfg := config.GetConfigEnvironment()

		dataSourceName := fmt.Sprintf("%s:root@tcp(%s:%s)/%s?charset=utf8",
			cfg.MySQL.Username, cfg.MySQL.Host, cfg.MySQL.Port, cfg.MySQL.DBName,
		)

		log.Println("Connecting to database...")
		db, connErr := sql.Open("mysql", dataSourceName)
		if connErr != nil {
			return
		}

		log.Println("Database Ping...")
		connErr = db.Ping()
		if connErr != nil {
			return
		}

		log.Println("Creating database...")
		if _, connErr = db.Exec("CREATE DATABASE IF NOT EXISTS sampamodas_db;"); connErr != nil {
			logger.Error("Create Database error: ", connErr)
			return
		}

		if _, connErr = db.Exec("USE sampamodas_db;"); connErr != nil {
			logger.Error("Default Database error: ", connErr)
			return
		}

		connErr = mysqlDriver.SetLogger(logger.GetLogger())
		if connErr != nil {
			return
		}
		logger.Info("Database successfully configured")

		logger.Info("Running the migrations")
		driver := darwin.NewGenericDriver(db, darwin.MySQLDialect{})

		d := darwin.New(driver, migrations.Migrations, nil)

		connErr = d.Migrate()
		if connErr != nil {
			logger.Error("Migrate Error: ", connErr)
			return
		}

		logger.Info("Migrations executed")

		conn = &mysqlConn{
			db: db,
		}
	})

	return conn, nil
}

//Category returns the category set
func (c *mysqlConn) Category() contract.CategoryRepo {
	return newCategoryRepo(c.db)
}

//Genre returns the category set
func (c *mysqlConn) Genre() contract.GenreRepo {
	return newGenreRepo(c.db)
}
