package go_bugs

import (
	"bytes"
	"fmt"
	"io"
)

func ExampleCrossAssert() {
	bf := bytes.NewBuffer(nil)
	var r io.Reader
	r = bf
	// 以下写法可以正常运行，原因是：类型断言是判断接口变量的动态类型(而非接口类型)定义的方法集合是否满足目标接口的定义
	// 对于空接口(runtime.eface)来说：也就是_type字段代表的动态类型是否实现了目标接口；
	// 对于非空接口(runtime.iface)来说:也就是tab._type字段代表的动态类型是否实现了目标接口；
	// type eface struct {
	//		_type *_type
	//		data  unsafe.Pointer
	// }
	//
	// type iface struct {
	//		tab  *itab
	//		data unsafe.Pointer
	// }
	// type itab struct {
	//		inter *interfacetype
	//		_type *_type
	//		hash  uint32 // copy of _type.hash. Used for type switches.
	//		_     [4]byte
	//		fun   [1]uintptr // variable sized. fun[0]==0 means _type does not implement inter.
	// }
	w, ok := r.(io.Writer)
	if !ok {
		return
	}

	n, _ := w.Write([]byte("helloworld"))
	fmt.Println(n)

	// output:
	// 10
}
