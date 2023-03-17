# Cache
___

In-memory cache implementation, Go package

```shell
go get github.com/AlexCorn999/inMemoryCache@v1.0.5
```
```go
...
  Set(key string, value any)
  Get(key string) (any, error)
  Delete(key string) error
...
```

# example 1

```go
package main

import (
	"fmt"
	"log"
	"time"

	cache "github.com/AlexCorn999/inMemoryCache"
)

func main() {
	cache := cache.New()

	cache.Set("userId", 42, time.Second*2)

	userId, err := cache.Get("userId")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(userId)

	time.Sleep(time.Second * 2)

	newId, err := cache.Get("userId")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(newId)
}
```
```text
go run main.go
42
value not found
```
___
