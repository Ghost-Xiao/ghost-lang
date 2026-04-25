package builtinclass

import (
	"github.com/Ghost-Xiao/ghost-lang/internal/errors"
	"github.com/Ghost-Xiao/ghost-lang/internal/frame"
	"github.com/Ghost-Xiao/ghost-lang/internal/object"
	"github.com/Ghost-Xiao/ghost-lang/internal/util"
)

// FunctionClass 表示 Function 内置类的类定义
var FunctionClass = initFunctionClass()

// initFunctionClass 初始化 Function 类
//
// 返回值:
//
//	*object.Class - 初始化后的 Function 类对象
func initFunctionClass() *object.Class {
	member := &object.Environment{
		Name:  "Function",
		Store: map[string]*object.Symbol{},
		Outer: nil,
	}

	member.Set("name", &object.Symbol{Name: "name", Value: &FUNCTIONNAME, IsConst: true})
	member.Set("arity", &object.Symbol{Name: "arity", Value: &FUNCTIONARITY, IsConst: true})

	return &object.Class{
		Name:   "Function",
		Parent: nil,
		Member: member,
	}
}

var (
	// FUNCTIONNAME 表示 Function 类的 name 方法
	FUNCTIONNAME = object.Method{
		Name: "name",
		Function: &object.BuiltinFunction{
			Name:         "name",
			Parameter:    []string{},
			DefaultValue: []object.Object{},
			HaveVariadic: false,
			Fn:           FunctionName,
		},
		Instance: nil,
	}
	// FUNCTIONARITY 表示 Function 类的 arity 方法
	FUNCTIONARITY = object.Method{
		Name: "arity",
		Function: &object.BuiltinFunction{
			Name:         "arity",
			Parameter:    []string{},
			DefaultValue: []object.Object{},
			HaveVariadic: false,
			Fn:           FunctionArity,
		},
		Instance: nil,
	}
)

// FunctionName 实现 Function 类的 name 方法
// 返回函数的名称
//
// 参数:
//
//	f - 当前调用栈
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数，第一个参数是 Function 实例
//
// 返回值:
//
//	object.Object - 函数的名称字符串
//	error - 可能出现的错误
func FunctionName(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this := args[0]
	if this == nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "method name() called without instance",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	functionObj, ok := this.(*object.Function)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "name() can only be called on Function instances",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	return &object.String{Value: functionObj.Name}, nil
}

// FunctionArity 实现 Function 类的 arity 方法
// 返回函数的参数数量
//
// 参数:
//
//	f - 当前调用栈
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数，第一个参数是 Function 实例
//
// 返回值:
//
//	object.Object - 函数的参数数量整数
//	error - 可能出现的错误
func FunctionArity(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this := args[0]
	if this == nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "method arity() called without instance",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	functionObj, ok := this.(*object.Function)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "arity() can only be called on Function instances",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	return &object.Int{Value: int64(len(functionObj.Parameter))}, nil
}
