package employee

type Employee struct {
	Id     int
	Name   string
	Salary float64
}

type EmployeeService interface { // interface defined
	Add(e *Employee)
	GetAll() []*Employee
}

func NewEmployeeService() EmployeeService {
	s := &employeeService{}
	s.empStore = make([]*Employee, 0)
	return s
}

type employeeService struct {
	empStore []*Employee
}

func (s *employeeService) Add(emp *Employee) {
	s.empStore = append(s.empStore, emp)
}

func (s *employeeService) GetAll() []*Employee {
	return s.empStore
}
