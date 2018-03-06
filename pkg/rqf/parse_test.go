package rqf

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parse", func() {

	Describe("ParseFilter", func() {

		It("Should parse fields from json", func() {
			query, err := ParseFilter("{\"fields\":{\"username\":true,\"email\":true}}")
			Expect(err).To(BeNil())
			Expect(len(query.Fields)).To(Equal(2))
			Expect(query.Fields["username"]).To(BeTrue())
			Expect(query.Fields["email"]).To(BeTrue())
			Expect(query.Fields["fail"]).To(BeFalse())
		})

		It("Should parse fields from encoded string", func() {
			query, err := ParseFilter("%7B%22fields%22%3A%7B%22username%22%3Atrue%2C%22email%22%3Atrue%7D%7D")
			Expect(err).To(BeNil())
			Expect(len(query.Fields)).To(Equal(2))
			Expect(query.Fields["username"]).To(BeTrue())
			Expect(query.Fields["email"]).To(BeTrue())
			Expect(query.Fields["fail"]).To(BeFalse())
		})

		It("Should parse where", func() {
			query, err := ParseFilter("{\"where\":{\"or\":[{\"id\":1},{\"name\":\"steven\"},{\"id\":20},{\"id\":21}]}}")
			Expect(err).To(BeNil())
			Expect(len(query.Where)).To(BeNumerically(">", 0))
			Expect(query.Where["or"]).ToNot(BeNil())
		})

		It("Should parse order", func() {
			query, err := ParseFilter("{\"order\":[\"username ASC\",\"email DESC\"]}")
			Expect(err).To(BeNil())
			Expect(len(query.Order)).To(Equal(2))
			Expect(query.Order).To(ContainElement("username ASC"))
			Expect(query.Order).To(ContainElement("email DESC"))
			Expect(query.Order).ToNot(ContainElement("fail ASC"))
		})

		It("Should parse offset", func() {
			query, err := ParseFilter("{\"offset\":100}")
			Expect(err).To(BeNil())
			Expect(query.Offset).To(Equal(100))
		})

		It("Should parse limit", func() {
			query, err := ParseFilter("{\"limit\":20}")
			Expect(err).To(BeNil())
			Expect(query.Limit).To(Equal(20))
		})
	})

	Describe("normalizeFilter", func() {

		It("Should strip filter from string", func() {

			data := make(map[string]string)
			data["{\"where\":{\"username\":\"john\",\"email\":\"callback@mydomain.com\"}}"] = "{\"where\":{\"username\":\"john\",\"email\":\"callback@mydomain.com\"}}"
			data["%7B%22where%22%3A%7B%22username%22%3A%22john%22%2C%22email%22%3A%22callback%40mydomain.com%22%7D%7D"] = "{\"where\":{\"username\":\"john\",\"email\":\"callback@mydomain.com\"}}"
			data["?filter={\"where\":{\"username\":\"john\",\"email\":\"callback@mydomain.com\"}}"] = "{\"where\":{\"username\":\"john\",\"email\":\"callback@mydomain.com\"}}"
			data["?filter=%7B%22where%22%3A%7B%22username%22%3A%22john%22%2C%22email%22%3A%22callback%40mydomain.com%22%7D%7D"] = "{\"where\":{\"username\":\"john\",\"email\":\"callback@mydomain.com\"}}"
			data["/api/users?filter={\"where\":{\"username\":\"john\",\"email\":\"callback@mydomain.com\"}}"] = "{\"where\":{\"username\":\"john\",\"email\":\"callback@mydomain.com\"}}"
			data["/api/users?filter=%7B%22where%22%3A%7B%22username%22%3A%22john%22%2C%22email%22%3A%22callback%40mydomain.com%22%7D%7D"] = "{\"where\":{\"username\":\"john\",\"email\":\"callback@mydomain.com\"}}"
			data["http://localhost:3000/api/users?filter={\"where\":{\"username\":\"john\",\"email\":\"callback@mydomain.com\"}}"] = "{\"where\":{\"username\":\"john\",\"email\":\"callback@mydomain.com\"}}"
			data["http://localhost:3000/api/users?filter=%7B%22where%22%3A%7B%22username%22%3A%22john%22%2C%22email%22%3A%22callback%40mydomain.com%22%7D%7D"] = "{\"where\":{\"username\":\"john\",\"email\":\"callback@mydomain.com\"}}"

			for key, value := range data {
				nf, err := normalizeFilter(key)
				Expect(err).To(BeNil())
				Expect(nf).To(Equal(value))
			}
		})
	})
})
