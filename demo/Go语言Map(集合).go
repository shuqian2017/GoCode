/*
@project: GoLand
@author: fke
@file: Go语言Map(集合).go
@time: 2019-09-19 10:02:09
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
//Map是一种无序的键值对集合，通过key来快速检索数据，key类似于索引，指向数据的值
func main() {
    var countryCapitaMap map[string]string   // 创建集合
    countryCapitaMap = make(map[string]string)

    // map插入key - value对
    countryCapitaMap ["France"] = "巴黎"
    countryCapitaMap ["Italy"] = "罗马"
    countryCapitaMap ["Japan"] = "东京"
    countryCapitaMap ["India"] = "新德里"

    // 使用键输出地图值
    for country := range  countryCapitaMap {
        fmt.Println(country, "首都是", countryCapitaMap [country])
    }

    // 查看元素在集合中是否存在
    capital, ok := countryCapitaMap ["American"]
    if(ok) {
        fmt.Println("American 的首都是", capital)
    } else {
        fmt.Println("American 的首都不存在")
    }
}
*/





/*
//delete()函数
//delete()函数用于删除集合的元素，参数为map和其对应的key
func main()  {
    // 创建map
    countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}
    fmt.Println("原始地图:")

    // 打印地图
    for country := range countryCapitalMap {
        fmt.Println(country, "首都是", countryCapitalMap [country])
    }

    // 删除元素
    delete(countryCapitalMap, "France")
    fmt.Println("\n法国条目被删除")
    fmt.Println("删除元素后地图\n")

    // 打印地图
    for country := range countryCapitalMap {
        fmt.Println(country, "首都是", countryCapitalMap [country])
    }
}
*/





// eg: 基于go实现简单HashMap，暂未做key值的校验
type HashMap struct {
    key string
    value string
    hashCode int
    next *HashMap
}

var table [16](*HashMap)


func initTable()  {
    for i := range table {
        table[i] = &HashMap{"", "", i, nil}
    }
}

func getInstance() [16](*HashMap) {
    if(table[0] == nil) {
        fmt.Println("列表为空，需要先进行初始化操作...")
        initTable()
    }
    fmt.Println("table -> ", table)
    return table
}

func genHashCode(k string) int {
    if len(k) == 0 {
        return 0
    }

    var hashCode int = 0
    var lastIndex int = len(k) - 1
    for i := range k {
        if i == lastIndex {
            hashCode += int(k[i])
            break
        }
        hashCode += (hashCode + int(k[i]))*31
    }
    return hashCode
}

func indexTable(hashCode int) int {
    return hashCode%16
}

func indexNode(hashCode int) int {
    return hashCode>>4
}

func put(k string, v string) string {
    var hashCode =genHashCode(k)
    var thisNode = HashMap{k, v, hashCode, nil}

    var tableIndex = indexTable(hashCode)
    var nodeIndex = indexNode(hashCode)

    var headPtr [16](*HashMap) = getInstance()
    var headNode = headPtr[tableIndex]

    if (*headNode).key == "" {
        *headNode = thisNode
        return ""
    }

    var lastNode *HashMap = headNode
    var nextNode *HashMap = (*headNode).next

    for nextNode != nil && (indexNode((*nextNode).hashCode) < nodeIndex) {
        lastNode = nextNode
        nextNode = (*nextNode).next
    }

    if (*lastNode).hashCode == thisNode.hashCode {
        var oldValue string = lastNode.value
        lastNode.value = thisNode.value
        return oldValue
    }

    if lastNode.hashCode < thisNode.hashCode {
        thisNode.next =nextNode
    }
    return ""
}

func get(k string) string {
    var hashCode = genHashCode(k)
    var tableIndex = indexTable(hashCode)

    var headPtr [16](*HashMap) = getInstance()
    var node *HashMap = headPtr[tableIndex]

    if (*node).key == k {
        return (*node).value
    }

    for (*node).next != nil {
        if k == (*node).key {
            return (*node).value
        }
        node = (*node).next
    }
    return ""
}

//example
func main() {
    getInstance()
    put("a", "a_put")
    put("b", "b_put")
    fmt.Println(get("a"))
    fmt.Println(get("b"))
    put("p", "p_put")
    fmt.Println(get("p"))
}
