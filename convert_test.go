package sqlx

import (
	"database/sql"
	"testing"
	"time"
)

var (
	now   = time.Now()
	str   = "ABC"
	nostr = "1"
	cases = []interface{}{
		int(1),
		uint(1),
		int8(1),
		uint8(1),
		int16(1),
		uint16(1),
		int32(1),
		uint32(1),
		int64(1),
		uint64(1),
		float32(1),
		float64(1),
		true,
		str,
		nostr,
		now,
		sql.RawBytes([]byte(str)),
		sql.RawBytes(nil),
		sql.NullBool{Valid: true, Bool: false},
		sql.NullBool{Valid: false, Bool: false},
		sql.NullFloat64{Valid: true, Float64: 1},
		sql.NullFloat64{Valid: false, Float64: 0},
		sql.NullInt64{Valid: true, Int64: 1},
		sql.NullInt64{Valid: false, Int64: 0},
		sql.NullString{Valid: true, String: str},
		sql.NullString{Valid: true, String: nostr},
		sql.NullString{Valid: false, String: ""},
	}
)

func TestToInt(t *testing.T) {
	ans := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 100, 1, 100, 100, 100, 0, 100, 1, 100, 1, 100, 100, 1, 100}

	for i, x := range cases {
		a := toInt(x, 100)
		if a != ans[i] {
			t.Errorf("%v toInt should be %d, but %d", x, ans[i], a)
		}
		//t.Logf("%d: %v => %d", i, x, a)
		//fmt.Printf("%d,", a)
	}
}

func TestToUint(t *testing.T) {
	ans := []uint{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 100, 1, 100, 100, 100, 0, 100, 1, 100, 1, 100, 100, 1, 100}

	for i, x := range cases {
		a := toUint(x, 100)
		if a != ans[i] {
			t.Errorf("%v toInt should be %d, but %d", x, ans[i], a)
		}
		//t.Logf("%d: %v => %d", i, x, a)
		//fmt.Printf("%d,", a)
	}
}

func TestToInt8(t *testing.T) {
	ans := []int8{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 100, 1, 100, 100, 100, 0, 100, 1, 100, 1, 100, 100, 1, 100}

	for i, x := range cases {
		a := toInt8(x, 100)
		if a != ans[i] {
			t.Errorf("%v toInt should be %d, but %d", x, ans[i], a)
		}
		//t.Logf("%d: %v => %d", i, x, a)
		//fmt.Printf("%d,", a)
	}
}

func TestToUint8(t *testing.T) {
	ans := []uint8{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 100, 1, 100, 100, 100, 0, 100, 1, 100, 1, 100, 100, 1, 100}

	for i, x := range cases {
		a := toUint8(x, 100)
		if a != ans[i] {
			t.Errorf("%v toInt should be %d, but %d", x, ans[i], a)
		}
		//t.Logf("%d: %v => %d", i, x, a)
		//fmt.Printf("%d,", a)
	}
}

func TestToInt16(t *testing.T) {
	ans := []int16{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 100, 1, 100, 100, 100, 0, 100, 1, 100, 1, 100, 100, 1, 100}

	for i, x := range cases {
		a := toInt16(x, 100)
		if a != ans[i] {
			t.Errorf("%v toInt should be %d, but %d", x, ans[i], a)
		}
		//t.Logf("%d: %v => %d", i, x, a)
		//fmt.Printf("%d,", a)
	}
}

func TestToUint16(t *testing.T) {
	ans := []uint16{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 100, 1, 100, 100, 100, 0, 100, 1, 100, 1, 100, 100, 1, 100}

	for i, x := range cases {
		a := toUint16(x, 100)
		if a != ans[i] {
			t.Errorf("%v toInt should be %d, but %d", x, ans[i], a)
		}
		//t.Logf("%d: %v => %d", i, x, a)
		//fmt.Printf("%d,", a)
	}
}

func TestToInt32(t *testing.T) {
	ans := []int32{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 100, 1, 100, 100, 100, 0, 100, 1, 100, 1, 100, 100, 1, 100}

	for i, x := range cases {
		a := toInt32(x, 100)
		if a != ans[i] {
			t.Errorf("%v toInt should be %d, but %d", x, ans[i], a)
		}
		//t.Logf("%d: %v => %d", i, x, a)
		//fmt.Printf("%d,", a)
	}
}

