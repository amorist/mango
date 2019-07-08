# mango

![mango](./logo/mango.svg)

use [mongo-go-driver](https://github.com/mongodb/mongo-go-driver) like [mgo](https://github.com/globalsign/mgo)

- [Installation](#installation)
- [Usage](#usage)
- [Doc](#doc)

## installation

`go get -u github.com/amorist/mango`

## usage

```go
import "github.com/amorist/mango"

var session *Session
session = session.New("mongodb://127.0.0.1")
err := session.Connect()
if err != nil {
    fmt.Println(err)
}
```

## doc

[mango](https://godoc.org/github.com/amorist/mango)
