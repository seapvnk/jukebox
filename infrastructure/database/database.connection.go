package db

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
    err := *new(error)
    DB, err = gorm.Open(sqlite.Open("./db.db"), &gorm.Config{})

    if err != nil {
        panic("can't create database connection")
    }
}