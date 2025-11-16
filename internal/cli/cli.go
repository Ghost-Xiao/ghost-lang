package cli

import (
	"flag"
	"os"
)

func Run() {
	// 定义命令行标志
	replMode := flag.Bool("r", false, "REPL")
	versionMode := flag.Bool("v", false, "Version")
	helpMode := flag.Bool("h", false, "Help")

	// 禁用自动退出
	flag.CommandLine.Init(flag.CommandLine.Name(), flag.ContinueOnError)
	originalStderr := os.Stderr
	nullFile, _ := os.OpenFile(os.DevNull, os.O_WRONLY, 0666)
	os.Stderr = nullFile
	defer func() {
		os.Stderr = originalStderr
		_ = nullFile.Close()
	}()

	// 执行解析
	flag.Parse()

	// 解析全局flag
	if *replMode {
		StartREPL()
		return
	}
	if *versionMode {
		PrintVersion()
		return
	}
	if *helpMode {
		PrintHelp()
		return
	}

	// 剩余未解析的参数
	args := flag.Args()
	// 参数验证：未指定任何模式且无输入文件时显示错误
	if len(args) == 0 {
		printError("ghost-lang: invalid command line arguments.")
		PrintHelp()
		return
	}

	// 分发子命令
	command := args[0]
	switch command {
	case "repl":
		// 启动REPL
		StartREPL()
		return
	case "run":
		// 运行文件
		RunFile(args[1])
		return
	default:
		// 显示错误
		printError("ghost-lang: unknown command.")
		PrintHelp()
		return
	}
}
