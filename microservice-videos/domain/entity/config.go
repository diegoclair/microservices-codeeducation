package entity

// InitialConfig entity
type InitialConfig struct {
	Username         string `json:"username,omitempty"`
	Password         string `json:"password,omitempty"`
	Host             string `json:"host,omitempty"`
	Port             string `json:"port,omitempty"`
	DBName           string `json:"db_name,omitempty"`
	MaxLifeInMinutes int    `json:"max_life_minutes,omitempty"`
	MaxIdleConns     int    `json:"max_idle_conns,omitempty"`
	MaxOpenConns     int    `json:"max_open_conns,omitempty"`
}
