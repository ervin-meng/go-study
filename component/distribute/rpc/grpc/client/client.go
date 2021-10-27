package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"go-study/component/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"time"
)

func main() {

	TracerCloser := InitTracer()

	defer TracerCloser()

	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())))
	//opts = append(opts, grpc.WithPerRPCCredentials())

	conn, e := grpc.Dial("127.0.0.1:9600", opts...)

	if e != nil {
		panic(e)
	}

	defer conn.Close()

	client := protobuf.NewStudyServiceClient(conn)

	md := metadata.Pairs("test", "test")

	ctx := metadata.NewOutgoingContext(context.Background(), md)

	ctx, _ = context.WithTimeout(ctx, 10*time.Second)

	r, e := client.Get(ctx, &protobuf.StudyRequest{Sys: 2})

	if e != nil {
		if st, ok := status.FromError(e); !ok {
			panic(e)
		} else {
			fmt.Println(st.Message())
			fmt.Println(st.Code())
		}
	} else {
		fmt.Println(r.Sys)
	}
}

func InitTracer() func() {
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

	return func() {
		closer.Close()
	}
}
