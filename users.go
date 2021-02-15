package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username 	string `json:"username"`
	Password 	string `json:"password"`
	Base_ID  	string `json:"base_ID"`
	Branch_ID 	string `json:"branch_ID"`
	Capsule_ID 	string `json:"capsule_ID"`
	Permissions []bool `json:"permissions"`
	Admin 		[]bool `json:"admin"`
}