/*
@project: GoLand
@author: fke
@file: 语言数组.go
@time: 2019-09-16 17:21:27
 code is far away from bugs with the god animal protecting
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
*/
package main

import (
    "fmt"
)

/*
func main() {
    var n [10] int   // n 是一个长度为10的数组
    var i, j int

    // 为数组n初始化元素
    for i = 0; i < 10; i++ {
        n[i] = i + 100
    }

    // 输出每个数组元素的值
    for j = 0; j < 10; j++ {
        fmt.Printf("Element[%d] = %d\n", j, n[j])
    }
}
*/



/*
//二维数组
func main() {
    // 数组 - 5 行 2 列
    var a = [5][2] int { {0,0}, {1,2}, {2,4}, {3,6}, {4,8} }
    var i, j int

    //输出数组元素
    for i = 0; i < 5; i++ {
        for j = 0; j < 2; j++ {
            fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
        }
    }
}
*/



/*
//eg: 多维数组初始化或赋值时需要注意Go语法规范
func main() {
    var a = [3][5] int { {1,2,3,4,5}, {0,9,8,7,6}, {3,4,5,6,7} }
    var i, j int
    for i = 0; i < 3; i++ {
        for j = 0; j < 5; j++ {
            fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
        }
    }
}
*/




/*
//向函数传递数组
func main()  {
    // 数组长度为5
    var balance = [5]int {1000, 2, 3, 17, 50}
    var avg float32

    //数组作为参数传递给函数
    avg = getAverage( balance, 5)

    // 输出返回的平均值
    fmt.Printf("平均值：%f", avg)
}

func getAverage(arr [5] int, size int) float32 {
    var i, sum int
    var avg float32

    for i = 0; i < size; i++ {
        sum += arr[i]
    }
    avg = float32(sum) / float32(size)
    return avg;
}
*/




//浮点数计算的时候，有时候会有偏差，可以设置固定精度
func main()  {
    a := 1690
    b := 1700
    c := a * b
    fmt.Println(c)
    fmt.Println(float64(c) / 1000000)   // 显示
}
