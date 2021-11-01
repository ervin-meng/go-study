package main

import (
	"errors"
	"fmt"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/alibaba/sentinel-golang/core/circuitbreaker"
	"github.com/alibaba/sentinel-golang/core/config"
	"github.com/alibaba/sentinel-golang/core/flow"
	"github.com/alibaba/sentinel-golang/logging"
	"github.com/alibaba/sentinel-golang/util"
	"log"
	"math/rand"
	"time"
)

func main() {
	//FLowLimitStrategyDirect()
	//FlowLimitStrategyWarmUp()
	Break()
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

//circuit breaker，慢调用比例，错误比例，错误数量

type stateChangeTestListener struct {
}

func (s *stateChangeTestListener) OnTransformToClosed(prev circuitbreaker.State, rule circuitbreaker.Rule) {
	fmt.Printf("rule.steategy: %+v, From %s to Closed, time: %d\n", rule.Strategy, prev.String(), util.CurrentTimeMillis())
}

func (s *stateChangeTestListener) OnTransformToOpen(prev circuitbreaker.State, rule circuitbreaker.Rule, snapshot interface{}) {
	fmt.Printf("rule.steategy: %+v, From %s to Open, snapshot: %d, time: %d\n", rule.Strategy, prev.String(), snapshot, util.CurrentTimeMillis())
}

func (s *stateChangeTestListener) OnTransformToHalfOpen(prev circuitbreaker.State, rule circuitbreaker.Rule) {
	fmt.Printf("rule.steategy: %+v, From %s to Half-Open, time: %d\n", rule.Strategy, prev.String(), util.CurrentTimeMillis())
}

func Break() {
	conf := config.NewDefaultConfig()
	// for testing, logging output to console
	conf.Sentinel.Log.Logger = logging.NewConsoleLogger()
	err := sentinel.InitWithConfig(conf)
	if err != nil {
		log.Fatal(err)
	}
	ch := make(chan struct{})
	// Register a state change listener so that we could observer the state change of the internal circuit breaker.
	circuitbreaker.RegisterStateChangeListeners(&stateChangeTestListener{})

	_, err = circuitbreaker.LoadRules([]*circuitbreaker.Rule{
		// Statistic time span=5s, recoveryTimeout=3s, maxErrorCount=50
		{
			Resource:                     "abc",
			Strategy:                     circuitbreaker.ErrorCount,
			RetryTimeoutMs:               3000,
			MinRequestAmount:             10,
			StatIntervalMs:               5000,
			StatSlidingWindowBucketCount: 10,
			Threshold:                    50,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	logging.Info("[CircuitBreaker ErrorCount] Sentinel Go circuit breaking demo is running. You may see the pass/block metric in the metric log.")

	go func() {
		for {
			e, b := sentinel.Entry("abc")
			if b != nil {
				// g1 blocked
				time.Sleep(time.Duration(rand.Uint64()%20) * time.Millisecond)
			} else {
				if rand.Uint64()%20 > 9 {
					// Record current invocation as error.
					sentinel.TraceError(e, errors.New("biz error"))
				}
				// g1 passed
				time.Sleep(time.Duration(rand.Uint64()%80+10) * time.Millisecond)
				e.Exit()
			}
		}
	}()

	go func() {
		for {
			e, b := sentinel.Entry("abc")
			if b != nil {
				// g2 blocked
				time.Sleep(time.Duration(rand.Uint64()%20) * time.Millisecond)
			} else {
				// g2 passed
				time.Sleep(time.Duration(rand.Uint64()%80) * time.Millisecond)
				e.Exit()
			}
		}
	}()
	<-ch
}
