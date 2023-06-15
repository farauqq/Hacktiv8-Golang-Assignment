package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

type Student struct {
	Id      string
	Name    string
	Address string
	Job     string
	Reason  string
}

var students []Student = []Student{
	{
		Id:      "1",
		Name:    "Farauq",
		Address: "Kudus",
		Job:     "Student",
		Reason:  "Karena Golang Memiliki Concurrency",
	},
	{
		Id:      "2",
		Name:    "Faresh",
		Address: "Semarang",
		Job:     "Student",
		Reason:  "Karena Golang memiliki standard library",
	},
	{
		Id:      "3",
		Name:    "Rahman",
		Address: "Solo",
		Job:     "Student",
		Reason:  "Karena Golang Memiliki Garbage Collector",
	},
	{
		Id:      "4",
		Name:    "Malik",
		Address: "Yogyakarta",
		Job:     "Student",
		Reason:  "Karena Eksekusi Cepat",
	},
}

func main() {
	var inputs = os.Args

	if !(len(inputs) >= 2) {
		log.Fatalln("Agar dapat menjalankan program menggunakan perintah go run main.go [id]")
	}

	result, err := FindStudent(inputs[1])

	if err != nil {
		log.Fatalln(err.Error())
	}
	
	fmt.Println("Data Student :")
	fmt.Printf("ID      : " + result.Id)
	fmt.Printf("\nNama    : " + result.Name)
	fmt.Printf("\nAddress : " + result.Address)
	fmt.Printf("\nJob     : " + result.Job)
	fmt.Printf("\nReason  : " + result.Reason)
}

func FindStudent(studentId string) (Student, error) {
	for _, value := range students {
		if value.Id == studentId {
			return value, nil
		}
	}

	return Student{}, errors.New("Oops... Data Student not found")
}



