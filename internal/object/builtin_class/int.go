package builtinclass

import (
	"fmt"

	"github.com/Ghost-Xiao/ghost-lang/internal/errors"
	"github.com/Ghost-Xiao/ghost-lang/internal/frame"
	"github.com/Ghost-Xiao/ghost-lang/internal/object"
	"github.com/Ghost-Xiao/ghost-lang/internal/util"
)

// IntClass 表示 Int 内置类的类定义
var IntClass = initIntClass()

// initIntClass 初始化 Int 类
//
// 返回值:
//
//	*object.Class - 初始化后的 Int 类对象
func initIntClass() *object.Class {
	member := &object.Environment{
		Name:  "Int",
		Store: map[string]*object.Symbol{},
		Outer: nil,
	}

	member.Set("init", &object.Symbol{Name: "init", Value: &INTINIT, IsConst: true})
	member.Set("isEven", &object.Symbol{Name: "isEven", Value: &ISEVEN, IsConst: true})
	member.Set("isOdd", &object.Symbol{Name: "isOdd", Value: &ISODD, IsConst: true})

	return &object.Class{
		Name:   "Int",
		Parent: nil,
		Member: member,
	}
}

var (
	// INTINIT 表示 Int 类的 init 方法
	INTINIT = object.Method{
		Name: "init",
		Function: &object.BuiltinFunction{
			Name:         "init",
			Parameter:    []string{"value"},
			DefaultValue: []object.Object{&object.Int{Value: 0}},
			HaveVariadic: false,
			Fn:           IntInit,
		},
		Instance: nil,
	}
	// ISEVEN 表示 Int 类的 isEven 方法
	ISEVEN = object.Method{
		Name: "isEven",
		Function: &object.BuiltinFunction{
			Name:         "isEven",
			Parameter:    []string{},
			DefaultValue: []object.Object{},
			HaveVariadic: false,
			Fn:           IntIsEven,
		},
		Instance: nil,
	}
	// ISODD 表示 Int 类的 isOdd 方法
	ISODD = object.Method{
		Name: "isOdd",
		Function: &object.BuiltinFunction{
			Name:         "isOdd",
			Parameter:    []string{},
			DefaultValue: []object.Object{},
			HaveVariadic: false,
			Fn:           IntIsOdd,
		},
		Instance: nil,
	}
)

// IntIsEven 实现 Int 类的 isEven 方法
// 判断整数是否为偶数
//
// 参数:
//
//	f - 当前调用栈
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数，第一个参数是 Int 实例
//
// 返回值:
//
//	object.Object - Bool 类型的结果，true 表示是偶数，false 表示不是
//	error - 可能出现的错误
func IntIsEven(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this := args[0]
	if this == nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "method isEven() called without instance",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	intObj, ok := this.(*object.Int)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "isEven() can only be called on Int instances",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	return &object.Bool{Value: intObj.Value%2 == 0}, nil
}

// IntIsOdd 实现 Int 类的 isOdd 方法
// 判断整数是否为奇数
//
// 参数:
//
//	f - 当前调用栈
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数，第一个参数是 Int 实例
//
// 返回值:
//
//	object.Object - Bool 类型的结果，true 表示是奇数，false 表示不是
//	error - 可能出现的错误
func IntIsOdd(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this := args[0]
	if this == nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "method isOdd() called without instance",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	intObj, ok := this.(*object.Int)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "isOdd() can only be called on Int instances",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	return &object.Bool{Value: intObj.Value%2 != 0}, nil
}

// IntInit 实现 Int 类的 init 构造方法
// 创建一个新的 Int 实例，支持从多种类型转换
//
// 参数:
//
//	f - 当前调用栈
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数，第一个参数是 Int 实例，第二个参数是要转换的值
//
// 返回值:
//
//	object.Object - 新创建的 Int 实例
//	error - 可能出现的错误
func IntInit(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	if len(args) < 1 {
		return &object.Int{Value: 0}, nil
	}

	value := args[0]

	switch v := value.(type) {
	case *object.Int:
		return &object.Int{Value: v.Value}, nil
	case *object.Float:
		return &object.Int{Value: int64(v.Value)}, nil
	case *object.String:
		var intValue int64
		_, err := fmt.Sscanf(v.Value, "%d", &intValue)
		if err != nil {
			return nil, &errors.TypeError{
				Frame:    f,
				Message:  fmt.Sprintf("cannot convert \"%s\" to Int.", value.String()),
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
		return &object.Int{Value: intValue}, nil
	case *object.Bool:
		if v.Value {
			return &object.Int{Value: 1}, nil
		}
		return &object.Int{Value: 0}, nil
	case *object.Null:
		return &object.Int{Value: 0}, nil
	default:
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  fmt.Sprintf("cannot convert %s to Int.", value.Type()),
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
}
