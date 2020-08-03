package entity

import (
	"fmt"

	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/go-playground/validator/v10"
)

// Category entity data
type Category struct {
	ID          int64  `json:"id"`
	Name        string `json:"name" validate:"required,max=3"`
	Description string `json:"description" validate:"max=8"`
	Active      bool   `json:"active"`
}

// Validate - to validate a struct
func (c *Category) Validate() resterrors.RestErr {

	validate := validator.New()
	err := validate.Struct(c)
	if err != nil {
		return resterrors.NewUnprocessableEntity("Validation error: " + fmt.Sprint(err))
	}
	return nil
}
