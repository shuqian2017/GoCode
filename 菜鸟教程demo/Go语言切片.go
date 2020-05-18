/*
@project: GoLand
@author: fke
@file: Go语言切片.go
@time: 2019-09-17 18:15:15
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

/*
//  make([]type, length, capacity)  也可以指定容量，其中capacity为可选参数
func main() {
    var numbers = make([]int,3,5)

    printSlice(numbers)
}

func printSlice(x []int) {
    fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
*/




/*
// 空切片(nil)
func main() {
    var numbers []int

    printSlice(numbers)
    if(numbers == nil) {
        fmt.Printf("切片是空的")
    }
}

func printSlice(x []int)  {
    fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
*/





/*
//切片截取
func main()  {
    //创建切片
    numbers := []int{0,1,2,3,4,5,6,7,8}
    printSlice(numbers)

    //打印原始切片
    fmt.Println("numbers ==", numbers)
    // 打印子切片索引1(包含)到索引4(不包含)
    fmt.Println("numbers[1:4] ==", numbers[1:4])
    // 默认下限为0
    fmt.Println("numbers[:3] ==", numbers[:3])
    // 默认上限为len(s)
    fmt.Println("numbers[4:] ==", numbers[4:])

    numbers1 := make([]int,0,5)
    printSlice(numbers1)

    // 打印子切片从索引[0,2}
    number2 := numbers[:2]
    printSlice(number2)

    //打印子切片从索引[2,5)
    number3 := numbers[2:5]
    printSlice(number3)
}

func printSlice(x []int) {
    fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
*/






//append()和copy()函数
// 如果想增加切片的容量，我们必须创建一个新的切片并把原来切片的内容都拷贝过来
func main()  {
    var numbers []int
    printSlice(numbers)

    //允许追加空切片
    numbers = append(numbers, 0)
    printSlice(numbers)

    // 向切片添加新的元素
    numbers = append(numbers, 1)
    printSlice(numbers)

    // 同时添加多个元素
    numbers = append(numbers, 2,3,4)
    printSlice(numbers)

    // 创建切片numbers1 是之前切片的两倍容量
    numbers1 := make([]int, len(numbers), (cap(numbers))*2)

    // 拷贝numbers 的内容到numbers1
    copy(numbers1, numbers)
    printSlice(numbers1)
}

func printSlice(x []int)  {
    fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

