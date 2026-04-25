package builtinclass

import (
	"github.com/Ghost-Xiao/ghost-lang/internal/errors"
	"github.com/Ghost-Xiao/ghost-lang/internal/frame"
	"github.com/Ghost-Xiao/ghost-lang/internal/object"
	"github.com/Ghost-Xiao/ghost-lang/internal/util"
)

// InstanceClass 表示 Instance 内置类的类定义
var InstanceClass = initInstanceClass()

// initInstanceClass 初始化 Instance 类
//
// 返回值:
//
//	*object.Class - 初始化后的 Instance 类对象
func initInstanceClass() *object.Class {
	member := &object.Environment{
		Name:  "Instance",
		Store: map[string]*object.Symbol{},
		Outer: nil,
	}

	member.Set("type", &object.Symbol{Name: "type", Value: &INSTANCETYPE, IsConst: true})
	member.Set("members", &object.Symbol{Name: "members", Value: &INSTANCEMEMBERS, IsConst: true})
	member.Set("hasMember", &object.Symbol{Name: "hasMember", Value: &INSTANCEHASMEMBER, IsConst: true})

	return &object.Class{
		Name:   "Instance",
		Parent: nil,
		Member: member,
	}
}

var (
	// INSTANCETYPE 表示 Instance 类的 type 方法
	INSTANCETYPE = object.Method{
		Name: "type",
		Function: &object.BuiltinFunction{
			Name:         "type",
			Parameter:    []string{},
			DefaultValue: []object.Object{},
			HaveVariadic: false,
			Fn:           InstanceType,
		},
		Instance: nil,
	}
	// INSTANCEMEMBERS 表示 Instance 类的 members 方法
	INSTANCEMEMBERS = object.Method{
		Name: "members",
		Function: &object.BuiltinFunction{
			Name:         "members",
			Parameter:    []string{},
			DefaultValue: []object.Object{},
			HaveVariadic: false,
			Fn:           InstanceMembers,
		},
		Instance: nil,
	}
	// INSTANCEHASMEMBER 表示 Instance 类的 hasMember 方法
	INSTANCEHASMEMBER = object.Method{
		Name: "hasMember",
		Function: &object.BuiltinFunction{
			Name:         "hasMember",
			Parameter:    []string{"name"},
			DefaultValue: []object.Object{nil},
			HaveVariadic: false,
			Fn:           InstanceHasMember,
		},
		Instance: nil,
	}
)

// InstanceType 实现 Instance 类的 type 方法
// 返回实例的类型
//
// 参数:
//
//	f - 当前调用栈
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数，第一个参数是 Instance 实例
//
// 返回值:
//
//	object.Object - 实例的类型对象
//	error - 可能出现的错误
func InstanceType(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this := args[0]
	if this == nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "method type() called without instance",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	instanceObj, ok := this.(*object.Instance)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "type() can only be called on Instance instances",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	return instanceObj.Class, nil
}

// InstanceMembers 实现 Instance 类的 members 方法
// 返回实例中所有成员名称的列表
//
// 参数:
//
//	f - 当前调用栈
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数，第一个参数是 Instance 实例
//
// 返回值:
//
//	object.Object - 包含所有成员名称的列表
//	error - 可能出现的错误
func InstanceMembers(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this := args[0]
	if this == nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "method members() called without instance",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	instanceObj, ok := this.(*object.Instance)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "members() can only be called on Instance instances",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	members := make([]object.Object, 0, len(instanceObj.Member.Store))
	for name := range instanceObj.Member.Store {
		members = append(members, &object.String{Value: name})
	}

	return &object.List{Elements: members}, nil
}

// InstanceHasMember 实现 Instance 类的 hasMember 方法
// 判断实例中是否包含指定名称的成员
//
// 参数:
//
//	f - 当前调用栈
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数，第一个参数是 Instance 实例，第二个参数是成员名称
//
// 返回值:
//
//	object.Object - Bool 类型的结果，true 表示包含该成员，false 表示不包含
//	error - 可能出现的错误
func InstanceHasMember(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this := args[0]
	if this == nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "method hasMember() called without instance",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	instanceObj, ok := this.(*object.Instance)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "hasMember() can only be called on Instance instances",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	if len(args) < 2 {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "hasMember() expects 1 argument",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	nameObj, ok := args[1].(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "hasMember() expects a String argument",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	_, exists := instanceObj.Member.Store[nameObj.Value]
	return &object.Bool{Value: exists}, nil
}
