package main

import (
	"encoding/gob"
	"log"
	"os"
)

// save saves the data in a file
func save(filename string, data []*DirEntry) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	enc := gob.NewEncoder(file)
	err = enc.Encode(data)
	if err != nil {
		log.Fatal(err)
	}
}