func TestToUint32(t *testing.T) {
	ans := []uint32{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 100, 1, 100, 100, 100, 0, 100, 1, 100, 1, 100, 100, 1, 100}

	for i, x := range cases {
		a := toUint32(x, 100)
		if a != ans[i] {
			t.Errorf("%v toInt should be %d, but %d", x, ans[i], a)
		}
		//t.Logf("%d: %v => %d", i, x, a)
		//fmt.Printf("%d,", a)
	}
}

func TestToInt64(t *testing.T) {
	ans := []int64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 100, 1, 100, 100, 100, 0, 100, 1, 100, 1, 100, 100, 1, 100}

	for i, x := range cases {
		a := toInt64(x, 100)
		if a != ans[i] {
			t.Errorf("%v toInt should be %d, but %d", x, ans[i], a)
		}
		//t.Logf("%d: %v => %d", i, x, a)
		//fmt.Printf("%d,", a)
	}
}

func TestToUint64(t *testing.T) {
	ans := []uint64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 100, 1, 100, 100, 100, 0, 100, 1, 100, 1, 100, 100, 1, 100}

	for i, x := range cases {
		a := toUint64(x, 100)
		if a != ans[i] {
			t.Errorf("%v toInt should be %d, but %d", x, ans[i], a)
		}
		//t.Logf("%d: %v => %d", i, x, a)
		//fmt.Printf("%d,", a)
	}
}

func TestToFloat32(t *testing.T) {
	ans := []float32{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 100, 1, 100, 100, 100, 0, 100, 1, 100, 1, 100, 100, 1, 100}

	for i, x := range cases {
		a := toFloat32(x, 100)
		if a != ans[i] {
			t.Errorf("%v toInt should be %v, but %v", x, ans[i], a)
		}
		//t.Logf("%d: %v => %d", i, x, a)
		//fmt.Printf("%d,", a)
	}
}

func TestToFloat64(t *testing.T) {
	ans := []float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 100, 1, 100, 100, 100, 0, 100, 1, 100, 1, 100, 100, 1, 100}

	for i, x := range cases {
		a := toFloat64(x, 100)
		if a != ans[i] {
			t.Errorf("%v toInt should be %v, but %v", x, ans[i], a)
		}
		//t.Logf("%d: %v => %d", i, x, a)
		//fmt.Printf("%d,", a)
	}
}

func TestToString(t *testing.T) {
	def := "xyz"
	ans := []string{"1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "true", str, nostr, now.String(), str, "", "false", def, "1", def, "1", def, str, nostr, def}

	for i, x := range cases {
		a := toString(x, def)
		if a != ans[i] {
			t.Errorf("%v toInt should be %v, but %v", x, ans[i], a)
		}
		//t.Logf("%d: %v => %v", i, x, a)
		//fmt.Printf("%q,", a)
	}
}

func TestToBool(t *testing.T) {
	def := false
	ans := []bool{true, true, true, true, true, true, true, true, true, true, true, true, true, false, true, false, false, false, false, false, true, false, true, false, false, true, false}

	for i, x := range cases {
		a := toBool(x, def)
		if a != ans[i] {
			t.Errorf("%v toInt should be %v, but %v", x, ans[i], a)
		}
		//t.Logf("%d: %v => %v", i, x, a)
		//fmt.Printf("%v,", a)
	}
}

func TestToTime(t *testing.T) {
	def := time.Now()
	d1 := time.Unix(1, 0)
	ans := []time.Time{
		d1, d1, //int, uint,
		def, def, def, def, // int8, uint8, int16, uint16
		d1, d1, d1, d1, // int32, uint32, int64, uint64
		def, def, // float32, float64
		def, def, def, // true, string, string
		now,
		def, def, // sql.RawBytes, sql.RawBytes
		def, def, def, def, // sql.NullBool, sql.NullBool, sql.NullFloat64, sql.NullFloat64
		d1, def, // sql.NullInt64, sql.NullInt64
		def, def, def, // sql.NullString, sql.NullString, sql.NullString
	}

	for i, x := range cases {
		a := toTime(x, def)
		if a != ans[i] {
			t.Errorf("%v toInt should be %v, but %v", x, ans[i], a)
		}
		//t.Logf("%d: %v (%T) => %v", i, x, x, a)
		//fmt.Printf("%v,", a)
	}
}
