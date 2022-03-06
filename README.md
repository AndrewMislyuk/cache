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
c = cache.New()
c.Set("first", 6)
c.Set("second", "March")
c.Get("first")
c.Delete("second")
```
