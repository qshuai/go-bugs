package go_bugs

import (
	"errors"
	"time"
)

// groutineLeakage goroutine泄漏的case，现象时groutine所在的函数已经结束，但是创建的groutine仍然阻塞在了一个channel上，
// 导致其后续无法被唤醒，其所占用的资源也得不到释放
func groutineLeakage(duration, timeout time.Duration) error {
	done := make(chan struct{}) // 设置为有缓冲的channel，会解决该问题（协程不会阻塞）
	go func() {
		time.Sleep(duration)
		done <- struct{}{}
	}()

	select {
	case <-done: // 如果timer先执行，那么scase对应的channel上用于阻塞了一个*sudog，会导致对应的groutine无法结束
		return nil
	case <-time.After(timeout):
		return errors.New("timeout")
	}
}
