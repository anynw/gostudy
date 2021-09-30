package main

import "fmt"

/**
1.Go 语言中数组可以存储同一类型的数据，但在结构体中我们可以为不同项定义不同的数据类型。
2.结构体是由一系列具有相同类型或不同类型的数据构成的数据集合。
3.类似Java的实体类
*/

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	// 创建一个新的结构体
	fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})

	// 使用key-value 的方式
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407})

	fmt.Println(Books{
		title:   "Go 语言",
		author:  "www.runoob.com",
		subject: "Go 语言教程",
		book_id: 6495407})

	// 忽略字段为0或者空
	fmt.Println(Books{title: "Go 语言", subject: "Go 语言教程", book_id: 6495407})

	// 访问结构体成员

	var book1 Books
	book1.title = "book1"
	book1.author = "wuhp"
	book1.subject = "数学"
	book1.book_id = 123

	fmt.Printf("title = %x, author = %x, subject = %x, book_id = %d\n", book1.title, book1.author, book1.subject, book1.book_id)

}
