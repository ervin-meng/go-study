package sentinel

import (
	"fmt"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/alibaba/sentinel-golang/core/config"
	"github.com/alibaba/sentinel-golang/core/flow"
	"github.com/alibaba/sentinel-golang/logging"
)

func main() {
	//FLowLimitStrategyDirect()
	//FlowLimitStrategyWarmUp()
}

func FLowLimitStrategyDirect() {
	_ = sentinel.InitDefault()
	//基于qps限流，Flow controller
	_, _ = flow.LoadRules([]*flow.Rule{
		{
			Resource:               "some-test",
			TokenCalculateStrategy: flow.Direct,     //直接计数
			ControlBehavior:        flow.Throttling, //匀速通过 //flow.Reject 不匀速
			Threshold:              10,
			StatIntervalInMs:       1000,
		},
	})

	_, b := sentinel.Entry("some-test", sentinel.WithTrafficType(base.Inbound))

	if b != nil {
		fmt.Println("限流了")
	}
}

func FlowLimitStrategyWarmUp() {
	_ = sentinel.InitDefault()

	_, _ = flow.LoadRules([]*flow.Rule{
		{
			Resource:               "some-test",
			WarmUpPeriodSec:        60,
			TokenCalculateStrategy: flow.WarmUp, //预热
			ControlBehavior:        flow.Reject,
			Threshold:              10,
		},
	})

	_, b := sentinel.Entry("some-test", sentinel.WithTrafficType(base.Inbound))

	if b != nil {
		fmt.Println("限流了")
	}
}

//circuit breaker
func Break() {

}
