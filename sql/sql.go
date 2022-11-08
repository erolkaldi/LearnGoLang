package main

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	connString := "server=erol\\sql19;user id=sa;password=1;database=summer"
	db, err := sql.Open("mssql", connString)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if err := db.Ping(); err != nil {
		fmt.Println(err.Error())
		return
	}

	rows, err := db.Query("Select top 10 recId,hotelName from Hotel")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var recId int
	var hotelName string

	for rows.Next() {
		rows.Scan(&recId, &hotelName)
		fmt.Println("Id:", recId, "Name:", hotelName)
	}

}
