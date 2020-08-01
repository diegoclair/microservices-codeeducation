package migrations

import "github.com/GuiaBolso/darwin"

//Migrations to execute our queries that changes database structure
var (
	Migrations = []darwin.Migration{
		{
			Version:     1,
			Description: "Creating table tab_categories",
			Script: `CREATE TABLE IF NOT EXISTS tab_categories (
				id INT AUTO_INCREMENT,
				name VARCHAR(300) NOT NULL,
				description VARCHAR(8000) NULL,
				active TINYINT(1) NOT NULL DEFAULT 1,
				created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

				PRIMARY KEY (id),
				UNIQUE INDEX ID_UNIQUE (id ASC)
			) ENGINE=InnoDB CHARACTER SET=utf8;
			`,
		},
	}
)
