package sqlx

import "fmt"

func ExampleCmd_Select() {

	cmd := NewCmd()

	cmd.Reset()
	cmd.Select("1 + 1")
	fmt.Println(cmd)

	cmd.Reset()
	cmd.Select("1 + 1").From("DUAL")
	fmt.Println(cmd)

	cmd.Reset()
	cmd.Select().From("t1").InnerJoinOn("t2", "t1.id=t2.id")
	fmt.Println(cmd)

	cmd.Reset()
	cmd.Select("CONCAT(last_name,', ',first_name) AS full_name").From("mytable").OrderBy("full_name")
	fmt.Println(cmd)

	cmd.Reset()
	cmd.Select("t1.name", "t2.salary").From("employee AS t1").From("info AS t2").Where("t1.name = t2.name")
	fmt.Println(cmd)

	cmd.Reset()
	cmd.Select("a", "COUNT(b)").From("test_table").GroupBy("a").OrderBy("NULL")
	fmt.Println(cmd)

	cmd.Reset()
	cmd.Select("COUNT(col1) AS col2").From("t").GroupBy("col2").Having("col2 = 2")
	fmt.Println(cmd)

	cmd.Reset()
	cmd.Select("REPEAT('a',1)").Union().Select("REPEAT('b',10)")
	fmt.Println(cmd)

	cmd.Reset()
	cmd.Select("college", "region", "seed").From("tournament").OrderBy("region").OrderBy("seed")
	fmt.Println(cmd)

	cmd.Reset()
	cmd.Select("user", "MAX(salary)").From("users").GroupBy("user").Having("MAX(salary) > 10")
	fmt.Println(cmd)

	cmd.Reset()
	cmd.Select().From("tbl").Limit(10, 5)
	fmt.Println(cmd)

	// Output:
	// SELECT 1 + 1
	// SELECT 1 + 1 FROM DUAL
	// SELECT * FROM t1 INNER JOIN t2 ON t1.id=t2.id
	// SELECT CONCAT(last_name,', ',first_name) AS full_name FROM mytable ORDER BY full_name
	// SELECT t1.name,t2.salary FROM employee AS t1,info AS t2 WHERE t1.name = t2.name
	// SELECT a,COUNT(b) FROM test_table GROUP BY a ORDER BY NULL
	// SELECT COUNT(col1) AS col2 FROM t GROUP BY col2 HAVING col2 = 2
	// SELECT REPEAT('a',1) UNION SELECT REPEAT('b',10)
	// SELECT college,region,seed FROM tournament ORDER BY region, seed
	// SELECT user,MAX(salary) FROM users GROUP BY user HAVING MAX(salary) > 10
	// SELECT * FROM tbl LIMIT 5,10
}

func ExampleCmd_Union() {
	cmd := NewCmd()
	sub := NewCmd()

	cmd.Reset()
	sub.Reset()
	cmd.Select("a").From("t1").WhereAnd("a=10", "B=1").OrderBy("a").Limit(10)
	sub.Select("a").From("t2").WhereAnd("a=11", "B=2").OrderBy("a").Limit(10)
	cmd.Union(sub)
	fmt.Println(cmd)

	cmd.Reset()
	sub.Reset()
	cmd.Select("a").From("t1").WhereAnd("a=10", "B=1")
	sub.Select("a").From("t2").WhereAnd("a=11", "B=2")
	cmd.Union(sub)
	cmd.OrderBy("a").Limit(10)
	fmt.Println(cmd)

	cmd.Reset()
	sub.Reset()
	cmd.Select("1 AS sort_col", "col1a", "col1b").From("t1")
	sub.Select("2", "col2a", "col2b").From("t2")
	cmd.Union(sub)
	cmd.OrderBy("sort_col").OrderBy("col1a")
	fmt.Println(cmd)

	// Output:
	// (SELECT a FROM t1 WHERE a=10 AND B=1 ORDER BY a LIMIT 10) UNION (SELECT a FROM t2 WHERE a=11 AND B=2 ORDER BY a LIMIT 10)
	// (SELECT a FROM t1 WHERE a=10 AND B=1) UNION (SELECT a FROM t2 WHERE a=11 AND B=2) ORDER BY a LIMIT 10
	// (SELECT 1 AS sort_col,col1a,col1b FROM t1) UNION (SELECT 2,col2a,col2b FROM t2) ORDER BY sort_col, col1a
}

func ExampleCmd_Join() {

	cmd := NewCmd()

	cmd.Reset()
	cmd.Select().From("t1").LeftJoin("t2", "t3", "t4").On("t2.a=t1.a", "t3.b=t1.b", "t4.c=t1.c")
	fmt.Println(cmd)

	cmd.Reset()
	cmd.Select().From("t1").RightJoinOn("t2", "t1.a=t2.a")
	fmt.Println(cmd)

	cmd.Reset()
	cmd.Select().From("table1").LeftJoinOn("table2", "table1.id=table2.id").LeftJoinOn("table3", "table2.id=table3.id")
	fmt.Println(cmd)

	// Output:
	// SELECT * FROM t1 LEFT JOIN (t2,t3,t4) ON (t2.a=t1.a AND t3.b=t1.b AND t4.c=t1.c)
	// SELECT * FROM t1 RIGHT JOIN t2 ON t1.a=t2.a
	// SELECT * FROM table1 LEFT JOIN table2 ON table1.id=table2.id LEFT JOIN table3 ON table2.id=table3.id
}

