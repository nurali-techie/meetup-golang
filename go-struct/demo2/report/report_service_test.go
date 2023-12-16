package report_test

import (
	"testing"

	"github.com/nurali-techie/meetup-golang/go-struct/demo2/employee"
	"github.com/nurali-techie/meetup-golang/go-struct/demo2/report"
)

func TestTotalSalary(t *testing.T) {
	employees := []*employee.Employee{
		{
			Salary: 50000,
		},
		{
			Salary: 20000,
		},
	}
	empServiceMock := employee.NewEmployeeServiceMock(employees) // use of mock instead actual is possible due to interface

	wantTotalSalary := 90000.0
	gotTotalSalary := report.TotalSalary(empServiceMock)

	if gotTotalSalary != wantTotalSalary {
		t.Errorf("want=%f, got=%f", wantTotalSalary, gotTotalSalary)
	}
}
