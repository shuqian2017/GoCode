/*
@project: GoLand
@author: fke
@file: Go语言范围.go
@time: 2019-09-18 18:13:07
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


// Go语言中range关键字用于for循环中迭代数组，切片，通道或集合的元素
func main() {
    // 这是我们使用range去求一个slice的和。使用数组跟这个很类似
    nums := []int{2,3,4}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)

    // 在数组上使用range将传入index和值两个变量。上面我们不需要使用该元素的序号，所以我们使用空白符"_"省略了。有时候我们需要知道它的索引
    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }

    // range 也可以用map的键值对上
    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

    //range也可以用来枚举unicode字符串，第一个参数是字符的索引，第二个是字符unicode的值本身
    for i, c := range "go" {
        fmt.Println(i, c)
    }
}

