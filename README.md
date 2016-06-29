# lru
a lru cache library

go library

```
➜  lru go version
go version go1.5 linux/amd64
➜  lru go test -v lru.go lru_test.go
=== RUN   Test_Get
--- PASS: Test_Get (0.00s)
        lru_test.go:31: test get sky success
        lru_test.go:37: test get sky1 success
        lru_test.go:41: test get sky2 success
        lru_test.go:48: test get sky success
        lru_test.go:54: test get sky1 success
PASS
ok      command-line-arguments  0.003s
➜  lru go test -test.bench=".*" lru_test.go lru.go 
PASS
BenchmarkAdd-2   3000000               421 ns/op
ok      command-line-arguments  1.695s
```
