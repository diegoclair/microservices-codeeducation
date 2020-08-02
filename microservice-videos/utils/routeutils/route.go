package routeutils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// GetAndValidateIntParam gets the integer param value and validates it, returning a validation error in case it's invalid
func GetAndValidateIntParam(c *gin.Context, paramName string, isZeroValid bool) (ID int64, err error) {
	IDStr := c.Param(paramName)
	errMessage := fmt.Sprintf("The param %s is invalid", paramName)

	if strings.TrimSpace(IDStr) == "" {
		err = fmt.Errorf(errMessage)
		return
	}

	ID, err = strconv.ParseInt(IDStr, 10, 64)
	if !isZeroValid && ID == 0 {
		err = fmt.Errorf(errMessage)
	}

	return
}
