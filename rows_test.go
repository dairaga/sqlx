package sqlx

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	gsql "database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	funcNames = []string{"GetBool", "GetFloat32", "GetFloat64", "GetInt", "GetInt16", "GetInt32", "GetInt64", "GetInt8", "GetString", "GetTime", "GetUint", "GetUint16", "GetUint32", "GetUint64", "GetUint8", "GetBytes"}
	funcs     = []interface{}{
		func(rs *Rows, name string, def bool) bool { return rs.GetBool(name, def) },
		func(rs *Rows, name string, def float32) float32 { return rs.GetFloat32(name, def) },
		func(rs *Rows, name string, def float64) float64 { return rs.GetFloat64(name, def) },
		func(rs *Rows, name string, def int) int { return rs.GetInt(name, def) },
		func(rs *Rows, name string, def int16) int16 { return rs.GetInt16(name, def) },
		func(rs *Rows, name string, def int32) int32 { return rs.GetInt32(name, def) },
		func(rs *Rows, name string, def int64) int64 { return rs.GetInt64(name, def) },
		func(rs *Rows, name string, def int8) int8 { return rs.GetInt8(name, def) },
		func(rs *Rows, name string, def string) string { return rs.GetString(name, def) },
		func(rs *Rows, name string, def time.Time) time.Time { return rs.GetTime(name, def) },
		func(rs *Rows, name string, def uint) uint { return rs.GetUint(name, def) },
		func(rs *Rows, name string, def uint16) uint16 { return rs.GetUint16(name, def) },
		func(rs *Rows, name string, def uint32) uint32 { return rs.GetUint32(name, def) },
		func(rs *Rows, name string, def uint64) uint64 { return rs.GetUint64(name, def) },
		func(rs *Rows, name string, def uint8) uint8 { return rs.GetUint8(name, def) },
		func(rs *Rows, name string, def []byte) []byte { return rs.GetBytes(name, def) },
	}
	defTime  = time.Now()
	defstr   = "xyz"
	defbytes = []byte{}
	defs     = []interface{}{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), defstr, defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), defbytes}
)

func TestZeroTime(t *testing.T) {
	t.Log(reflect.Zero(timeType))
}
func Test1(t *testing.T) {
	db, err := gsql.Open("mysql", dsn)

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	rows, err := db.Query("select * from test1")
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()

	cols, err := rows.ColumnTypes()
	if err != nil {
		t.Fatal(err)
	}

	for i, col := range cols {
		canNull, ok1 := col.Nullable()
		len, ok2 := col.Length()
		p, s, ok3 := col.DecimalSize()

		t.Logf("\n%d: %s\t%s\t%v\n\tnull: %v, %v\n\tlen: %v, %v\n\tdecimal size: %v, %v, %v", i, col.Name(), col.DatabaseTypeName(), col.ScanType(), canNull, ok1, len, ok2, p, s, ok3)

	}

	rows2, err := db.Query("select * from test2")
	if err != nil {
		t.Fatal(err)
	}
	defer rows2.Close()

	cols, err = rows2.ColumnTypes()
	if err != nil {
		t.Fatal(err)
	}

	for i, col := range cols {
		canNull, ok1 := col.Nullable()
		len, ok2 := col.Length()
		p, s, ok3 := col.DecimalSize()

		t.Logf("\n%d: %s\t%s\t%v\n\tnull: %v, %v\n\tlen: %v, %v\n\tdecimal size: %v, %v, %v", i, col.Name(), col.DatabaseTypeName(), col.ScanType(), canNull, ok1, len, ok2, p, s, ok3)

	}
}

