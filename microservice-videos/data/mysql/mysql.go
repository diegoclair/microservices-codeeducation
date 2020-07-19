package mysql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/GuiaBolso/darwin"
	"github.com/diegoclair/microservices-codeeducation/tree/master/microservice-videos/data/migrations"
	"github.com/diegoclair/microservices-codeeducation/tree/master/microservice-videos/domain/contract"
	"github.com/diegoclair/microservices-codeeducation/tree/master/microservice-videos/infra/config"
)

// DBManager is the MySQL connection manager
type DBManager struct {
	db *sql.DB
}

//Instance retunrs an instance of a RepoManager
func Instance() (contract.RepoManager, error) {
	cfg := config.GetDBConfig()

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName,
	)

	log.Println("Connecting to database...")

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Database successfully configured")

	log.Println("Running the migrations")
	driver := darwin.NewGenericDriver(db, darwin.MySQLDialect{})

	d := darwin.New(driver, migrations.Migrations, nil)

	err = d.Migrate()
	if err != nil {
		return nil, err
	}

	log.Println("Migrations executed")

	instance := &DBManager{
		db: db,
	}

	return instance, nil
}
