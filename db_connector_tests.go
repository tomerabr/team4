package main

import (
	"fmt"
	"database/sql"
)

// func InsertBase(db *sql.DB) {
// 	baseName := "Gidonim"
// 	insert_result := Insert("bases", []string{"name"}, []string{baseName}, db)
// 	if insert_result != "" {
// 		panic("Failed to insert base Gidonim")
// 	}
// }

func TestInsert(db *sql.DB) {
	baseName := "Gidonim"
	insert_result := Insert("bases", []string{"name"}, []string{baseName}, db)
	fmt.Printf(insert_result)
	
	// table_name := "users"
	// ID := 8421633
	// username = "Roy"
	// password := "12345"
	// base_ID := 0
	// branch_ID := 0
	// capsule_ID := 0
	// permissions := []bool{true, true, true, true, true, true}
	// admin := []bool{false, false, false, false, false, false}

	// columns = "ID, username, password, base_ID, branch_ID, capsule_ID, permissions, admin"
	// values := []string[string(ID), username, password, string(base_ID),
	// 	string(branch_ID), string(capsule_ID), string(permissions), string(admin)]
	// result = Insert(table_name, columns, values, db)
	// fmt.Printf(result)
}