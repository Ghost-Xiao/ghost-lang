package builtinclass

import (
	"github.com/Ghost-Xiao/ghost-lang/internal/errors"
	"github.com/Ghost-Xiao/ghost-lang/internal/frame"
	"github.com/Ghost-Xiao/ghost-lang/internal/object"
	"github.com/Ghost-Xiao/ghost-lang/internal/util"
)

// ModuleClass 表示 Module 内置类的类定义
var ModuleClass = initModuleClass()

// initModuleClass 初始化 Module 类
//
// 返回值:
//
//	*object.Class - 初始化后的 Module 类对象
func initModuleClass() *object.Class {
	member := &object.Environment{
		Name:  "Module",
		Store: map[string]*object.Symbol{},
		Outer: nil,
	}

	member.Set("path", &object.Symbol{Name: "path", Value: &MODULEPATH, IsConst: true})
	member.Set("name", &object.Symbol{Name: "name", Value: &MODULENAME, IsConst: true})

	return &object.Class{
		Name:   "Module",
		Parent: nil,
		Member: member,
	}
}

var (
	// MODULEPATH 表示 Module 类的 path 方法
	MODULEPATH = object.BuiltinFunction{
		Name:         "path",
		Parameter:    []string{},
		DefaultValue: []object.Object{},
		HaveVariadic: false,
		Fn:           ModulePath,
	}
	// MODULENAME 表示 Module 类的 name 方法
	MODULENAME = object.BuiltinFunction{
		Name:         "name",
		Parameter:    []string{},
		DefaultValue: []object.Object{},
		HaveVariadic: false,
		Fn:           ModuleName,
	}
)

// ModulePath 实现 Module 类的 path 方法
// 返回模块的路径
//
// 参数:
//
//	f - 当前调用栈
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数，第一个参数是 Module 实例
//
// 返回值:
//
//	object.Object - 模块的路径字符串
//	error - 可能出现的错误
func ModulePath(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this := args[0]
	if this == nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "method path() called without instance",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	moduleObj, ok := this.(*object.Module)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "path() can only be called on Module instances",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	return &object.String{Value: moduleObj.Name}, nil
}

// ModuleName 实现 Module 类的 name 方法
// 返回模块的名称
//
// 参数:
//
//	f - 当前调用栈
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数，第一个参数是 Module 实例
//
// 返回值:
//
//	object.Object - 模块的名称字符串
//	error - 可能出现的错误
func ModuleName(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this := args[0]
	if this == nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "method name() called without instance",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	moduleObj, ok := this.(*object.Module)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "name() can only be called on Module instances",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	return &object.String{Value: moduleObj.Name}, nil
}
