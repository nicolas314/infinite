# infinite
Infinite web page server

This small piece of code defines a misbehaving web server. Any request will
return an infinite web page, which is a good test for HTTP clients. Can
also be used to keep network mappers busy for a while.

Build and run:

```
go build infinite.go
./infinite :8080
```

-- 
