package rqf_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/stevenceuppens/go-rest-query-filter/pkg/rqf"
)

var _ = Describe("Query", func() {

	Describe("NewQuery", func() {

		q := NewQuery()

		It("Fields should be empty by default", func() {
			Expect(len(q.Fields)).To(Equal(0))
		})

		It("Where should be empty by default", func() {
			Expect(len(q.Where)).To(Equal(0))
		})

		It("Order should be empty by default", func() {
			Expect(len(q.Order)).To(Equal(0))
		})

		It("Offset should be empty by default", func() {
			Expect(q.Offset).To(Equal(0))
		})

		It("Limit should be empty by default", func() {
			Expect(q.Limit).To(Equal(-1))
		})
	})
})
