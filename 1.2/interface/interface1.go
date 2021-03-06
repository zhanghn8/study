package main

import (
	"fmt"
)

func main() {
	var a interface{} = func(a int) string {
		return fmt.Sprintf("d:% d", a)
	}
	switch b := a.(type) { // 判断变量b 是类型转换后的结果
	case nil:
		println("nil")
	case *int:
		println(*b)
	case func(int2 int) string:
		println(b(88))
	case fmt.Stringer:
		fmt.Println(b)
	default:
		println("unknown")
	}
}
