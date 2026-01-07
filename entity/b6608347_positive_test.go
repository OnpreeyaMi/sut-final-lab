package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)
func TestPositive(t *testing.T){
	g :=NewGomegaWithT(t)
	t.Run("Positive case",func(t *testing.T){
		employees := Employees{
			Name: "mivfds",
			Salary: 15001,
			EmployeeCode: "AS",
		}
	ok, err:=govalidator.ValidateStruct(employees)
	g.Expect(ok).To(BeTrue())
	g.Expect(err).To(BeNil())
	})
	
}