package main

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"time"
)

func main() {
	V1()
}

func V1() {
	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{ //采样
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "127.0.0.1:6831",
		},
		ServiceName: "study",
	}

	tracer, closer, _ := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))

	opentracing.SetGlobalTracer(tracer)

	defer closer.Close()

	//singleSpan(tracer)
	//multiSpan(tracer)
	multiSpanWithParent(tracer)
}

func singleSpan(t opentracing.Tracer) {
	span := t.StartSpan("single")
	time.Sleep(500 * time.Millisecond)
	span.Finish()
}

func multiSpan(t opentracing.Tracer) {
	spanA := t.StartSpan("multi-a")
	time.Sleep(500 * time.Millisecond)
	spanA.Finish()
	spanB := t.StartSpan("multi-b")
	time.Sleep(1000 * time.Millisecond)
	spanB.Finish()
}

func multiSpanWithParent(t opentracing.Tracer) {
	parentSpan := t.StartSpan("parent")
	spanA := t.StartSpan("multi-a", opentracing.ChildOf(parentSpan.Context()))
	time.Sleep(500 * time.Millisecond)
	spanA.Finish()
	spanB := t.StartSpan("multi-b", opentracing.ChildOf(parentSpan.Context()))
	time.Sleep(1000 * time.Millisecond)
	spanB.Finish()
	parentSpan.Finish()
}

func V2() {
	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{ //采样
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "127.0.0.1:6831",
		},
		ServiceName: "study",
	}

	tracer, closer, _ := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))

	//设置全局tracer
	opentracing.SetGlobalTracer(tracer)

	defer closer.Close()

	span := opentracing.StartSpan("single")
	time.Sleep(500 * time.Millisecond)
	span.Finish()
}