func TestToRowsAllNotNull(t *testing.T) {

	db, err := gsql.Open("mysql", dsn)

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	rs, err := WrapRows(db.Query("select * from test1"))
	if err != nil {
		t.Fatal(err)
	}
	defer rs.Close()

	ansTime1 := time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)
	ansTime1Bytes, _ := ansTime1.MarshalBinary()

	ansTime2 := time.Date(2014, time.September, 23, 10, 01, 02, 0, time.UTC)
	ansTime2Bytes, _ := ansTime2.MarshalBinary()

	ans := [][][]interface{}{
		{
			{true, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), string([]byte{1}), defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), []byte{1}},
			{true, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), string([]byte{1}), defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), []byte{1}},
			{true, float32(-2), float64(-2), int(-2), int16(-2), int32(-2), int64(-2), int8(-2), "-2", defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), []byte{254}},
			{true, float32(2), float64(2), int(2), int16(2), int32(2), int64(2), int8(2), "2", defTime, uint(2), uint16(2), uint32(2), uint64(2), uint8(2), []byte{2}},
			{true, float32(1), float64(1), int(1), int16(1), int32(1), int64(1), int8(1), "1", defTime, uint(1), uint16(1), uint32(1), uint64(1), uint8(1), []byte{1}},
			{true, float32(-4), float64(-4), int(-4), int16(-4), int32(-4), int64(-4), int8(-4), "-4", defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), []byte{252, 255}},
			{true, float32(4), float64(4), int(4), int16(4), int32(4), int64(4), int8(4), "4", defTime, uint(4), uint16(4), uint32(4), uint64(4), uint8(4), []byte{4, 0}},
			{true, float32(-5), float64(-5), int(-5), int16(-5), int32(-5), int64(-5), int8(-5), "-5", time.Unix(-5, 0), uint(100), uint16(100), uint32(100), uint64(100), uint8(100), []byte{251, 255, 255, 255}},
			{true, float32(5), float64(5), int(5), int16(5), int32(5), int64(5), int8(5), "5", time.Unix(5, 0), uint(5), uint16(5), uint32(5), uint64(5), uint8(5), []byte{5, 0, 0, 0}},
			{true, float32(-6), float64(-6), int(-6), int16(-6), int32(-6), int64(-6), int8(-6), "-6", time.Unix(-6, 0), uint(100), uint16(100), uint32(100), uint64(100), uint8(100), []byte{250, 255, 255, 255}},
			{true, float32(6), float64(6), int(6), int16(6), int32(6), int64(6), int8(6), "6", time.Unix(6, 0), uint(6), uint16(6), uint32(6), uint64(6), uint8(6), []byte{6, 0, 0, 0}},
			{true, float32(-7), float64(-7), int(-7), int16(-7), int32(-7), int64(-7), int8(-7), "-7", time.Unix(-7, 0), uint(100), uint16(100), uint32(100), uint64(100), uint8(100), []byte{249, 255, 255, 255, 255, 255, 255, 255}},
			{true, float32(7), float64(7), int(7), int16(7), int32(7), int64(7), int8(7), "7", time.Unix(7, 0), uint(7), uint16(7), uint32(7), uint64(7), uint8(7), []byte{7, 0, 0, 0, 0, 0, 0, 0}},
			{true, float32(-8), float64(-8), int(-8), int16(-8), int32(-8), int64(-8), int8(-8), "-8", defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), defbytes},
			{true, float32(8), float64(8), int(8), int16(8), int32(8), int64(8), int8(8), "8", defTime, uint(8), uint16(8), uint32(8), uint64(8), uint8(8), defbytes},
			{true, float32(-9), float64(-9), int(-9), int16(-9), int32(-9), int64(-9), int8(-9), "-9", defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), defbytes},
			{true, float32(9), float64(9), int(9), int16(9), int32(9), int64(9), int8(9), "9", defTime, uint(9), uint16(9), uint32(9), uint64(9), uint8(9), defbytes},
			{true, float32(-10), float64(-10), int(-10), int16(-10), int32(-10), int64(-10), int8(-10), "-10", defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), defbytes},
			{true, float32(10), float64(10), int(10), int16(10), int32(10), int64(10), int8(10), "10", defTime, uint(10), uint16(10), uint32(10), uint64(10), uint8(10), defbytes},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), ansTime1.String(), ansTime1, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), ansTime1Bytes},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), "123:04:05", defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), []byte("123:04:05")},
			{true, float32(2019), float64(2019), int(2019), int16(2019), int32(2019), int64(2019), int8(-29), "2019", defTime, uint(2019), uint16(2019), uint32(2019), uint64(2019), uint8(227), []byte{227, 7}},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), ansTime2.String(), ansTime2, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), ansTime2Bytes},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), ansTime2.String(), ansTime2, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), ansTime2Bytes},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), "C16", defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), []byte("C16")},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), "C17", defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), []byte("C17")},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), string([]byte{'C', '1', '8', 0, 0, 0, 0, 0, 0, 0}), defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), []byte{'C', '1', '8', 0, 0, 0, 0, 0, 0, 0}},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), "C19", defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), []byte("C19")},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), "C20", defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), []byte("C20")},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), "C21", defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), []byte("C21")},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), "C22", defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), []byte("C22")},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), "C23", defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), []byte("C23")},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), "C24", defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), []byte("C24")},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), "C25", defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), []byte("C25")},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), "C26", defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), []byte("C26")},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), "C27", defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), []byte("C27")},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), "a", defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), []byte("a")},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), "b", defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), []byte("b")},
		},
	}

	count := 0
	for rs.Next() {
		for i, name := range rs.names {

			if i < len(ans[count]) {
				for j, f := range funcs {
					a := ans[count][i][j]
					b := reflect.ValueOf(f).Call([]reflect.Value{reflect.ValueOf(rs), reflect.ValueOf(name), reflect.ValueOf(defs[j])})[0].Interface()
					switch av := a.(type) {
					case []byte:
						if string(av) != string(b.([]byte)) {
							t.Errorf("%d: %s %s should be %v, but %v", i, name, funcNames[j], av, b.([]byte))
						}
					default:
						if a != b {
							t.Errorf("%d: %s %s should be %v, but %v", i, name, funcNames[j], a, b)
						}
					}

				}
			} else {
				t.Logf("%d: %s: bool:    %v", i, name, rs.GetBool(name, false))
				t.Logf("%d: %s: float32: %v", i, name, rs.GetFloat32(name, 100.0))
				t.Logf("%d: %s: float64: %v", i, name, rs.GetFloat64(name, 100.0))
				t.Logf("%d: %s: int:     %v", i, name, rs.GetInt(name, -100))
				t.Logf("%d: %s: int16:   %v", i, name, rs.GetInt16(name, -100))
				t.Logf("%d: %s: int32:   %v", i, name, rs.GetInt32(name, -100))
				t.Logf("%d: %s: int64:   %v", i, name, rs.GetInt64(name, -100))
				t.Logf("%d: %s: int8:    %v", i, name, rs.GetInt8(name, -100))
				t.Logf("%d: %s: string:  %v", i, name, rs.GetString(name, "xyz"))
				t.Logf("%d: %s: time:    %v", i, name, rs.GetTime(name, defTime))
				t.Logf("%d: %s: uint:    %v", i, name, rs.GetUint(name, 100))
				t.Logf("%d: %s: uint16:  %v", i, name, rs.GetUint16(name, 100))
				t.Logf("%d: %s: uint32:  %v", i, name, rs.GetUint32(name, 100))
				t.Logf("%d: %s: uint64:  %v", i, name, rs.GetUint64(name, 100))
				t.Logf("%d: %s: uint8:   %v", i, name, rs.GetUint8(name, 100))
				t.Logf("%d: %s: bytes:   %v", i, name, rs.GetBytes(name, defbytes))
			}

		}

		count++
	}

	if err := rs.Err(); err != nil {
		t.Errorf("rs err: %v", err)
	}
}

