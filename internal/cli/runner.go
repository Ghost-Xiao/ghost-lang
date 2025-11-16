package cli

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/Ghost-Xiao/ghost-lang/internal/evaluator"
	"github.com/Ghost-Xiao/ghost-lang/internal/frame"
	"github.com/Ghost-Xiao/ghost-lang/internal/lexer"
	"github.com/Ghost-Xiao/ghost-lang/internal/object"
	"github.com/Ghost-Xiao/ghost-lang/internal/parser"
)

// RunFile 执行指定的.gh文件
//
// 参数:
//
//	fileName - 要执行的文件路径
func RunFile(fileName string) {
	// 捕获中断信号 (Ctrl+C)，跨平台处理
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		// 等待中断信号
		<-sigChan
		// 打印退出信息
		printInfo("\nExecution stopped by user.")
		// 刷新标准输出缓冲区
		_ = os.Stdout.Sync()
		// 退出
		os.Exit(0)
	}()

	// 验证文件扩展名
	slice := strings.Split(fileName, ".")
	if (len(slice) > 1 && slice[len(slice)-1] != "gh") || len(slice) <= 1 {
		printError(fmt.Sprintf("ghost-lang: invalid file extension: \"%s\".", fileName))
		return
	}

	// 读取文件内容
	data, err := os.ReadFile(fileName)
	if err != nil {
		printError(fmt.Sprintf("ghost-lang: file not found: \"%s\".", fileName))
		return
	}

	// 获取绝对路径
	absPath, err := filepath.Abs(fileName)
	if err != nil {
		printError(fmt.Sprintf("ghost-lang: failed to resolve absolute path: \"%s\".", fileName))
		return
	}

	// 显示版本和文件信息
	printInfo(fmt.Sprintf("ghost-lang %s | %s/%s | built %s.", Version, Platform, Arch, BuildTime))
	printInfo(fmt.Sprintf("Running file \"%s\".", absPath))

	// 记录开始时间
	startTime := time.Now()

	// 执行文件内容
	code := strings.ReplaceAll(string(data), "\t", "    ")
	baseName := filepath.Base(absPath)
	l := lexer.NewLexer(baseName, code)
	p, err2 := parser.NewParser(l)
	if err2 != nil {
		printError(err2)
		return
	}
	program := p.ParseProgram()
	if p.Err != nil {
		printError(p.Err)
		return
	}
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
	f := &frame.Frame{
		FuncName: baseName,
		PosStart: nil,
		PosEnd:   nil,
		Parent:   nil,
	}
	e := evaluator.NewEvaluator(f)
	e.Eval(program, env)
	if e.Err != nil {
		printError(e.Err)
		return
	}

	// 记录结束时间并计算执行时间
	endTime := time.Now()
	executionTime := endTime.Sub(startTime)

	// 输出执行时间
	// 根据时间长度决定是否显示组合格式
	if executionTime.Nanoseconds() == 0 {
		// 时间为0，直接显示0 ns
		printInfo("Execution time: 0 ns")
	} else if executionTime < time.Second {
		// 毫秒、微秒、纳秒级别，显示秒数和换算单位
		printInfo(fmt.Sprintf("Execution time: %.9f s (%s)", executionTime.Seconds(), formatDuration(executionTime)))
	} else if executionTime < 60*time.Second {
		// 1-60秒之间，只显示秒数
		printInfo(fmt.Sprintf("Execution time: %.9f s", executionTime.Seconds()))
	} else {
		// 60秒及以上，显示秒数和组合格式
		printInfo(fmt.Sprintf("Execution time: %.9f s (%s)", executionTime.Seconds(), formatDuration(executionTime)))
	}
}

// formatDuration 根据时间长短自动选择合适的单位格式化持续时间
func formatDuration(d time.Duration) string {
	// 定义时间单位常量
	hour := time.Hour
	day := 24 * hour
	week := 7 * day   // 一周7天
	month := 30 * day // 近似值
	year := 365 * day // 近似值

	switch {
	case d == 0:
		return "0 ns"
	case d < time.Microsecond:
		return fmt.Sprintf("%d ns", d.Nanoseconds())
	case d < time.Millisecond:
		return fmt.Sprintf("%d µs", d.Microseconds())
	case d < time.Second:
		return fmt.Sprintf("%d ms", d.Milliseconds())
	case d < hour:
		// 分钟和秒组合
		minutes := int(d.Minutes())
		seconds := int(d.Seconds()) % 60
		return fmt.Sprintf("%d min %d s", minutes, seconds)
	case d < day:
		// 小时和分钟组合
		hours := int(d.Hours())
		minutes := int(d.Minutes()) % 60
		return fmt.Sprintf("%d h %d min", hours, minutes)
	case d < week:
		// 天和小时组合
		days := int(d.Hours() / 24)
		hours := int(d.Hours()) % 24
		return fmt.Sprintf("%d d %d h", days, hours)
	case d < month:
		// 周和天组合
		weeks := int(d.Hours() / (24 * 7))
		days := int(d.Hours()/24) % 7
		return fmt.Sprintf("%d wk %d d", weeks, days)
	case d < year:
		// 月和天组合 (每月按30天近似)
		months := int(d.Hours() / (24 * 30))
		days := int(d.Hours()/24) % 30
		return fmt.Sprintf("%d mth %d d", months, days)
	default:
		// 年和月组合 (每年按365天近似)
		years := int(d.Hours() / (24 * 365))
		months := int(d.Hours()/(24*30)) % 12
		return fmt.Sprintf("%d yr %d mth", years, months)
	}
}
