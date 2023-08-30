# bunmicro
A query hook for [uptrace/bun](https://github.com/uptrace/bun) that logs with [go-micro/logger](https://github.com/go-micro/go-micro/tree/master/logger).

```bash
$ go get github.com/alexlast/bunmicro
```

All errors will be logged at error level with the hook enabled, everything else will be logged as debug. If `SlowDuration` is defined, only operations taking longer than the defined duration will be logged.

## Usage
```go
db := bun.NewDB()
db.AddQueryHook(bunmicro.NewQueryHook(bunmicro.QueryHookOptions{
    Logger:       logger.DefaultLogger,
    SlowDuration: 200 * time.Millisecond, // Omit to log all operations as debug
}))
```
