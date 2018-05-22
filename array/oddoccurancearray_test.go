package array_test

import "github.com/ibnumasud/codility/array"

var _ = Describe("Oddoccurancearray", func() {
	Describe("CyclycRotation normal", func() {
		given := []int{9, 3, 9, 3, 9, 7, 9}
		expected := 7
		Context("Dimasukkan array [3, 8, 9, 7, 6] harus dirotate kedepan ", func() {
			It("Dirotate jadi [6, 3, 8, 9, 7]", func() {
				Expect(array.OddOccuranceArray(given)).To(Equal(expected))
			})
		})
	})

})