func TestToRowsAllNull(t *testing.T) {

	db, err := gsql.Open("mysql", dsn)

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	rs, err := WrapRows(db.Query("select * from test2"))
	if err != nil {
		t.Fatal(err)
	}
	defer rs.Close()

	timeAns := time.Date(2014, time.September, 23, 10, 01, 02, 0, time.UTC)
	timeAnsBytes, _ := timeAns.MarshalBinary()

	ans := [][][]interface{}{
		{
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), "", defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), defbytes},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), "", defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), defbytes},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), defstr, defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), defbytes},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), defstr, defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), defbytes},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), defstr, defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), defbytes},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), defstr, defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), defbytes},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), defstr, defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), defbytes},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), defstr, defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), defbytes},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), defstr, defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), defbytes},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), defstr, defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), defbytes},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), defstr, defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), defbytes},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), defstr, defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), defbytes},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), defstr, defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), defbytes},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), defstr, defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), defbytes},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), defstr, defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), defbytes},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), defstr, defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), defbytes},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), defstr, defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), defbytes},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), defstr, defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), defbytes},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), defstr, defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), defbytes},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), defstr, defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), defbytes},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), "", defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), defbytes},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), defstr, defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), defbytes},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), defstr, defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), defbytes},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), timeAns.String(), timeAns, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), timeAnsBytes},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), defstr, defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), defbytes},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), defstr, defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), defbytes},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), "", defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), defbytes},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), "", defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), defbytes},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), "", defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), defbytes},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), defstr, defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), defbytes},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), "", defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), defbytes},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), defstr, defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), defbytes},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), "", defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), defbytes},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), defstr, defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), defbytes},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), "", defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), defbytes},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), defstr, defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), defbytes},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), defstr, defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), defbytes},
			{false, float32(100), float64(100), int(-100), int16(-100), int32(-100), int64(-100), int8(-100), defstr, defTime, uint(100), uint16(100), uint32(100), uint64(100), uint8(100), defbytes},
		},
	}

	count := 0
	for rs.Next() {
		for i, name := range rs.names {

			if i < len(ans[count]) {
				for j, f := range funcs {
					a := ans[count][i][j]
					b := reflect.ValueOf(f).Call([]reflect.Value{reflect.ValueOf(rs), reflect.ValueOf(name), reflect.ValueOf(defs[j])})[0].Interface()
					switch av := a.(type) {
					case []byte:
						if string(av) != string(b.([]byte)) {
							t.Errorf("%d: %s %s should be %v, but %v", i, name, funcNames[j], av, b.([]byte))
						}
					default:
						if a != b {
							t.Errorf("%d: %s %s should be %v, but %v", i, name, funcNames[j], a, b)
						}
					}
				}
			} else {
				t.Logf("%d: %s: bool:    %v", i, name, rs.GetBool(name, false))
				t.Logf("%d: %s: float32: %v", i, name, rs.GetFloat32(name, 100.0))
				t.Logf("%d: %s: float64: %v", i, name, rs.GetFloat64(name, 100.0))
				t.Logf("%d: %s: int:     %v", i, name, rs.GetInt(name, -100))
				t.Logf("%d: %s: int16:   %v", i, name, rs.GetInt16(name, -100))
				t.Logf("%d: %s: int32:   %v", i, name, rs.GetInt32(name, -100))
				t.Logf("%d: %s: int64:   %v", i, name, rs.GetInt64(name, -100))
				t.Logf("%d: %s: int8:    %v", i, name, rs.GetInt8(name, -100))
				t.Logf("%d: %s: string:  %v", i, name, rs.GetString(name, "xyz"))
				t.Logf("%d: %s: time:    %v", i, name, rs.GetTime(name, defTime))
				t.Logf("%d: %s: uint:    %v", i, name, rs.GetUint(name, 100))
				t.Logf("%d: %s: uint16:  %v", i, name, rs.GetUint16(name, 100))
				t.Logf("%d: %s: uint32:  %v", i, name, rs.GetUint32(name, 100))
				t.Logf("%d: %s: uint64:  %v", i, name, rs.GetUint64(name, 100))
				t.Logf("%d: %s: uint8:   %v", i, name, rs.GetUint8(name, 100))
				t.Logf("%d: %s: bytes:   %v", i, name, rs.GetBytes(name, defbytes))
			}

		}

		count++
	}

	if err := rs.Err(); err != nil {
		t.Errorf("rs err: %v", err)
	}
}

