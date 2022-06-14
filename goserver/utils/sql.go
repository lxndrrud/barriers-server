package utils

import "database/sql"

func NewNullInt(intValue int64) sql.NullInt64 {
	if intValue == 0 {
		return sql.NullInt64{}
	} else {
		return sql.NullInt64{
			Int64: intValue,
			Valid: true,
		}
	}
}
