package sqlx

import (
	"database/sql"
	"reflect"
	"time"
)

var (
	timeType    = reflect.TypeOf(time.Time{})
	stringType  = reflect.TypeOf("")
	boolType    = reflect.TypeOf(false)
	int64Type   = reflect.TypeOf(int64(0))
	float64Type = reflect.TypeOf(float64(0.0))

	rawBytesType = reflect.TypeOf(sql.RawBytes{})
	bytesType    = reflect.TypeOf([]byte{})

	nullStringType  = reflect.TypeOf(sql.NullString{})
	nullBoolType    = reflect.TypeOf(sql.NullBool{})
	nullInt64Type   = reflect.TypeOf(sql.NullInt64{})
	nullFloat64Type = reflect.TypeOf(sql.NullFloat64{})
)
