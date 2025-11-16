package cli

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/Ghost-Xiao/ghost-lang/internal/evaluator"
	"github.com/Ghost-Xiao/ghost-lang/internal/frame"
	"github.com/Ghost-Xiao/ghost-lang/internal/lexer"
	"github.com/Ghost-Xiao/ghost-lang/internal/object"
	"github.com/Ghost-Xiao/ghost-lang/internal/parser"
)

// StartREPL 启动repl，提供即时代码执行环境
func StartREPL() {
	// 捕获中断信号 (Ctrl+C)，跨平台处理
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	// 添加退出标志
	var exitRequested bool
	go func() {
		// 等待中断信号
		<-sigChan
		// 首先设置退出标志
		exitRequested = true
		// 打印退出信息
		printInfo("\nBye!")
		// 刷新标准输出缓冲区
		_ = os.Stdout.Sync()
	}()
	// 显示版本和欢迎信息
	printInfo(fmt.Sprintf("ghost-lang %s | %s/%s | built %s.", Version, Platform, Arch, BuildTime))
	printInfo("Welcome to the Ghost REPL.")
	printInfo("Press Ctrl+C to exit.")
	// 创建解释器环境
	env := &object.Environment{
		Store: make(map[string]*object.Symbol),
		Outer: nil,
	}
	// 加载内置函数
	for name, builtin := range object.Builtins {
		env.Store[name] = &object.Symbol{
			Name:    name,
			Value:   builtin,
			IsConst: true,
		}
	}
	// 创建调用栈
	f := &frame.Frame{
		FuncName: "<stdin>",
		PosStart: nil,
		PosEnd:   nil,
		Parent:   nil,
	}
	scanner := bufio.NewScanner(os.Stdin)
	// 交互式输入循环
	for !exitRequested {
		fmt.Print(">>> ")
		// 刷新标准输出缓冲区
		_ = os.Stdout.Sync()
		// 多行输入处理
		var lines []string
		scannerOK := false
		// 重复读取输入
		for scanner.Scan() && !exitRequested {
			lines = append(lines, scanner.Text())
			source := strings.ReplaceAll(strings.Join(lines, "\n"), "\t", "    ")
			// 尝试解析，词法分析
			l := lexer.NewLexer("<stdin>", source)
			// 语法分析
			p, err := parser.NewParser(l)
			if err != nil {
				if !shouldContinue(err) {
					printError(err)
					scannerOK = true
					break
				} else {
					fmt.Print("... ")
					// 刷新标准输出缓冲区
					_ = os.Stdout.Sync()
					continue
				}
			}
			program := p.ParseProgram()
			if p.Err != nil {
				if !shouldContinue(p.Err) {
					var syntaxError *parser.SyntaxError
					ok := errors.As(p.Err, &syntaxError)
					if ok && syntaxError.Message == "expected \"SEMICOLON\", but got \"EOF\"." {
						// 重试解析为表达式
						l2 := lexer.NewLexer("<stdin>", source)
						p2, err2 := parser.NewParser(l2)
						if err2 != nil {
							if !shouldContinue(err2) {
								printError(err2)
								scannerOK = true
								break
							} else {
								fmt.Print("... ")
								// 刷新标准输出缓冲区
								_ = os.Stdout.Sync()
								continue
							}
						}
						expr := p2.ParseExpression(parser.LOWEST)
						if p2.Err != nil {
							if !shouldContinue(p2.Err) {
								printError(p2.Err)
								scannerOK = true
								break
							} else {
								fmt.Print("... ")
								// 刷新标准输出缓冲区
								_ = os.Stdout.Sync()
								continue
							}
						}
						// 执行表达式并输出结果
						e := evaluator.NewEvaluator(f)
						ret := e.Eval(expr, env)
						if e.Err != nil {
							printError(e.Err)
							scannerOK = true
							break
						}
						if ret != nil {
							fmt.Print("::: ")
							fmt.Println(ret)
							// 刷新标准输出缓冲区
							_ = os.Stdout.Sync()
						}
						scannerOK = true
						break
					}
					printError(p.Err)
					scannerOK = true
					break
				} else {
					fmt.Print("... ")
					// 刷新标准输出缓冲区
					_ = os.Stdout.Sync()
					continue
				}
			}
			// 执行程序
			e := evaluator.NewEvaluator(f)
			res := e.Eval(program, env)
			if e.Err != nil {
				printError(e.Err)
				scannerOK = true
				break
			}
			if res != nil {
				fmt.Print("::: ")
				fmt.Println(res)
				// 刷新标准输出缓冲区
				_ = os.Stdout.Sync()
			}
			scannerOK = true
			break
		}
		if !scannerOK && !exitRequested {
			if err := scanner.Err(); err != nil {
				printError("ghost-lang: failed to read input.")
			} else {
				// 打印退出信息
				printInfo("\nBye!")
				// 刷新标准输出缓冲区
				_ = os.Stdout.Sync()
				return
			}
		}
	}
	// 确保退出前刷新缓冲区
	_ = os.Stdout.Sync()
}

// 判断是否需要继续解析
func shouldContinue(err error) bool {
	msg := err.Error()
	if strings.Contains(msg, "expected \"SEMICOLON\", but got \"EOF\".") {
		return false
	}
	return strings.Contains(msg, "\"*/\" is expected.") ||
		strings.Contains(msg, "unterminated string literal.") ||
		strings.Contains(msg, "unexpected \"EOF\".") ||
		strings.Contains(msg, "but got \"EOF\"")
}
