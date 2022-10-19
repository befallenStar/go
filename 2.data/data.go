package main

import (
	"fmt"
)

func main() {
	// 1、布尔型：常量true或者false
	// var b bool = false
	// 2、数字类型：包括整形int和浮点型float，支持uint{8, 16, 32, 64}，int{8, 16, 32, 64}，float{32, 64}，complex{32, 64}
	// 其他数字类型包括byte(uint8)、rune(uint16)、uint(uint32或uint64)、int(uint)、uintptr(指向无符号整形的指针)
	// var b int = 100
	// 3、字符串类型：字符连接的字符序列
	// var b string = "123"
	// 4、派生类型：
	// 4.1、指针类型（Pointer）
	// var a int = 1
	// var p *int = &a
	// 4.2、数组类型
	// var arr = [10]int{1, 2, 3}
	// 4.3、结构体类型(struct)
	// type a struct {
	// 	member1 int
	// 	member2 string
	// }
	// b := a{1, "1"}
	// b.member1 = 2
	// 4.4、联合体类型 (union)?
	// 4.5、函数类型
	// func a(){}
	// 4.6、切片类型：维护一个指针属性，指向对应的底层数组，本质是[length/capacity]0xADDR
	// 和数组的不同是[]内没有大小或者...，
	// var s = []int{3,3}
	// var s = make([]int, 3, 5)
	// 对切片扩容超出capacity后会创建新的底层数组，新的切片指向新的数组
	// 4.7、接口类型（interface）
	// type in interface {
	// 	in()
	// }
	// 4.8、Map 类型
	// var a map[string]int
	// 使用map前必须先make，不然会报空指针错
	// a = make(map[string]int)
	// 4.9、Channel 类型
	// var c chan int
	// Channel本质上是一个固定大小的队列，使用前make初始化大小
	// c = make(chan int, 3)
	fmt.Println()
}
