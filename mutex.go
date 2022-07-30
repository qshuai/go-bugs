package go_bugs

import (
	"sync"
)

type MyClass struct {
	count *int
	sync.Mutex
}

// Incr 学生加入 func(go_bugs.MyClass) 在该方法调用时，会使用值类型的MyClass传入参数，相当于拷贝了sync.Mutex,
// 那么在并发执行Incr()、Decr()两个方法时，用到的锁已经不是同一把了，也起不到保护临界区的作用。
// 而拷贝的sync.Mutex可能已经是处于Lock状态了，之后在调用Lock()会阻塞，当程序中剩下的goroutine都阻塞时，那么就会
// 报死锁错误(fatal error: all goroutines are asleep - deadlock!)
func (c MyClass) Incr() {
	c.Lock()
	defer c.Unlock()

	*c.count++
}

// Decr 学生离开 func(*go_bugs.MyClass)
func (c *MyClass) Decr() {
	c.Lock()
	defer c.Unlock()

	*c.count--
}

// fatal error: all goroutines are asleep - deadlock!
func mutexDeadLock() {
	class := &MyClass{count: new(int)}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(no int) {
			defer wg.Done()

			if no%2 == 0 {
				class.Incr()
			} else {
				class.Decr()
			}
		}(i)
	}

	wg.Wait()
}
