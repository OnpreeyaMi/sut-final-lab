package entity

import (
	
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)
func TestEmployeeCode(t *testing.T){
	g:=NewGomegaWithT(t)
	t.Run("EmloyeeCode",func(t *testing.T){
		employees := Employees{
			Name: "dd",
			Salary: 15001,
			EmployeeCode: "asd",
		}
		ok, err:=govalidator.ValidateStruct(employees)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("EmployeeCode must be 2 uppercase English letters (A-Z) followed by '-' and 4 digits (0-9)"))
	})
}