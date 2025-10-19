package main

import (
	"html/template"
	"log"
	"os"
)

type WeekDay struct {
	Num   string
	Title string
}
type ForHTML struct {
	WeekDays     []WeekDay
	WorkingHours []string
}

func main() {
	test := ForHTML{
		WeekDays: []WeekDay{
			{
				Num:   "1",
				Title: "Mn",
			},
			{
				Num:   "2",
				Title: "Tu",
			},
			{
				Num:   "3",
				Title: "Wd",
			},
		},
		WorkingHours: []string{"1", "2", "3"},
	}
	tmpl, err := template.ParseFiles("./internal/templates/schedule1.html")
	if err != nil {
		panic(err)
	}

	outputFile, err := os.Create("output.html")
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}
	defer outputFile.Close()

	if err := tmpl.Execute(outputFile, test); err != nil {
		panic(err)
	}

}
