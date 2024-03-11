package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

type employee struct {
	Id     int     `json :"id"`
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
}

type comp_employee struct {
	Emp []employee `json:"employee"`
}

func (emp *comp_employee) UpdateSalary() {

	avg_salary := 800.0

	for ind := range emp.Emp {

		if float64(emp.Emp[ind].Salary) < avg_salary {

			emp.Emp[ind].Salary = emp.Emp[ind].Salary + (0.1)*emp.Emp[ind].Salary
		}
	}

}

func (emp *comp_employee) save(file string) {

	data, err := json.MarshalIndent(emp, "", " ")

	if err != nil {
		log.Fatal(err)
		return
	}

	err = ioutil.WriteFile(file, data, 0644)

	if err != nil {
		log.Fatal(err)
		return
	}
}

func main() {

	// opening the json file

	json_file, err := os.Open("emp.json")

	if err != nil {

		log.Fatal(err)
	}

	fmt.Println("Successfully Opened emp.json")
	defer json_file.Close() // closing file

	b, _ := io.ReadAll(json_file)

	var emp_list comp_employee // initializing

	json.Unmarshal(b, &emp_list)

	fmt.Println(emp_list)

	for _, val := range emp_list.Emp {

		fmt.Println(val)
	}

	fmt.Println("==========After Update=========")

	emp_list.UpdateSalary()

	for _, val := range emp_list.Emp {

		fmt.Println(val)
	}

	emp_list.save("emp.json")
}
