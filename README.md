# Cache

Simple in-memory cache

Getting Cache

```sh
go get -u github.com/AndrewMislyuk/cache
```

Import

```go
import "github.com/AndrewMislyuk/cache"
```

Usage

```go
c := cache.New()
c.Set("userId", 42, time.Second*2)
userId, err := c.Get("userId")
if err != nil { // err == nil
	log.Fatal(err)
}
fmt.Println(userId) // Output: 42

time.Sleep(time.Second * 3) // прошло 3 секунд

_, err = c.Get("userId")
if err != nil { // err != nil
	log.Fatal(err) // сработает этот код
}
```
