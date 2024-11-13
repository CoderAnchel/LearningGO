package main

import (
	"encoding/csv"
	"fmt"
	"github.com/gocarina/gocsv"
	"os"
)

type Person struct {
	Name   string `csv:"name"`
	Age    int    `csv:"age"`
	Gender string `csv:"gender"`
}

func main() {
	// Creating files using go and checking err
	file, err := os.Create("data.csv")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	//creating a writer for the file
	writer := csv.NewWriter(file)
	defer writer.Flush()
	//creating data to put in
	headers := []string{"Name", "Surname", "age"}
	data := [][]string{
		{"Jhon", "Doe", "14"},
		{"Jhon", "Doe2", "18"},
	}

	writer.Write(headers)

	for _, row := range data {
		writer.Write(row)
	}

	// Reading csv files and checking err
	file2, err := os.Open("data2.csv")

	if err != nil {
		panic(err)
	}

	defer file2.Close()

	reader := csv.NewReader(file2)
	data2, _ := reader.Read()

	for _, row := range data2 {
		for _, column := range row {
			fmt.Printf("%s", column)
		}
	}

	//you can add a row by itself
	data3 := []string{"pedro", "gonzalez", "28"}
	writer.Write(data3)

	//CREATING FILES USING GOCSV
	file3, err := os.Create("data3.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	people := []*Person{
		{"Alice", 25, "Female"},
		{"Bob", 30, "Male"},
		{"Charlie", 35, "Male"},
	}

	if err := gocsv.MarshalFile(&people, file3); err != nil {
		panic(err)
	}

	//READING FILES USING GOCSV
	data4, err := os.Open("data3.csv")
	if err != nil {
		panic(err)
	}

	defer data4.Close()

	var personHolder []Person
	if err := gocsv.UnmarshalFile(data4, &personHolder); err != nil {
		panic(err)
	}

	for _, person := range personHolder {
		fmt.Printf("Name: %s, Gender: %s, Age: %d", person.Name, person.Gender, person.Age)
	}

	//APPENDING DATA
	data5, err := os.OpenFile("data3.csv", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer data5.Close()

	writer4 := csv.NewWriter(data5)
	defer writer4.Flush()

	person := []string{"Samuel", "7", "Male"}

	err = writer4.Write(person)
	if err != nil {
		panic(err)
	}
}
