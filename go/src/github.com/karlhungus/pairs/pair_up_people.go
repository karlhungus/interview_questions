package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

type Person struct {
	Name  string
	Email string
}

func main() {
	csvfile, err := os.Open("test.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	rand.Seed(time.Now().Unix())
	records := []Person{}
	//Parse the file
	r := csv.NewReader(csvfile)
	for i := 0; true; i++ {
		record, err := r.Read()
		if i == 0 {
			continue
		}

		// Read each record from csv
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		records = append(records, Person{Name: record[0], Email: record[1]})
	}

	selections := make(map[string]string)
	for i := 0; i < len(records); i++ {
		assigned := records[i].Email
		var selection string
		//determine assignee
		for {
			selectionNo := rand.Intn(len(records))
			selection = records[selectionNo].Email

			// validate selection is not receiving
			var alreadyRcv bool
			for _, v := range selections {
				if v == selection {
					alreadyRcv = true
					break
				}
			}

			//Stop if you've selected another non selected person
			if selection != assigned && !alreadyRcv {
				selections[assigned] = selection
				break
			}

			//If assigned is the last one in the list, and there are no other selected people swap with the first selection
			if len(records) == i+1 && selection == assigned {
				first := selections[records[0].Email]
				tmp := selections[first]
				selections[assigned] = tmp
				selections[first] = assigned
				break
			}

		}

	}

	for k, v := range selections {
		fmt.Printf("%s\t\t\t -> \t\t%s\n", k, v)
	}

}
