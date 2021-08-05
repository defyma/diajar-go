# go

`example.com/greetings` was imported in `hello/hello.go`

because `example.com/greetings` already in local, tell go to use local module by call this command
```shell
$ go mod edit -replace example.com/greetings=../greetings
$ go mod tidy
```

see hello/go.mod

run code
```shell
$ cd hello
$ go run .
```