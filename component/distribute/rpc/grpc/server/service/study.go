package service

import (
	"context"
	"fmt"
	"go-study/component/protobuf"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type StudyService struct {
}

func (s *StudyService) Get(ctx context.Context, request *protobuf.StudyRequest) (*protobuf.StudyResponse, error) {

	//time.Sleep(5 * time.Second)

	md, ok := metadata.FromIncomingContext(ctx)

	if ok {
		fmt.Println(md)
	}

	//err := status.Error(codes.InvalidArgument, "invalid request")

	return &protobuf.StudyResponse{
		Sys:     int32(request.Sys),
		AddTime: timestamppb.New(time.Now()),
	}, nil
}
