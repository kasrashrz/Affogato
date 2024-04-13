package mysql_utils

import (
	"github.com/go-sql-driver/mysql"
	"github.com/kasrashrz/Affogato/utils/errors"
	"strings"
)

const (
	ErrorNoRows = "no rows in result"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return errors.NotFoundError("no record matching given id")
		}
		return errors.InternalServerError("error parsing database response")
	}
	switch sqlErr.Number {
	case 1062:
		return errors.BadRequestError("The data has inserted before")
	}
	return errors.InternalServerError("error parsing request")
}
