package countingelement_test

import (
	"github.com/ibnumasud/codility/countingelement"
)

var _ = Describe("Countingelement", func() {
	Describe("Tape equlibirium", func() {
		given := []int{1, 3, 1, 4, 2, 3, 5, 4}
		expected := 6
		Context("diberi array bagisesuai jumlah yg paling kecil", func() {
			It("return 1", func() {
				Expect(expected).To(Equal(countingelement.FrogRiverOne(5, given)))
			})
		})
	})
	Describe("Missing Integer", func() {
		Context("Normal case", func() {
			given := []int{1, 3, 6, 4, 1, 2}
			expected := 5
			It("int yang hilang", func() {
				Expect(expected).To(Equal(countingelement.MissingInteger(given)))
			})
		})
		Context("tidak ada yg hilang", func() {
			given := []int{1, 2, 3}
			expected := 4
			It("int next +1", func() {
				Expect(expected).To(Equal(countingelement.MissingInteger(given)))
			})
		})
		Context("Negative number", func() {
			given := []int{-1, -3}
			expected := 1
			It("int 1", func() {
				Expect(expected).To(Equal(countingelement.MissingInteger(given)))
			})
		})
	})

})
