package builtinclass

import (
	"fmt"
	"strconv"

	"github.com/Ghost-Xiao/ghost-lang/internal/errors"
	"github.com/Ghost-Xiao/ghost-lang/internal/frame"
	"github.com/Ghost-Xiao/ghost-lang/internal/object"
	"github.com/Ghost-Xiao/ghost-lang/internal/util"
)

// BoolClass 表示 Bool 内置类的类定义
var BoolClass = initBoolClass()

// initBoolClass 初始化 Bool 类
//
// 返回值:
//
//	*object.Class - 初始化后的 Bool 类对象
func initBoolClass() *object.Class {
	member := &object.Environment{
		Name:  "Bool",
		Store: map[string]*object.Symbol{},
		Outer: nil,
	}

	member.Set("init", &object.Symbol{Name: "init", Value: &BOOLINIT, IsConst: true})

	return &object.Class{
		Name:   "Bool",
		Parent: nil,
		Member: member,
	}
}

var (
	// BOOLINIT 表示 Bool 类的 init 方法
	BOOLINIT = object.Method{
		Name: "init",
		Function: &object.BuiltinFunction{
			Name:         "init",
			Parameter:    []string{"value"},
			DefaultValue: []object.Object{&object.Bool{Value: false}},
			HaveVariadic: false,
			Fn:           BoolInit,
		},
		Instance: nil,
	}
)

// BoolInit 实现 Bool 类的 init 构造方法
// 创建一个新的 Bool 实例，支持从多种类型转换
//
// 参数:
//
//	f - 当前调用栈
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数，第一个参数是 Bool 实例，第二个参数是要转换的值
//
// 返回值:
//
//	object.Object - 新创建的 Bool 实例
//	error - 可能出现的错误
func BoolInit(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	if len(args) < 1 {
		return &object.Bool{Value: false}, nil
	}

	value := args[0]

	switch v := value.(type) {
	case *object.Bool:
		return &object.Bool{Value: v.Value}, nil
	case *object.Int:
		return &object.Bool{Value: v.Value != 0}, nil
	case *object.Float:
		return &object.Bool{Value: v.Value != 0.0}, nil
	case *object.String:
		boolValue, err := strconv.ParseBool(v.Value)
		if err != nil {
			return nil, &errors.TypeError{
				Frame:    f,
				Message:  fmt.Sprintf("cannot convert \"%s\" to Bool.", value.String()),
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
		return &object.Bool{Value: boolValue}, nil
	case *object.Null:
		return &object.Bool{Value: false}, nil
	default:
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  fmt.Sprintf("cannot convert %s to Bool.", value.Type()),
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
}
