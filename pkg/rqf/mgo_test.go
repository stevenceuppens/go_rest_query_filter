package rqf_test

import (
	"encoding/json"
	"fmt"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	. "github.com/stevenceuppens/go-rest-query-filter/pkg/rqf"
)

func connectMgo() *mgo.Session {
	session, err := mgo.Dial("mongo/rqf")

	if err != nil {
		panic(err)
	}

	return session
}

type Meta struct {
	Active  bool       `bson:"active" json:"active,omitempty"`
	Created *time.Time `bson:"created" json:"created,omitempty"`
}
type Book struct {
	ID   bson.ObjectId `bson:"_id" json:"id,omitempty"`
	Name string        `bson:"name" json:"name,omitempty"`
	ISBN string        `bson:"isbn" json:"isbn,omitempty"`
	Meta *Meta         `bson:"meta" json:"meta,omitempty"`
}
type BookList []*Book

// Now ...
func Now() *time.Time {
	time := time.Now()
	return &time
}

var _ = Describe("Mgo", func() {

	BeforeSuite(func() {
		mgoSession := connectMgo().Clone()
		defer mgoSession.Close()

		err := mgoSession.DB("").C("fields_test").Insert(
			&Book{
				ID:   bson.NewObjectId(),
				Name: "Book1",
				ISBN: "C_ISBN",
				Meta: &Meta{
					Active:  true,
					Created: Now(),
				},
			},
			&Book{
				ID:   bson.NewObjectId(),
				Name: "Book2",
				ISBN: "A_ISBN",
				Meta: &Meta{
					Active:  true,
					Created: Now(),
				},
			},
			&Book{
				ID:   bson.NewObjectId(),
				Name: "Book3",
				ISBN: "B_ISBN",
				Meta: &Meta{
					Active:  true,
					Created: Now(),
				},
			})

		if err != nil {
			panic(err)
		}
	})

	Describe("MgoFields", func() {

		It("Should convert fields (no filter)", func() {

			mgoSession := connectMgo().Clone()
			defer mgoSession.Close()

			var data BookList

			q := mgoSession.DB("").C("fields_test").Find(nil)
			q.All(&data)

			Expect(len(data)).To(Equal(3))
			for _, item := range data {
				Expect(item.ID).ToNot(BeEmpty())
				Expect(item.Name).ToNot(BeEmpty())
				Expect(item.ISBN).ToNot(BeEmpty())
			}

			fmt.Fprintln(GinkgoWriter, "#####################################################")
			fmt.Fprintln(GinkgoWriter, "No fields filter")
			json, _ := json.MarshalIndent(data, "", "  ")
			fmt.Fprintln(GinkgoWriter, string(json))
		})

		It("Should convert fields (basic filter)", func() {

			filter, err := ParseFilter("{\"fields\":{\"ID\":true,\"name\":true}}")
			Expect(err).To(BeNil())
			mgoSession := connectMgo().Clone()
			defer mgoSession.Close()

			var data BookList

			q := mgoSession.DB("").C("fields_test").Find(nil)
			filter.MgoFields(q)
			q.All(&data)

			Expect(len(data)).To(Equal(3))
			for _, item := range data {
				Expect(item.ID).ToNot(BeEmpty())
				Expect(item.Name).ToNot(BeEmpty())
				Expect(item.ISBN).To(BeEmpty())
			}

			fmt.Fprintln(GinkgoWriter, "#####################################################")
			fmt.Fprintln(GinkgoWriter, "Fields filter: {\"fields\":{\"ID\":true,\"name\":true}}")
			json, _ := json.MarshalIndent(data, "", "  ")
			fmt.Fprintln(GinkgoWriter, string(json))
		})

		It("Should convert fields (sub document filter)", func() {

			filter, err := ParseFilter("{\"fields\":{\"ID\":true,\"meta.active\":true}}")
			Expect(err).To(BeNil())
			mgoSession := connectMgo().Clone()
			defer mgoSession.Close()

			var data BookList

			q := mgoSession.DB("").C("fields_test").Find(nil)
			filter.MgoFields(q)
			q.All(&data)

			Expect(len(data)).To(Equal(3))
			for _, item := range data {
				Expect(item.ID).ToNot(BeEmpty())
				Expect(item.Name).To(BeEmpty())
				Expect(item.ISBN).To(BeEmpty())
				Expect(item.Meta).ToNot(BeNil())
			}

			fmt.Fprintln(GinkgoWriter, "#####################################################")
			fmt.Fprintln(GinkgoWriter, "Fields filter: {\"fields\":{\"ID\":true,\"meta.active\":true}}")
			json, _ := json.MarshalIndent(data, "", "  ")
			fmt.Fprintln(GinkgoWriter, string(json))
		})
	})

	Describe("MgoOrder", func() {

		It("Should order results (order isbn ASC)", func() {
			filter, err := ParseFilter("{\"order\":[\"isbn ASC\"]}")
			Expect(err).To(BeNil())
			mgoSession := connectMgo().Clone()
			defer mgoSession.Close()

			var data BookList

			q := mgoSession.DB("").C("fields_test").Find(nil)
			filter.MgoOrder(q)
			q.All(&data)

			Expect(data[0].ISBN).To(Equal("A_ISBN"))
			Expect(data[1].ISBN).To(Equal("B_ISBN"))
			Expect(data[2].ISBN).To(Equal("C_ISBN"))

			fmt.Fprintln(GinkgoWriter, "#####################################################")
			fmt.Fprintln(GinkgoWriter, "Order filter: {\"order\":[\"isbn ASC\"]}")
			json, _ := json.MarshalIndent(data, "", "  ")
			fmt.Fprintln(GinkgoWriter, string(json))
		})

		It("Should order results (order isbn)", func() {
			filter, err := ParseFilter("{\"order\":[\"isbn\"]}")
			Expect(err).To(BeNil())
			mgoSession := connectMgo().Clone()
			defer mgoSession.Close()

			var data BookList

			q := mgoSession.DB("").C("fields_test").Find(nil)
			filter.MgoOrder(q)
			q.All(&data)

			Expect(data[0].ISBN).To(Equal("A_ISBN"))
			Expect(data[1].ISBN).To(Equal("B_ISBN"))
			Expect(data[2].ISBN).To(Equal("C_ISBN"))

			fmt.Fprintln(GinkgoWriter, "#####################################################")
			fmt.Fprintln(GinkgoWriter, "Order filter: {\"order\":[\"isbn\"]}")
			json, _ := json.MarshalIndent(data, "", "  ")
			fmt.Fprintln(GinkgoWriter, string(json))
		})

		It("Should order results (order isbn DESC)", func() {
			filter, err := ParseFilter("{\"order\":[\"isbn DESC\"]}")
			Expect(err).To(BeNil())
			mgoSession := connectMgo().Clone()
			defer mgoSession.Close()

			var data BookList

			q := mgoSession.DB("").C("fields_test").Find(nil)
			filter.MgoOrder(q)
			q.All(&data)

			Expect(data[0].ISBN).To(Equal("C_ISBN"))
			Expect(data[1].ISBN).To(Equal("B_ISBN"))
			Expect(data[2].ISBN).To(Equal("A_ISBN"))

			fmt.Fprintln(GinkgoWriter, "#####################################################")
			fmt.Fprintln(GinkgoWriter, "Order filter: {\"order\":[\"isbn DESC\"]}")
			json, _ := json.MarshalIndent(data, "", "  ")
			fmt.Fprintln(GinkgoWriter, string(json))
		})
	})

	Describe("MgoLimit", func() {

		It("Should limit results (limit 2)", func() {
			filter, err := ParseFilter("{\"limit\":2}")
			Expect(err).To(BeNil())
			mgoSession := connectMgo().Clone()
			defer mgoSession.Close()

			var data BookList

			q := mgoSession.DB("").C("fields_test").Find(nil)
			filter.MgoLimit(q)
			q.All(&data)

			Expect(len(data)).To(Equal(2))
			Expect(data[0].Name).To(Equal("Book1"))
			Expect(data[1].Name).To(Equal("Book2"))
		})
	})

	Describe("MgoOffset", func() {

		It("Should skip results (offset 2)", func() {
			filter, err := ParseFilter("{\"offset\":2}")
			Expect(err).To(BeNil())
			mgoSession := connectMgo().Clone()
			defer mgoSession.Close()

			var data BookList

			q := mgoSession.DB("").C("fields_test").Find(nil)
			filter.MgoOffset(q)
			q.All(&data)

			Expect(len(data)).To(Equal(1))
			Expect(data[0].Name).To(Equal("Book3"))
		})
	})
})
