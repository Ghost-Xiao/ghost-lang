package builtinclass

import (
	"github.com/Ghost-Xiao/ghost-lang/internal/errors"
	"github.com/Ghost-Xiao/ghost-lang/internal/frame"
	"github.com/Ghost-Xiao/ghost-lang/internal/object"
	"github.com/Ghost-Xiao/ghost-lang/internal/util"
)

// ClassClass 表示 Class 内置类的类定义
var ClassClass = initClassClass()

// initClassClass 初始化 Class 类
//
// 返回值:
//
//	*object.Class - 初始化后的 Class 类对象
func initClassClass() *object.Class {
	member := &object.Environment{
		Name:  "Class",
		Store: map[string]*object.Symbol{},
		Outer: nil,
	}

	member.Set("name", &object.Symbol{Name: "name", Value: &CLASSNAME, IsConst: true})
	member.Set("superclass", &object.Symbol{Name: "superclass", Value: &CLASSSUPERCLASS, IsConst: true})
	member.Set("methods", &object.Symbol{Name: "methods", Value: &CLASSMETHODS, IsConst: true})
	member.Set("isSubclassOf", &object.Symbol{Name: "isSubclassOf", Value: &CLASSISSUBCLASSOF, IsConst: true})

	return &object.Class{
		Name:   "Class",
		Parent: nil,
		Member: member,
	}
}

var (
	// CLASSNAME 表示 Class 类的 name 方法
	CLASSNAME = object.BuiltinFunction{
		Name:         "name",
		Parameter:    []string{},
		DefaultValue: []object.Object{},
		HaveVariadic: false,
		Fn:           ClassName,
	}
	// CLASSSUPERCLASS 表示 Class 类的 superclass 方法
	CLASSSUPERCLASS = object.BuiltinFunction{
		Name:         "superclass",
		Parameter:    []string{},
		DefaultValue: []object.Object{},
		HaveVariadic: false,
		Fn:           ClassSuperclass,
	}
	// CLASSMETHODS 表示 Class 类的 methods 方法
	CLASSMETHODS = object.BuiltinFunction{
		Name:         "methods",
		Parameter:    []string{},
		DefaultValue: []object.Object{},
		HaveVariadic: false,
		Fn:           ClassMethods,
	}
	// CLASSISSUBCLASSOF 表示 Class 类的 isSubclassOf 方法
	CLASSISSUBCLASSOF = object.BuiltinFunction{
		Name:         "isSubclassOf",
		Parameter:    []string{"cls"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           ClassIsSubclassOf,
	}
)

// ClassName 实现 Class 类的 name 方法
// 返回类的名称
//
// 参数:
//
//	f - 当前调用栈
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数，第一个参数是 Class 实例
//
// 返回值:
//
//	object.Object - 类的名称字符串
//	error - 可能出现的错误
func ClassName(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this := args[0]
	if this == nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "method name() called without instance",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	classObj, ok := this.(*object.Class)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "name() can only be called on Class instances",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	return &object.String{Value: classObj.Name}, nil
}

// ClassSuperclass 实现 Class 类的 superclass 方法
// 返回类的父类
//
// 参数:
//
//	f - 当前调用栈
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数，第一个参数是 Class 实例
//
// 返回值:
//
//	object.Object - 父类对象，如果没有父类则返回 null
//	error - 可能出现的错误
func ClassSuperclass(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this := args[0]
	if this == nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "method superclass() called without instance",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	classObj, ok := this.(*object.Class)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "superclass() can only be called on Class instances",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	if classObj.Parent == nil {
		return &object.Null{}, nil
	}

	return classObj.Parent, nil
}

// ClassMethods 实现 Class 类的 methods 方法
// 返回类中所有方法名称的列表
//
// 参数:
//
//	f - 当前调用栈
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数，第一个参数是 Class 实例
//
// 返回值:
//
//	object.Object - 包含所有方法名称的列表
//	error - 可能出现的错误
func ClassMethods(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this := args[0]
	if this == nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "method methods() called without instance",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	classObj, ok := this.(*object.Class)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "methods() can only be called on Class instances",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	methods := make([]object.Object, 0, len(classObj.Member.Store))
	for name, symbol := range classObj.Member.Store {
		if _, ok := symbol.Value.(*object.Function); ok {
			methods = append(methods, &object.String{Value: name})
		}
		if _, ok := symbol.Value.(*object.BuiltinFunction); ok {
			methods = append(methods, &object.String{Value: name})
		}
	}

	return &object.List{Elements: methods}, nil
}

// ClassIsSubclassOf 实现 Class 类的 isSubclassOf 方法
// 判断当前类是否是指定类的子类
//
// 参数:
//
//	f - 当前调用栈
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数，第一个参数是 Class 实例，第二个参数是要检查的类
//
// 返回值:
//
//	object.Object - Bool 类型的结果，true 表示是子类，false 表示不是
//	error - 可能出现的错误
func ClassIsSubclassOf(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this := args[0]
	if this == nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "method isSubclassOf() called without instance",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	classObj, ok := this.(*object.Class)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "isSubclassOf() can only be called on Class instances",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	if len(args) < 2 {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "isSubclassOf() expects 1 argument",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	parentCls, ok := args[1].(*object.Class)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "isSubclassOf() expects a Class argument",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	current := classObj
	for current != nil {
		if current.Name == parentCls.Name {
			return &object.Bool{Value: true}, nil
		}
		current = current.Parent
	}

	return &object.Bool{Value: false}, nil
}
