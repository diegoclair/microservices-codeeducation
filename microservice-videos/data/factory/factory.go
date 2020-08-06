package factory

import (
	"strconv"

	"github.com/GuiaBolso/darwin"
)

//FakeData to include fake data on database
func FakeData(migrationsList []darwin.Migration, howManyLines int) []darwin.Migration {
	init := len(migrationsList)
	size := len(migrationsList) + howManyLines
	var values string
	for i := init; i < size; i++ {
		if (i + 1) == size {

			values += `('test ` + strconv.Itoa(i) + `','description ` + strconv.Itoa(i) + `');`
			break
		}
		values += `('test ` + strconv.Itoa(i) + `','description ` + strconv.Itoa(i) + `'), `
	}

	var (
		migration = darwin.Migration{
			Version:     float64(init + 1),
			Description: "Inserting data tab_categories",
			Script: `
					INSERT INTO tab_categories 
						(name,description) 
					VALUES ` + values,
		}
	)

	migrationsList = append(migrationsList, migration)

	return migrationsList
}
