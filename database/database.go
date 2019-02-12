package database

import "log"

var DB *FileDatabase

func Init() {
	log.Println("Database: init")
	DB = NewFileDatabase()
}
