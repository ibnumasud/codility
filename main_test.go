package main_test

var _ = Describe("hello world", func() {
	Describe("test dalam describe", func() {
		Context("keluar hello world", func() {
			It("keluar hello world", func() {
				// Expect().To(Equal("hello go"))
			})
		})
	})
})
