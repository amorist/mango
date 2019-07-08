# mango: like mgo api but use [mongo-go-driver](https://github.com/mongodb/mongo-go-driver) for golang mongodb driver

## Usage

```go
var session *Session
session = session.New("mongodb://127.0.0.1")
err := session.Connect()
if err != nil {
    fmt.Println(err)
}
```

## Doc

[mango](https://godoc.org/github.com/amorist/mango)
