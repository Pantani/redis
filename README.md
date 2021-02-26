[![codecov](https://codecov.io/gh/Pantani/redis/branch/master/graph/badge.svg?token=8IHUB3K2Q7)](https://codecov.io/gh/Pantani/redis)

# Simple abstraction for [Go-Redis](github.com/go-redis/redis).

Simple abstraction using generic interfaces for [Go-Redis](github.com/go-redis/redis).

Initialize the database:
```go
import "github.com/Pantani/redis"

cache := redis.New("localhost:6379", "password", 0)
if err != nil {
    panic("Cannot initialize the redis storage")
}
if !storage.IsReady() {
    panic("redis storage is not ready")
}
```

 ## Fetching Objects:

- Get value:
```go
var result CustomObject
err := s.GetObject("key", &result)
```

- Add value
```go
data := CustomObject{Name: "name", Id: "id"}
err := s.AddObject("key", data, 0)

// with expiration time
err := s.AddObject("key", data, 10 * time.Second)
```

- Delete value:
```go
err := s.DeleteObject("table", "key")
```


### Hash Map

Redis hash map abstraction

- Get all values from a hash map table:
```go
result, err := s.GetAllHMObjects("table")
```

- Get value from a hash map table:
```go
var result CustomObject
err := s.GetHMObject("table", "key", &result)
```

- Add value to a hash map table:
```go
data := CustomObject{Name: "name", Id: "id"}
err := s.AddHMObject("table", "key", data)
```

- Delete value from a hash map table:
```go
err := s.DeleteHMObject("table", "key")
```
