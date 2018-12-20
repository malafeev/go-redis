package otredigo

import (
	"github.com/gomodule/redigo/redis"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
)

type OTConn struct {
	conn redis.Conn
}

func (c *OTConn) Close() error {
	return c.conn.Close()
}

func (c *OTConn) Err() error {
	return c.conn.Err()
}

func (c *OTConn) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	sp := buildSpan(commandName)
	r, e := c.conn.Do(commandName, args...)
	if e != nil {
		ext.Error.Set(sp, true)
		sp.LogFields(log.Error(e))
	}
	sp.Finish()

	return r, e
}

func (c *OTConn) Send(commandName string, args ...interface{}) error {
	sp := buildSpan(commandName)
	e := c.conn.Send(commandName, args...)
	if e != nil {
		ext.Error.Set(sp, true)
		sp.LogFields(log.Error(e))
	}
	sp.Finish()

	return e
}

func (c *OTConn) Flush() error {
	return c.conn.Flush()
}

func (c *OTConn) Receive() (reply interface{}, err error) {
	return c.conn.Receive()
}

func buildSpan(commandName string) opentracing.Span {
	sp := opentracing.StartSpan(commandName)
	ext.SpanKindRPCClient.Set(sp)
	ext.Component.Set(sp, "go-redis")
	ext.DBType.Set(sp, "redis")
	return sp
}
