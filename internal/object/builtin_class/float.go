package builtinclass

import (
	"fmt"
	"strconv"

	"github.com/Ghost-Xiao/ghost-lang/internal/errors"
	"github.com/Ghost-Xiao/ghost-lang/internal/frame"
	"github.com/Ghost-Xiao/ghost-lang/internal/object"
	"github.com/Ghost-Xiao/ghost-lang/internal/util"
)

// FloatClass 表示 Float 内置类的类定义
var FloatClass = initFloatClass()

// initFloatClass 初始化 Float 类
//
// 返回值:
//
//	*object.Class - 初始化后的 Float 类对象
func initFloatClass() *object.Class {
	member := &object.Environment{
		Name:  "Float",
		Store: map[string]*object.Symbol{},
		Outer: nil,
	}

	member.Set("init", &object.Symbol{Name: "init", Value: &FLOATINIT, IsConst: true})

	return &object.Class{
		Name:   "Float",
		Parent: nil,
		Member: member,
	}
}

var (
	// FLOATINIT 表示 Float 类的 init 方法
	FLOATINIT = object.BuiltinFunction{
		Name:         "init",
		Parameter:    []string{"value"},
		DefaultValue: []object.Object{&object.Float{Value: 0.0}},
		HaveVariadic: false,
		Fn:           FloatInit,
	}
)

// FloatInit 实现 Float 类的 init 构造方法
// 创建一个新的 Float 实例，支持从多种类型转换
//
// 参数:
//
//	f - 当前调用栈
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数，第一个参数是 Float 实例，第二个参数是要转换的值
//
// 返回值:
//
//	object.Object - 新创建的 Float 实例
//	error - 可能出现的错误
func FloatInit(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	if len(args) < 1 {
		return &object.Float{Value: 0.0}, nil
	}

	value := args[0]

	switch v := value.(type) {
	case *object.Float:
		return &object.Float{Value: v.Value}, nil
	case *object.Int:
		return &object.Float{Value: float64(v.Value)}, nil
	case *object.String:
		floatValue, err := strconv.ParseFloat(v.Value, 64)
		if err != nil {
			return nil, &errors.TypeError{
				Frame:    f,
				Message:  fmt.Sprintf("cannot convert \"%s\" to Float.", value.String()),
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
		return &object.Float{Value: floatValue}, nil
	case *object.Bool:
		if v.Value {
			return &object.Float{Value: 1.0}, nil
		}
		return &object.Float{Value: 0.0}, nil
	case *object.Null:
		return &object.Float{Value: 0.0}, nil
	default:
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  fmt.Sprintf("cannot convert %s to Float.", value.Type()),
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
}
