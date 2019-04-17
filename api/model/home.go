package model

import (
	"encoding/json"
	"github.com/g8y3e/mangabeast/database"
	"log"
)

type Home struct {
	Title string `json:"title"`
}

func GetHome() (*Home, error) {
	home := Home{}

	data := database.DB.GetData("home")
	err := json.Unmarshal(data, &home)
	if err != nil {
		log.Println("Model Home: error unmarshal data:", err.Error())
		return nil, err
	}

	return &home, nil
}
