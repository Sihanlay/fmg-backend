package current_limiting

import (
	"grpc-demo/utils"
	"time"
)

var TokenBucket chan int

// 启动令牌桶
func StartTokenBucket() {

	TokenBucket = make(chan int, utils.GlobalConfig.Server.TokenBucketCapacity)
	// 放置令牌
	go func() {
		t := time.NewTimer(time.Second)
		for {
			<-t.C
			for i := 0; i < utils.GlobalConfig.Server.TokenBucketOutputPerSecond; i++ {
				TokenBucket <- 1
			}
			t.Reset(time.Second)
		}
	}()
}



