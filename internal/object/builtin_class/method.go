package builtinclass

import (
	"github.com/Ghost-Xiao/ghost-lang/internal/errors"
	"github.com/Ghost-Xiao/ghost-lang/internal/frame"
	"github.com/Ghost-Xiao/ghost-lang/internal/object"
	"github.com/Ghost-Xiao/ghost-lang/internal/util"
)

// MethodClass 表示 Method 内置类的类定义
var MethodClass = initMethodClass()

// initMethodClass 初始化 Method 类
//
// 返回值:
//
//	*object.Class - 初始化后的 Method 类对象
func initMethodClass() *object.Class {
	member := &object.Environment{
		Name:  "Method",
		Store: map[string]*object.Symbol{},
		Outer: nil,
	}

	member.Set("name", &object.Symbol{Name: "name", Value: &METHODNAME, IsConst: true})
	member.Set("owner", &object.Symbol{Name: "owner", Value: &METHODOWNER, IsConst: true})

	return &object.Class{
		Name:   "Method",
		Parent: nil,
		Member: member,
	}
}

var (
	// METHODNAME 表示 Method 类的 name 方法
	METHODNAME = object.Method{
		Name: "name",
		Function: &object.BuiltinFunction{
			Name:         "name",
			Parameter:    []string{},
			DefaultValue: []object.Object{},
			HaveVariadic: false,
			Fn:           MethodName,
		},
		Instance: nil,
	}
	// METHODOWNER 表示 Method 类的 owner 方法
	METHODOWNER = object.Method{
		Name: "owner",
		Function: &object.BuiltinFunction{
			Name:         "owner",
			Parameter:    []string{},
			DefaultValue: []object.Object{},
			HaveVariadic: false,
			Fn:           MethodOwner,
		},
		Instance: nil,
	}
)

// MethodName 实现 Method 类的 name 方法
// 返回方法的名称
//
// 参数:
//
//	f - 当前调用栈
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数，第一个参数是 Method 实例
//
// 返回值:
//
//	object.Object - 方法的名称字符串
//	error - 可能出现的错误
func MethodName(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this := args[0]
	if this == nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "method name() called without instance",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	methodObj, ok := this.(*object.Method)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "name() can only be called on Method instances",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	return &object.String{Value: methodObj.Name}, nil
}

// MethodOwner 实现 Method 类的 owner 方法
// 返回方法的拥有者
//
// 参数:
//
//	f - 当前调用栈
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数，第一个参数是 Method 实例
//
// 返回值:
//
//	object.Object - 方法的拥有者对象
//	error - 可能出现的错误
func MethodOwner(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this := args[0]
	if this == nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "method owner() called without instance",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	methodObj, ok := this.(*object.Method)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "owner() can only be called on Method instances",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	if methodObj.Instance == nil {
		return &object.Null{}, nil
	}

	return methodObj.Instance, nil
}
