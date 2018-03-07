# Golang REST Query Filter (RQF)

This library can parse a json formatted filter (JSON format) to allow inspection, and conversion into go mgo query (Go Mongo driver)

## Install
```
$ go get github.com/stevenceuppens/go_rest_query_filter
```

## Import
```golang
import "github.com/stevenceuppens/go_rest_query_filter/pkg/rqf"
```

## Use
main function to use is
```golang
rqf.ParseFilter( url_string|query_string|filter_value )
```

This function accepts the json object (both encoded as decoded) in following formats
```
- {json filter}
- ?filter={json filter}
- /api/v1/books?filter={json filter}
- http://xxx:3000/api/v1/books?filter={json filter}
```
---

## Field selection

Allows to select which fields should be returned

JSON object to filter fields
```json
{"fields":{"id":true,"name":true}}
```

URL encoded
```
http://localhost:3000/api/books?filter=%7B%22fields%22%3A%7B%22id%22%3Atrue%2C%22name%22%3Atrue%7D%7D
```

Code to parse and convert JSON filter into Mgo
```golang
// parse the json
filter, err := rqf.ParseFilter( url_string|query_string|filter_value )

// basic query
query := mgoSession.DB("myDB").C("myCollection").Find(nil)

// auto inject field selection in mgo query
filter.MgoFields(query)

// execute
query.All(&data)
```

Result
```json
[
  {
    "id": "5a9ea44d7cb20300b70daaae",
    "name": "Book1"
  },
  {
    "id": "5a9ea44d7cb20300b70daaaf",
    "name": "Book2"
  },
  {
    "id": "5a9ea44d7cb20300b70daab0",
    "name": "Book3"
  }
]
```

## Order results

Allows to order the results ASC or DESC.

To add an order:
- add the key name to the order array
- add the sort type ASC|DESC (ASC by default)

JSON object to filter fields
```json
{"order":["isbn ASC"]}
```
Equivalent of
```json
{"order":["isbn"]}
```

URL encoded
```
http://localhost:3000/api/books?filter=%7B%22order%22%3A%5B%22isbn%20ASC%22%5D%7D
```

Code to parse and convert JSON filter into Mgo
```golang
// parse the json
filter, err := rqf.ParseFilter( url_string|query_string|filter_value )

// basic query
query := mgoSession.DB("myDB").C("myCollection").Find(nil)

// auto inject ordering in mgo query
filter.MgoOrder(query)

// execute
query.All(&data)
```

Result
```json
[
  {
    "id": "5a9ea44d7cb20300b70daaaf",
    "name": "Book2",
    "isbn": "A_ISBN",
    "meta": {
      "active": true,
      "created": "2018-03-06T14:23:09.607Z"
    }
  },
  {
    "id": "5a9ea44d7cb20300b70daab0",
    "name": "Book3",
    "isbn": "B_ISBN",
    "meta": {
      "active": true,
      "created": "2018-03-06T14:23:09.607Z"
    }
  },
  {
    "id": "5a9ea44d7cb20300b70daaae",
    "name": "Book1",
    "isbn": "C_ISBN",
    "meta": {
      "active": true,
      "created": "2018-03-06T14:23:09.607Z"
    }
  }
]
```

## Page results

Allows to page the results.

JSON object to add paging
```json
{"limit":20,"offset":40}
```

URL encoded
```
http://localhost:3000/api/books?filter=%7B%22limit%22%3A20%2C%22offset%22%3A40%7D
```

Code to parse and convert JSON filter into Mgo
```golang
// parse the json
filter, err := rqf.ParseFilter( url_string|query_string|filter_value )

// basic query
query := mgoSession.DB("myDB").C("myCollection").Find(nil)

// auto inject paging in mgo query
filter.MgoLimit(query)
filter.MgoOffset(query)

// execute
query.All(&data)
```
---