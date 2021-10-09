package main

import "fmt"

// 定义结构体
type DivideError struct {
	dividee int
	divider int
}

// 定义错误接口
func (de *DivideError) Error() string {
	strFormat := `Cannot proceed,this divider is zoro
		dividee: %d
		divicer: 0
	`
	return fmt.Sprintf(strFormat, de.dividee)
}

// 定义int类型除法运算函数
func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg := dData.Error()
		return 0, errorMsg
	} else {
		return varDividee / varDivider, ""
	}

}

func main() {
	// 正常情况
	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}
	// 当除数为零的时候会返回错误信息
	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}
}
