package main

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "github.com/gin-gonic/gin"
    // "encoding/json"
//    cors "github.com/rs/cors/wrapper/gin"

)

func main() {
    fmt.Println("hello world!")

    db, err := sql.Open("mysql", 
        "team4@team4-mysql-db:Aa123456123456@tcp(team4-mysql-db.mysql.database.azure.com:3306)/team4")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        panic(err)
    }
    

    // gin.SetMode(gin.ReleaseMode)
    r := gin.Default()
    // r.Use(cors.Default())

    r.POST("/login", func(c *gin.Context) {
		
	})
    r.Run()

}