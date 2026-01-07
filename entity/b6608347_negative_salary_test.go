package entity

import (
	
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)
func TestSalary(t *testing.T){
	g:=NewGomegaWithT(t)
	t.Run("EmloyeeCode",func(t *testing.T){
		employees := Employees{
			Name: "mi",
			Salary: 14000,
			EmployeeCode: "AS",
		}
		ok, err:=govalidator.ValidateStruct(employees)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Salary must be between 15000 and 200000"))
	})
}