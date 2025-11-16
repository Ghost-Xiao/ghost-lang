package object

import (
	"fmt"
	"math"

	"github.com/Ghost-Xiao/ghost-lang/internal/frame"
	"github.com/Ghost-Xiao/ghost-lang/internal/util"
)

// Float 表示浮点数值类型，实现了Number和Object接口
// 支持的操作包括算术运算、比较运算等

type Float struct {
	Value float64
}

// Type 返回值的类型
//
// 返回值:
//
//	string - 值的类型
func (f *Float) Type() string {
	return "Float"
}

// String 返回值的字符串表示
//
// 返回值:
//
//	string - 格式化的字符串表示
func (f *Float) String() string {
	return fmt.Sprintf("%f", f.Value)
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
func (f *Float) Negative(*util.Pos, *util.Pos, *frame.Frame) (Object, error) {
	return &Float{Value: -f.Value}, nil
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
func (f *Float) BitNot(posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	return nil, &OperationError{
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
func (f *Float) Not(posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (f *Float) Add(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	switch o := other.(type) {
	case *Int:
		return &Float{Value: f.Value + float64(o.Value)}, nil
	case *Float:
		return &Float{Value: f.Value + o.Value}, nil
	default:
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
func (f *Float) Subtract(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	switch o := other.(type) {
	case *Int:
		return &Float{Value: f.Value - float64(o.Value)}, nil
	case *Float:
		return &Float{Value: f.Value - o.Value}, nil
	default:
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
//	other - 右操作数，可以是Int或Float类型
//	posStart - 节点起始位置
//	posEnd - 节点结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	Object - 运算结果
//	error - 可能出现的错误
func (f *Float) Multiply(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	switch o := other.(type) {
	case *Int:
		// 浮点数 * 整数: 将整数转换为浮点数后相乘
		return &Float{Value: f.Value * float64(o.Value)}, nil
	case *Float:
		// 浮点数 * 浮点数: 直接相乘
		return &Float{Value: f.Value * o.Value}, nil
	default:
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
func (f *Float) Divide(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	switch o := other.(type) {
	case *Int:
		// 浮点数 / 整数: 检查除数是否为0，然后将整数转换为浮点数后相除
		if o.Value == 0 {
			return nil, &MathError{
				Frame:    frame,
				Message:  "division by zero.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
		return &Float{Value: f.Value / float64(o.Value)}, nil
	case *Float:
		// 浮点数 / 浮点数: 检查除数是否为0，然后直接相除
		if o.Value == 0 {
			return nil, &MathError{
				Frame:    frame,
				Message:  "division by zero.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
		return &Float{Value: f.Value / o.Value}, nil
	default:
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
//	other - 右操作数，可以是Int或Float类型
//	posStart - 节点起始位置
//	posEnd - 节点结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	Object - 运算结果
//	error - 可能出现的错误
func (f *Float) Mod(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	switch o := other.(type) {
	case *Int:
		// 浮点数 % 整数: 检查除数是否为0，然后将整数转换为浮点数后取模
		if o.Value == 0 {
			return nil, &MathError{
				Frame:    frame,
				Message:  "division by zero.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
		return &Float{Value: math.Mod(f.Value, float64(o.Value))}, nil
	case *Float:
		// 浮点数 % 浮点数: 检查除数是否为0，然后直接取模
		if o.Value == 0 {
			return nil, &MathError{
				Frame:    frame,
				Message:  "division by zero.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
		return &Float{Value: math.Mod(f.Value, o.Value)}, nil
	default:
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
//	other - 右操作数，可以是Int或Float类型
//	posStart - 节点起始位置
//	posEnd - 节点结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	Object - 运算结果
//	error - 可能出现的错误
func (f *Float) Equal(other Object, _, _ *util.Pos, _ *frame.Frame) (Object, error) {
	switch o := other.(type) {
	case *Int:
		// 浮点数 == 整数: 将整数转换为浮点数后比较
		return &Bool{Value: f.Value == float64(o.Value)}, nil
	case *Float:
		// 浮点数 == 浮点数: 直接比较
		return &Bool{Value: f.Value == o.Value}, nil
	case *Null:
		// 浮点数 == null: 始终返回false
		return &Bool{Value: false}, nil
	default:
		// 与其他类型比较：返回false
		return &Bool{Value: false}, nil
	}
}

// NotEqual 判断当前值与另一个值是否不相等
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
func (f *Float) NotEqual(other Object, _, _ *util.Pos, _ *frame.Frame) (Object, error) {
	switch o := other.(type) {
	case *Int:
		// 浮点数 != 整数: 将整数转换为浮点数后比较
		return &Bool{Value: f.Value != float64(o.Value)}, nil
	case *Float:
		// 浮点数 != 浮点数: 直接比较
		return &Bool{Value: f.Value != o.Value}, nil
	case *Null:
		// 浮点数 != null: 始终返回true
		return &Bool{Value: true}, nil
	default:
		// 与其他类型比较：返回true
		return &Bool{Value: true}, nil
	}
}

// LessThan 对值进行小于比较
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
func (f *Float) LessThan(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	switch o := other.(type) {
	case *Int:
		// 浮点数 < 整数: 将整数转换为浮点数后比较
		return &Bool{Value: f.Value < float64(o.Value)}, nil
	case *Float:
		// 浮点数 < 浮点数: 直接比较
		return &Bool{Value: f.Value < o.Value}, nil
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
//	other - 右操作数，可以是Int或Float类型
//	posStart - 节点起始位置
//	posEnd - 节点结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	Object - 运算结果
//	error - 可能出现的错误
func (f *Float) GreaterThan(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	switch o := other.(type) {
	case *Int:
		// 浮点数 > 整数: 将整数转换为浮点数后比较
		return &Bool{Value: f.Value > float64(o.Value)}, nil
	case *Float:
		// 浮点数 > 浮点数: 直接比较
		return &Bool{Value: f.Value > o.Value}, nil
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
//	other - 右操作数，可以是Int或Float类型
//	posStart - 节点起始位置
//	posEnd - 节点结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	Object - 运算结果
//	error - 可能出现的错误
func (f *Float) LessThanOrEqual(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	switch o := other.(type) {
	case *Int:
		// 浮点数 <= 整数: 将整数转换为浮点数后比较
		return &Bool{Value: f.Value <= float64(o.Value)}, nil
	case *Float:
		// 浮点数 <= 浮点数: 直接比较
		return &Bool{Value: f.Value <= o.Value}, nil
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
//	other - 右操作数，可以是Int或Float类型
//	posStart - 节点起始位置
//	posEnd - 节点结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	Object - 运算结果
//	error - 可能出现的错误
func (f *Float) GreaterThanOrEqual(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	switch o := other.(type) {
	case *Int:
		// 浮点数 >= 整数: 将整数转换为浮点数后比较
		return &Bool{Value: f.Value >= float64(o.Value)}, nil
	case *Float:
		// 浮点数 >= 浮点数: 直接比较
		return &Bool{Value: f.Value >= o.Value}, nil
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
//	other - 另一个操作数
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	Object - 运算结果
//	error - 可能出现的错误
func (f *Float) BitAnd(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	return nil, &OperationError{
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
func (f *Float) BitOr(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	return nil, &OperationError{
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
func (f *Float) Xor(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	return nil, &OperationError{
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
func (f *Float) LeftShift(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	return nil, &OperationError{
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
func (f *Float) RightShift(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	return nil, &OperationError{
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
func (f *Float) And(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
//	other - 另一个操作数
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	Object - 运算结果
//	error - 可能出现的错误
func (f *Float) Or(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (f *Float) Index(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	return nil, &TypeError{
		Frame:    frame,
		Message:  "index expression not supported for this type.",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}
