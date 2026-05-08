package object

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/Ghost-Xiao/ghost-lang/internal/errors"
	"github.com/Ghost-Xiao/ghost-lang/internal/frame"
	"github.com/Ghost-Xiao/ghost-lang/internal/util"
)

// BuiltinFunction 表示内建函数类型，实现了Object接口
// 支持的操作包括调用函数等
type BuiltinFunction struct {
	Name         string                                                                                             // 函数名
	Parameter    []string                                                                                           // 参数名
	DefaultValue []Object                                                                                           // 默认参数值
	HaveVariadic bool                                                                                               // 是否为可变参数
	Fn           func(f *frame.Frame, env *Environment, posStart, posEnd *util.Pos, args ...Object) (Object, error) // 函数体
}

// Type 返回值的类型
//
// 返回值:
//
//	string - 值的类型
func (bf *BuiltinFunction) Type() string {
	return "BuiltinFunction"
}

// String 返回值的字符串表示
//
// 返回值:
//
//	string - 格式化的字符串表示
func (bf *BuiltinFunction) String() string {
	var sb strings.Builder
	sb.WriteString("func ")
	sb.WriteString(bf.Name)
	sb.WriteString("(")
	for i, param := range bf.Parameter {
		if bf.HaveVariadic && i == len(bf.Parameter)-1 {
			sb.WriteString("...")
		}
		sb.WriteString(param)
		if bf.DefaultValue[i] != nil {
			sb.WriteString("=")
			sb.WriteString(bf.DefaultValue[i].String())
		}
		if i < len(bf.Parameter)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString(") { [builtin code] }")
	return sb.String()
}

// Negative 对值进行负运算
//
// 参数:
//
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	Object - 运算结果
//	error - 可能出现的错误
func (bf *BuiltinFunction) Negative(posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"-\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// BitNot 对值进行按位非运算
//
// 参数:
//
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	Object - 运算结果
//	error - 可能出现的错误
func (bf *BuiltinFunction) BitNot(posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"~\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// Not 对值进行逻辑非运算
//
// 参数:
//
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	Object - 运算结果
//	error - 可能出现的错误
func (bf *BuiltinFunction) Not(posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"!\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// Add 对值进行加法运算
//
// 参数:
//
//	other - 另一个操作数
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	Object - 运算结果
//	error - 可能出现的错误
func (bf *BuiltinFunction) Add(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"+\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// Subtract 对值进行减法运算
//
// 参数:
//
//	other - 另一个操作数
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	Object - 运算结果
//	error - 可能出现的错误
func (bf *BuiltinFunction) Subtract(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"-\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// Multiply 对值进行乘法运算
//
// 参数:
//
//	other - 另一个操作数
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	Object - 运算结果
//	error - 可能出现的错误
func (bf *BuiltinFunction) Multiply(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"*\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// Divide 对值进行除法运算
//
// 参数:
//
//	other - 另一个操作数
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	Object - 运算结果
//	error - 可能出现的错误
func (bf *BuiltinFunction) Divide(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"/\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// Mod 对值进行取模运算
//
// 参数:
//
//	other - 另一个操作数
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	Object - 运算结果
//	error - 可能出现的错误
func (bf *BuiltinFunction) Mod(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"%\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// Equal 判断当前函数与另一个值是否相等
//
// 参数:
//
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	布尔值，表示比较结果；无错误
//
// 比较规则:
//
//	引用性比较
func (bf *BuiltinFunction) Equal(other Object, _, _ *util.Pos, _ *frame.Frame) (Object, error) {
	// 函数相等比较规则: 比较引用是否相等
	otherFunc, ok := other.(*BuiltinFunction)
	if !ok {
		return &Bool{Value: false}, nil
	}
	return &Bool{Value: bf == otherFunc}, nil
}

// NotEqual 判断当前函数与另一个值是否不相等
//
// 参数:
//
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	布尔值，表示比较结果；无错误
//
// 比较规则:
//
//	引用性比较
func (bf *BuiltinFunction) NotEqual(other Object, _, _ *util.Pos, _ *frame.Frame) (Object, error) {
	// 函数不等比较规则: 比较引用是否不等
	otherFunc, ok := other.(*BuiltinFunction)
	if !ok {
		return &Bool{Value: true}, nil
	}
	return &Bool{Value: bf != otherFunc}, nil
}

// LessThan 对值进行小于比较
//
// 参数:
//
//	other - 另一个操作数
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	Object - 比较结果
func (bf *BuiltinFunction) LessThan(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"<\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// GreaterThan 对值进行大于比较
//
// 参数:
//
//	other - 另一个操作数
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	Object - 比较结果
func (bf *BuiltinFunction) GreaterThan(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \">\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// LessThanOrEqual 对值进行小于等于比较
//
// 参数:
//
//	other - 另一个操作数
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	Object - 比较结果
func (bf *BuiltinFunction) LessThanOrEqual(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"<=\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// GreaterThanOrEqual 对值进行大于等于比较
//
// 参数:
//
//	other - 另一个操作数
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	Object - 比较结果
func (bf *BuiltinFunction) GreaterThanOrEqual(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \">=\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// BitAnd 对值进行按位与运算
//
// 参数:
//
//	other - 另一个操作数
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	Object - 运算结果
//	error - 可能出现的错误
func (bf *BuiltinFunction) BitAnd(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"&\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// BitOr 对值进行按位或运算
//
// 参数:
//
//	other - 另一个操作数
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	Object - 运算结果
//	error - 可能出现的错误
func (bf *BuiltinFunction) BitOr(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"|\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// Xor 对值进行异或运算
//
// 参数:
//
//	other - 另一个操作数
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	Object - 运算结果
//	error - 可能出现的错误
func (bf *BuiltinFunction) Xor(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"^\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// LeftShift 对值进行左移运算
//
// 参数:
//
//	other - 另一个操作数
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	Object - 运算结果
//	error - 可能出现的错误
func (bf *BuiltinFunction) LeftShift(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"<<\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// RightShift 对值进行右移运算
//
// 参数:
//
//	other - 另一个操作数
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	Object - 运算结果
//	error - 可能出现的错误
func (bf *BuiltinFunction) RightShift(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \">>\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// And 对值进行逻辑与运算
//
// 参数:
//
//	other - 另一个操作数
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	Object - 运算结果
//	error - 可能出现的错误
func (bf *BuiltinFunction) And(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"&&\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// Or 对值进行逻辑或运算
//
// 参数:
//
//	other - 另一个操作数
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	Object - 运算结果
//	error - 可能出现的错误
func (bf *BuiltinFunction) Or(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"||\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// Index 执行索引运算
//
// 参数:
//
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	Object - 运算结果
//	error - 可能出现的错误
func (bf *BuiltinFunction) Index(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	return nil, &errors.TypeError{
		Frame:    frame,
		Message:  "index expression not supported for this type.",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// Builtins 是内建函数和常量
var Builtins = map[string]Object{
	// inf常量
	"inf": &Float{Value: math.Inf(1)},
	// nan常量
	"nan": &Float{Value: math.NaN()},
	// print函数
	"print": &BuiltinFunction{
		Name:      "print",
		Parameter: []string{"a"},
		DefaultValue: []Object{
			nil,
		},
		HaveVariadic: true,
		Fn: func(f *frame.Frame, env *Environment, posStart, posEnd *util.Pos, args ...Object) (Object, error) {
			for i, arg := range args[0].(*List).Elements {
				if i > 0 {
					fmt.Print(" ")
				}
				fmt.Print(arg.String())
			}
			// 刷新缓冲区
			_ = os.Stdout.Sync()
			return &Null{}, nil
		},
	},
	// println函数
	"println": &BuiltinFunction{
		Name:      "println",
		Parameter: []string{"a"},
		DefaultValue: []Object{
			nil,
		},
		HaveVariadic: true,
		Fn: func(f *frame.Frame, env *Environment, posStart, posEnd *util.Pos, args ...Object) (Object, error) {
			for i, arg := range args[0].(*List).Elements {
				if i > 0 {
					fmt.Print(" ")
				}
				fmt.Print(arg.String())
			}
			fmt.Println()
			// 刷新缓冲区
			_ = os.Stdout.Sync()
			return &Null{}, nil
		},
	},
	// input函数
	"input": &BuiltinFunction{
		Name:      "input",
		Parameter: []string{"prompt"},
		DefaultValue: []Object{
			&String{Value: ""},
		},
		HaveVariadic: false,
		Fn: func(f *frame.Frame, env *Environment, posStart, posEnd *util.Pos, args ...Object) (Object, error) {
			// 打印提示
			prompt, ok := args[0].(*String)
			if !ok {
				return nil, &errors.TypeError{
					Frame:    f,
					Message:  "input() argument must be a string.",
					PosStart: posStart,
					PosEnd:   posEnd,
				}
			}
			fmt.Print(prompt.String())
			// 刷新缓冲区
			_ = os.Stdout.Sync()

			// 从标准输入读取一行
			var line string
			_, err := fmt.Scanln(&line)
			if err != nil {
				// 如果是 EOF 或空输入，返回空字符串
				line = ""
			}

			return &String{Value: line}, nil
		},
	},
	// len函数
	"len": &BuiltinFunction{
		Name:      "len",
		Parameter: []string{"a"},
		DefaultValue: []Object{
			nil,
		},
		HaveVariadic: false,
		Fn: func(f *frame.Frame, env *Environment, posStart, posEnd *util.Pos, args ...Object) (Object, error) {
			a := args[0]
			indexable, ok := a.(Indexable)
			if !ok {
				return nil, &errors.TypeError{
					Frame:    f,
					Message:  "len() argument must be a sequence or collection.",
					PosStart: posStart,
					PosEnd:   posEnd,
				}
			}
			return &Int{Value: indexable.Length()}, nil
		},
	},
	// power函数
	"power": &BuiltinFunction{
		Name:      "power",
		Parameter: []string{"a", "n"},
		DefaultValue: []Object{
			nil,
		},
		HaveVariadic: false,
		Fn: func(f *frame.Frame, env *Environment, posStart, posEnd *util.Pos, args ...Object) (Object, error) {
			a := args[0]
			n := args[1]
			var base, exp float64
			switch a := a.(type) {
			case *Int:
				base = float64(a.Value)
			case *Float:
				base = a.Value
			default:
				return nil, &errors.TypeError{
					Frame:    f,
					Message:  "power() arguments must be integers or floats.",
					PosStart: posStart,
					PosEnd:   posEnd,
				}
			}
			switch n := n.(type) {
			case *Int:
				exp = float64(n.Value)
			case *Float:
				exp = n.Value
			default:
				return nil, &errors.TypeError{
					Frame:    f,
					Message:  "power() arguments must be integers or floats.",
					PosStart: posStart,
					PosEnd:   posEnd,
				}
			}
			return &Float{Value: math.Pow(base, exp)}, nil
		},
	},
	// typeof函数
	"typeof": &BuiltinFunction{
		Name:      "typeof",
		Parameter: []string{"a"},
		DefaultValue: []Object{
			nil,
		},
		HaveVariadic: false,
		Fn: func(f *frame.Frame, env *Environment, posStart, posEnd *util.Pos, args ...Object) (Object, error) {
			a := args[0]
			return &String{Value: a.Type()}, nil
		},
	},
	// isInstanceOf函数
	"isInstanceOf": &BuiltinFunction{
		Name:      "isInstanceOf",
		Parameter: []string{"obj", "cls"},
		DefaultValue: []Object{
			nil,
		},
		HaveVariadic: false,
		Fn: func(f *frame.Frame, env *Environment, posStart, posEnd *util.Pos, args ...Object) (Object, error) {
			obj := args[0]
			cls, ok := args[1].(*Class)
			if !ok {
				return nil, &errors.TypeError{
					Frame:    f,
					Message:  "isInstanceOf() second argument must be a class.",
					PosStart: posStart,
					PosEnd:   posEnd,
				}
			}
			switch o := obj.(type) {
			case *Instance:
				return &Bool{Value: o.Class == cls}, nil
			case *Int, *Float, *Bool, *String, *List, *Map, *Namespace, *Class, *Function, *Method, *Module:
				return &Bool{Value: o.Type() == cls.Name}, nil
			default:
				return &Bool{Value: false}, nil
			}
		},
	},
	// Error 类
	"Error": ErrorClass,
	// OperationError 类
	"OperationError": OperationErrorClass,
	// MathError 类
	"MathError": MathErrorClass,
	// TypeError 类
	"TypeError": TypeErrorClass,
	// IndexError 类
	"IndexError": IndexErrorClass,
	// VariableError 类
	"VariableError": VariableErrorClass,
	// ArgumentError 类
	"ArgumentError": ArgumentErrorClass,
	// ModuleError 类
	"ModuleError": ModuleErrorClass,
}

// ErrorClass 表示 Error 内置基类的类定义
var ErrorClass = initErrorClass()

// initErrorClass 初始化 Error 基类
func initErrorClass() *Class {
	member := &Environment{
		Name:  "Error",
		Store: map[string]*Symbol{},
		Outer: nil,
	}

	member.Set("message", &Symbol{Name: "message", Value: &String{Value: ""}, IsConst: false})
	member.Set("init", &Symbol{Name: "init", Value: &ERRORINIT, IsConst: true})

	return &Class{
		Name:   "Error",
		Parent: nil,
		Member: member,
	}
}

// OperationErrorClass 表示 OperationError 内置类的类定义
var OperationErrorClass = initOperationErrorClass()

// initOperationErrorClass 初始化 OperationError 类
func initOperationErrorClass() *Class {
	member := &Environment{
		Name:  "OperationError",
		Store: map[string]*Symbol{},
		Outer: nil,
	}

	member.Set("message", &Symbol{Name: "message", Value: &String{Value: ""}, IsConst: false})
	member.Set("init", &Symbol{Name: "init", Value: &OPERATIONERRORINIT, IsConst: true})

	return &Class{
		Name:   "OperationError",
		Parent: ErrorClass,
		Member: member,
	}
}

// MathErrorClass 表示 MathError 内置类的类定义
var MathErrorClass = initMathErrorClass()

// initMathErrorClass 初始化 MathError 类
func initMathErrorClass() *Class {
	member := &Environment{
		Name:  "MathError",
		Store: map[string]*Symbol{},
		Outer: nil,
	}

	member.Set("message", &Symbol{Name: "message", Value: &String{Value: ""}, IsConst: false})
	member.Set("init", &Symbol{Name: "init", Value: &MATHERRORINIT, IsConst: true})

	return &Class{
		Name:   "MathError",
		Parent: ErrorClass,
		Member: member,
	}
}

// TypeErrorClass 表示 TypeError 内置类的类定义
var TypeErrorClass = initTypeErrorClass()

// initTypeErrorClass 初始化 TypeError 类
func initTypeErrorClass() *Class {
	member := &Environment{
		Name:  "TypeError",
		Store: map[string]*Symbol{},
		Outer: nil,
	}

	member.Set("message", &Symbol{Name: "message", Value: &String{Value: ""}, IsConst: false})
	member.Set("init", &Symbol{Name: "init", Value: &TYPEERRORINIT, IsConst: true})

	return &Class{
		Name:   "TypeError",
		Parent: ErrorClass,
		Member: member,
	}
}

// IndexErrorClass 表示 IndexError 内置类的类定义
var IndexErrorClass = initIndexErrorClass()

// initIndexErrorClass 初始化 IndexError 类
func initIndexErrorClass() *Class {
	member := &Environment{
		Name:  "IndexError",
		Store: map[string]*Symbol{},
		Outer: nil,
	}

	member.Set("message", &Symbol{Name: "message", Value: &String{Value: ""}, IsConst: false})
	member.Set("init", &Symbol{Name: "init", Value: &INDEXERRORINIT, IsConst: true})

	return &Class{
		Name:   "IndexError",
		Parent: ErrorClass,
		Member: member,
	}
}

// VariableErrorClass 表示 VariableError 内置类的类定义
var VariableErrorClass = initVariableErrorClass()

// initVariableErrorClass 初始化 VariableError 类
func initVariableErrorClass() *Class {
	member := &Environment{
		Name:  "VariableError",
		Store: map[string]*Symbol{},
		Outer: nil,
	}

	member.Set("message", &Symbol{Name: "message", Value: &String{Value: ""}, IsConst: false})
	member.Set("init", &Symbol{Name: "init", Value: &VARIABLEERRORINIT, IsConst: true})

	return &Class{
		Name:   "VariableError",
		Parent: ErrorClass,
		Member: member,
	}
}

// ArgumentErrorClass 表示 ArgumentError 内置类的类定义
var ArgumentErrorClass = initArgumentErrorClass()

// initArgumentErrorClass 初始化 ArgumentError 类
func initArgumentErrorClass() *Class {
	member := &Environment{
		Name:  "ArgumentError",
		Store: map[string]*Symbol{},
		Outer: nil,
	}

	member.Set("message", &Symbol{Name: "message", Value: &String{Value: ""}, IsConst: false})
	member.Set("init", &Symbol{Name: "init", Value: &ARGUMENTERRORINIT, IsConst: true})

	return &Class{
		Name:   "ArgumentError",
		Parent: ErrorClass,
		Member: member,
	}
}

// ModuleErrorClass 表示 ModuleError 内置类的类定义
var ModuleErrorClass = initModuleErrorClass()

// initModuleErrorClass 初始化 ModuleError 类
func initModuleErrorClass() *Class {
	member := &Environment{
		Name:  "ModuleError",
		Store: map[string]*Symbol{},
		Outer: nil,
	}

	member.Set("message", &Symbol{Name: "message", Value: &String{Value: ""}, IsConst: false})
	member.Set("init", &Symbol{Name: "init", Value: &MODULEERRORINIT, IsConst: true})

	return &Class{
		Name:   "ModuleError",
		Parent: ErrorClass,
		Member: member,
	}
}

var (
	// ERRORINIT 表示 Error 基类的 init 方法
	ERRORINIT = BuiltinFunction{
		Name:         "init",
		Parameter:    []string{"message"},
		DefaultValue: []Object{&String{Value: ""}},
		HaveVariadic: false,
		Fn:           ErrorInit,
	}
	// OPERATIONERRORINIT 表示 OperationError 类的 init 方法
	OPERATIONERRORINIT = BuiltinFunction{
		Name:         "init",
		Parameter:    []string{"message"},
		DefaultValue: []Object{&String{Value: ""}},
		HaveVariadic: false,
		Fn:           OperationErrorInit,
	}
	// MATHERRORINIT 表示 MathError 类的 init 方法
	MATHERRORINIT = BuiltinFunction{
		Name:         "init",
		Parameter:    []string{"message"},
		DefaultValue: []Object{&String{Value: ""}},
		HaveVariadic: false,
		Fn:           MathErrorInit,
	}
	// TYPEERRORINIT 表示 TypeError 类的 init 方法
	TYPEERRORINIT = BuiltinFunction{
		Name:         "init",
		Parameter:    []string{"message"},
		DefaultValue: []Object{&String{Value: ""}},
		HaveVariadic: false,
		Fn:           TypeErrorInit,
	}
	// INDEXERRORINIT 表示 IndexError 类的 init 方法
	INDEXERRORINIT = BuiltinFunction{
		Name:         "init",
		Parameter:    []string{"message"},
		DefaultValue: []Object{&String{Value: ""}},
		HaveVariadic: false,
		Fn:           IndexErrorInit,
	}
	// VARIABLEERRORINIT 表示 VariableError 类的 init 方法
	VARIABLEERRORINIT = BuiltinFunction{
		Name:         "init",
		Parameter:    []string{"message"},
		DefaultValue: []Object{&String{Value: ""}},
		HaveVariadic: false,
		Fn:           VariableErrorInit,
	}
	// ARGUMENTERRORINIT 表示 ArgumentError 类的 init 方法
	ARGUMENTERRORINIT = BuiltinFunction{
		Name:         "init",
		Parameter:    []string{"message"},
		DefaultValue: []Object{&String{Value: ""}},
		HaveVariadic: false,
		Fn:           ArgumentErrorInit,
	}
	// MODULEERRORINIT 表示 ModuleError 类的 init 方法
	MODULEERRORINIT = BuiltinFunction{
		Name:         "init",
		Parameter:    []string{"message"},
		DefaultValue: []Object{&String{Value: ""}},
		HaveVariadic: false,
		Fn:           ModuleErrorInit,
	}
)

// ErrorInit 实现 Error 基类的 init 构造方法
func ErrorInit(f *frame.Frame, env *Environment, posStart, posEnd *util.Pos, args ...Object) (Object, error) {
	this, ok := args[0].(*Instance)
	if this == nil || !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "method init() called without instance.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	message := args[1].(*String)
	if message == nil {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "message must be a string.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	// 将 message 存储到实例的成员中
	this.Member.Set("message", &Symbol{
		Name:    "message",
		Value:   message,
		IsConst: false,
	})

	return &Null{}, nil
}

// OperationErrorInit 实现 OperationError 类的 init 构造方法
func OperationErrorInit(f *frame.Frame, env *Environment, posStart, posEnd *util.Pos, args ...Object) (Object, error) {
	return ErrorInit(f, env, posStart, posEnd, args...)
}

// MathErrorInit 实现 MathError 类的 init 构造方法
func MathErrorInit(f *frame.Frame, env *Environment, posStart, posEnd *util.Pos, args ...Object) (Object, error) {
	return ErrorInit(f, env, posStart, posEnd, args...)
}

// TypeErrorInit 实现 TypeError 类的 init 构造方法
func TypeErrorInit(f *frame.Frame, env *Environment, posStart, posEnd *util.Pos, args ...Object) (Object, error) {
	return ErrorInit(f, env, posStart, posEnd, args...)
}

// IndexErrorInit 实现 IndexError 类的 init 构造方法
func IndexErrorInit(f *frame.Frame, env *Environment, posStart, posEnd *util.Pos, args ...Object) (Object, error) {
	return ErrorInit(f, env, posStart, posEnd, args...)
}

// VariableErrorInit 实现 VariableError 类的 init 构造方法
func VariableErrorInit(f *frame.Frame, env *Environment, posStart, posEnd *util.Pos, args ...Object) (Object, error) {
	return ErrorInit(f, env, posStart, posEnd, args...)
}

// ArgumentErrorInit 实现 ArgumentError 类的 init 构造方法
func ArgumentErrorInit(f *frame.Frame, env *Environment, posStart, posEnd *util.Pos, args ...Object) (Object, error) {
	return ErrorInit(f, env, posStart, posEnd, args...)
}

// ModuleErrorInit 实现 ModuleError 类的 init 构造方法
func ModuleErrorInit(f *frame.Frame, env *Environment, posStart, posEnd *util.Pos, args ...Object) (Object, error) {
	return ErrorInit(f, env, posStart, posEnd, args...)
}

// ConvertErrorToInstance 将错误类转换为实例
//
// 参数:
//
//	err - 错误
//	env - 执行环境
//
// 返回值:
//
//	Object - 转换后的实例
func ConvertErrorToInstance(err error, env *Environment) Object {
	switch e := err.(type) {
	case *errors.OperationError:
		return &Instance{
			Class: OperationErrorClass,
			Member: &Environment{
				Name: "instance of class OperationError",
				Store: map[string]*Symbol{
					"message": {
						Name:    "message",
						Value:   &String{Value: e.Message},
						IsConst: false,
					},
					"init": {
						Name:    "init",
						Value:   &OPERATIONERRORINIT,
						IsConst: true,
					},
				},
				Outer: env,
			},
		}
	case *errors.MathError:
		return &Instance{
			Class: MathErrorClass,
			Member: &Environment{
				Name: "instance of class MathError",
				Store: map[string]*Symbol{
					"message": {
						Name:    "message",
						Value:   &String{Value: e.Message},
						IsConst: false,
					},
					"init": {
						Name:    "init",
						Value:   &MATHERRORINIT,
						IsConst: true,
					},
				},
				Outer: env,
			},
		}
	case *errors.TypeError:
		return &Instance{
			Class: TypeErrorClass,
			Member: &Environment{
				Name: "instance of class TypeError",
				Store: map[string]*Symbol{
					"message": {
						Name:    "message",
						Value:   &String{Value: e.Message},
						IsConst: false,
					},
					"init": {
						Name:    "init",
						Value:   &TYPEERRORINIT,
						IsConst: true,
					},
				},
				Outer: env,
			},
		}
	case *errors.IndexError:
		return &Instance{
			Class: IndexErrorClass,
			Member: &Environment{
				Name: "instance of class IndexError",
				Store: map[string]*Symbol{
					"message": {
						Name:    "message",
						Value:   &String{Value: e.Message},
						IsConst: false,
					},
					"init": {
						Name:    "init",
						Value:   &INDEXERRORINIT,
						IsConst: true,
					},
				},
				Outer: env,
			},
		}
	case *errors.VariableError:
		return &Instance{
			Class: VariableErrorClass,
			Member: &Environment{
				Name: "instance of class VariableError",
				Store: map[string]*Symbol{
					"message": {
						Name:    "message",
						Value:   &String{Value: e.Message},
						IsConst: false,
					},
					"init": {
						Name:    "init",
						Value:   &VARIABLEERRORINIT,
						IsConst: true,
					},
				},
				Outer: env,
			},
		}
	case *errors.ArgumentError:
		return &Instance{
			Class: ArgumentErrorClass,
			Member: &Environment{
				Name: "instance of class ArgumentError",
				Store: map[string]*Symbol{
					"message": {
						Name:    "message",
						Value:   &String{Value: e.Message},
						IsConst: false,
					},
					"init": {
						Name:    "init",
						Value:   &ARGUMENTERRORINIT,
						IsConst: true,
					},
				},
				Outer: env,
			},
		}
	case *errors.ModuleError:
		return &Instance{
			Class: ModuleErrorClass,
			Member: &Environment{
				Name: "instance of class ModuleError",
				Store: map[string]*Symbol{
					"message": {
						Name:    "message",
						Value:   &String{Value: e.Message},
						IsConst: false,
					},
					"init": {
						Name:    "init",
						Value:   &MODULEERRORINIT,
						IsConst: true,
					},
				},
				Outer: env,
			},
		}
	case *UserError:
		return e.Err
	default:
		return &Instance{
			Class: ErrorClass,
			Member: &Environment{
				Name: "instance of class Error",
				Store: map[string]*Symbol{
					"message": {
						Name:    "message",
						Value:   &String{Value: e.Error()},
						IsConst: false,
					},
					"init": {
						Name:    "init",
						Value:   &ERRORINIT,
						IsConst: true,
					},
				},
				Outer: env,
			},
		}
	}
}
