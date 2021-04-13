package config

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"time"
)

type JaegerConfig struct {
	IP string
	Port int64
	Addr string
}

func NewJaegerConfig() *JaegerConfig {
	c := new(JaegerConfig)

	c.IP = "192.168.18.128"
	c.Port = 6831
	c.Addr = fmt.Sprintf("%s:%d", c.IP, c.Port)

	return c
}

// 创建链路追踪
func NewTracer(serviceName string, addr string) (opentracing.Tracer, io.Closer, error) {
	config2 := &config.Configuration{
		ServiceName: serviceName,
		Sampler: &config.SamplerConfig{
			Type:                     jaeger.SamplerTypeConst,
			Param:                    1,
		},
		Reporter: &config.ReporterConfig{
			BufferFlushInterval:        1 * time.Second,
			LogSpans:                   true,
			LocalAgentHostPort:         addr,
		},
	}

	return config2.NewTracer()
}
