package data

import (
	"github.com/diegoclair/microservices-codeeducation/tree/master/microservice-videos/data/mysql"
	"github.com/diegoclair/microservices-codeeducation/tree/master/microservice-videos/domain/contract"
)

// Connect returns a instace of cassandra db
func Connect() (contract.RepoManager, error) {
	return mysql.Instance()
}
