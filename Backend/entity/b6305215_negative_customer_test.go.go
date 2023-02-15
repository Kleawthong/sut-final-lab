package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestReacorddata(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run("Name can not be blan", func(t *testing.T) {
		cu := Customer{
			Name:       "sdfsdfsfd",
			Email:      "asdasdasd@gmail.com",
			CustomerID: "L12345", //ผิด
		}

		ok, err := govalidator.ValidateStruct(cu)
		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("not valid"))
	})

	t.Run("Name can not be blan", func(t *testing.T) {
		cu := Customer{
			Name:       "sdfsdfsfd",
			Email:      "asdasdasd@gmail.com",
			CustomerID: "12345", //ผิด
		}

		ok, err := govalidator.ValidateStruct(cu)
		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("not valid"))
	})
	t.Run("Name can not be blan", func(t *testing.T) {
		cu := Customer{
			Name:       "sdfsdfsfd",
			Email:      "asdasdasd@gmail.com",
			CustomerID: "sdfsdfsdf", //ผิด
		}

		ok, err := govalidator.ValidateStruct(cu)
		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("not valid"))
	})

}
