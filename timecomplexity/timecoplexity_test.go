package timecomplexity_test

import (
	"github.com/ibnumasud/codility/timecomplexity"
)

var _ = Describe("Timecoplexity", func() {
	Describe("Frogjump", func() {
		x := 10
		y := 80
		d := 30

		expected := 3
		Context("jump frog normal", func() {
			It("return 3", func() {
				Expect(timecomplexity.FrogJump(x, y, d)).To(Equal(expected))
			})
		})
	})
	Describe("oddoccurancyarray", func() {
		given := []int{2, 3, 1, 5}
		expected := 4
		Context("diberi array yg missing", func() {
			It("return missing integer", func() {
				Expect(timecomplexity.PermMissElement(given)).To(Equal(expected))
			})
		})
	})
})
