# fever-go

> Fever API for Golang
> Fever is a service for reading RSS feeds.

## Install

```shell
go get github.com/song940/fever-go
```

## Example

```go
package main

import (
  "net/http"

  "github.com/song940/song940/fever-go/fever"
)

func main() {
  reader, err := reader.NewReader()
  if err != nil {
    panic(err)
  }
  api := fever.New(reader)
  http.Handle("/fever/", api)
  err = http.ListenAndServe(":8080", nil)
  if err != nil {
    panic(err)
  }
}
```

## License

MIT