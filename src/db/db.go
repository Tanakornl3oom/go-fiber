package db

import (	"database/sql"
			_ "github.com/go-sql-driver/mysql"
			"fmt"
			"encoding/json"
			"strconv"
		)
// struct name should starts with Capital letter to export
type Usersdata struct {
	ID              int   	`json:"id"`
	Firstname       string	`json:"firstName"`
	Lastname        string	`json:"lastName"`
	Age             int		`json:"age"`
}

type User struct {
	Firstname       string
	Lastname        string
	Age             string
}

type UserDB struct {
	Firstname       string
	Lastname        string
	Age             int
}

type TableInDB struct {
	TablesInDatabase     string
}

type Response struct {
	Status       bool
	Response     interface{}
}

func ConnectDB() *sql.DB{
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/database")
	result, _ := db.Query("SHOW TABLES LIKE '%USER%'")
	for result.Next() {
		var table TableInDB
		err := result.Scan(&table.TablesInDatabase)
		if err != nil {
			fmt.Println(err)
		}
		// check table is exist
		if table == (TableInDB{}){
			fmt.Println(table)
			_, err =  db.Exec("CREATE TABLE USER (ID INT AUTO_INCREMENT,first_name VARCHAR (50) NOT NULL,last_name VARCHAR (50) NOT NULL,age INT NOT NULL,PRIMARY KEY (ID))")
			if err != nil {
				fmt.Println("Can't create table user")
				fmt.Println(err)
			}
		}
	}
	
	if err != nil {
		fmt.Println("connect fail")
		fmt.Println(err)
	} else {
		fmt.Println("connect success")
	}
	return db
}

func GetUsers(db *sql.DB) []interface{} {
	result, _ := db.Query("SELECT * FROM USER")
	var userDataList []interface{}
	if result == nil {
		return userDataList
	}
	for result.Next() {
		var user Usersdata
		err := result.Scan(
			&user.ID,
			&user.Firstname,
			&user.Lastname,
			&user.Age,
		)

		if err != nil {
			panic(err.Error())
		}
		userDataList = append(userDataList, PrettyPrint(user))
	}
	return userDataList
}

func GetUserById(db *sql.DB,id int) Usersdata {
	result, _ := db.Query("SELECT * FROM USER WHERE id = "+strconv.Itoa(id))

	var userById Usersdata
	if result == nil {
		return Usersdata{}
	}
	for result.Next() {
	err := result.Scan(
		&userById.ID,
		&userById.Firstname,
		&userById.Lastname,
		&userById.Age,
	)

	if err != nil {
		panic(err.Error())
	}
}

	return userById
}

func AddUser(db *sql.DB,user UserDB) bool {
	statement, err := db.Prepare("INSERT INTO USER ( first_name, last_name, age) VALUES (?,?,?)")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	_, err = statement.Exec(user.Firstname,user.Lastname,user.Age)

	if err != nil {
		panic(err.Error())
		return false
	}
	return true
}

func UpdateUser(db *sql.DB,user Usersdata) Response {
	statement, err := db.Prepare("UPDATE USER SET first_name= ?,last_name= ? ,age=? WHERE id = ?")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	_, err = statement.Exec(user.Firstname,user.Lastname,user.Age,user.ID)

	var response Response
	if err != nil {
		panic(err.Error())
		response.Status = false
		response.Response = "Can't update user"
		return response
	}else {
		response.Status = true
		response.Response = PrettyPrint(user)
		return response
	}
}

func DeleteUser(db *sql.DB,id int) Response {
	statement, _ := db.Prepare(`DELETE FROM USER WHERE id = ?`)

	_, err := statement.Exec(id)

	var response Response
	if err != nil {
		panic(err.Error())
		response.Status = false
		response.Response = "can't delete user"
		return response
	}else{
		response.Status = true
		response.Response = "Success delete user"
		return response
	}
	
}

func PrettyPrint(i interface{}) string {
    s, _ := json.MarshalIndent(i, "", "\t")
    return string(s)
}
