package main

import (
	"fmt"
	"my_project/internal/adapter/in_memory_db/employee_in_memory_db"
	"my_project/internal/adapter/in_memory_db/position_in_memory_db"
	"my_project/internal/domain/model"
	"my_project/internal/usecase/employee_usecase"
	"my_project/internal/usecase/position_usecase"
)

func main() {
	employeeRepo := employee_in_memory_db.NewRepoEmployees()
	employeeUseCase := employee_usecase.NewUseCaseEmployee(employeeRepo)
	positionRepo := position_in_memory_db.NewRepoPosition()
	positionUseCase := position_usecase.NewUseCasePosition(positionRepo)

	employee, err := employeeUseCase.AddEmployee(employee_usecase.AddEmployeeReq{
		NameEmployee:    "Denis",
		SurnameEmployee: "Popov",
		PhoneEmployee:   "+79995398037",
		EmailEmployee:   "denpopov.m@gmail.com",
		EmployeeRole:    model.LeaderEmployeeRole,
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(employee)

	position, err2 := positionUseCase.AddPosition(position_usecase.AddPositionReq{
		NamePosition:    "Болгарка Sturm AG90125E",
		BarcodePosition: "",
		PricePosition:   4376,
		PositionType:    model.BasicProductPositionType,
	})
	if err2 != nil {
		fmt.Println(err)
	}

	fmt.Println(position)
}
