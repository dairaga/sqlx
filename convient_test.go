package sqlx

import "testing"

const (
	t1 = "a_id_name_age"
	t2 = "AIdNameAge"
)

func TestToCamel(t *testing.T) {
	test := toCamel(t1)
	if test != t2 {
		t.Errorf("toCamel %s, %s", test, t2)
	}
}

func TestToSnake(t *testing.T) {
	test := toSnake(t2)
	if test != t1 {
		t.Errorf("toSnake %s, %s", test, t1)
	}
}
