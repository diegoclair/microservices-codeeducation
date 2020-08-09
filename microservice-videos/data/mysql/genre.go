package mysql

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/diegoclair/go_utils-lib/logger"
	"github.com/diegoclair/go_utils-lib/mysqlutils"
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/microservices-codeeducation/tree/master/microservice-videos/domain/entity"
)

type genreRepo struct {
	db *sql.DB
}

// newGenreRepo returns a instance of genre repo
func newGenreRepo(db *sql.DB) *genreRepo {
	return &genreRepo{
		db: db,
	}
}

//GetGenres - return a list of all genres
func (r *genreRepo) GetGenres() (*[]entity.Genre, resterrors.RestErr) {

	query := `
		SELECT 	tc.id,
				tc.uuid,
				tc.name,
				tc.active

		FROM 	tab_genres 	tc
		WHERE  	tc.deleted_at 	IS NULL;;`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		errorCode := "Error 0001 - "
		logger.Error(fmt.Sprintf("%sGetGenres: ", errorCode), err)
		return nil, resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		errorCode := "Error 0002 - "
		logger.Error(fmt.Sprintf("%sGetGenres: ", errorCode), err)
		return nil, resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer rows.Close()

	var genres []entity.Genre
	for rows.Next() {
		var genre entity.Genre
		err = rows.Scan(
			&genre.ID,
			&genre.UUID,
			&genre.Name,
			&genre.Active,
		)
		if err != nil {
			errorCode := "Error 0003 - "
			logger.Error(fmt.Sprintf("%sGetGenres: ", errorCode), err)
			errMessage := mysqlutils.HandleMySQLError(err)

			return nil, resterrors.NewRestError(fmt.Sprintf("%s", errorCode)+errMessage.Message(), errMessage.StatusCode(), errMessage.Error(), errMessage.Causes())
		}

		genres = append(genres, genre)
	}

	if len(genres) == 0 {
		return nil, resterrors.NewNotFoundError(fmt.Sprintf("No genres found"))
	}

	return &genres, nil
}

func (r *genreRepo) GetGenreByUUID(uuid string) (*entity.Genre, resterrors.RestErr) {

	query := `
		SELECT 	tc.id,
				tc.uuid,
				tc.name,
				tc.active

		FROM 	tab_genres 	tc
		WHERE 	tc.uuid			= ?
		  AND 	tc.deleted_at 	IS NULL;`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		errorCode := "Error 0001: "
		logger.Error(fmt.Sprintf("%sGetGenreByID", errorCode), err)
		return nil, resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	result := stmt.QueryRow(uuid)
	var genre entity.Genre
	err = result.Scan(
		&genre.ID,
		&genre.UUID,
		&genre.Name,
		&genre.Active,
	)

	if err != nil {
		errorCode := "Error 0002: "
		logger.Error(fmt.Sprintf("%sGetGenreByID", errorCode), err)
		return nil, mysqlutils.HandleMySQLError(err)
	}

	return &genre, nil
}

func (r *genreRepo) CreateGenre(genre entity.Genre) resterrors.RestErr {

	query := `
		INSERT INTO tab_genres 
				(uuid, name) 
		VALUES	(?, ?, ?);
		`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		errorCode := "Error 0001 - "
		logger.Error(fmt.Sprintf("%sCreateGenre: ", errorCode), err)
		return resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	_, err = stmt.Exec(genre.UUID, genre.Name)
	if err != nil {
		errorCode := "Error 0002 - "
		logger.Error(fmt.Sprintf("%sCreateGenre: ", errorCode), err)
		return mysqlutils.HandleMySQLError(err)
	}

	return nil
}

func (r *genreRepo) UpdateGenreByID(uuid string, genre entity.Genre) resterrors.RestErr {

	query := `
		UPDATE tab_genres 
			SET name 		= ?

		WHERE 	uuid 			= ?
		  AND 	deleted_at 	IS NULL;
		`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		errorCode := "Error 0001 - "
		logger.Error(fmt.Sprintf("%sUpdateGenreByID: ", errorCode), err)
		return resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	result, err := stmt.Exec(genre.Name, uuid)
	if err != nil {
		errorCode := "Error 0002 - "
		logger.Error(fmt.Sprintf("%sUpdateGenreByID: ", errorCode), err)
		return mysqlutils.HandleMySQLError(err)
	}

	rowsAffected, err := result.RowsAffected()

	if rowsAffected == 0 {
		return resterrors.NewNotFoundError(fmt.Sprintf("No genres found with the parameter id"))
	}

	return nil
}

func (r *genreRepo) DeleteGenreByID(uuid string) resterrors.RestErr {
	query := `
		UPDATE tab_genres 
			SET deleted_at	= ?

		WHERE 	uuid 		= ?;
		`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		errorCode := "Error 0001 - "
		logger.Error(fmt.Sprintf("%sError when trying to prepare the query statement in DeleteGenreByID: ", errorCode), err)
		return resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	_, err = stmt.Exec(time.Now(), uuid)
	if err != nil {
		errorCode := "Error 0002 - "
		logger.Error(fmt.Sprintf("%sError when trying to execute Query in DeleteGenreByID: ", errorCode), err)
		return mysqlutils.HandleMySQLError(err)
	}

	return nil
}
