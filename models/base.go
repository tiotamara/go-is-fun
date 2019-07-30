package models

import (
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "github.com/jinzhu/gorm"
    "fmt"
    "github.com/joho/godotenv"
    "os"
)

var db *gorm.DB

func init() {
    fmt.Println("init is called here")

    if e := godotenv.Load(); e != nil {
        fmt.Print(e)
    }

    db_username := os.Getenv("db_username_local")
    db_password := os.Getenv("db_password_local")
    db_host := os.Getenv("db_host_local")
    db_name := os.Getenv("db_name_local")
    db_port := os.Getenv("db_port_local")

    dbUri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", db_username, db_password, db_host, db_port, db_name)
    conn, err := gorm.Open("mysql", dbUri)

    if err != nil {
        fmt.Print(err)
    }

    db = conn
    db.Debug().AutoMigrate(&Todo{}, &Account{})
}

func GetDB() *gorm.DB {
    return db
}