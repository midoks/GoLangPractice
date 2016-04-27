package main

import (
	"database/sql"
	"fmt"

	"time"

	_ "github.com/go-sql-driver/mysql"
)

//type TestMysql struct {
//	db *sql.DB
//}

//db.Query("create table t(id int primary key auto_increment, name varchar(255) not null, ts timestamp)")
func insertTable(db *sql.DB) {
	stmt, err := db.Prepare("insert into t(name,ts) values(?,?)")

	if err != nil {
		fmt.Println("err insert:", err.Error())
	}

	ts, _ := time.Parse("2006-01-02 15:04:05", "2014-08-28 15:04:00")
	stmt.Exec("test", ts)
}

func queryTable(db *sql.DB) {

	rows, err := db.Query("select * from t")

	if err != nil {
		fmt.Println("query err:", err.Error())
	}
	defer rows.Close()

	cols, _ := rows.Columns()
	for i := range cols {
		fmt.Print(cols[i])
		fmt.Print("\t")
	}

	fmt.Println("\n-----------------------")
	var id int
	var name string
	var age int
	for rows.Next() {

		if err := rows.Scan(&id, &name, &age); err != nil {

			fmt.Print(id, "\t", name, "\t", age, "\n")
		}
	}
}

func updateTable(db *sql.DB) {

	stmt, err := db.Prepare("update t set name=? where id = 1")
	if err != nil {
		fmt.Println("update err", err.Error())
	}

	defer stmt.Close()
	//rnd, _ := rand.Int(rand.Reader, big.NewInt(100))

	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	tm_string := tm.Format("2006-01-02-03-04-05")
	name := "test_" + string(tm_string)

	if result, err := stmt.Exec(name); err == nil {
		//fmt.Println(result, err)
		if c, err := result.RowsAffected(); err == nil {
			fmt.Println("update count:", c, result)
		}
	} else {
		//fmt.Println(result, err)
	}

}

func deleteTable(db *sql.DB) {
	stmt, err := db.Prepare("delete from t where id = ?")
	if err != nil {
		fmt.Println("delete err", err.Error())
	}
	if result, err := stmt.Exec(1); err == nil {
		if c, err := result.RowsAffected(); err == nil {
			fmt.Println("delete count:", c, result)
		}
	}

}

func main() {

	db, err := sql.Open("mysql", "root:123123@tcp(127.0.0.1:3306)/test?charset=utf8")

	if err != nil {
		fmt.Println("Link MySQL fail:", err.Error())
	} else {
		fmt.Println("Link MySQL ok")

	}
	//fmt.Println(db)
	//insertTable(db)
	//queryTable(db)
	//updateTable(db)
	deleteTable(db)
}
