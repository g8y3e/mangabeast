package model

import (
	"encoding/json"
	"github.com/g8y3e/mangabeast/database"
	"log"
)

type HeaderInfo struct {
}

type FooterInfo struct {
}

type Layout struct {
	Header HeaderInfo `json:"header"`
	Footer FooterInfo `json:"footer"`
}

func GetLayout() (*Layout, error) {
	layout := Layout{}

	data := database.DB.GetData("layout")
	err := json.Unmarshal(data, &layout)
	if err != nil {
		log.Println("Model Layout: error unmarshal data:", err.Error())
		return nil, err
	}

	return &layout, nil
}
