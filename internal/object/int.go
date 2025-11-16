package object

import (
	"fmt"
	"math"
	"strings"

	"github.com/Ghost-Xiao/ghost-lang/internal/frame"
	"github.com/Ghost-Xiao/ghost-lang/internal/util"
)

// Int 整数类型结构体，表示运行时的整数数值
// 实现Number接口和Object接口，支持各种整数运算

type Int struct {
	Value int64 // 整数实际值
}

// Type 返回值的类型
//
// 返回值:
//
//	string - 值的类型
func (i *Int) Type() string {
	return "Int"
}

// String 返回值的字符串表示
//
// 返回值:
//
//	string - 格式化的字符串表示
func (i *Int) String() string {
	return fmt.Sprintf("%d", i.Value)
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
func (i *Int) Negative(*util.Pos, *util.Pos, *frame.Frame) (Object, error) {
	return &Int{Value: -i.Value}, nil
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
func (i *Int) BitNot(*util.Pos, *util.Pos, *frame.Frame) (Object, error) {
	return &Int{Value: ^i.Value}, nil
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
func (i *Int) Not(posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	return nil, &OperationError{
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
//	other - 右操作数，可以是Int或Float类型
//	posStart - 节点起始位置
//	posEnd - 节点结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	Object - 运算结果
//	error - 可能出现的错误
func (i *Int) Add(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	// 根据右操作数类型执行不同加法逻辑
	switch o := other.(type) {
	case *Int:
		// 整数+整数=整数
		return &Int{Value: i.Value + o.Value}, nil
	case *Float:
		// 整数+浮点数=浮点数
		return &Float{Value: float64(i.Value) + o.Value}, nil
	default:
		// 不支持的操作数类型
		return nil, &OperationError{
			Frame:    frame,
			Message:  "invalid operation \"+\".",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
}

// Subtract 对值进行减法运算
//
// 参数:
//
//	other - 右操作数，可以是Int或Float类型
//	posStart - 节点起始位置
//	posEnd - 节点结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	Object - 运算结果
//	error - 可能出现的错误
func (i *Int) Subtract(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	// 根据右操作数类型执行不同减法逻辑
	switch o := other.(type) {
	case *Int:
		// 整数-整数=整数
		return &Int{Value: i.Value - o.Value}, nil
	case *Float:
		// 整数-浮点数=浮点数
		return &Float{Value: float64(i.Value) - o.Value}, nil
	default:
		// 不支持的操作数类型
		return nil, &OperationError{
			Frame:    frame,
			Message:  "invalid operation \"-\".",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
}

// Multiply 对值进行乘法运算
//
// 参数:
//
//	other - 右操作数，可以是Int、Float或String类型
//	posStart - 节点起始位置
//	posEnd - 节点结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	Object - 运算结果
//	error - 可能出现的错误
func (i *Int) Multiply(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	// 根据右操作数类型执行不同乘法逻辑
	switch o := other.(type) {
	case *Int:
		// 整数*整数=整数
		return &Int{Value: i.Value * o.Value}, nil
	case *Float:
		// 整数*浮点数=浮点数
		return &Float{Value: float64(i.Value) * o.Value}, nil
	case *String:
		// 整数*字符串=重复字符串(仅支持非负整数)
		if i.Value < 0 {
			return nil, &OperationError{
				Frame:    frame,
				Message:  "invalid operation \"*\".",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
		// 检查整数是否超出最大可表示范围
		if i.Value > math.MaxInt {
			return nil, &OperationError{
				Frame:    frame,
				Message:  "invalid operation \"*\".",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
		// 重复字符串指定次数
		return &String{Value: strings.Repeat(o.Value, int(i.Value))}, nil
	default:
		// 不支持的操作数类型
		return nil, &OperationError{
			Frame:    frame,
			Message:  "invalid operation \"*\".",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
}

// Divide 对值进行除法运算
//
// 参数:
//
//	other - 右操作数，可以是Int或Float类型
//	posStart - 节点起始位置
//	posEnd - 节点结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	Object - 运算结果
//	error - 可能出现的错误
func (i *Int) Divide(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	// 根据右操作数类型执行不同除法逻辑
	switch o := other.(type) {
	case *Int:
		// 整数除法，除数为0时返回错误
		if o.Value == 0 {
			return nil, &MathError{
				Frame:    frame,
				Message:  "division by zero.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
		// 整数/整数=浮点数
		return &Float{Value: float64(i.Value) / float64(o.Value)}, nil
	case *Float:
		// 浮点数除法，除数为0时返回错误
		if o.Value == 0 {
			return nil, &MathError{
				Frame:    frame,
				Message:  "division by zero.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
		// 整数/浮点数=浮点数
		return &Float{Value: float64(i.Value) / o.Value}, nil
	default:
		// 不支持的操作数类型
		return nil, &OperationError{
			Frame:    frame,
			Message:  "invalid operation \"/\".",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
}

// Mod 对值进行取模运算
//
// 参数:
//
//	other - 要取模的右侧值
//	posStart - 节点起始位置
//	posEnd - 节点结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	Object - 运算结果
//	error - 可能出现的错误
func (i *Int) Mod(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	// 根据右操作数类型执行不同取模逻辑
	switch o := other.(type) {
	case *Int:
		// 整数取模，除数为0时返回错误
		if o.Value == 0 {
			return nil, &MathError{
				Frame:    frame,
				Message:  "division by zero.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
		// 整数%整数=整数，结果符号与被除数相同
		return &Int{Value: i.Value % o.Value}, nil
	case *Float:
		// 浮点数取模，除数为0时返回错误
		if o.Value == 0 {
			return nil, &MathError{
				Frame:    frame,
				Message:  "division by zero.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
		// 整数%浮点数=浮点数
		return &Float{Value: math.Mod(float64(i.Value), o.Value)}, nil
	default:
		// 不支持的操作数类型
		return nil, &OperationError{
			Frame:    frame,
			Message:  "invalid operation \"%\".",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
}

// Equal 判断当前值与另一个值是否相等
//
// 参数:
//
//	other - 要比较的右侧值
//	posStart - 节点起始位置
//	posEnd - 节点结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	Object - 运算结果
//	error - 可能出现的错误
func (i *Int) Equal(other Object, _, _ *util.Pos, _ *frame.Frame) (Object, error) {
	switch o := other.(type) {
	case *Int:
		// 与整数比较：直接比较整数值
		return &Bool{Value: i.Value == o.Value}, nil
	case *Float:
		// 与浮点数比较：将整数转换为浮点数后比较
		return &Bool{Value: float64(i.Value) == o.Value}, nil
	default:
		// 与其他类型比较：返回false
		return &Bool{Value: false}, nil
	}
}

// NotEqual 判断当前值与另一个值是否不相等
//
// 参数:
//
//	other - 要比较的右侧值
//	posStart - 节点起始位置
//	posEnd - 节点结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	Object - 运算结果
//	error - 可能出现的错误
func (i *Int) NotEqual(other Object, _, _ *util.Pos, _ *frame.Frame) (Object, error) {
	switch o := other.(type) {
	case *Int:
		// 与整数比较：直接比较整数值
		return &Bool{Value: i.Value != o.Value}, nil
	case *Float:
		// 与浮点数比较：将整数转换为浮点数后比较
		return &Bool{Value: float64(i.Value) != o.Value}, nil
	default:
		// 与其他类型比较：返回true
		return &Bool{Value: true}, nil
	}
}

// LessThan 对值进行小于比较
//
// 参数:
//
//	other - 要比较的右侧值
//	posStart - 节点起始位置
//	posEnd - 节点结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	Object - 运算结果
//	error - 可能出现的错误
func (i *Int) LessThan(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	switch o := other.(type) {
	case *Int:
		// 与整数比较：直接比较整数值
		return &Bool{Value: i.Value < o.Value}, nil
	case *Float:
		// 与浮点数比较：将整数转换为浮点数后比较
		return &Bool{Value: float64(i.Value) < o.Value}, nil
	default:
		return nil, &OperationError{
			Frame:    frame,
			Message:  "invalid operation \"<\".",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
}

// GreaterThan 对值进行大于比较
//
// 参数:
//
//	other - 要比较的右侧值
//	posStart - 节点起始位置
//	posEnd - 节点结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	Object - 运算结果
//	error - 可能出现的错误
func (i *Int) GreaterThan(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	switch o := other.(type) {
	case *Int:
		// 与整数比较：直接比较整数值
		return &Bool{Value: i.Value > o.Value}, nil
	case *Float:
		// 与浮点数比较：将整数转换为浮点数后比较
		return &Bool{Value: float64(i.Value) > o.Value}, nil
	default:
		return nil, &OperationError{
			Frame:    frame,
			Message:  "invalid operation \">\".",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
}

// LessThanOrEqual 对值进行小于等于比较
//
// 参数:
//
//	other - 要比较的右侧值
//	posStart - 节点起始位置
//	posEnd - 节点结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	Object - 运算结果
//	error - 可能出现的错误
func (i *Int) LessThanOrEqual(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	switch o := other.(type) {
	case *Int:
		// 与整数比较：直接比较整数值
		return &Bool{Value: i.Value <= o.Value}, nil
	case *Float:
		// 与浮点数比较：将整数转换为浮点数后比较
		return &Bool{Value: float64(i.Value) <= o.Value}, nil
	default:
		return nil, &OperationError{
			Frame:    frame,
			Message:  "invalid operation \"<=\".",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
}

// GreaterThanOrEqual 对值进行大于等于比较
//
// 参数:
//
//	other - 要比较的右侧值
//	posStart - 节点起始位置
//	posEnd - 节点结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	Object - 运算结果
//	error - 可能出现的错误
func (i *Int) GreaterThanOrEqual(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	switch o := other.(type) {
	case *Int:
		// 与整数比较：直接比较整数值
		return &Bool{Value: i.Value >= o.Value}, nil
	case *Float:
		// 与浮点数比较：将整数转换为浮点数后比较
		return &Bool{Value: float64(i.Value) >= o.Value}, nil
	default:
		return nil, &OperationError{
			Frame:    frame,
			Message:  "invalid operation \">=\".",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
}

// BitAnd 对值进行按位与运算
//
// 参数:
//
//	other - 右侧整数值
//	posStart - 节点起始位置
//	posEnd - 节点结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	新的整数值，表示按位与的结果；若操作失败则返回运行时错误
//	error - 可能出现的错误
//
// 注意事项:
//
//	仅支持与*Int类型进行按位与操作，其他类型将返回错误
func (i *Int) BitAnd(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	// 检查右侧操作数是否为整数类型
	if o, ok := other.(*Int); ok {
		// 执行按位与操作并返回结果
		return &Int{Value: i.Value & o.Value}, nil
	} else {
		// 类型不支持，返回操作错误
		return nil, &OperationError{
			Frame:    frame,
			Message:  "invalid operation \"&\".",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
}

// BitOr 对值进行按位或运算
//
// 参数:
//
//	other - 右侧整数值
//	posStart - 节点起始位置
//	posEnd - 节点结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	新的整数值，表示按位或的结果；若操作失败则返回运行时错误
//	error - 可能出现的错误
//
// 注意事项:
//
//	仅支持与*Int类型进行按位或操作，其他类型将返回错误
func (i *Int) BitOr(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	// 检查右侧操作数是否为整数类型
	if o, ok := other.(*Int); ok {
		// 执行按位或操作并返回结果
		return &Int{Value: i.Value | o.Value}, nil
	} else {
		// 类型不支持，返回操作错误
		return nil, &OperationError{
			Frame:    frame,
			Message:  "invalid operation \"|\".",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
}

// Xor 对值进行异或运算
//
// 参数:
//
//	other - 右侧整数值
//	posStart - 节点起始位置
//	posEnd - 节点结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	新的整数值，表示按位异或的结果；若操作失败则返回运行时错误
//	error - 可能出现的错误
//
// 注意事项:
//
//	仅支持与*Int类型进行按位异或操作，其他类型将返回错误
func (i *Int) Xor(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	// 检查右侧操作数是否为整数类型
	if o, ok := other.(*Int); ok {
		// 执行按位异或操作并返回结果
		return &Int{Value: i.Value ^ o.Value}, nil
	} else {
		// 类型不支持，返回操作错误
		return nil, &OperationError{
			Frame:    frame,
			Message:  "invalid operation \"^\".",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
}

// LeftShift 对值进行异或运算
//
// 参数:
//
//	other - 左移的位数
//	posStart - 节点起始位置
//	posEnd - 节点结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	新的整数值，表示左移后的结果；若操作失败则返回运行时错误
//	error - 可能出现的错误
//
// 注意事项:
//
//  1. 仅支持与*Int类型进行左移操作，其他类型将返回错误
//  2. 右操作数不能为负数，否则返回错误
//
// error - 可能出现的错误
func (i *Int) LeftShift(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	// 检查右侧操作数是否为整数类型
	if o, ok := other.(*Int); ok {
		// 检查右操作数是否为负数
		if o.Value < 0 {
			return nil, &OperationError{
				Frame:    frame,
				Message:  "invalid operation \"<<\".",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
		// 执行左移操作并返回结果
		return &Int{Value: i.Value << o.Value}, nil
	} else {
		// 类型不支持，返回操作错误
		return nil, &OperationError{
			Frame:    frame,
			Message:  "invalid operation \"<<\".",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
}

// RightShift 对值进行右移运算
//
// 参数:
//
//	other - 右移的位数
//	posStart - 节点起始位置
//	posEnd - 节点结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	新的整数值，表示右移后的结果；若操作失败则返回运行时错误
//	error - 可能出现的错误
//
// 注意事项:
//
//  1. 仅支持与*Int类型进行右移操作，其他类型将返回错误
//  2. 右操作数不能为负数，否则返回错误
func (i *Int) RightShift(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	// 检查右侧操作数是否为整数类型
	if o, ok := other.(*Int); ok {
		// 检查右操作数是否为负数
		if o.Value < 0 {
			return nil, &OperationError{
				Frame:    frame,
				Message:  "invalid operation \">>\".",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
		// 执行右移操作并返回结果
		return &Int{Value: i.Value >> o.Value}, nil
	} else {
		// 类型不支持，返回操作错误
		return nil, &OperationError{
			Frame:    frame,
			Message:  "invalid operation \">>\".",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
}

// And 对值进行逻辑与运算
//
// 参数:
//
//	other - 右侧表达式
//	posStart - 节点起始位置
//	posEnd - 节点结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	Object - 运算结果
//	error - 可能出现的错误
func (i *Int) And(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	return nil, &OperationError{
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
//	other - 右侧表达式
//	posStart - 节点起始位置
//	posEnd - 节点结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	Object - 运算结果
//	error - 可能出现的错误
func (i *Int) Or(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	return nil, &OperationError{
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
func (i *Int) Index(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	return nil, &TypeError{
		Frame:    frame,
		Message:  "index expression not supported for this type.",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}
