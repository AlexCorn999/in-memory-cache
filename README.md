# Cache
___

In-memory cache implementation, Go package

```shell
go get github.com/AlexCorn999/inMemoryCache@v1.0.3
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

	cache "github.com/AlexCorn999/inMemoryCache"
)

func main() {
	cache := cache.New()

	cache.Set("userId", 42)

	userId, err := cache.Get("userId")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(userId)

	cache.Delete("userId")
	newId, err := cache.Get("userId")
	if err != nil {
		fmt.Println(err)
		return
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
