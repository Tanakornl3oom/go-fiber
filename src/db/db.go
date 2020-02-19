package db

import (	"database/sql"
			_ "github.com/go-sql-driver/mysql"
			"fmt"
		)
// struct name should starts with Capital letter to export
type Usersdata struct {
	ID              int
	Fristname       string
	Lastname        string
	Age             int
}

func ConnectDB() bool{
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/testsck")
	defer db.Close()
	if err != nil {
		fmt.Println("connect fail")
		return false
	} else {
		fmt.Println("connect success")
		return true
	}
}

