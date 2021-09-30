package main

import "fmt"

// 不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对
func main() {
	//声明变量，默认map是nil
	//var map_variable map[key_data_type]value_data_type
	// 使用make函数
	//map_variable := make(map[key_data_type]value_data_type)

	var map1 map[string]string
	map1 = make(map[string]string)
	map1["a"] = "11"
	map1["b"] = "22"
	map1["c"] = "33"

	for key := range map1 {
		fmt.Println("key:", key, "value:", map1[key])
	}

	// delete函数
	delete(map1, "a")
	fmt.Println("删除元素a")
	for key := range map1 {
		fmt.Println("key:", key, "value:", map1[key])
	}

}
