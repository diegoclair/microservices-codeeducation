package factory

import (
	"strconv"

	"github.com/GuiaBolso/darwin"
	uuid "github.com/satori/go.uuid"
)

//FakeData to include fake data on database
func FakeData(migrationsList []darwin.Migration, howManyLines int) []darwin.Migration {
	init := len(migrationsList)
	size := len(migrationsList) + howManyLines
	var values, UUID string
	for i := init; i < size; i++ {
		UUID = uuid.NewV4().String()
		if (i + 1) == size {

			values += `('` + UUID + `', 'test ` + strconv.Itoa(i) + `','description ` + strconv.Itoa(i) + `');`
			break
		}
		values += `('` + UUID + `', 'test ` + strconv.Itoa(i) + `','description ` + strconv.Itoa(i) + `'), `
	}
	var (
		migration = darwin.Migration{
			Version:     float64(init + 1),
			Description: "Inserting data tab_categories",
			Script: `
					INSERT INTO tab_categories 
						(uuid,name,description) 
					VALUES ` + values,
		}
	)

	migrationsList = append(migrationsList, migration)

	return migrationsList
}
