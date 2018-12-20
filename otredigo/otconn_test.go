package otredigo

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/mocktracer"
	"testing"
)

// Test requires running local instance of Redis server
func TestRedigo(t *testing.T) {
	tracer := mocktracer.New()
	opentracing.InitGlobalTracer(tracer)

	conn, e := redis.DialURL("redis://localhost:6379")
	if e != nil {
		fmt.Println(e)
		return
	}

	otConn := OTConn{conn}

	n, e := otConn.Do("SET", "key", "value")
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println(n)

	n, e = redis.String(otConn.Do("GET", "key"))
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println(n)

	spans := tracer.FinishedSpans()
	fmt.Println(spans)

	if len(spans) != 2 {
		t.Errorf("Expected 2 finished span. Found: %d", len(spans))
	}

}
