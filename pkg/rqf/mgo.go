package rqf

import (
	"strings"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MgoFilter will add all filters to the Mongo Query
func (q *Filter) MgoFilter(mq *mgo.Query) *mgo.Query {
	tq := q.MgoFields(mq)
	tq = q.MgoWhere(tq)
	tq = q.MgoOrder(tq)
	tq = q.MgoLimit(tq)
	tq = q.MgoOffset(tq)
	return tq
}

// MgoFields will add field filter to the Mongo Query
func (q *Filter) MgoFields(mq *mgo.Query) *mgo.Query {
	mgoSelect := bson.M{}
	for field := range q.Fields {
		mgoSelect[field] = 1
	}
	return mq.Select(mgoSelect)
}

// MgoWhere will add where filter to the Mongo Query
func (q *Filter) MgoWhere(mq *mgo.Query) *mgo.Query {
	// TODO
	return mq
}

// MgoOrder will add order filter to the Mongo Query
func (q *Filter) MgoOrder(mq *mgo.Query) *mgo.Query {
	sort := []string{}
	for _, value := range q.Order {
		strippedValue := strings.Split(value, " ")
		mgoValue := strippedValue[0]
		if len(strippedValue) > 1 && strippedValue[1] == "DESC" {
			mgoValue = "-" + strippedValue[0]
		}
		sort = append(sort, mgoValue)
	}
	return mq.Sort(sort...)
}

// MgoLimit will add limit field filter to the Mongo Query
func (q *Filter) MgoLimit(mq *mgo.Query) *mgo.Query {
	return mq.Limit(q.Limit)
}

// MgoOffset will add offset filter to the Mongo Query
func (q *Filter) MgoOffset(mq *mgo.Query) *mgo.Query {
	return mq.Skip(q.Offset)
}
