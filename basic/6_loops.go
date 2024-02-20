// 循环语句

package main

import (
	"fmt"
	"time"
)

func loops() {
	// Go支持for循环以及嵌套循环

	// for init; condition; post { }
	//  init： 一般为赋值表达式，给控制变量赋初值；（可选）
	//  condition： 关系表达式或逻辑表达式，循环控制条件；（可选）
	//  post： 一般为赋值表达式，给控制变量增量或减量。
	// 运行逻辑
	// 1、先对表达式 1 赋初值；
	// 2、判别赋值表达式 init 是否满足给定条件，若其值为真，满足循环条件，则执行循环体内语句，然后执行 post，进入第二次循环，
	//    再判别 condition；否则判断 condition 的值为假，不满足条件，就终止for循环，执行循环体外语句。
	//（与C语言中for一致）
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}

	// for condition { }
	//（与C语言中while一致）
	a := 10
	for a > 0 {
		fmt.Println(a)
		a--
	}

	// 无限循环：for { }
	//（与C语言中for(::)一致）
	//for {
	//	fmt.Println("test")
	//}

	// For-each range 循环
	// for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环。（与python类似）
	m := map[string]int{
		"11": 1,
		"22": 2,
	}
	for key, value := range m {
		fmt.Printf("%s:%d\n", key, value)
	}
	for key := range m {
		fmt.Printf("%s\n", key)
	}
	for _, value := range m {
		fmt.Printf("%d\n", value)
	}
	for key, _ := range m {
		fmt.Printf("%s\n", key)
	}

}

func loopControl() {
	// break
	// 在 Go 语言中，break 语句用于终止当前循环或者 switch 语句的执行，并跳出该循环或者 switch 语句的代码块。
	// break 语句可以用于以下几个方面：
	// - 用于循环语句中跳出循环，并开始执行循环之后的语句。
	// - break 在 switch 语句中在执行一条 case 后跳出语句的作用。
	// - break 可应用在 select 语句中。
	// - 在多重循环中，可以用标号 label 标出想 break 的循环。

	//1. 单独在select中使用break和不使用break没有区别。
	//2. 单独在表达式switch语句，并且没有fallthough，使用break和不使用break没有区别。
	//3. 单独在表达式switch语句，并且有fallthough，使用break能够终止fallthough后面的case语句的执行。
	//4. 带标签的break，可以跳出多层select/ switch作用域。让break更加灵活，写法更加简单灵活，不需要使用控制变量一层一层跳出循环，没有带break的只能跳出当前语句块。

	for i := 0; i < 10; i++ {
		if i == 5 {
			break // 当 i 等于 5 时跳出循环
		}
		fmt.Println(i)
	}

SELECT:
	for {
		select {
		case <-time.After(time.Second):
			fmt.Println("一秒后退出")
			//break 跳出select
			break SELECT //带标签的break，实际上跳出到select外层的for语句块
		case <-time.After(time.Second * 10):
			fmt.Println("十秒后退出")
			break
		}
	}
	fmt.Println("select 语句结束")

	// continue		跳过当前循环的剩余语句，然后继续进行下一轮循环。
	// Go 语言的 continue 语句 有点像 break 语句。但是 continue 不是跳出循环，而是跳过当前循环执行下一次循环语句。
	// for 循环中，执行 continue 语句会触发 for 增量语句的执行。
	// 在多重循环中，可以用标号 label 标出想 continue 的循环。
	var a int = 10

	/* for 循环 */
	for a < 20 {
		if a == 15 {
			/* 跳过此次循环 */
			a = a + 1
			continue
		}
		fmt.Printf("a 的值为 : %d\n", a)
		a++
	}

	// goto（不推荐）
	// Go 语言的 goto 语句可以无条件地转移到过程中指定的行。
	// goto 语句通常与条件语句配合使用。可用来实现条件转移， 构成循环，跳出循环体等功能。
	a = 10

	/* 循环 */
LOOP:
	for a < 20 {
		if a == 15 {
			/* 跳过迭代 */
			a = a + 1
			goto LOOP
		}
		fmt.Printf("a的值为 : %d\n", a)
		a++
	}
}

func main() {
	loops()
	loopControl()
}
