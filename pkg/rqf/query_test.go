package rqf_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/stevenceuppens/go-rest-query-filter/pkg/rqf"
)

var _ = Describe("Query", func() {

	Describe("NewQuery", func() {

		filter := NewFilter()

		It("Fields should be empty by default", func() {
			Expect(len(filter.Fields)).To(Equal(0))
		})

		It("Where should be empty by default", func() {
			Expect(len(filter.Where)).To(Equal(0))
		})

		It("Order should be empty by default", func() {
			Expect(len(filter.Order)).To(Equal(0))
		})

		It("Offset should be empty by default", func() {
			Expect(filter.Offset).To(Equal(0))
		})

		It("Limit should be empty by default", func() {
			Expect(filter.Limit).To(Equal(0))
		})
	})
})
