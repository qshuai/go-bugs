package go_bugs

// deferReturn 程序执行顺序：
// 1. r = n -> r=1
// 2. 执行defer（闭包函数）n -> 2；但是r的值不变
func deferReturn() (r int) {
	n := 1
	defer func() {
		n++
	}()

	return n
}

// deferReturn2 defer函数捕获了局部变量，虽然执行顺序和 deferReturn 一样，
// 但是修改的是同一份内存地址，所以defer执行影响到了返回值
func deferReturn2() (n int) {
	n = 1
	defer func() {
		n++
	}()

	return
}
