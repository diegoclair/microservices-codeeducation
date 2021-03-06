package routeutils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/labstack/echo/v4"
)

// GetAndValidateIntParam gets the integer param value and validates it, returning a validation error in case it's invalid
func GetAndValidateIntParam(c echo.Context, paramName string, isZeroValid bool) (ID int64, restErr resterrors.RestErr) {
	IDStr := c.Param(paramName)

	if strings.TrimSpace(IDStr) == "" {
		return ID, resterrors.NewBadRequestError(fmt.Sprintf("The param %s is needed to proceed with this request", paramName))
	}

	ID, err := strconv.ParseInt(IDStr, 10, 64)
	if err != nil {
		return ID, resterrors.NewBadRequestError(fmt.Sprintf("The param %s is should be a integer number", paramName))
	}
	if !isZeroValid && ID == 0 {
		return ID, resterrors.NewBadRequestError(fmt.Sprintf("The param %s can't be zero", paramName))
	}

	return
}
