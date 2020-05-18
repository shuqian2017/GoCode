/**
@project: GoCode
@author: fke
@file: 语言条件语句.go
@time: 2019-05-31 12:39:51
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

package main



/*
// if else 嵌套
func main() {
	// 定义局部变量
	a, b := 100, 200
	if a == 100 {
		if b == 200 {
			fmt.Printf("a 的值为100, b的值为 200\n")
		}
	}
	fmt.Printf("a 值为: %d\n", a)
	fmt.Printf("b 值为：%d\n", b)
}
*/


/*
// switch 语句
func main() {
	// 定义局部变量
	var grade string = "B"
	var name string
	var marks int
	fmt.Scanf("%s  %d", &name, &marks)

	switch marks {
		case 90: grade = "A"
		case 80: grade = "B"
		case 50, 60, 70: grade = "C"
		default: grade = "D"
	}

	switch {
		case grade == "A":
			fmt.Printf("优秀！\n")
		case grade == "B", grade == "C":
			fmt.Printf("良好\n")
		case grade == "D":
			fmt.Printf("及格\n")
		case grade == "F":
			fmt.Printf("不及格\n")
		default:
			fmt.Printf("差\n")
	}
	fmt.Printf("你的等级是 %s\n", grade)
}
*/


// switch 语句还可以用于type-switch来判断某个interface变量中实际存储变量类型


/*
// select select 是 Go 中的一个控制结构，类似于用于通信的 switch 语句
//如果有 default 子句，则执行该语句。
//如果没有 default 子句，select 将阻塞，直到某个通信可以运行；Go 不会重新对 channel 或值进行求值。
import "fmt"

func main() {
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
		case i1 = <-c1:
			fmt.Printf("received", i1, "from c1\n")
		case c2 <- i2:
			fmt.Printf("send", i2, "to c2\n")
		case i3, ok := (<-c3): // same as: i3, ok := <-c3
			if ok {
				fmt.Printf("received", i3, "from c3\n")
			} else {
				fmt.Printf("c3 is closed\n")
			}
		default:
			fmt.Printf("no communication\n")
	}
}
*/

// for
/*
func main() {
	for i := 0; i < 5; i++ {
		println(i)
	}

	for i := 4; i >= 0; i-- {
		println(i)
	}
}
*/



/*
// 在迭代遍历时，for...range除元素外，还可以返回索引
func main() {
	x := [] int{100, 101, 102}

	for i, n := range x {
		fmt.Println(i, ":", n)
	}
}
*/

/*
//函数可以定义多个返回值，甚至对其命名
import "errors"

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}

	return a / b, nil     // nil 空切片
}

func main() {
	a, b := 10, 2		// 定义多个变量
	c, err := div(a, b)	// 接手多个返回值

	fmt.Println(c, err)
}
*/

/*
//函数是第一类型， 可作为参数或返回值
func test (x int) func() {
	return func() {
		fmt.Println(x)
	}
}

func main() {
	x := 100
	f := test(x)
	f()
}
*/


/*
//用defer定义延迟调用，无论函数是否出错，它都确保结束前被调用
func test(a, b int) {
	defer println("dispose...")   	// 常用来释放资源，解除锁定，或执行清理操作
	println(a / b)							// 可定义多个defer， 按FILO顺序执行
}

func main() {
	test(10, 0)
}
*/



// 数据
//切片(slice) 可实现类似动态数组的功能
import "fmt"

func main() {
	x := make([] int, 0, 5)   // []int 标识其元素类型为int的切片  len()获取切片的长度， cap()测量切片的最大容量

	for i := 0; i < 8; i++ {
		x = append(x, i)	// 追加数据，当超出容量限制时，自动分配更大的内存空间
	}

	fmt.Println(x)
	fmt.Println(x[5:7])   // 类似Python的切片
}






