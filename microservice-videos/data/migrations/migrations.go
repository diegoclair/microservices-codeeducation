package migrations

import "github.com/GuiaBolso/darwin"

//Migrations to execute our queries that changes database structure
var (
	Migrations = []darwin.Migration{
		{
			Version:     1,
			Description: "Creating table tab_categories",
			Script: `
				CREATE TABLE IF NOT EXISTS tab_categories (
					id INT AUTO_INCREMENT,
					uuid CHAR(36) NOT NULL,
					name VARCHAR(300) NOT NULL,
					description VARCHAR(8000) NULL,
					active TINYINT(1) NOT NULL DEFAULT 1,
					updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
					created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
					deleted_at TIMESTAMP NULL,

					PRIMARY KEY (id),
					UNIQUE INDEX UUID_UNIQUE (uuid ASC),
					UNIQUE INDEX ID_UNIQUE (id ASC)
				) ENGINE=InnoDB CHARACTER SET=utf8;
			`,
		},
		{
			Version:     2,
			Description: "Creating table tab_genres",
			Script: `
				CREATE TABLE IF NOT EXISTS tab_genres (
					id INT AUTO_INCREMENT,
					uuid CHAR(36) NOT NULL,
					name VARCHAR(300) NOT NULL,
					active TINYINT(1) NOT NULL DEFAULT 1,
					updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
					created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
					deleted_at TIMESTAMP NULL,

					PRIMARY KEY (id),
					UNIQUE INDEX UUID_UNIQUE (uuid ASC),
					UNIQUE INDEX ID_UNIQUE (id ASC)
				) ENGINE=InnoDB CHARACTER SET=utf8;
			`,
		},
	}
)
