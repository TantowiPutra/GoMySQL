package belajar_golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	// "fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_golang_mysql")

	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

func PrintRows(data *sql.Rows) {
	for data.Next() {
		var id, name string

		err := data.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Println("Id: ", id)
		fmt.Println("Name: ", name)
	}
}

func TestInsertCustomer(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(id, name) VALUES('123', 'Putra')"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success Insert data Into Database")
}

func TestQueryContext(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT * FROM customer"
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}

	PrintRows(rows)

	defer rows.Close()
}

func RecoverPanic() {
	rcv := recover()
	fmt.Println(rcv)
}

func TestDeferPanic(t *testing.T) {
	defer RecoverPanic()

	panic("Terjadi Error")
}

func TestLoopFor(t *testing.T) {
	// ! Beware Infinite Looping :O
	array := [5]int{1, 2, 3, 4, 5}
	for {
		fmt.Println(array)
	}
}
