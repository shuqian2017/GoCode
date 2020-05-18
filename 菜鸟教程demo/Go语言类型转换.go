/*
@project: GoLand
@author: fke
@file: Go语言类型转换.go
@time: 2019-09-19 14:37:19
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
import "fmt"

// 将实例中的整型转化为浮点型，并计算结果
func main() {
    var sum int = 17
    var count int = 5
    var mean float32
    mean = float32(sum)/float32(count)
    fmt.Printf("mean 的值为： %f\n", mean)
}
