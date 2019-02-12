package main

import (
	"github.com/g8y3e/mangabeast/api"
	"github.com/g8y3e/mangabeast/database"
)

func main() {
	database.Init()

	api.Init()
}
