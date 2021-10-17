package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

type Student struct {
	Name           string `json:"Student Name"`
	Email          string `json:"Student Email"`
	Institution    string `json:"Institution"`
	Enrolment_date string `json:"Enrolment Date & Time"`
	Status         string `json:"Google Cloud Skills Boost (previously Qwiklabs) Profile URL"`
	T1             string `json:"# of Skill Badges Completed in Track 1"`
	T2             string `json:"# of Skill Badges Completed in Track 2"`
}

func main() {
	csv_file, err := os.Open("sample.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csv_file.Close()

	r := csv.NewReader(csv_file)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var st Student
	var Students []Student

	for _, rec := range records {
		st.Name = rec[0]
		st.Email = rec[1]
		st.Institution = rec[2]
		st.Enrolment_date = rec[3]
		st.Status = rec[4]
		st.T1 = rec[5]
		st.T2 = rec[6]
		Students = append(Students, st)
	}
	// Convert to JSON
	json_data, err := json.Marshal(Students)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//print json data
	// fmt.Println(string(json_data))

	//create json file
	json_file, err := os.Create("sample.json")
	if err != nil {
		fmt.Println(err)
	}
	defer json_file.Close()

	json_file.Write(json_data)
	json_file.Close()
}
