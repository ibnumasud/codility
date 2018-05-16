package add_test

import (
	"github.com/ibnumasud/codility/add"
)

var _ = Describe("Add", func() {
	Describe("test dalam nambah ", func() {
		Context("ingeter positif", func() {
			It("bisa nambah integer posifit", func() {
				Expect(add.LikeThis(1, 2)).To(Equal(3))
			})
		})
	})
})
