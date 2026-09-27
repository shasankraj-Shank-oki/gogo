package service

import "example.com/employee-management/model"

type EmployeeService interface {
	AddEmployee(employee model.Employee)
	GetEmployee(id int)
	GetAllEmployees()
	UpdateEmployee(employee model.Employee)
	DeleteEmployee(id int)
}
