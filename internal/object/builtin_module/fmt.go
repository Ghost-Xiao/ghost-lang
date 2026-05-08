package builtin_module

import (
	"fmt"

	"github.com/Ghost-Xiao/ghost-lang/internal/errors"
	"github.com/Ghost-Xiao/ghost-lang/internal/frame"
	"github.com/Ghost-Xiao/ghost-lang/internal/object"
	"github.com/Ghost-Xiao/ghost-lang/internal/util"
)

// FmtModule 表示 fmt 内置模块
var FmtModule = initFmtModule()

// initFmtModule 初始化 fmt 模块
//
// 返回值:
//
//	*object.Module - 初始化后的 fmt 模块
func initFmtModule() *object.Module {
	env := &object.Environment{
		Name:  "fmt",
		Store: map[string]*object.Symbol{},
		Outer: nil,
	}

	env.Set("print", &object.Symbol{Name: "print", Value: &PRINT, IsConst: true})
	env.Set("println", &object.Symbol{Name: "println", Value: &PRINTLN, IsConst: true})
	env.Set("printf", &object.Symbol{Name: "printf", Value: &PRINTF, IsConst: true})
	env.Set("sprint", &object.Symbol{Name: "sprint", Value: &SPRINT, IsConst: true})
	env.Set("sprintf", &object.Symbol{Name: "sprintf", Value: &SPRINTF, IsConst: true})

	return &object.Module{
		Name: "fmt",
		Env:  env,
	}
}

var (
	// PRINT 打印函数（不换行）
	PRINT = object.BuiltinFunction{
		Name:         "print",
		Parameter:    []string{"x"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: true,
		Fn:           FmtPrint,
	}
	// PRINTLN 打印函数（换行）
	PRINTLN = object.BuiltinFunction{
		Name:         "println",
		Parameter:    []string{"x"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: true,
		Fn:           FmtPrintln,
	}
	// PRINTF 格式化打印函数
	PRINTF = object.BuiltinFunction{
		Name:         "printf",
		Parameter:    []string{"format", "a"},
		DefaultValue: []object.Object{nil, nil},
		HaveVariadic: true,
		Fn:           FmtPrintf,
	}
	// SPRINT 格式化输出为字符串（不换行）
	SPRINT = object.BuiltinFunction{
		Name:         "sprint",
		Parameter:    []string{"x"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: true,
		Fn:           FmtSprint,
	}
	// SPRINTF 格式化输出为字符串
	SPRINTF = object.BuiltinFunction{
		Name:         "sprintf",
		Parameter:    []string{"format", "a"},
		DefaultValue: []object.Object{nil, nil},
		HaveVariadic: true,
		Fn:           FmtSprintf,
	}
)

// FmtPrint 实现 print 函数
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 函数参数
//
// 返回值:
//
//	object.Object - 函数返回值
//	error - 可能出现的错误
func FmtPrint(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	for _, arg := range args[0].(*object.List).Elements {
		switch v := arg.(type) {
		case *object.String:
			fmt.Print(v.Value)
		case *object.Int:
			fmt.Print(v.Value)
		case *object.Float:
			fmt.Print(v.Value)
		case *object.Bool:
			fmt.Print(v.Value)
		case *object.Null:
			fmt.Print("null")
		default:
			fmt.Print(arg.String())
		}
	}
	return &object.Null{}, nil
}

// FmtPrintln 实现 println 函数
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 函数参数
//
// 返回值:
//
//	object.Object - 空结果
//	error - 可能出现的错误
func FmtPrintln(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	for i, arg := range args[0].(*object.List).Elements {
		if i > 0 {
			fmt.Print(" ")
		}
		switch v := arg.(type) {
		case *object.String:
			fmt.Print(v.Value)
		case *object.Int:
			fmt.Print(v.Value)
		case *object.Float:
			fmt.Print(v.Value)
		case *object.Bool:
			fmt.Print(v.Value)
		case *object.Null:
			fmt.Print("null")
		default:
			fmt.Print(arg.String())
		}
	}
	fmt.Println()
	return &object.Null{}, nil
}

// FmtPrintf 实现 printf 函数
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 函数参数
//
// 返回值:
//
//	object.Object - 空结果
//	error - 可能出现的错误
func FmtPrintf(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	if len(args) == 0 {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "printf requires at least one argument.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	format, ok := args[0].(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "printf() first argument must be a string.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	formatStr := format.Value
	var goArgs []interface{}
	for i := 0; i < len(args[1].(*object.List).Elements); i++ {
		arg := args[1].(*object.List).Elements[i]
		switch v := arg.(type) {
		case *object.String:
			goArgs = append(goArgs, v.Value)
		case *object.Int:
			goArgs = append(goArgs, v.Value)
		case *object.Float:
			goArgs = append(goArgs, v.Value)
		case *object.Bool:
			goArgs = append(goArgs, v.Value)
		case *object.Null:
			goArgs = append(goArgs, "null")
		default:
			goArgs = append(goArgs, arg.String())
		}
	}

	fmt.Printf(formatStr, goArgs...)
	return &object.Null{}, nil
}

// FmtSprint 实现 sprint 函数
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 函数参数
//
// 返回值:
//
//	object.Object - 字符串结果
//	error - 可能出现的错误
func FmtSprint(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	var result string
	for _, arg := range args[0].(*object.List).Elements {
		switch v := arg.(type) {
		case *object.String:
			result += v.Value
		case *object.Int:
			result += fmt.Sprintf("%d", v.Value)
		case *object.Float:
			result += fmt.Sprintf("%g", v.Value)
		case *object.Bool:
			result += fmt.Sprintf("%t", v.Value)
		case *object.Null:
			result += "null"
		default:
			result += arg.String()
		}
	}
	return &object.String{Value: result}, nil
}

// FmtSprintf 实现 sprintf 函数
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 函数参数
//
// 返回值:
//
//	object.Object - 字符串结果
//	error - 可能出现的错误
func FmtSprintf(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	if len(args) == 0 {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "sprintf requires at least one argument.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	format, ok := args[0].(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "sprintf() first argument must be a string.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	formatStr := format.Value
	var goArgs []interface{}
	for i := 0; i < len(args[1].(*object.List).Elements); i++ {
		arg := args[1].(*object.List).Elements[i]
		switch v := arg.(type) {
		case *object.String:
			goArgs = append(goArgs, v.Value)
		case *object.Int:
			goArgs = append(goArgs, v.Value)
		case *object.Float:
			goArgs = append(goArgs, v.Value)
		case *object.Bool:
			goArgs = append(goArgs, v.Value)
		case *object.Null:
			goArgs = append(goArgs, "null")
		default:
			goArgs = append(goArgs, arg.String())
		}
	}

	result := fmt.Sprintf(formatStr, goArgs...)
	return &object.String{Value: result}, nil
}
