package main

import (
	"context"
	"fmt"
	"os"
	"strings"
)

func main() {
	// e1.1 输出命令名字
	fmt.Println(strings.Join(os.Args[0:], " "))
	fmt.Println(" ========================== ")
	// e1.2 输出参数和索引和值 每行一个
	for i, arg := range os.Args[1:] {
		fmt.Println(i, " ", arg)
	}
	// e1.3 程序执行时间差异
	context.TODO()

}
