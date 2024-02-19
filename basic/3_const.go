// 常量

package main

import (
	"fmt"
	"unsafe"
)

// 常量是一个简单值的标识符，在程序运行时，不会被修改的量。
// 常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。
// 显式类型定义
const ca int = 1

// 隐式类型定义
const cb = 2

// 多个相同类型的声明
const cc, cd = 3, 4

func demo() {
	const length = 10
	const width = 5
	var area int
	const aa, ab, ac = 1, 2, false

	area = length * width
	fmt.Printf("area: %d\n", area)
	fmt.Println(aa, ab, ac)
}

// 枚举常量
// 常量可以用len(), cap(), unsafe.Sizeof()函数计算表达式的值
// 常量表达式中，函数必须是内置函数，否则编译不过
const (
	ba = "abc"
	bb = len(ba)
	bc = unsafe.Sizeof(ba)
)

// iota
// 特殊常量，可以认为是一个可以被编译器修改的常量。
// iota在const关键字出现时将被重置为0(const内部的第一行之前)，const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。

// da=0, db=1, dc=2可简写为如下形式
const (
	da = iota
	db
	dc
)

func demoIota() {
	const (
		a = iota //0
		b        //1
		c        //2
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数
		i        //8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)
}

func main() {
	demo()
	fmt.Println(ba, bb, bc)
	demoIota()
}
