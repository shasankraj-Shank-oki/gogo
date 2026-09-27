package service

import (
	"fmt"

	"example.com/employee-management/model"
	"example.com/employee-management/repository"
)

type EmployeeServiceImpl struct {
	repository repository.EmployeeRepository
}

func NewEmployeeService(repository repository.EmployeeRepository) EmployeeService {
	return &EmployeeServiceImpl{repository: repository}
}

func (s *EmployeeServiceImpl) AddEmployee(employee model.Employee) {
	fmt.Println("Hello from Service - Add Employee")
	s.repository.Save(employee)
}

func (s *EmployeeServiceImpl) GetEmployee(id int) {
	fmt.Println("Hello from Service - Get Employee")
	s.repository.FindByID(id)
}

func (s *EmployeeServiceImpl) GetAllEmployees() {
	fmt.Println("Hello from Service - Get All Employees")
	s.repository.FindAll()
}

func (s *EmployeeServiceImpl) UpdateEmployee(employee model.Employee) {
	fmt.Println("Hello from Service - Update Employee")
	s.repository.Update(employee)
}

func (s *EmployeeServiceImpl) DeleteEmployee(id int) {
	fmt.Println("Hello from Service - Delete Employee")
	s.repository.Delete(id)
}
