package sqlx

import (
	"database/sql"
	"fmt"
	"reflect"
	"time"
)

var (
	zeroTime = time.Time{}
)

// Rows ...
type Rows struct {
	x     *sql.Rows
	s     reflect.Type
	names []string
	types []reflect.Type
	r     *Row
}

// WrapRows ...
func WrapRows(rows *sql.Rows, err error) (*Rows, error) {
	if err != nil {
		return nil, err
	}

	return toRows(rows)
}

func toRows(rows *sql.Rows) (*Rows, error) {
	cols, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}
	size := len(cols)
	rs := &Rows{
		x:     rows,
		names: make([]string, size),
		types: make([]reflect.Type, size),
	}

	fields := make([]reflect.StructField, size)

	for i, col := range cols {
		rs.names[i] = col.Name()
		f := reflect.StructField{
			Name: toCamel(rs.names[i]),
			Tag:  reflect.StructTag(fmt.Sprintf(`json:"%s"`, rs.names[i])),
		}

		switch col.DatabaseTypeName() {
		case "DECIMAL":
			if null, _ := col.Nullable(); null {
				rs.types[i] = nullFloat64Type
			} else {
				rs.types[i] = float64Type
			}
			f.Type = float64Type
		case "TEXT", "VARCHAR", "CHAR", "NCHAR", "NVARCHAR":
			if null, _ := col.Nullable(); null {
				rs.types[i] = nullStringType
			} else {
				rs.types[i] = stringType
			}
			f.Type = stringType
		case "DATE", "DATETIME", "TIMESTAMP":
			rs.types[i] = col.ScanType()
			f.Type = timeType
		default:
			t := col.ScanType()
			rs.types[i] = t
			if t.AssignableTo(nullBoolType) {
				f.Type = boolType
			} else if t.AssignableTo(nullFloat64Type) {
				f.Type = float64Type
			} else if t.AssignableTo(nullInt64Type) {
				f.Type = int64Type
			} else if t.AssignableTo(nullStringType) {
				f.Type = stringType
			} else if t.AssignableTo(rawBytesType) {
				f.Type = bytesType
			} else {
				f.Type = t
			}
		}

		fields[i] = f
	}
	rs.s = reflect.StructOf(fields)

	return rs, nil
}

// Next ...
func (rs *Rows) Next() bool {
	//if rs.lastErr != nil {
	//	return false
	//}

	if !rs.x.Next() {
		return false
	}

	rs.r = &Row{rows: rs}
	if err := rs.r.scan(); err != nil {
		return false
	}

	return true
}

// Close ...
func (rs *Rows) Close() error {
	return rs.x.Close()
}

// Err returns error of Rows.
func (rs *Rows) Err() error {
	return rs.x.Err()
}

// AllRow ...
func (rs *Rows) AllRow() ([]*Row, error) {
	defer rs.Close()
	var ret []*Row
	for rs.Next() {
		ret = append(ret, rs.r)
	}
	return ret, rs.Err()
}

// All ...
func (rs *Rows) All() ([]interface{}, error) {
	tmp, err := rs.AllRow()

	if err != nil {
		return nil, err
	}

	var ret []interface{}
	for _, r := range tmp {
		s, err := r.Data()
		if err != nil {
			return nil, err
		}
		ret = append(ret, s)
	}
	return ret, nil
}

// UnmarshalAll ...
func (rs *Rows) UnmarshalAll(x interface{}) error {
	defer rs.Close()
	rv := reflect.ValueOf(x)

	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return fmt.Errorf("x must be Ptr of Slice")
	}

	rv = rv.Elem()

	if rv.Kind() != reflect.Slice {
		return fmt.Errorf("x must be Ptr of Slice")
	}

	newV := reflect.MakeSlice(rv.Type(), 0, 0)

	for rs.Next() {
		d := reflect.New(rv.Type().Elem())
		if err := rs.r.Unmarshal(d.Interface()); err != nil {
			return err
		}

		newV = reflect.Append(newV, d.Elem())
	}

	rv.Set(newV)

	return nil
}

// GetInt ...
func (rs *Rows) GetInt(name string, def ...int) int {
	return rs.r.GetInt(name, def...)
}

// GetUint ...
func (rs *Rows) GetUint(name string, def ...uint) uint {
	return rs.r.GetUint(name, def...)
}

// GetInt8 ...
func (rs *Rows) GetInt8(name string, def ...int8) int8 {
	return rs.r.GetInt8(name, def...)
}

// GetUint8 ...
func (rs *Rows) GetUint8(name string, def ...uint8) uint8 {
	return rs.r.GetUint8(name, def...)
}

// GetInt16 ...
func (rs *Rows) GetInt16(name string, def ...int16) int16 {
	return rs.r.GetInt16(name, def...)
}

// GetUint16 ...
func (rs *Rows) GetUint16(name string, def ...uint16) uint16 {
	return rs.r.GetUint16(name, def...)
}

