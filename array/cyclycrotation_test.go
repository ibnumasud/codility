package array_test

import (
	"github.com/ibnumasud/codility/array"
)

var _ = Describe("Array", func() {
	Describe("CyclycRotation normal", func() {
		given := []int{3, 8, 9, 7, 6}
		expected := []int{9, 7, 6, 3, 8}
		Context("Dimasukkan array [3, 8, 9, 7, 6] harus dirotate kedepan ", func() {
			It("Dirotate jadi [6, 3, 8, 9, 7]", func() {
				Expect(array.CyclycRotation(given, 3)).To(Equal(expected))
			})
		})
	})

})
