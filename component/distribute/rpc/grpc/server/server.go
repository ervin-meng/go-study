package main

import (
	"context"
	"go-study/component/distribute/rpc/grpc/server/service"
	"go-study/component/protobuf"
	"google.golang.org/grpc"
	"net"
	"time"
)

func main() {

	lis, err := net.Listen("tcp", ":9600")

	if err != nil {
		panic(err)
	}

	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

		println(time.Now().Format("2006-01-02 15:04:05"))
		res, err := handler(ctx, req)
		println(time.Now().Format("2006-01-02 15:04:05"))
		return res, err
	}

	g := grpc.NewServer(grpc.UnaryInterceptor(interceptor))

	protobuf.RegisterStudyServiceServer(g, &service.StudyService{})

	err = g.Serve(lis)

	if err != nil {
		panic(err)
	}
}

//服务端流模式，例子是客户端向服务端发送一个股票代码，服务端就把该股票的实时数据源源不断返回给客户端
//客户端流模式，例子是物联网终端向服务器报送数据
//双向流模式，例子是聊天机器人

//interceptor 拦截器

//validator 验证器 还不稳定

//metadata
//每一次的RPC调用中，都可能会有一些有用的数据，这些数据可以通过metadata传递

//status
