package mysql

import (
	"database/sql"
	"fmt"

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
func (r *categoryRepo) GetCategories() (categories []*entity.Category, restErr resterrors.RestErr) {

	query := `
		SELECT 	tc.id,
				tc.name,
				tc.description,
				tc.active

		FROM 	tab_categories 		tc;`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		errorCode := "Error 0001 - "
		logger.Error(fmt.Sprintf("%sError when trying to prepare the query statement in GetCategories: ", errorCode), err)
		return nil, resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		errorCode := "Error 0002 - "
		logger.Error(fmt.Sprintf("%sError when trying to execute Query in GetCategories: ", errorCode), err)
		return nil, resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer rows.Close()

	for rows.Next() {
		var category entity.Category
		err = rows.Scan(
			&category.ID,
			&category.Name,
			&category.Description,
			&category.Active,
		)
		if err != nil {
			errorCode := "Error 0003 - "
			logger.Error(fmt.Sprintf("%sError when trying to do For Scan in the Rows GetCategories: ", errorCode), err)
			errMessage := mysqlutils.HandleMySQLError(err)

			return nil, resterrors.NewRestError(fmt.Sprintf("%s", errorCode)+errMessage.Message(), errMessage.StatusCode(), errMessage.Error())
		}

		categories = append(categories, &category)
	}

	if len(categories) == 0 {
		return nil, resterrors.NewNotFoundError(fmt.Sprintf("No categories found"))
	}

	return categories, nil
}

func (r *categoryRepo) GetCategoryByID(id int64) (*entity.Category, resterrors.RestErr) {

	query := `
		SELECT 	tc.id,
				tc.name,
				tc.description,
				tc.active

		FROM 	tab_categories 	tc
		WHERE 	tc.id 			= ?;`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		errorCode := "Error 0001: "
		logger.Error(fmt.Sprintf("%sError when trying to prepare the query statement in GetCategoryByID", errorCode), err)
		return nil, resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	result := stmt.QueryRow(id)
	var category entity.Category
	err = result.Scan(
		&category.ID,
		&category.Name,
		&category.Description,
		&category.Active,
	)

	if err != nil {
		errorCode := "Error 0002: "
		logger.Error(fmt.Sprintf("%sError when trying to execute QueryRow in GetByID", errorCode), err)
		return nil, mysqlutils.HandleMySQLError(err)
	}

	return &category, nil
}

func (r *categoryRepo) CreateCategory(category entity.Category) (*entity.Category, resterrors.RestErr) {

	query := `
		INSERT INTO tab_categories 
				(name, description) 
		VALUES	(?, ?);
		`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		errorCode := "Error 0001 - "
		logger.Error(fmt.Sprintf("%sError when trying to prepare the query statement in CreateCategory: ", errorCode), err)
		return nil, resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(category.Name, category.Description)
	if err != nil {
		errorCode := "Error 0002 - "
		logger.Error(fmt.Sprintf("%sError when trying to execute Query in CreateCategory: ", errorCode), err)
		return nil, mysqlutils.HandleMySQLError(err)
	}

	categoryID, err := insertResult.LastInsertId()
	if err != nil {
		errorCode := "Error 0003 - "
		logger.Error(fmt.Sprintf("%sError when trying to get LastInsertId in CreateCategory: ", errorCode), err)
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

		WHERE 	id 			= ?;
		`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		errorCode := "Error 0001 - "
		logger.Error(fmt.Sprintf("%sError when trying to prepare the query statement in CreateCategory: ", errorCode), err)
		return resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	_, err = stmt.Exec(category.Name, category.Description, id)
	if err != nil {
		errorCode := "Error 0002 - "
		logger.Error(fmt.Sprintf("%sError when trying to execute Query in CreateCategory: ", errorCode), err)
		return mysqlutils.HandleMySQLError(err)
	}

	return nil
}
