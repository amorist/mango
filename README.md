# mango [mongo-go-driver](https://github.com/mongodb/mongo-go-driver)  toolkit for golang

## Usage

```go
var session *Session
session = session.New("mongodb://127.0.0.1")
err := session.Connect()
if err != nil {
    fmt.Println(err)
}
```

## session.DB

```go
session := mango.New("mongodb://127.0.0.1")
session.SetPoolLimit(10)
if err := session.Connect(); err != nil {
    fmt.Println(err)
    return
}
var result []Person
if err := session.DB("test").Collection("articles").Find(bson.M{}).All(&result); err != nil {
    fmt.Println(err)
}

for _, r := range result {
    fmt.Println(r.Title)
}
```

## SetDB

```go
session := mango.New("mongodb://127.0.0.1")
session.SetPoolLimit(10)
session.SetDB("test")
if err := session.Connect(); err != nil {
    fmt.Println(err)
    return
}
var result []Person
if err := session.Collection("articles").Find(bson.M{}).All(&result); err != nil {
    fmt.Println(err)
}
for _, r := range result {
    fmt.Println(r.Title)
}
```