// GetInt32 ...
func (rs *Rows) GetInt32(name string, def ...int32) int32 {
	return rs.r.GetInt32(name, def...)
}

// GetUint32 ...
func (rs *Rows) GetUint32(name string, def ...uint32) uint32 {
	return rs.r.GetUint32(name, def...)
}

// GetInt64 ...
func (rs *Rows) GetInt64(name string, def ...int64) int64 {
	return rs.r.GetInt64(name, def...)
}

// GetUint64 ...
func (rs *Rows) GetUint64(name string, def ...uint64) uint64 {
	return rs.r.GetUint64(name, def...)
}

// GetFloat32 ...
func (rs *Rows) GetFloat32(name string, def ...float32) float32 {
	return rs.r.GetFloat32(name, def...)
}

// GetFloat64 ...
func (rs *Rows) GetFloat64(name string, def ...float64) float64 {
	return rs.r.GetFloat64(name, def...)
}

// GetString ...
func (rs *Rows) GetString(name string, def ...string) string {
	return rs.r.GetString(name, def...)
}

// GetTime ...
func (rs *Rows) GetTime(name string, def ...time.Time) time.Time {
	return rs.r.GetTime(name, def...)
}

// GetBool ...
func (rs *Rows) GetBool(name string, def ...bool) bool {
	return rs.r.GetBool(name, def...)
}

// GetBytes ...
func (rs *Rows) GetBytes(name string, def ...[]byte) []byte {
	return rs.r.GetBytes(name, def...)
}

// GetDuration ...
func (rs *Rows) GetDuration(name string, def ...time.Duration) time.Duration {
	return rs.r.GetDuration(name, def...)
}

// ----------------------------------------------------------------------------

// Row ...
type Row struct {
	rows *Rows
	err  error
	data map[string]interface{}
	s    interface{}
}

// WrapRow ...
func WrapRow(rs *Rows, err error) *Row {
	r := &Row{rows: rs, err: err}
	if err == nil {
		if rs.Next() {
			r.scan()
		} else {
			r.err = sql.ErrNoRows
		}
		rs.Close()
	}

	return r
}

// Err ...
func (r *Row) Err() error {
	if r.err != nil {
		return r.err
	}

	return nil
}

func (r *Row) scan() error {
	if r.err != nil {
		return r.err
	}

	size := len(r.rows.types)
	vals := make([]interface{}, size)
	for i, typ := range r.rows.types {
		vals[i] = reflect.New(typ).Interface()
	}

	if err := r.rows.x.Scan(vals...); err != nil {
		r.err = err
		return err
	}

	r.data = make(map[string]interface{}, size)

	for i, x := range vals {
		r.data[r.rows.names[i]] = reflect.ValueOf(x).Elem().Interface()
	}
	return nil
}

func (r *Row) get(name string) interface{} {
	if r.err != nil {
		return nil
	}

	x, ok := r.data[name]
	if !ok {
		x, ok = r.data[toSnake(name)]
	}

	if !ok {
		return nil
	}
	return x
}

// GetInt ...
func (r *Row) GetInt(name string, def ...int) int {
	v := int(0)
	if len(def) > 0 {
		v = def[0]
	}

	x := r.get(name)
	if x == nil {
		return v
	}
	return toInt(x, v)
}

// GetUint ...
func (r *Row) GetUint(name string, def ...uint) uint {
	v := uint(0)
	if len(def) > 0 {
		v = def[0]
	}

	x := r.get(name)
	if x == nil {
		return v
	}
	return toUint(x, v)
}

// GetInt8 ...
func (r *Row) GetInt8(name string, def ...int8) int8 {
	v := int8(0)
	if len(def) > 0 {
		v = def[0]
	}

	x := r.get(name)
	if x == nil {
		return v
	}
	return toInt8(x, v)
}

// GetUint8 ...
func (r *Row) GetUint8(name string, def ...uint8) uint8 {
	v := uint8(0)
	if len(def) > 0 {
		v = def[0]
	}
	x := r.get(name)
	if x == nil {
		return v
	}
	return toUint8(x, v)
}

// GetInt16 ...
func (r *Row) GetInt16(name string, def ...int16) int16 {
	v := int16(0)
	if len(def) > 0 {
		v = def[0]
	}
	x := r.get(name)
	if x == nil {
		return v
	}
	return toInt16(x, v)
}

// GetUint16 ...
func (r *Row) GetUint16(name string, def ...uint16) uint16 {
	v := uint16(0)
	if len(def) > 0 {
		v = def[0]
	}
	x := r.get(name)
	if x == nil {
		return v
	}
	return toUint16(x, v)
}

// GetInt32 ...
func (r *Row) GetInt32(name string, def ...int32) int32 {
	v := int32(0)
	if len(def) > 0 {
		v = def[0]
	}
	x := r.get(name)
	if x == nil {
		return v
	}
	return toInt32(x, v)
}

