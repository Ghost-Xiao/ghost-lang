package builtinclass

import (
	"github.com/Ghost-Xiao/ghost-lang/internal/errors"
	"github.com/Ghost-Xiao/ghost-lang/internal/frame"
	"github.com/Ghost-Xiao/ghost-lang/internal/object"
	"github.com/Ghost-Xiao/ghost-lang/internal/util"
)

// NamespaceClass 表示 Namespace 内置类的类定义
var NamespaceClass = initNamespaceClass()

// initNamespaceClass 初始化 Namespace 类
//
// 返回值:
//
//	*object.Class - 初始化后的 Namespace 类对象
func initNamespaceClass() *object.Class {
	member := &object.Environment{
		Name:  "Namespace",
		Store: map[string]*object.Symbol{},
		Outer: nil,
	}

	member.Set("members", &object.Symbol{Name: "members", Value: &NAMESPACEMEMBERS, IsConst: true})
	member.Set("has", &object.Symbol{Name: "has", Value: &NAMESPACEHAS, IsConst: true})

	return &object.Class{
		Name:   "Namespace",
		Parent: nil,
		Member: member,
	}
}

var (
	// NAMESPACEMEMBERS 表示 Namespace 类的 members 方法
	NAMESPACEMEMBERS = object.Method{
		Name: "members",
		Function: &object.BuiltinFunction{
			Name:         "members",
			Parameter:    []string{},
			DefaultValue: []object.Object{},
			HaveVariadic: false,
			Fn:           NamespaceMembers,
		},
		Instance: nil,
	}
	// NAMESPACEHAS 表示 Namespace 类的 has 方法
	NAMESPACEHAS = object.Method{
		Name: "has",
		Function: &object.BuiltinFunction{
			Name:         "has",
			Parameter:    []string{"name"},
			DefaultValue: []object.Object{nil},
			HaveVariadic: false,
			Fn:           NamespaceHas,
		},
		Instance: nil,
	}
)

// NamespaceMembers 实现 Namespace 类的 members 方法
// 返回命名空间中所有成员名称的列表
//
// 参数:
//
//	f - 当前调用栈
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数，第一个参数是 Namespace 实例
//
// 返回值:
//
//	object.Object - 包含所有成员名称的列表
//	error - 可能出现的错误
func NamespaceMembers(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this := args[0]
	if this == nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "method members() called without instance",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	nsObj, ok := this.(*object.Namespace)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "members() can only be called on Namespace instances",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	members := make([]object.Object, 0, len(nsObj.Member.Store))
	for name := range nsObj.Member.Store {
		members = append(members, &object.String{Value: name})
	}

	return &object.List{Elements: members}, nil
}

// NamespaceHas 实现 Namespace 类的 has 方法
// 判断命名空间中是否包含指定名称的成员
//
// 参数:
//
//	f - 当前调用栈
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数，第一个参数是 Namespace 实例，第二个参数是成员名称
//
// 返回值:
//
//	object.Object - Bool 类型的结果，true 表示包含该成员，false 表示不包含
//	error - 可能出现的错误
func NamespaceHas(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this := args[0]
	if this == nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "method has() called without instance",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	nsObj, ok := this.(*object.Namespace)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "has() can only be called on Namespace instances",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	if len(args) < 2 {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "has() expects 1 argument",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	nameObj, ok := args[1].(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "has() expects a String argument",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	_, exists := nsObj.Member.Store[nameObj.Value]
	return &object.Bool{Value: exists}, nil
}
