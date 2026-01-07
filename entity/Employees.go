package entity

import "gorm.io/gorm"

type Employees struct{
	gorm.Model 
	Name      string	`valid:"matches(^[A-Za-z]~not good"`
	Salary    float64  `valid:"range(15000|200000)~Salary must be between 15000 and 200000"` 
	EmployeeCode string	`valid:"matches(^[A-Z]{2}$)~EmployeeCode must be 2 uppercase English letters (A-Z) followed by '-' and 4 digits (0-9)"`
}