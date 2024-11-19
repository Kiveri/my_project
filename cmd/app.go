package main

import (
	"fmt"
	"my_project/internal/adapter/in_memory_db/employee_in_memory_db"
	"my_project/internal/domain/model"
	"my_project/internal/usecase/employee_usecase"
)

func main() {
	employeeRepo := employee_in_memory_db.NewRepo()
	employeeUseCase := employee_usecase.NewUseCase(employeeRepo)

	employee, err := employeeUseCase.AddEmployee(employee_usecase.AddEmployeeReq{
		Name:         "Denis",
		Surname:      "Popov",
		Phone:        "+79995398037",
		Email:        "denpopov.m@gmail.com",
		EmployeeRole: model.LeaderEmployeeRole,
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(employee)
}
