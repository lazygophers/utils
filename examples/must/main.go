package main

import (
	"fmt"
	"os"

	"github.com/lazygophers/utils"
)

func main() {
	// 示例1: Must - 处理 (value, error) 返回值
	content := utils.Must(os.ReadFile("example.txt"))
	fmt.Printf("文件内容: %s\n", content)

	// 示例2: MustOk - 处理 (value, bool) 返回值
	m := map[string]int{"key": 42}
	value := utils.MustOk(m["key"])
	fmt.Printf("Map值: %d\n", value)

	// 示例3: MustSuccess - 仅检查错误
	utils.MustSuccess(os.MkdirAll("temp", 0755))
	fmt.Println("目录创建成功")

	// 示例4: Ignore - 忽略错误或第二个返回值
	result := utils.Ignore(someFunction())
	fmt.Printf("结果: %s\n", result)
}

func someFunction() (string, error) {
	return "Hello, World!", nil
}
