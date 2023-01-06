package main

import (
	"database/sql"
	"fmt"
	"log"
)

func DeleteUser(sql *sql.DB) {
	var email string
	fmt.Println("Write your email:")
	fmt.Scan(&email)
	// Delete
	deleteStmt := fmt.Sprintf("delete from users where email='%s'", email)
	_, e := sql.Exec(deleteStmt)
	CheckError(e)
}

func UpdateUser(sql *sql.DB) {
	var nameDB, surenameDB, email string
	fmt.Println("Write your email")
	fmt.Scan(&email)
	fmt.Print("New name:")
	fmt.Scan(&nameDB)
	fmt.Print("New surename:")
	fmt.Scan(&surenameDB)
	updateStmt := fmt.Sprintf("update users set name='%s', surname='%s' where email='%s'", nameDB, surenameDB, email)
	_, e := sql.Exec(updateStmt)
	CheckError(e)
}

func OneUser(db *sql.DB) {
	var id int
	fmt.Println("Write your your id")
	fmt.Scan(&id)
	fmt.Println("test1")
	row := db.QueryRow("select Name from users where user_id = ?", id)
	fmt.Println("test2")
	var name string
	if err := row.Scan(&name); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(name)
}
