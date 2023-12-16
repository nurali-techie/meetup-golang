package employee

type employeeServiceMock struct {
	empStore []*Employee
}

func NewEmployeeServiceMock(employees []*Employee) EmployeeService {
	mock := &employeeServiceMock{}
	mock.empStore = employees
	return mock
}

func (mock *employeeServiceMock) Add(*Employee) {
}

func (mock *employeeServiceMock) GetAll() []*Employee {
	return mock.empStore
}
