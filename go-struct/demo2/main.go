package main

import (
	"fmt"

	"github.com/nurali-techie/meetup-golang/go-struct/demo2/employee"
	"github.com/nurali-techie/meetup-golang/go-struct/demo2/report"
)

func main() {
	fmt.Println("demo2") // demo2 shows usage of go interface
	empService := employee.NewEmployeeService()

	e1 := &employee.Employee{
		Id:     1,
		Name:   "amit",
		Salary: 100000.0,
	}
	empService.Add(e1)

	e2 := &employee.Employee{
		Id:     2,
		Name:   "hemal",
		Salary: 200000.0,
	}
	empService.Add(e2)

	totalSalary := report.TotalSalary(empService)
	fmt.Println(totalSalary)
}
