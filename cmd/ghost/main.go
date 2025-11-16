// Package main 是Ghost语言解释器的入口程序，负责命令行参数解析、文件执行和交互式Studio模式
package main

import "github.com/Ghost-Xiao/ghost-lang/internal/cli"

// main 程序入口函数，处理命令行参数并分发到相应模式
func main() {
	cli.Run()
}
