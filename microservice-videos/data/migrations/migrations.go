package migrations

import "github.com/GuiaBolso/darwin"

//Migrations to execute our queries that changes database structure
var (
	Migrations = []darwin.Migration{
		{
			Version:     1,
			Description: "Creating table users",
			Script: `CREATE TABLE IF NOT EXISTS users (
				id INT AUTO_INCREMENT,
				first_name VARCHAR(30) NOT NULL,
				last_name VARCHAR(30) NOT NULL,
				email VARCHAR(50) NOT NULL,
				password VARCHAR(100) NOT NULL,
				created_at TIMESTAMP,

				PRIMARY KEY (id),
				UNIQUE INDEX ID_UNIQUE (id ASC),
				UNIQUE INDEX EMAIL_UNIQUE (email ASC)
			) ENGINE=InnoDB CHARACTER SET=utf8;`,
		},
		{
			Version:     2,
			Description: "Adding column status to table users",
			Script:      "ALTER TABLE users ADD COLUMN status VARCHAR(30) NOT NULL DEFAULT 'active' AFTER password;",
		},
	}
)
