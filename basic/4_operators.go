// 运算符

package main

import "fmt"

// 算数运算符
func arithmeticOperator() {
	//	+	相加
	//	-	相减
	//	*	相乘
	//	/	相除
	//	%	求余
	//	++	自增
	//	--	自减
	const (
		a = 10
		b = 20
	)
	var c int
	var d = 30

	c = a + b
	fmt.Printf("%d + %d = %d\n", a, b, c)
	c = a - b
	fmt.Printf("%d - %d = %d\n", a, b, c)
	c = a * b
	fmt.Printf("%d * %d = %d\n", a, b, c)
	c = a / b
	fmt.Printf("%d / %d = %d\n", a, b, c)
	c = a % b
	fmt.Printf("%d 取余 %d = %d\n", a, b, c)
	d++
	fmt.Printf("30自增=%d\n", d)
	d = 30
	d--
	fmt.Printf("30自减=%d\n", d)
}

// 关系运算符
func relationalOperator() {
	//	==	检查两个值是否相等，如果相等返回 True 否则返回 False。
	//	!=	检查两个值是否不相等，如果不相等返回 True 否则返回 False。
	//	>	检查左边值是否大于右边值，如果是返回 True 否则返回 False。
	//	<	检查左边值是否小于右边值，如果是返回 True 否则返回 False。
	//	>=	检查左边值是否大于等于右边值，如果是返回 True 否则返回 False。
	//	<=	检查左边值是否小于等于右边值，如果是返回 True 否则返回 False。
	aa, ab := 10, 20
	var ac bool
	ac = aa == ab
	fmt.Println(ac)
	ac = aa != ab
	fmt.Println(ac)
	ac = aa > ab
	fmt.Println(ac)
	ac = aa < ab
	fmt.Println(ac)
	ac = aa >= ab
	fmt.Println(ac)
	ac = aa <= ab
	fmt.Println(ac)
}

// 逻辑运算符
func logicalOperator() {
	// &&	逻辑 AND 运算符。 如果两边的操作数都是 True，则条件 True，否则为 False。
	// ||	逻辑 OR 运算符。 如果两边的操作数有一个 True，则条件 True，否则为 False。
	// !	逻辑 NOT 运算符。 如果条件为 True，则逻辑 NOT 条件 False，否则为 True。
	const (
		a = true
		b = false
	)
	var c bool
	c = a && b
	fmt.Println(c)
	c = a || b
	fmt.Println(c)
	c = !a
	fmt.Println(c)
	c = !b
	fmt.Println(c)
}

// 位运算符
func bitwiseOperators() {
	//	p	q	p & q	p | q	p ^ q
	//	0	0	  0		  0		  0
	//	0	1	  0		  1		  1
	//	1	1	  1		  1		  0
	//	1	0	  0		  1		  1
	// ======================================
	//	&	按位与运算符"&"是双目运算符。 其功能是参与运算的两数各对应的二进位相与
	//	|	按位或运算符"|"是双目运算符。 其功能是参与运算的两数各对应的二进位相或
	//	^	按位异或运算符"^"是双目运算符。 其功能是参与运算的两数各对应的二进位相异或
	//	<<	左移运算符"<<"是双目运算符。左移n位就是乘以2的n次方
	//	>>	右移运算符">>"是双目运算符。右移n位就是除以2的n次方
	const (
		a uint = 60 // 0011 1100
		b uint = 13 // 0000 1101
	)
	var c uint = 0
	c = a & b
	fmt.Println(c) // 0011 1100 & 0000 1101 = 0000 1100 (12)
	c = a | b
	fmt.Println(c) // 0011 1100 | 0000 1101 = 0011 1101 (61)
	c = a ^ b
	fmt.Println(c) // 0011 1100 ^ 0000 1101 = 0011 0001 (49)
	c = a << 2
	fmt.Println(c) // 0011 1100 << 2 = 1111 0000 (240)
	c = a >> 2
	fmt.Println(c) // 0011 1100 >> 2 = 0000 1111 (15)
}

// 赋值运算符
func assignmentOperators() {
	//	=	简单的赋值运算符，将一个表达式的值赋给一个左值	C = A + B 将 A + B 表达式结果赋值给 C
	//	+=	相加后再赋值	C += A 等于 C = C + A
	//	-=	相减后再赋值	C -= A 等于 C = C - A
	//	*=	相乘后再赋值	C *= A 等于 C = C * A
	//	/=	相除后再赋值	C /= A 等于 C = C / A
	//	%=	求余后再赋值	C %= A 等于 C = C % A
	//	<<=	左移后赋值	C <<= 2 等于 C = C << 2
	//	>>=	右移后赋值	C >>= 2 等于 C = C >> 2
	//	&=	按位与后赋值	C &= 2 等于 C = C & 2
	//	^=	按位异或后赋值	C ^= 2 等于 C = C ^ 2
	//	|=	按位或后赋值	C |= 2 等于 C = C | 2
	const a = 21
	var c int
	c = a
	fmt.Println(c)
	c += a
	fmt.Println(c)
	c -= a
	fmt.Println(c)
	c *= a
	fmt.Println(c)
	c /= a
	fmt.Println(c)
	c %= a
	fmt.Println(c)

	c = 200
	c <<= 2
	fmt.Println(c)
	c >>= 2
	fmt.Println(c)
	c &= 2
	fmt.Println(c)
	c ^= 2
	fmt.Println(c)
	c |= 2
	fmt.Println(c)
}

// 其他运算符
func otherOperators() {
	//	&	返回变量存储地址。	&a; 将给出变量的实际地址。
	//	*	指针变量。	*a; 是一个指针变量
	var a int = 4
	var b int32
	var c float32
	var ptr *int

	fmt.Printf("第 1 行 - a 变量类型为 = %T\n", a)
	fmt.Printf("第 2 行 - b 变量类型为 = %T\n", b)
	fmt.Printf("第 3 行 - c 变量类型为 = %T\n", c)

	/*  & 和 * 运算符实例 */
	ptr = &a /* 'ptr' 包含了 'a' 变量的地址 */
	fmt.Printf("a 的值为  %d\n", a)
	fmt.Printf("*ptr 为 %d\n", *ptr)
}

func main() {
	fmt.Println("算数运算符")
	arithmeticOperator()
	fmt.Println("关系运算符")
	relationalOperator()
	fmt.Println("逻辑运算符")
	logicalOperator()
	fmt.Println("位运算符")
	bitwiseOperators()
	fmt.Println("赋值运算符")
	assignmentOperators()
	fmt.Println("其他运算符")
	otherOperators()
}
