# Golang REST Query Filter (RQF)

This library can parse a json formatted filter (JSON format) to allow inspection, and conversion into go mgo query (Go Mongo driver)

---

## Field selection

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