// GetUint32 ...
func (r *Row) GetUint32(name string, def ...uint32) uint32 {
	v := uint32(0)
	if len(def) > 0 {
		v = def[0]
	}
	x := r.get(name)
	if x == nil {
		return v
	}
	return toUint32(x, v)
}

// GetInt64 ...
func (r *Row) GetInt64(name string, def ...int64) int64 {
	v := int64(0)
	if len(def) > 0 {
		v = def[0]
	}
	x := r.get(name)
	if x == nil {
		return v
	}
	return toInt64(x, v)
}

// GetUint64 ...
func (r *Row) GetUint64(name string, def ...uint64) uint64 {
	v := uint64(0)
	if len(def) > 0 {
		v = def[0]
	}
	x := r.get(name)
	if x == nil {
		return v
	}
	return toUint64(x, v)
}

// GetFloat32 ...
func (r *Row) GetFloat32(name string, def ...float32) float32 {
	v := float32(0)
	if len(def) > 0 {
		v = def[0]
	}
	x := r.get(name)
	if x == nil {
		return v
	}
	return toFloat32(x, v)
}

// GetFloat64 ...
func (r *Row) GetFloat64(name string, def ...float64) float64 {
	v := float64(0)
	if len(def) > 0 {
		v = def[0]
	}
	x := r.get(name)
	if x == nil {
		return v
	}
	return toFloat64(x, v)
}

// GetString ...
func (r *Row) GetString(name string, def ...string) string {
	v := ""
	if len(def) > 0 {
		v = def[0]
	}
	x := r.get(name)
	if x == nil {
		return v
	}
	return toString(x, v)
}

// GetTime ...
func (r *Row) GetTime(name string, def ...time.Time) time.Time {
	v := zeroTime
	if len(def) > 0 {
		v = def[0]
	}
	x := r.get(name)
	if x == nil {
		return v
	}
	return toTime(x, v)
}

// GetBool ...
func (r *Row) GetBool(name string, def ...bool) bool {
	v := false
	if len(def) > 0 {
		v = def[0]
	}
	x := r.get(name)
	if x == nil {
		fmt.Printf("GetBool %s is nil", name)
		return v
	}
	return toBool(x, v)
}

// GetBytes ...
func (r *Row) GetBytes(name string, def ...[]byte) []byte {
	var v []byte
	if len(def) > 0 {
		v = def[0]
	}
	x := r.get(name)
	if x == nil {
		return v
	}
	return toBytes(x, v)
}

// GetDuration ...
func (r *Row) GetDuration(name string, def ...time.Duration) time.Duration {
	v := time.Duration(0)
	if len(def) > 0 {
		v = def[0]
	}

	x := r.get(name)
	if x == nil {
		return v
	}
	return toDuration(x, v)
}

// Unmarshal ...
func (r *Row) Unmarshal(x interface{}) error {
	return r._unmarshal("", x)
}

// Data ...
func (r *Row) Data() (interface{}, error) {
	if r.err != nil {
		return nil, r.err
	}

	x := reflect.New(r.rows.s).Interface()
	if err := r.Unmarshal(x); err != nil {
		return nil, err
	}
	return x, nil
}

func (r *Row) _unmarshal(prefix string, d interface{}) (err error) {
	if r.err != nil {
		return r.err
	}
	defer func() {
		if pr := recover(); pr != nil {
			err = fmt.Errorf("%v", pr)
		}
	}()

	v := reflect.ValueOf(d)

	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("d must be pointer")
	}

	v = v.Elem()
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("d must be pointer of struct")
	}

	typ := v.Type()
	size := typ.NumField()

	for i := 0; i < size; i++ {
		f := v.Field(i)
		t := typ.Field(i)
		if !f.CanSet() {
			continue
		}
		name := ""
		if tag := t.Tag.Get("sqlx"); tag != "" && tag != "-" {
			name = tag
		} else if tag := t.Tag.Get("json"); tag != "" && tag != "-" {
			name = tag
		} else {
			name = toSnake(t.Name)
		}

		if prefix != "" {
			name = fmt.Sprintf("%s.%s", prefix, name)
		}

		switch f.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			f.SetInt(r.GetInt64(name))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			f.SetUint(r.GetUint64(name))
		case reflect.String:
			f.SetString(r.GetString(name))
		case reflect.Float32, reflect.Float64:
			f.SetFloat(r.GetFloat64(name))
		case reflect.Bool:
			f.SetBool(r.GetBool(name))
		case reflect.Ptr:
			if f.Type().Elem().Kind() == reflect.Struct {
				nextV := reflect.New(f.Type().Elem())
				if err := r._unmarshal(name, nextV.Interface()); err != nil {
					return err
				}
			}
		default:
			switch f.Interface().(type) {
			case []byte:
				f.Set(reflect.ValueOf(r.GetBytes(name)))
			case time.Time:
				f.Set(reflect.ValueOf(r.GetTime(name)))
			case time.Duration:
				f.Set(reflect.ValueOf(r.GetDuration(name)))
			}
		}

	}
	return nil
}
