package main

import (
	"database/sql"
	"fmt"
)

func main() {
	db, err := sql.Open("mysql", "")
	checkErr(err)

	stmt, err := db.Prepare("")
	checkErr(err)

	res, err := stmt.Exec("")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}