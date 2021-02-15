package main

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    // "github.com/gin-gonic/gin"
    // "net/http"
    // "encoding/json"
	// cors "github.com/rs/cors/wrapper/gin"
)

var DB_TYPE = "mysql"
var DB_URL  = "team4@team4-mysql-db:Aa123456123456@tcp(team4-mysql-db.mysql.database.azure.com:3306)/team4"

func main() {
    fmt.Println("hello world!")

    db, err := sql.Open(DB_TYPE, DB_URL)
    if err != nil {
        panic(err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        panic(err)
    }
    
	TestInsert(db)

    // // gin.SetMode(gin.ReleaseMode)
    // r := gin.Default()
    // // r.Use(cors.Default())

    // r.POST("/login", func(c *gin.Context) { //id, pass
	// 	var u User
	// 	c.BindJSON(&u)
	// 	valid, reason, name,base,capsule,branch,admin,permissions := addUser(db, u)
		
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"valid":  valid,
	// 		"reason": reason,
    //         "name": name,
    //         "base": base,
    //         "capsule": capsule,
    //         "branch": branch,
    //         "admin": admin,
    //         "permissions": permissions,
	// 	})
	// })
    // r.POST("/register", func(c *gin.Context) { //id, pass
	// 	var u User
	// 	c.BindJSON(&u)
	// 	valid, reason := registerUser(db, u)
		
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"valid":  valid,
	// 		"reason": reason,
	// 	})
	// })
    // r.POST("/change/pass", func(c *gin.Context) { //id, pass
	// 	var u User
	// 	c.BindJSON(&u)
	// 	valid, reason := changePass(db, u)
		
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"valid":  valid,
	// 		"reason": reason,
    //         "newpass": u.Password,
	// 	})
	// })
    // r.POST("/change/branch", func(c *gin.Context) { //id, pass
	// 	var u User
	// 	c.BindJSON(&u)
	// 	valid, reason, name := changeBranch(db, u)
		
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"valid":  valid,
	// 		"reason": reason,
    //         "newbranch": name,
	// 	})
	// })
    // r.POST("/change/base", func(c *gin.Context) { //id, pass
	// 	var u User
	// 	c.BindJSON(&u)
	// 	valid, reason,name := changeBase(db, u)
		
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"valid":  valid,
	// 		"reason": reason,
    //         "newbase": name,
	// 	})
	// })
    // r.POST("/change/capsule", func(c *gin.Context) { //id, pass
	// 	var u User
	// 	c.BindJSON(&u)
	// 	valid, reason, id := changeCapsule(db, u)
		
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"valid":  valid,
	// 		"reason": reason,
    //         "newcapsule": id,
	// 	})
	// })
    // r.POST("/update/admin", func(c *gin.Context) { //id, pass
	// 	var u User
	// 	c.BindJSON(&u)
	// 	valid, reason := updateAdmin(db, u)
		
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"valid":  valid,
	// 		"reason": reason,
	// 	})
	// })
    // r.POST("/update/permission", func(c *gin.Context) { //id, pass
	// 	var u User
	// 	c.BindJSON(&u)
	// 	valid, reason := updatePermission(db, u)
		
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"valid":  valid,
	// 		"reason": reason,
	// 	})
	// })
    
    
    // r.Run()

}