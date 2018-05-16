package looping_test

import (
	"github.com/ibnumasud/codility/looping"
)

var _ = Describe("Binarygap", func() {
	Describe("mencari jumlah loongest binary gap", func() {
		Context("dimasukan angka 1041", func() {
			It("keluar gap dair binary 1041", func() {
				Expect(looping.BinaryGap(1041)).To(Equal(5))
			})
		})
		Context("dimasukan angka yg ga ada gap", func() {
			It("keluar 0", func() {
				Expect(looping.BinaryGap(1)).To(Equal(0))
			})
		})
		Context("binary gap ada di ahir", func() {
			It("keluar binary gap", func() {
				Expect(looping.BinaryGap(6)).To(Equal(1))
			})
		})
	})
})
