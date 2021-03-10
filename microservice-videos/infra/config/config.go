package config

import (
	"sync"

	"github.com/diegoclair/go_utils-lib/logger"
	"github.com/spf13/viper"
)

var (
	config *EnvironmentVariables
	once   sync.Once
)

// EnvironmentVariables is environment variables configs
type EnvironmentVariables struct {
	MySQL mysqlConfig
}

// GetConfigEnvironment to read initial config
func GetConfigEnvironment() *EnvironmentVariables {
	once.Do(func() {

		viper.SetConfigFile(".env")
		viper.AutomaticEnv()

		err := viper.ReadInConfig()
		if err != nil {
			logger.Error("Error to read configs: ", err)
			panic(err)
		}

		config = &EnvironmentVariables{}
		config.MySQL = getMySQLConfig()
	})

	return config
}

type mysqlConfig struct {
	Username         string
	Password         string
	Host             string
	Port             string
	DBName           string
	MaxLifeInMinutes int
	MaxIdleConns     int
	MaxOpenConns     int
}

func getMySQLConfig() (mysql mysqlConfig) {

	mysql.Username = viper.GetString("DB_USER")
	mysql.Password = viper.GetString("DB_PASSWORD")
	mysql.Host = viper.GetString("DB_HOST")
	mysql.Port = viper.GetString("DB_PORT")
	mysql.DBName = viper.GetString("DB_NAME")
	mysql.MaxLifeInMinutes = viper.GetInt("MAX_LIFE_IN_MINUTES")
	mysql.MaxIdleConns = viper.GetInt("MAX_IDLE_CONNS")
	mysql.MaxOpenConns = viper.GetInt("MAX_OPEN_CONNS")

	return mysql
}
