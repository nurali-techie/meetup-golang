package report

import "github.com/nurali-techie/meetup-golang/go-struct/demo2/employee"

func TotalSalary(employeeService employee.EmployeeService) float64 {
	employees := employeeService.GetAll()
	total := 0.0
	for _, emp := range employees {
		total += emp.Salary
	}
	return total
}
