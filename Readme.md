# Golang REST Query Filter (RQF)

This library can parse a json formatted filter (JSON format) to allow inspection, and conversion into go mgo query (Go Mongo driver)

---

## Field selection

Allows to select which fields should be returned

JSON object to filter fields
```json
{"fields":{"ID":true,"name":true}}
```

URL encoded
```
http://localhost:3000/api/books?filter=%7B%22fields%22%3A%7B%22ID%22%3Atrue%2C%22name%22%3Atrue%7D%7D
```

Code to parse and convert JSON filter into Mgo
```golang
// parse the json
query, err := rqf.ParseFilter( url_string|query_string|filter_value )

// basic query
q := mgoSession.DB("myDB").C("myCollection").Find(nil)

// auto inject field selection in mgo query
query.MgoFields(q)

// execute
q.All(&data)
```

Result
```json
[
  {
    "ID": "5a9ea44d7cb20300b70daaae",
    "Name": "Book1"
  },
  {
    "ID": "5a9ea44d7cb20300b70daaaf",
    "Name": "Book2"
  },
  {
    "ID": "5a9ea44d7cb20300b70daab0",
    "Name": "Book3"
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
query, err := rqf.ParseFilter( url_string|query_string|filter_value )

// basic query
q := mgoSession.DB("myDB").C("myCollection").Find(nil)

// auto inject field selection in mgo query
query.MgoOrder(q)

// execute
q.All(&data)
```

Result
```json
[
  {
    "ID": "5a9ea44d7cb20300b70daaaf",
    "Name": "Book2",
    "ISBN": "A_ISBN",
    "Meta": {
      "Active": true,
      "Created": "2018-03-06T14:23:09.607Z"
    }
  },
  {
    "ID": "5a9ea44d7cb20300b70daab0",
    "Name": "Book3",
    "ISBN": "B_ISBN",
    "Meta": {
      "Active": true,
      "Created": "2018-03-06T14:23:09.607Z"
    }
  },
  {
    "ID": "5a9ea44d7cb20300b70daaae",
    "Name": "Book1",
    "ISBN": "C_ISBN",
    "Meta": {
      "Active": true,
      "Created": "2018-03-06T14:23:09.607Z"
    }
  }
]

```