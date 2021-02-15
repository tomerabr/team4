package main

import (
	"fmt"
	"strings"
	"database/sql"
)

var SELECT_QUERY = "SELECT %s from %s WHERE %s"
var INSERT_QUERY = "INSERT INTO %s (%s) VALUES (%s)"
var UPDATE_QUERY = "UPDATE %s SET %s WHERE %s"
var DELETE_QUERY = "DELETE FROM %s WHERE %s"

var VALUES_COLUMNS_MISMATCH = "ERROR: number of values is different to the number of columns"
var INSERT_SUCCESS = ""
var UPDATE_SUCCESS = ""
var DELETE_SUCCESS = ""

func Select(table_name string, cols []string, conditions string, db *sql.DB) *sql.Row {
	cols_string := strings.join(cols, ", ")	
	query 		:= fmt.Sprintf(SELECT_QUERY, cols_string, table_name, conditions)

	rows, err := db.Query(query)

	defer rows.Close()

	if err != nil {
		return err
	}

	return rows
}

func Insert(table_name string, cols []string, values []string, db *sql.DB) string {
	if len(cols) != len(values) {
		return VALUES_COLUMNS_MISMATCH
	}

	cols_string   := strings.Join(cols, ", ")
	values_string := strings.Join(values, ", ")
	query  		  := fmt.Sprintf(INSERT_QUERY, table_name, cols_string, values_string)

	_, err := db.Exec(query)

	if err != nil {
		return err.Error()
	}

	return INSERT_SUCCESS
}

func makeValueSetString(cols []string, newVals []string) string {
	valueSetList = make([]string, len(cols))
	for _, i := range cols {
		set_str := cols[i] + " = " newVals[i])
		append(valueSetList, set_str)
	}
	valuesSetStr := strings.Join(set_str, ", ")

	return valuesSetStr
}

func Update(table_name string, cols []string, newVals []string, conditions string, db *sql.DB) string {
	if len(cols) != len(newVals) {
		return VALUES_COLUMNS_MISMATCH
	}

	newValsSet := makeValueSetString(cols, newVals)
	query 	   := fmt.Sprintf(UPDATE_QUERY, table_name, newValsSet, conditions)

	_, err = db.Exec(query)

	if err != nil {
		return err.Error()
	}

	return UPDATE_SUCCESS
}

func Delete(table_name string, conditions string, db *sql.DB) string {
	query := fmt.Sprintf(DELETE_QUERY, table_name, conditions)

	_, err = db.Exec(query)

	if err != nil {
		return err.Error()
	}

	return DELETE_SUCCESS
}