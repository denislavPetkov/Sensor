package customflags

import (
	"fmt"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Arrayflag", func() {

	testData := StringSliceFlag{"test1", "test2"}

	Describe("String()", func() {
		Context("When called", func() {
			expected := strings.Join(testData, " ")
			It("Returns a single string", func() {
				actual := testData.String()
				Expect(actual).To(Equal(expected))
			})
		})

		Context("When slice is empty", func() {
			var emptySlice StringSliceFlag
			expected := ""
			It("Returns empty string", func() {
				actual := emptySlice.String()
				fmt.Println(actual)
				Expect(actual).To(Equal(expected))
			})
		})
	})

	Describe("Set()", func() {
		Context("When called", func() {
			expected := StringSliceFlag{"test1", "test2", "test3"}
			It("Adds the string to the slice", func() {
				err := testData.Set("test3")
				Expect(err).NotTo(HaveOccurred())
				Expect(testData).To(Equal(expected))
			})
		})
	})
})
