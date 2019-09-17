/**
@project: GoCode
@author: fke
@file: 语言运算符.go
@time: 2019-05-30 10:31:04
# code is far away from bugs with the god animal protecting
    I love animals. They taste delicious.
              ┏┓      ┏┓
            ┏┛┻━━━┛┻┓
            ┃      ☃      ┃
            ┃  ┳┛  ┗┳  ┃
            ┃      ┻      ┃
            ┗━┓      ┏━┛
                ┃      ┗━━━┓
                ┃  神兽保佑    ┣┓
                ┃　永无BUG。   ┏┛
                ┗┓┓┏━┳┓┏┛
                  ┃┫┫  ┃┫┫
                  ┗┻┛  ┗┻┛
**/

/*
  算术运算符
  关系运算符
  逻辑运算符
  位运算符
  赋值运算符
  其他运算符
*/

package main

import "fmt"

/*
//  算术运算
func main() {
	a, b := 21, 10
	var c int

	c = a + b
	fmt.Printf("第一行 -c 的值为 %d\n", c)
	c = a - b
	fmt.Printf("第二行 -c 的值为 %d\n", c)
	c = a * b
	fmt.Printf("第三行 -c 的值为 %d\n", c)
	c = a / b
	fmt.Printf("第四行 -c 的值为 %d\n", c)
	c = a % b
	fmt.Printf("第五行 -c 的值为 %d\n", c)
	a++
	fmt.Printf("第六行 -a 的值为 %d\n", a)
	a =21   	// 为了方便测试这里重新赋值为21
	a--
	fmt.Printf("第七行 -a 的值为 %d\n", a)
}
*/


/*
// 关系运算符
func main() {
	a, b := 21, 10

	if ( a == b) {
		fmt.Printf("第一行 - a 等于 b\n")
	} else {
		fmt.Printf("第一行 - a 不等于 b\n")
	}
	if ( a < b ) {
		fmt.Printf("第二行 - a 小于 b\n")
	} else {
		fmt.Printf("第二行 - a 不小于 b\n")
	}

	if ( a > b ) {
		fmt.Printf("第三行 - a 大于 b\n")
	} else {
		fmt.Printf("第三行 - a 不大于 b\n")
	}

	// Lets change value of a and b
	a, b = 5, 20
	if ( a <= b ) {
		fmt.Printf("第四行 - a 小于等于 b\n")
	}
	if ( b >= a) {
		fmt.Printf("第五行 - b 大于等于 a\n")
	}
}

*/

/*
// 逻辑运算
func main() {
	a, b := true, false
	if ( a && b) {
		fmt.Printf("第一行 - 条件为 true\n")
	}
	if ( a || b ) {
		fmt.Printf("第二行 - 条件为 true\n")
	}

	// 修改 a 和 b 的值
	a, b = false, true
	if ( a && b ) {
		fmt.Printf("第三行 - 条件为 true\n")
	} else {
		fmt.Printf("第三行 - 条件为 false\n")
	}
	if ( !(a && b)) {
		fmt.Printf("第四行 - 条件为 true\n")
	}
}
*/


// 位运算
// TODO

/*
// 赋值运算
func main() {
	var a int = 21
	var c int

	c = a
	fmt.Printf("第1行 -= 运算符实例, c 值为 = %d\n", c)

	c += a
	fmt.Printf("第2行 - += 运算符实例, c 值为 = %d\n", c)

	c -= a
	fmt.Printf("第3行 - -= 运算符实例, c 值为 = %d\n", c)

	c *= a
	fmt.Printf("第4行 - *= 运算符实例, c 值为 = %d\n", c)

	c /= a
	fmt.Printf("第5行 - /= 运算符实例, c 值为 = %d\n", c)

	c = 200
	c <<= 2
	fmt.Printf("第6行 - <<= 运算符实例, c 值为 = %d\n", c)

	c >>= 2
	fmt.Printf("第7行 - >>= 运算符实例, c 值为 = %d\n", c)

	c &= 2
	fmt.Printf("第8行 - &= 运算符实例, c 值为 = %d\n", c)

	c ^= 2
	fmt.Printf("第9行 - ^= 运算符实例, c 值为 = %d\n", c)

	c |= 2
	fmt.Printf("第10行 - |= 运算符实例, c 值为 = %d\n", c)
}
*/

/*
// 其他运算符
func main() {
	var a int = 4
	var b int32
	var c float32
	var ptr *int

	// 运算符实例
	fmt.Printf("第1行 - a 变量类型为 = %T\n", a)
	fmt.Printf("第2行 - b 变量类型为 = %T\n", b)
	fmt.Printf("第3行 - c 变量类型为 = %T\n", c)

	// & 和 * 运算符实例
	ptr = &a    //  'ptr' 包含了'a' 变量的地址
	fmt.Printf("a 的值为 %d\n", a)
	fmt.Printf("*ptr 为 %d\n", *ptr)
}
*/


// 运算符优先级
func main() {
	a, b, c, d := 20, 10, 15, 5
	var e int

	e = (a + b) * c / d
	fmt.Printf("(a + b) * a / d = %d\n", e)

	e = ((a + b) * c) / d
	fmt.Printf("((a + b) * c) / d = %d\n", e)

	e = a + (b * c) / d
	fmt.Printf("a + (b * c) / d = %d\n", e)
}