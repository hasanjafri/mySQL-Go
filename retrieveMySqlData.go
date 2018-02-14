package main

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "fmt"

func main() {
	db, err := sql.Open("mysql", "root:newpassword@/ZatiqDB")
	checkErr(err)

	stmt, err := db.Prepare("INSERT ZatiqBusinesses SET businessEmail=?,businessPassword=?,businessName=?")
	checkErr(err)

	res, err := stmt.Exec("go", "golang", "golang")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	stmt, err = db.Prepare("UPDATE ZatiqBusinesses SET businessEmail=? WHERE id=?")
	checkErr(err)

	res, err = stmt.Exec("go@golang.com", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	rows, err := db.Query("SELECT * FROM ZatiqBusinesses")
	checkErr(err)

	for rows.Next() {
		var id int
		var businessName string
		var businessEmail string
		var businessPassword string
		err = rows.Scan(&id, &businessName, &businessEmail, &businessPassword)
		checkErr(err)
		fmt.Println(id)
		fmt.Println(businessName)
		fmt.Println(businessEmail)
		fmt.Println(businessPassword)
	}

	stmt, err = db.Prepare("DELETE FROM ZatiqBusinesses where id=?")
	checkErr(err)

	res, err = stmt.Exec(20)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
