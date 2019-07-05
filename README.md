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
