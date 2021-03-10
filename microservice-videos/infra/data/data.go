package data

import (
	"database/sql"
	"log"

	"github.com/diegoclair/microservices-codeeducation/microservice-videos/contract"
	"github.com/diegoclair/microservices-codeeducation/microservice-videos/infra/data/mysql"
)

//we can add here more than one database
type data struct {
	mysqlRepo contract.MySQLRepo
	mysqlDB   *sql.DB
}

// Connect returns a instace of cassandra db
func Connect() (contract.DataManager, error) {
	repo := new(data)
	return &data{
		mysqlRepo: repo.MySQL(),
	}, nil
}

func (d *data) MySQL() contract.MySQLRepo {
	mysqlRepo, err := mysql.Instance()
	if err != nil {
		log.Fatalf("Error to start mysql instance: %v", err)
	}
	return mysqlRepo
}
