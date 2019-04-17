package database

import (
	"io/ioutil"
	"log"
	"os"
)

type FileDatabase struct {
	data map[string][]byte
}

func NewFileDatabase() *FileDatabase {
	db := FileDatabase{}
	err := db.getAllData("./database/data/")
	if err != nil {
		panic("File Database: error init file database: " + err.Error())
	}

	return &db
}

func (db *FileDatabase) getAllData(dirname string) error {
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}

	db.data = make(map[string][]byte)

	for _, f := range files {
		log.Println(f.Name())
		jsonFile, err := os.Open("./database/data/" + f.Name())
		if err != nil {
			log.Println("File Database: error read data file:", f.Name(), "; error:", err.Error())
			return err
		}

		byteValue, _ := ioutil.ReadAll(jsonFile)
		db.data[f.Name()] = byteValue

		jsonFile.Close()
	}
	return nil
}

func (db *FileDatabase) GetData(filename string) []byte {
	byteValue, ok := db.data[filename+".json"]
	if !ok {
		return []byte{}
	}

	return byteValue
}
