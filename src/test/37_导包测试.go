package main

import (
	_ "test/lib1" // 匿名
	//	lib "test/lib2" // 别名
	. "test/lib2" // 通过"."导包，直接输入包中的方法名
)

func main() {
	//lib1.Lib1Test()
	//	lib.Lib2Test()
	Lib2Test()
}
