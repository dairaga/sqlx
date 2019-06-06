package sqlx

import (
	"encoding/json"
	"testing"
)

func TestDBXData(t *testing.T) {
	db, err := OpenDB("mysql", dsn)

	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	d1, err := db.Reset().Select().From("test1").Data()
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

	d2, err := db.Reset().Select().From("test2").Data()
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

func TestDBXUnmarshal(t *testing.T) {
	db, err := OpenDB("mysql", dsn)

	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	d1 := &TestStruct1{}
	err = db.Reset().Select().From("test1").Unmarshal(d1)
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

	err = db.Reset().Select().From("test2").Unmarshal(d1)
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

func TestDBXAll(t *testing.T) {
	db, err := OpenDB("mysql", dsn)
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	all, err := db.Reset().Select().From("test1").All()

	if err != nil {
		t.Errorf("to internal data: %v", err)
	}
	tmp, err := json.Marshal(all)
	if err != nil {
		t.Errorf("marshal: %v", err)
	}
	if string(tmp) != test1Rows {
		t.Errorf("json string should be \n%s\n but \n%s", test1Rows, string(tmp))
	}

	all, err = db.Reset().Select().From("test2").All()
	if err != nil {
		t.Errorf("to internal data: %v", err)
	}
	tmp, err = json.Marshal(all)
	if err != nil {
		t.Errorf("marshal: %v", err)
	}
	if string(tmp) != test2Rows {
		t.Errorf("json string should be \n%s\n but \n%s", test2Rows, string(tmp))
	}
}

func TestDBXUnmarshalAll(t *testing.T) {
	db, err := OpenDB("mysql", dsn)

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	d1 := []TestStruct1{}

	err = db.Reset().Select().From("test1").UnmarshalAll(&d1)
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

	err = db.Reset().Select().From("test2").UnmarshalAll(&d1)
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
