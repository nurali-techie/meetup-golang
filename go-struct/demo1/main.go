package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type AuditLog struct {
	CreatedDate  time.Time `json:"created_date"`
	ModifiedDate time.Time `json:"modified_date"`
}

type Employee struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
	AuditLog
}

func (e *Employee) Increment(perc float64) {
	e.Salary = e.Salary + (e.Salary * perc / 100)
}

func (e Employee) ToJson() string {
	content, _ := json.Marshal(e)
	return string(content)
}

func (e *Employee) ToString() string {
	return fmt.Sprintf("id=%d, name=%s, salary=%f, created=%v, modified=%v",
		e.Id, e.Name, e.Salary, e.CreatedDate, e.ModifiedDate)
}

func Sanitize(emp *Employee) {
	if emp.CreatedDate.IsZero() {
		emp.CreatedDate = time.Now()
	}
	if emp.ModifiedDate.IsZero() {
		emp.ModifiedDate = time.Now()
	}
}

func main() {
	fmt.Println("demo1") // demo1 shows different aspect of go struct

	e1 := &Employee{
		Id:     1,
		Name:   "amit",
		Salary: 100000,
	}

	fmt.Printf("before, e1=%v\n", e1.ToString()) // check CreatedDate, not set
	Sanitize(e1)
	fmt.Printf("after date set, e1=%v\n", e1.ToString()) // check CreatedDate, set to today
	e1.Increment(5.0)
	fmt.Printf("after incr, e1=%v\n", e1.ToJson()) // check salary increased and json output
}
