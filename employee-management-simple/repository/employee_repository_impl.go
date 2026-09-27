package repository

import (
	"fmt"

	"example.com/employee-management/model"
)

type EmployeeRepositoryImpl struct{}

func NewEmployeeRepository() EmployeeRepository {
	return &EmployeeRepositoryImpl{}
}

func (r *EmployeeRepositoryImpl) Save(employee model.Employee) {
	fmt.Println("Hello from Repository - Save Employee")
}

func (r *EmployeeRepositoryImpl) FindByID(id int) {
	fmt.Println("Hello from Repository - Find Employee")
}

func (r *EmployeeRepositoryImpl) FindAll() {
	fmt.Println("Hello from Repository - Find All Employees")
}

func (r *EmployeeRepositoryImpl) Update(employee model.Employee) {
	fmt.Println("Hello from Repository - Update Employee")
}

func (r *EmployeeRepositoryImpl) Delete(id int) {
	fmt.Println("Hello from Repository - Delete Employee")
}
