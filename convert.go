package sqlx

import (
	"database/sql"
	"database/sql/driver"
	"encoding/binary"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/spf13/cast"
)

var errNil = errors.New("input is nil")

func _cast(x, def interface{}, castFunc interface{}) interface{} {
	if x == nil {
		return def
	}

	switch v := x.(type) {
	case driver.Valuer:
		dv, err := v.Value()
		if err != nil {
			return def
		}
		if dv == nil {
			return def
		}
		return _cast(dv, def, castFunc)
	case sql.RawBytes:
		return _cast([]byte(v), def, castFunc)
	default:
		result := reflect.ValueOf(castFunc).Call([]reflect.Value{reflect.ValueOf(x)})
		if result[1].IsNil() {
			return result[0].Interface()
		}
		return def
	}
}

func toInt(x interface{}, def int) int {
	return _cast(x, def, cast.ToIntE).(int)
}

func toUint(x interface{}, def uint) uint {
	return _cast(x, def, cast.ToUintE).(uint)
}

func toInt8(x interface{}, def int8) int8 {
	return _cast(x, def, cast.ToInt8E).(int8)
}

func toUint8(x interface{}, def uint8) uint8 {
	return _cast(x, def, cast.ToUint8E).(uint8)
}

func toInt16(x interface{}, def int16) int16 {
	return _cast(x, def, cast.ToInt16E).(int16)
}

func toUint16(x interface{}, def uint16) uint16 {
	return _cast(x, def, cast.ToUint16E).(uint16)
}

func toInt32(x interface{}, def int32) int32 {
	return _cast(x, def, cast.ToInt32E).(int32)
}

func toUint32(x interface{}, def uint32) uint32 {
	return _cast(x, def, cast.ToUint32E).(uint32)
}

func toInt64(x interface{}, def int64) int64 {
	return _cast(x, def, cast.ToInt64E).(int64)
}

func toUint64(x interface{}, def uint64) uint64 {
	return _cast(x, def, cast.ToUint64E).(uint64)
}

func toFloat32(x interface{}, def float32) float32 {
	return _cast(x, def, cast.ToFloat32E).(float32)
}

func toFloat64(x interface{}, def float64) float64 {
	return _cast(x, def, cast.ToFloat64E).(float64)
}

func toString(x interface{}, def string) string {
	return _cast(x, def, cast.ToStringE).(string)
}

func toTime(x interface{}, def time.Time) time.Time {
	return _cast(x, def, cast.ToTimeE).(time.Time)
}

func toDuration(x interface{}, def time.Duration) time.Duration {
	return _cast(x, def, cast.ToDurationE).(time.Duration)
}

func _castToBoolE(x interface{}) (bool, error) {
	if x == nil {
		return false, errNil
	}

	switch v := x.(type) {
	case bool:
		return v, nil
	case int:
		return !(0 == v), nil
	case uint:
		return !(0 == v), nil
	case int8:
		return !(0 == v), nil
	case uint8:
		return !(0 == v), nil
	case int16:
		return !(0 == v), nil
	case uint16:
		return !(0 == v), nil
	case int32:
		return !(0 == v), nil
	case uint32:
		return !(0 == v), nil
	case int64:
		return !(0 == v), nil
	case uint64:
		return !(0 == v), nil
	case float32:
		return !(0 == v), nil
	case float64:
		return !(0 == v), nil
	case string:
		return strconv.ParseBool(v)
	case []byte:
		if len(v) == 1 {
			return !(0 == v[0]), nil
		}
		return false, fmt.Errorf("%v can not convert to bool", v)
	default:
		return false, fmt.Errorf("%v can not convert to bool", v)
	}
}

func toBool(x interface{}, def bool) bool {
	return _cast(x, def, _castToBoolE).(bool)
}

func _castToBytesE(x interface{}) ([]byte, error) {
	if x == nil {
		return nil, errNil
	}

	switch v := x.(type) {
	case sql.RawBytes:
		return []byte(v), nil
	case []byte:
		return v, nil
	case uint:
		ret := make([]byte, 8)
		binary.LittleEndian.PutUint64(ret, uint64(v))
		return ret, nil
	case int:
		ret := make([]byte, 8)
		binary.LittleEndian.PutUint64(ret, uint64(v))
		return ret, nil
	case uint8:
		return []byte{v}, nil
	case int8:
		return []byte{uint8(v)}, nil
	case uint16:
		ret := make([]byte, 2)
		binary.LittleEndian.PutUint16(ret, v)
		return ret, nil
	case int16:
		ret := make([]byte, 2)
		binary.LittleEndian.PutUint16(ret, uint16(v))
		return ret, nil
	case uint32:
		ret := make([]byte, 4)
		binary.LittleEndian.PutUint32(ret, v)
		return ret, nil
	case int32:
		ret := make([]byte, 4)
		binary.LittleEndian.PutUint32(ret, uint32(v))
		return ret, nil
	case uint64:
		ret := make([]byte, 8)
		binary.LittleEndian.PutUint64(ret, v)
		return ret, nil
	case int64:
		ret := make([]byte, 8)
		binary.LittleEndian.PutUint64(ret, uint64(v))
		return ret, nil
	case string:
		return []byte(v), nil
	case bool:
		if v {
			return []byte{1}, nil
		}
		return []byte{0}, nil
	case time.Time:
		return v.MarshalBinary()
	default:
		return nil, fmt.Errorf("%v can convert to bytes", v)
	}
}

func toBytes(x interface{}, def []byte) []byte {
	return _cast(x, def, _castToBytesE).([]byte)
}
