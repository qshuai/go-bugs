package go_bugs

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func Test_groutineLeakage(t *testing.T) {
	go func() {
		// 启动3个groutine，在实际应用中，可能是有http handler来启动
		for i := 0; i < 3; i++ {
			// 模拟过期时间小于程序执行时间的设置
			go groutineLeakage(100*time.Millisecond, 50*time.Millisecond)
		}
	}()

	for i := 0; i < 3; i++ {
		// 每秒钟打印一次groutine的数量
		time.Sleep(200 * time.Millisecond)
		fmt.Println(runtime.NumGoroutine())
	}

	// output:
	// 5
	// 5
	// 5
}