func TestData(t *testing.T) {
	db, err := gsql.Open("mysql", dsn)

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	rs1, err := WrapRows(db.Query("select * from test1"))
	if err != nil {
		t.Fatal(err)
	}
	defer rs1.Close()

	if !rs1.Next() {
		t.Errorf("no data")
		return
	}

	d1, err := rs1.r.Data()
	if err != nil {
		t.Errorf("to internal data: %v", err)
	}
	tmp, err := json.Marshal(d1)
	if err != nil {
		t.Errorf("marshal: %v", err)
	}
	if string(tmp) != test1Row {
		t.Errorf("json string should be \n%s\n but \n%s", test1Row, string(tmp))
	}

	rs2, err := WrapRows(db.Query("select * from test2"))
	if err != nil {
		t.Fatal(err)
	}
	defer rs2.Close()

	if !rs2.Next() {
		t.Errorf("no data")
		return
	}

	d2, err := rs2.r.Data()
	if err != nil {
		t.Errorf("to internal data: %v", err)
	}
	tmp, err = json.Marshal(d2)
	if err != nil {
		t.Errorf("marshal: %v", err)
	}
	if string(tmp) != test2Row {
		t.Errorf("json string should be \n%s\n but \n%s", test2Row, string(tmp))
	}
}

func TestMarshal(t *testing.T) {
	db, err := gsql.Open("mysql", dsn)

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	rs1, err := WrapRows(db.Query("select * from test1"))
	if err != nil {
		t.Fatal(err)
	}
	defer rs1.Close()

	if !rs1.Next() {
		t.Errorf("no data")
		return
	}

	d1 := &TestStruct1{}
	err = rs1.r.Unmarshal(d1)
	if err != nil {
		t.Errorf("to internal data: %v", err)
	}
	tmp, err := json.Marshal(d1)
	if err != nil {
		t.Errorf("marshal: %v", err)
	}
	if string(tmp) != s1 {
		t.Errorf("json string should be \n%s\n but \n%s", s1, string(tmp))
	}

	rs2, err := WrapRows(db.Query("select * from test2"))
	if err != nil {
		t.Fatal(err)
	}
	defer rs2.Close()

	if !rs2.Next() {
		t.Errorf("no data")
		return
	}

	err = rs2.r.Unmarshal(d1)
	if err != nil {
		t.Errorf("to internal data: %v", err)
	}
	tmp, err = json.Marshal(d1)
	if err != nil {
		t.Errorf("marshal: %v", err)
	}
	if string(tmp) != s2 {
		t.Errorf("json string should be \n%s\n but \n%s", s2, string(tmp))
	}
}

func TestMarshalAll(t *testing.T) {
	db, err := gsql.Open("mysql", dsn)

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	rs1, err := WrapRows(db.Query("select * from test1"))
	if err != nil {
		t.Fatal(err)
	}
	defer rs1.Close()

	d1 := []TestStruct1{}

	err = rs1.UnmarshalAll(&d1)
	if err != nil {
		t.Errorf("to internal data: %v", err)
	}
	tmp, err := json.Marshal(d1)
	if err != nil {
		t.Errorf("marshal: %v", err)
	}
	if string(tmp) != s1Slice {
		t.Errorf("json string should be \n%s\n but \n%s", s1Slice, string(tmp))
	}

	rs2, err := WrapRows(db.Query("select * from test2"))
	if err != nil {
		t.Fatal(err)
	}
	defer rs2.Close()

	err = rs2.UnmarshalAll(&d1)
	if err != nil {
		t.Errorf("to internal data: %v", err)
	}
	tmp, err = json.Marshal(d1)
	if err != nil {
		t.Errorf("marshal: %v", err)
	}
	if string(tmp) != s2Slice {
		t.Errorf("json string should be \n%s\n but \n%s", s2Slice, string(tmp))
	}
}
