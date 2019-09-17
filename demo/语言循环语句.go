package main

import "fmt"

/*
// 普通循环语句
func main() {

	var b int = 15
	var a int

	numbers := [6]int{1, 2, 3, 5}

	// for 循环
	for a := 0; a < 10; a++ {
		fmt.Printf("a 的值为：%d\n", a)
	}

	for a < b {
		a++
		fmt.Printf("a 的值为：%d\n", a)
	}

	for i, x:= range numbers {
		fmt.Printf("第%d 位 x 的值 = %d\n", i, x)
	}
}

*/



/*
// 语言循环嵌套
// eg: 以下实例使用循环嵌套来输出2到100间的素数
func main() {
	// 定义局部变量
	var i, j int

	for i = 2; i < 100; i++ {
		for j = 2; j <= (i/j); j++ {
			if(i%j ==0) {
				break;  // 如果发现因子， 则不是素数
			}
		}
		if(j > (i/j)) {
			fmt.Printf("%d 是素数\n", i);
		}
	}
}
*/



/*
// 语言break语句
func main() {
	// 定义局部变量
	var a int = 10

	// for 循环
	for a < 20 {
		fmt.Printf("a 的值为 ：%d\n", a);
		a++;
		if a > 15 {
			// 使用break 语句跳出循环
			break;
		}
	}
}
*/


/*
// 语言continue语句
func main() {
	// 定义局部变量
	var a int = 10

	// for循环
	for a < 20 {
		if a == 15 {
			// 跳过此次循环
			a += 1;
			continue
		}
		fmt.Printf("a 的值为 %d\n", a);
		a++;
	}
}
*/




/*
// 语言goto语句
//eg: 在变量a等于15的时候跳过本次循环并回到循环的开始语句LOOP处：
func main() {
	// 定义局部变量
	var a int = 10

	// 循环
	LOOP: for a < 20 {
		if  a == 15 {
			// 跳过迭代
			fmt.Printf("由于特殊原因,本次取值被跳过... 当前a的值为： 【%d】\n", a)
			a += 1;
			goto LOOP
		}
		fmt.Printf("a 的值为 ： %d\n", a)
		a++;
	}
}
*/

// example
// 打印九九乘法表
func main() {
	// print9x()
	gotoTag()
}

//嵌套for循环打印九九乘法表
func print9x() {
	for m:= 1; m < 10; m++ {
		for n := 1; n < m; n++ {
			fmt.Printf("%dx%d=%d", n, m, m*n)
		}
		fmt.Printf("")
	}
}

// for循环配合goto打印九九乘法表
func gotoTag() {
	for m:= 1; m < 10; m++ {
		n:= 1
		LOOP: if n <= m {
			fmt.Printf("%dx%d=%d ", n, m, m*n)
			n++
			goto LOOP
		}else {
			fmt.Printf("\n")
		}
	}
}




