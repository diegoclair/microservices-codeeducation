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

type categoryRepo struct {
	db *sql.DB
}

// newCategoryRepo returns a instance of category repo
func newCategoryRepo(db *sql.DB) *categoryRepo {
	return &categoryRepo{
		db: db,
	}
}

//GetCategories - return a list of all categories
func (r *categoryRepo) GetCategories() (*[]entity.Category, resterrors.RestErr) {

	query := `
		SELECT 	tc.id,
				tc.uuid,
				tc.name,
				tc.description,
				tc.active

		FROM 	tab_categories 	tc
		WHERE  	tc.deleted_at 	IS NULL;;`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		errorCode := "Error 0001 - "
		logger.Error(fmt.Sprintf("%sGetCategories: ", errorCode), err)
		return nil, resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		errorCode := "Error 0002 - "
		logger.Error(fmt.Sprintf("%sGetCategories: ", errorCode), err)
		return nil, resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer rows.Close()

	var categories []entity.Category
	for rows.Next() {
		var category entity.Category
		err = rows.Scan(
			&category.ID,
			&category.UUID,
			&category.Name,
			&category.Description,
			&category.Active,
		)
		if err != nil {
			errorCode := "Error 0003 - "
			logger.Error(fmt.Sprintf("%sGetCategories: ", errorCode), err)
			errMessage := mysqlutils.HandleMySQLError(err)

			return nil, resterrors.NewRestError(fmt.Sprintf("%s", errorCode)+errMessage.Message(), errMessage.StatusCode(), errMessage.Error())
		}

		categories = append(categories, category)
	}

	if len(categories) == 0 {
		return nil, resterrors.NewNotFoundError(fmt.Sprintf("No categories found"))
	}

	return &categories, nil
}

func (r *categoryRepo) GetCategoryByID(id int64) (*entity.Category, resterrors.RestErr) {

	query := `
		SELECT 	tc.id,
				tc.uuid,
				tc.name,
				tc.description,
				tc.active

		FROM 	tab_categories 	tc
		WHERE 	tc.id 			= ?
		  AND 	tc.deleted_at 	IS NULL;`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		errorCode := "Error 0001: "
		logger.Error(fmt.Sprintf("%sGetCategoryByID", errorCode), err)
		return nil, resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	result := stmt.QueryRow(id)
	var category entity.Category
	err = result.Scan(
		&category.ID,
		&category.UUID,
		&category.Name,
		&category.Description,
		&category.Active,
	)

	if err != nil {
		errorCode := "Error 0002: "
		logger.Error(fmt.Sprintf("%sGetCategoryByID", errorCode), err)
		return nil, mysqlutils.HandleMySQLError(err)
	}

	return &category, nil
}

func (r *categoryRepo) CreateCategory(category entity.Category) (*entity.Category, resterrors.RestErr) {

	query := `
		INSERT INTO tab_categories 
				(uuid, name, description) 
		VALUES	(?, ?, ?);
		`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		errorCode := "Error 0001 - "
		logger.Error(fmt.Sprintf("%sCreateCategory: ", errorCode), err)
		return nil, resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(category.UUID, category.Name, category.Description)
	if err != nil {
		errorCode := "Error 0002 - "
		logger.Error(fmt.Sprintf("%sCreateCategory: ", errorCode), err)
		return nil, mysqlutils.HandleMySQLError(err)
	}

	categoryID, err := insertResult.LastInsertId()
	if err != nil {
		errorCode := "Error 0003 - "
		logger.Error(fmt.Sprintf("%sCreateCategory: ", errorCode), err)
		return nil, mysqlutils.HandleMySQLError(err)
	}

	category.ID = categoryID

	return &category, nil
}

func (r *categoryRepo) UpdateCategoryByID(id int64, category entity.Category) resterrors.RestErr {

	query := `
		UPDATE tab_categories 
			SET name 		= ?, 
				description	= ?

		WHERE 	uuid 			= ?
		  AND 	deleted_at 	IS NULL;
		`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		errorCode := "Error 0001 - "
		logger.Error(fmt.Sprintf("%sUpdateCategoryByID: ", errorCode), err)
		return resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	result, err := stmt.Exec(category.Name, category.Description, id)
	if err != nil {
		errorCode := "Error 0002 - "
		logger.Error(fmt.Sprintf("%sUpdateCategoryByID: ", errorCode), err)
		return mysqlutils.HandleMySQLError(err)
	}

	rowsAffected, err := result.RowsAffected()

	if rowsAffected == 0 {
		return resterrors.NewNotFoundError(fmt.Sprintf("No categories found with the parameter id"))
	}

	return nil
}

func (r *categoryRepo) DeleteCategoryByID(id int64) resterrors.RestErr {
	query := `
		UPDATE tab_categories 
			SET deleted_at	= ?

		WHERE 	id 			= ?;
		`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		errorCode := "Error 0001 - "
		logger.Error(fmt.Sprintf("%sError when trying to prepare the query statement in DeleteCategoryByID: ", errorCode), err)
		return resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	_, err = stmt.Exec(time.Now(), id)
	if err != nil {
		errorCode := "Error 0002 - "
		logger.Error(fmt.Sprintf("%sError when trying to execute Query in DeleteCategoryByID: ", errorCode), err)
		return mysqlutils.HandleMySQLError(err)
	}

	return nil
}
