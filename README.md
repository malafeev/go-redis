# OpenTracing support for Redis Client in Go

## Installation

```
go get github.com/opentracing-contrib/go-redis
```

## Documentation

See the basic usage examples below and the [package documentation on
godoc.org](https://godoc.org/github.com/opentracing-contrib/go-redis).

## Usage

```go
// You must have some sort of OpenTracing Tracer instance on hand
var tracer opentracing.Tracer = ...

// Set Tracer as global 
opentracing.SetGlobalTracer(tracer)
```

### Redigo

The `otredigo` package makes it easy to add OpenTracing support for Redigo.

```go
// Create Redigo Conn
conn, e := redis.DialURL("redis://localhost:6379")

// Decorate  Conn with OTConn 
otConn := OTConn{conn}

// use OTConn to send redis commands
n, e := otConn.Do("SET", "key", "value")

```

## License

[Apache 2.0 License](./LICENSE).
