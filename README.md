# ulib

A collection of production-ready Go utility packages used across
[uopensail](https://github.com/uopensail) services.

**Module:** `github.com/uopensail/ulib`
**Requires:** Go 1.26+

---

## Packages

### `commonconfig` — Shared configuration structs

Common configuration types shared across services. All structs support
JSON, TOML, and YAML tags.

| Type | Purpose |
|---|---|
| `HttpServerConfig` | HTTP server (port, timeouts, header limits) |
| `EtcdConfig` | etcd client (endpoints, credentials, dial timeout) |
| `RegisterDiscoveryConfig` | Service registration / discovery via etcd |
| `ServerConfig` | Top-level server config embedding the above |
| `MongoClientConfig` | MongoDB connection pool settings |
| `RedisConfig` | Redis connection and pool settings |
| `GRPCClientConfig` | gRPC client connection settings |
| `DatabaseConfig` | Composite database config (Mongo + Redis) |
| `ClientConfig` | Composite client config (gRPC + Redis + Mongo) |

---

### `zlog` — Structured logging

Thin wrapper around [go.uber.org/zap](https://pkg.go.dev/go.uber.org/zap)
that exposes two package-level loggers and a helper to initialise them
with file rotation.

```go
// Package-level loggers, available after init().
zlog.LOG  // *zap.Logger
zlog.SLOG // *zap.SugaredLogger

// Custom initialisation with lumberjack rotation.
zlog.InitLogger("myapp", debug, "/var/log/myapp")
// Writes to /var/log/myapp/myapp_release.log (or _debug.log in debug mode).
// debug=true also mirrors output to stdout and adds stack traces on Warn+.
```

---

### `prome` — Prometheus metrics

Drop-in Prometheus exporter with built-in latency histograms, percentiles,
and QPS tracking. Metrics are accumulated in a background goroutine and
collected every 30 seconds.

```go
// Record a call with automatic latency measurement.
stat := prome.NewStat("my_rpc").MarkErr()
defer stat.End()
// ... do work ...
stat.MarkOk()

// Expose /metrics on port 9090.
exporter := prome.NewExporter("myapp")
go exporter.Start(9090)
```

Per-metric data exported to Prometheus:

| Metric | Description |
|---|---|
| `qps` | Queries per second |
| `avg_cost_time` | Mean latency (ms) |
| `p90/p95/p99_cost_time` | Percentile latencies (ms) |
| `max_cost_time` | Maximum observed latency (ms) |
| `avg_counter` | Mean counter value |
| `cost` | Cumulative latency histogram |

---

### `dmutex` — Distributed mutex

Distributed lock backed by etcd using `concurrency.Session` / `concurrency.Mutex`.
The internal session is created lazily and is safe for concurrent use.

```go
cli, _ := etcdclient.New(etcdclient.Config{Endpoints: []string{"localhost:2379"}})
m := dmutex.NewDMutexEtcd(cli, 10) // 10-second timeout

mu, err := m.Lock("/my/lock/key")
if err != nil { ... }
defer m.Unlock(mu)

// Non-blocking variant:
mu, err = m.TryLock("/my/lock/key") // returns concurrency.ErrLocked if held
```

---

### `sample` — High-performance feature storage

Arena-allocated, zero-copy feature store for heterogeneous typed values.
See [`sample/readme.md`](sample/readme.md) for full documentation.

```go
arena := sample.NewArena()
features, _ := sample.NewImmutableFeaturesFromMap(map[string]any{
    "user_id": int64(42),
    "score":   float32(9.8),
    "tags":    []string{"vip", "active"},
}, arena)

// Zero-copy read
id := features.Get("user_id").GetInt64Unsafe()

// Range iteration (Go 1.23+)
for key, f := range features.All() {
    fmt.Printf("%s = %v\n", key, f.Get())
}
```

Supported types: `int64`, `float32`, `string`, `[]int64`, `[]float32`, `[]string`.

---

### `datastruct` — Generic data structures

#### `BitMap` / `DoubleBitMap`

Compact bit arrays backed by a `[]byte` slice.
`DoubleBitMap` stores two bits per slot (values 0–3).

```go
bm := datastruct.CreateBitMap(1024)
bm.MarkTrue(42)
bm.Check(42)    // true
bm.Clear()
bm.And(other)   // bitwise AND in-place
bm.Or(other)    // bitwise OR in-place

dbm := datastruct.CreateDoubleBitMap(512)
dbm.Mark(10, true)              // set slot 10
dbm.Check(10)                   // BitMapStatus
```

#### `CombineSliceBuilder`

Appends elements from many small slices into a single per-type backing array,
returning sub-slices that share that memory. Avoids a heap allocation per slice.

```go
var b datastruct.CombineSliceBuilder

ids   := b.CombineInt64s([]int64{1, 2, 3})    // view into shared int64 array
tags  := b.CombineStrings([]string{"a", "b"}) // view into shared string array
raw   := b.CombineStringByte([]byte("hello")) // zero-copy string from byte array
```

#### `Tuple`

Generic ordered pair with JSON support.

```go
t := datastruct.Tuple[int64, float32]{First: 1, Second: 9.5}
fmt.Println(t.String()) // (1, 9.5)

data, _ := json.Marshal(&t) // {"first":1,"second":9.5}
```

Type constraints available: `Ordered`, `Integer`, `Signed`, `Unsigned`, `Float`.

---

### `targz` — Tar/Gzip utilities

Compress a directory tree or extract a `.tar.gz` archive with path traversal protection.

```go
// Compress src/ into archive.tar.gz
err := targz.Compress("src/", "archive.tar.gz")

// Extract archive.tar.gz into dst/
err = targz.Extract("archive.tar.gz", "dst/")
```

---

### `utils` — General utilities

#### String conversion (`str.go`)

Null-safe conversions between strings and numeric/bool types.
All functions return the zero value on parse failure.

```go
utils.String2Int64("123")          // 123
utils.String2Float32("3.14")       // 3.14
utils.String2Int64List("1,2,3", ",") // []int64{1, 2, 3}
utils.Int642String(int64(42))      // "42"
utils.StringSplit("a,b,c", ",")    // ["a","b","c"] — nil on empty input
```

#### IP utilities (`ip.go`)

```go
ip, err := utils.GetLocalIP()                   // first non-loopback IPv4
ips, err := utils.GetAllLocalIPs()              // all non-loopback IPv4
ip, err = utils.GetLocalIPByInterface("eth0")  // specific interface
ip, err = utils.GetOutboundIP()                // preferred outbound IP (UDP trick)
utils.IsValidIPv4("192.168.1.1")               // true
```

#### File utilities (`file.go`)

```go
utils.FilePathExists(path)         // existence check
utils.GetFileModifyTime(path)      // Unix mtime, -1 on error
utils.GetFileSize(path)            // bytes, -1 on error
utils.GetMD5(path)                 // hex MD5 string
utils.ListDir(dir)                 // []string of full child paths
utils.CopyFile(src, dst)           // copies, creating dst directories
utils.MoveFile(src, dst)           // os.Rename with mkdirall
```

#### GORM JSON types (`db.go`)

Drop-in GORM column types that serialize Go values as JSON in SQL columns.
Use `JSONB` on PostgreSQL, `NVARCHAR(MAX)` on SQL Server, `JSON` elsewhere.

```go
type Model struct {
    Tags   utils.StringMap   `gorm:"column:tags"`    // map[string]string
    Ids    utils.Int64Slice   `gorm:"column:ids"`     // []int64
    Scores utils.Float32Slice `gorm:"column:scores"`  // []float32
    Counts utils.IntSlice     `gorm:"column:counts"`  // []int
    Vals   utils.Float64Slice `gorm:"column:vals"`    // []float64
}
```

#### Reference counting (`misc.go`)

```go
ref := &utils.Reference{CloseHandler: resource.Close}
ref.Retain()
// ... hand to goroutine ...
ref.Release()

ref.Free()               // blocks until count == 0, then calls CloseHandler
ref.LazyFree(5)          // waits 5 s, then waits for count == 0 (non-blocking)
```

#### Distributed job runner (`metux_job.go`)

Runs a function exclusively across a cluster using an etcd distributed lock.
Only one node executes the job at a time; others return immediately.

```go
runner := utils.NewMetuxJobUtil("myprefix", reg, etcdCli, 10)
runner.TryRun(func() {
    // Only one node in the cluster executes this at a time.
})
```

#### Service discovery helpers (`misc.go`)

```go
conn, err := utils.NewKratosGrpcConn(rdConf)  // gRPC via etcd discovery
cli, err  := utils.NewKratosHttpConn(rdConf)  // HTTP via etcd discovery
```

---

## Installation

```bash
go get github.com/uopensail/ulib
```

## License

See [LICENSE](LICENSE).