func ExampleCmd_Delete() {
	cmd := NewCmd()

	cmd.Reset()
	cmd.Delete().From("somelog").Where("user='jcole'").OrderBy("timestamp_column", DESC).Limit(1)
	fmt.Println(cmd)

	cmd.Reset()
	cmd.Delete("t1", "t2").From("t1").InnerJoinOn("t2", "t1.id=t2.id").InnerJoinOn("t3", "t2.id=t3.id")
	fmt.Println(cmd)

	cmd.Reset()
	cmd.Delete("t1").From("t1").LeftJoinOn("t2", "t1.id=t2.id").Where("t2.id IS NULL")
	fmt.Println(cmd)

	cmd.Reset()
	cmd.Delete("a1", "a2").From("t1 AS a1").InnerJoin("t2 AS a2").Where("a1.id=a2.id")
	fmt.Println(cmd)

	// Output:
	// DELETE FROM somelog WHERE user='jcole' ORDER BY timestamp_column DESC LIMIT 1
	// DELETE t1,t2 FROM t1 INNER JOIN t2 ON t1.id=t2.id INNER JOIN t3 ON t2.id=t3.id
	// DELETE t1 FROM t1 LEFT JOIN t2 ON t1.id=t2.id WHERE t2.id IS NULL
	// DELETE a1,a2 FROM t1 AS a1 INNER JOIN t2 AS a2 WHERE a1.id=a2.id
}

func ExampleCmd_Insert() {

	cmd := NewCmd()

	cmd.Reset()
	cmd.Insert("tbl_name", "col1", "col2").Values()
	fmt.Println(cmd)

	cmd.Reset()
	cmd.Insert("tbl_name", "a", "b", "c").Values().Values().Values()
	fmt.Println(cmd)

	cmd.Reset()
	cmd.Insert("tbl_temp2", "fld_id").Select("tbl_temp1.fld_order_id").From("tbl_temp1").Where("tbl_temp1.fld_order_id > 100")
	fmt.Println(cmd)

	cmd.Reset()
	cmd.Insert("t1", "a", "b", "c").Values().Duplicate("c=c+1")
	fmt.Println(cmd)

	cmd.Reset()
	cmd.Insert("t1", "a", "b", "c").Values().DuplicateValues("a", "b", "c")
	fmt.Println(cmd)

	cmd.Reset()
	cmd.Insert("t1", "a", "b", "c").Values().Values().Duplicate("c=VALUES(a)+VALUES(b)")
	fmt.Println(cmd)

	cmd.Reset()
	cmd.Insert("t1", "a", "b").Select("c", "d").From("t2").Union().Select("e", "f").From("t3").Duplicate("b=b+c")
	fmt.Println(cmd)

	sub := NewCmd()
	sub.Select("c", "d").From("t2").Union().Select("e", "f").From("t3")

	cmd.Reset()
	cmd.Insert("t1", "a", "b").Select().SubQuery(sub, "dt").Duplicate("b=b+c")
	fmt.Println(cmd)

	// Output:
	// INSERT INTO tbl_name (col1,col2) VALUES (?,?)
	// INSERT INTO tbl_name (a,b,c) VALUES (?,?,?),(?,?,?),(?,?,?)
	// INSERT INTO tbl_temp2 (fld_id) SELECT tbl_temp1.fld_order_id FROM tbl_temp1 WHERE tbl_temp1.fld_order_id > 100
	// INSERT INTO t1 (a,b,c) VALUES (?,?,?) ON DUPLICATE KEY UPDATE c=c+1
	// INSERT INTO t1 (a,b,c) VALUES (?,?,?) ON DUPLICATE KEY UPDATE a=VALUES(a),b=VALUES(b),c=VALUES(c)
	// INSERT INTO t1 (a,b,c) VALUES (?,?,?),(?,?,?) ON DUPLICATE KEY UPDATE c=VALUES(a)+VALUES(b)
	// INSERT INTO t1 (a,b) SELECT c,d FROM t2 UNION SELECT e,f FROM t3 ON DUPLICATE KEY UPDATE b=b+c
	// INSERT INTO t1 (a,b) SELECT * FROM (SELECT c,d FROM t2 UNION SELECT e,f FROM t3) AS dt ON DUPLICATE KEY UPDATE b=b+c
}

func ExampleCmd_Replace() {

	cmd := NewCmd()
	cmd.Reset()
	cmd.Replace("test", "a", "b", "c").Values()
	fmt.Println(cmd)

	// Output:
	// REPLACE INTO test(a,b,c) VALUES (?,?,?)
}

func ExampleCmd_Update() {

	cmd := NewCmd()

	cmd.Reset()
	cmd.Update("t1").Set("col1=col1+1")
	fmt.Println(cmd)

	cmd.Reset()
	cmd.Update("t1").SetFields("col1")
	fmt.Println(cmd)

	cmd.Reset()
	cmd.Update("t1").Set("col1=col1+1", "col2=col1")
	fmt.Println(cmd)

	cmd.Reset()
	cmd.Update("t").Set("id=id+1").OrderBy("id", DESC)
	fmt.Println(cmd)

	cmd.Reset()
	cmd.Update("items", "month").Set("items.price=month.price").Where("items.id=month.id")
	fmt.Println(cmd)

	// Output:
	// UPDATE t1 SET col1=col1+1
	// UPDATE t1 SET col1=?
	// UPDATE t1 SET col1=col1+1,col2=col1
	// UPDATE t SET id=id+1 ORDER BY id DESC
	// UPDATE items,month SET items.price=month.price WHERE items.id=month.id
}
