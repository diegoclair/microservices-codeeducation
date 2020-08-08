package mysql

import (
	"database/sql"
	"fmt"

	"github.com/GuiaBolso/darwin"
	"github.com/diegoclair/go_utils-lib/logger"
	"github.com/diegoclair/microservices-codeeducation/tree/master/microservice-videos/data/migrations"
	"github.com/diegoclair/microservices-codeeducation/tree/master/microservice-videos/domain/contract"
	"github.com/diegoclair/microservices-codeeducation/tree/master/microservice-videos/infra/config"
	"github.com/go-sql-driver/mysql"
)

// DBManager is the MySQL connection manager
type DBManager struct {
	db *sql.DB
}

//Instance returns an instance of a RepoManager
func Instance() (contract.RepoManager, error) {
	cfg := config.GetDBConfig()

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName,
	)

	logger.Info("Connecting to database...")

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	err = mysql.SetLogger(logger.GetLogger())
	if err != nil {
		return nil, err
	}

	logger.Info("Database successfully configured")

	logger.Info("Running the migrations")
	driver := darwin.NewGenericDriver(db, darwin.MySQLDialect{})

	d := darwin.New(driver, migrations.Migrations, nil)
	err = d.Migrate()
	if err != nil {
		return nil, err
	}

	logger.Info("Migrations executed")

	instance := &DBManager{
		db: db,
	}

	return instance, nil
}

//Category returns the category set
func (c *DBManager) Category() contract.CategoryRepo {
	return newCategoryRepo(c.db)
}
