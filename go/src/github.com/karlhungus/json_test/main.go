package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Jsons struct {
	Jsons []Json `json:"jsons"`
}

type Json struct {
	foo string  `json:"foo"`
	bar float64 `json:"bar"`
}

func main() {
	jsonFile, err := os.Open("test.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		panic("failed to read json file")
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var jsons Jsons

	json.Unmarshal(byteValue, &jsons)
	fmt.Printf("jsons: %+v", jsons)
}
