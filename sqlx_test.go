package sqlx_test

import (
	"fmt"
	"reflect"
	"time"

	"github.com/dairaga/sqlx"
)

const dsnFmt = "%s:%s@tcp(%s:%d)/%s?%s=%s&%s=%s"

var (
	dsn = fmt.Sprintf(dsnFmt, "test", "test", "127.0.0.1", 3306, "mytest", "charset", "utf8mb4,utf8", "parseTime", "true")
)

var (
	timeType = reflect.TypeOf(time.Time{})
)

const (
	test1Row  = `{"c01":"AQ==","c01_1":"AQ==","c02_s":-2,"c02_u":2,"c03":1,"c04_s":-4,"c04_u":4,"c05_s":-5,"c05_u":5,"c06_s":-6,"c06_u":6,"c07_s":-7,"c07_u":7,"c08_s":-8,"c08_u":8,"c09_s":-9,"c09_u":9,"c10_s":-10,"c10_u":10,"c11":"2019-01-01T00:00:00Z","c12":"MTIzOjA0OjA1","c13":2019,"c14":"2014-09-23T10:01:02Z","c15":"2014-09-23T10:01:02Z","c16":"C16","c17":"C17","c18":"QzE4AAAAAAAAAA==","c19":"QzE5","c20":"QzIw","c21":"C21","c22":"QzIy","c23":"C23","c24":"QzI0","c25":"C25","c26":"QzI2","c27":"C27","c28":"a","c29":"b"}`
	test1Rows = `[{"c01":"AQ==","c01_1":"AQ==","c02_s":-2,"c02_u":2,"c03":1,"c04_s":-4,"c04_u":4,"c05_s":-5,"c05_u":5,"c06_s":-6,"c06_u":6,"c07_s":-7,"c07_u":7,"c08_s":-8,"c08_u":8,"c09_s":-9,"c09_u":9,"c10_s":-10,"c10_u":10,"c11":"2019-01-01T00:00:00Z","c12":"MTIzOjA0OjA1","c13":2019,"c14":"2014-09-23T10:01:02Z","c15":"2014-09-23T10:01:02Z","c16":"C16","c17":"C17","c18":"QzE4AAAAAAAAAA==","c19":"QzE5","c20":"QzIw","c21":"C21","c22":"QzIy","c23":"C23","c24":"QzI0","c25":"C25","c26":"QzI2","c27":"C27","c28":"a","c29":"b"}]`

	test2Row  = `{"c01":null,"c01_1":null,"c02_s":0,"c02_u":0,"c03":0,"c04_s":0,"c04_u":0,"c05_s":0,"c05_u":0,"c06_s":0,"c06_u":0,"c07_s":0,"c07_u":0,"c08_s":0,"c08_u":0,"c09_s":0,"c09_u":0,"c10_s":0,"c10_u":0,"c11":"0001-01-01T00:00:00Z","c12":null,"c13":0,"c14":"0001-01-01T00:00:00Z","c15":"2014-09-23T10:01:02Z","c16":"","c17":"","c18":null,"c19":null,"c20":null,"c21":"","c22":null,"c23":"","c24":null,"c25":"","c26":null,"c27":"","c28":"","c29":""}`
	test2Rows = `[{"c01":null,"c01_1":null,"c02_s":0,"c02_u":0,"c03":0,"c04_s":0,"c04_u":0,"c05_s":0,"c05_u":0,"c06_s":0,"c06_u":0,"c07_s":0,"c07_u":0,"c08_s":0,"c08_u":0,"c09_s":0,"c09_u":0,"c10_s":0,"c10_u":0,"c11":"0001-01-01T00:00:00Z","c12":null,"c13":0,"c14":"0001-01-01T00:00:00Z","c15":"2014-09-23T10:01:02Z","c16":"","c17":"","c18":null,"c19":null,"c20":null,"c21":"","c22":null,"c23":"","c24":null,"c25":"","c26":null,"c27":"","c28":"","c29":""}]`

	s1      = `{"id":6,"name":"C16","birthday":"2014-09-23T10:01:02Z"}`
	s1Slice = `[{"id":6,"name":"C16","birthday":"2014-09-23T10:01:02Z"}]`

	s2      = `{"id":0,"name":"","birthday":"0001-01-01T00:00:00Z"}`
	s2Slice = `[{"id":0,"name":"","birthday":"0001-01-01T00:00:00Z"}]`
)

type TestStruct1 struct {
	ID       int64     `sqlx:"c06_u" json:"id"`
	Name     string    `sqlx:"c16" json:"name"`
	Birthday time.Time `sqlx:"c14" json:"birthday"`
}

func OpenDB(d, dsn string) (*sqlx.DB, error) {
	return sqlx.OpenDB("mysql", dsn)
}
