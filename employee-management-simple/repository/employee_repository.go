package repository

import "example.com/employee-management/model"

type EmployeeRepository interface {
	Save(employee model.Employee)
	FindByID(id int)
	FindAll()
	Update(employee model.Employee)
	Delete(id int)
}
