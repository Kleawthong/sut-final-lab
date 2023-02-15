package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestNegativeName(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run("name canonot be blank", func(t *testing.T) {
		cu := Customer{
			Name:       "", //ผิด
			Email:      "asdasdasd@gmail.com",
			CustomerID: "L1234567",
		}

		ok, err := govalidator.ValidateStruct(cu)
		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("name canonot be blank"))
	})
}
