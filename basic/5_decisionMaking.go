// 条件语句

package main

import "fmt"

const (
	dma = 10
	dmb = 5
)

func ifDecisionMaking() {
	//（1） 不需使用括号将条件包含起来
	//（2） 大括号{}必须存在，即使只有一行语句
	//（3） 左括号必须在if或else的同一行
	//（4） 在if之后，条件语句之前，可以添加变量初始化语句，使用；进行分隔
	//（5） 在有返回值的函数中，最终的return不能在条件语句中
	// if 包含另外一种形式，它存在一个 statement 可选语句部分，使用 ; 分隔，该组件在条件判断之前运行。
	if v := 200; v > dma {
		fmt.Println(true)
	}
	// if else
	if dma == dmb {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
	// 嵌套if
	if dma == dmb {
		if dma > 20 {
			fmt.Println("dma>20")
		} else {
			fmt.Println("dma<20")
		}
	}
}
func switchDecisionMaking() {
	// switch
	// switch 语句用于基于不同条件执行不同动作，每一个 case 分支都是唯一的，从上至下逐一测试，直到匹配为止。
	// switch 语句执行的过程从上至下，直到找到匹配项，匹配项后面也不需要再加 break。
	// switch 默认情况下 case 最后自带 break 语句，匹配成功后就不会执行其他 case，如果我们需要执行后面的 case，可以使用 fallthrough 。
	switch dma / dmb {
	case 2:
		fmt.Println("2")
		fallthrough // 不中断, 使用 fallthrough 会强制执行后面的 case 语句，fallthrough 不会判断下一条 case 的表达式结果是否为 true。
	case 0, 3, 5: // 多条件匹配
		fmt.Println("0")
	default:
		// 缺省分支
		fmt.Println("default")
	}

	// 判断接口变量返回类型
	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Println("%T", i)
	case int:
		fmt.Println("int")
	case bool:
		fmt.Println("bool")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("unknown")
	}
}

func selectDecisionMaking() {
	//select 类似于 switch 语句。
	//select 语句只能用于通道操作，每个 case 必须是一个通道操作，要么是发送要么是接收。
	//select 语句会监听所有指定的通道上的操作，一旦其中一个通道准备好就会执行相应的代码块。
	//如果多个通道都准备好，那么 select 语句会随机选择一个通道执行。如果所有通道都没有准备好，那么执行 default 块中的代码。

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func(numCh chan string) {
		numCh <- "one"
	}(ch1)
	go func(numCh chan string) {
		numCh <- "two"
	}(ch2)

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
			//default:
			//	// 如果两个通道都没有可用的数据，则执行这里的语句
			//	fmt.Println("no message received")
		}
	}
}

func main() {
	ifDecisionMaking()
	switchDecisionMaking()
	selectDecisionMaking()
}
