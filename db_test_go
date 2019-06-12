package sqlx

import (
	"database/sql"
	"encoding/json"
	"testing"
)

func TestDB(t *testing.T) {
	db, err := WrapDB(sql.Open("mysql", dsn))

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	result, err := db.Data("select * from test1")
	if err != nil {
		t.Fatal(err)
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		t.Fatal(err)
	}

	if string(resultBytes) != test1Row {
		t.Errorf("json string should be \n%s\n but \n%s", test1Row, string(resultBytes))
	}

	data := TestStruct1{}

	if err := db.Unmarshal(&data, "select * from test1"); err != nil {
		t.Fatal(err)
	}

	resultBytes, err = json.Marshal(data)
	if string(resultBytes) != s1 {
		t.Errorf("json string should be \n%s\n but \n%s", s1, string(resultBytes))
	}

	result, err = db.Data("select * from test2")
	if err != nil {
		t.Fatal(err)
	}

	resultBytes, err = json.Marshal(result)
	if err != nil {
		t.Fatal(err)
	}

	if string(resultBytes) != test2Row {
		t.Errorf("json string should be \n%s\n but \n%s", test2Row, string(resultBytes))
	}

	if err := db.Unmarshal(&data, "select * from test2"); err != nil {
		t.Fatal(err)
	}

	resultBytes, err = json.Marshal(data)
	if string(resultBytes) != s2 {
		t.Errorf("json string should be \n%s\n but \n%s", s2, string(resultBytes))
	}

}

func TestDBAll(t *testing.T) {
	db, err := WrapDB(sql.Open("mysql", dsn))

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	all, err := db.All("select * from test1")
	if err != nil {
		t.Fatal(err)
	}

	allbytes, err := json.Marshal(all)
	if err != nil {
		t.Fatal(err)
	}

	if string(allbytes) != test1Rows {
		t.Errorf("json string should be \n%s\n but \n%s", test1Rows, string(allbytes))
	}

	data := []TestStruct1{}

	if err := db.UnmarshalAll(&data, "select * from test1"); err != nil {
		t.Fatal(err)
	}

	allbytes, err = json.Marshal(data)
	if string(allbytes) != s1Slice {
		t.Errorf("json string should be \n%s\n but \n%s", s1Slice, string(allbytes))
	}

	all, err = db.All("select * from test2")
	if err != nil {
		t.Fatal(err)
	}

	allbytes, err = json.Marshal(all)
	if err != nil {
		t.Fatal(err)
	}

	if string(allbytes) != test2Rows {
		t.Errorf("json string should be \n%s\n but \n%s", test2Rows, string(allbytes))
	}

	if err := db.UnmarshalAll(&data, "select * from test2"); err != nil {
		t.Fatal(err)
	}

	allbytes, err = json.Marshal(data)
	if string(allbytes) != s2Slice {
		t.Errorf("json string should be \n%s\n but \n%s", s2Slice, string(allbytes))
	}

}
