package cli

import (
	"fmt"
	"os"
)

// printError 打印带红色高亮的错误信息并刷新标准输出缓冲区
//
// 参数:
//
//	message - 错误文本内容
func printError(message any) {
	fmt.Printf("\033[31m%s\033[0m\n", message)
	// 刷新标准输出缓冲区
	_ = os.Stdout.Sync()
}

// printInfo 打印带蓝色高亮的信息文本并刷新标准输出缓冲区
//
// 参数:
//
//	message - 信息文本内容
func printInfo(message string) {
	fmt.Printf("\033[34m%s\033[0m\n", message)
	// 刷新标准输出缓冲区
	_ = os.Stdout.Sync()
}
