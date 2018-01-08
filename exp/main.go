package main

import (
	"html/template"
	"os"
)

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := struct {
		Name        string
		Time        int
		Friends     []string
		FriendsJobs map[string]string
	}{"Scott Hedrick", 12, []string{"Jim", "Bob", "Greg"}, map[string]string{
		"Jim":  "Cook",
		"Bob":  "Doctor",
		"Greg": "Nurse",
	}}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
