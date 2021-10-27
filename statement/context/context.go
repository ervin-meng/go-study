package main

import (
	"context"
	"fmt"
	"time"
)

type request struct {
	sys      int
	deviceId string
	cityId   int
}

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, request{}, request{1, "device", 1})
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	task(ctx)
}

func task(ctx context.Context) {
	taskStep(ctx, "task Step 1", 1*time.Second)
	taskStep(ctx, "task Step 2", 2*time.Second)
	ctx1 := context.WithValue(ctx, "sub", "1234")
	smallTask(ctx1, "small task ")
	go func() {
		backgroundTask(context.Background(), "background task")
	}()
}

func taskStep(ctx context.Context, name string, d time.Duration) {
	fmt.Println(name)
	select {
	case <-time.After(d):
		requestParams := ctx.Value(request{}).(request)
		fmt.Println("Done", requestParams.cityId)
	case <-ctx.Done():
		fmt.Println("cancel")
	}
}

//子任务
func smallTask(ctx context.Context, name string) {

}

//后台任务
func backgroundTask(ctx context.Context, name string) {

}
