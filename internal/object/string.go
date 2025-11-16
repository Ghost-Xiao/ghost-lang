package object

import (
	"math"
	"strings"

	"github.com/Ghost-Xiao/ghost-lang/internal/frame"
	"github.com/Ghost-Xiao/ghost-lang/internal/util"
)

// String 表示字符串值类型，实现了Object接口
// 用于存储和操作文本数据

type String struct {
	Value string // 字符串的实际值
}

// Type 返回值的类型
//
// 返回值:
//
//	string - 值的类型
func (s *String) Type() string {
	return "String"
}

// String 返回值的字符串表示
//
// 返回值:
//
//	string - 格式化的字符串表示
func (s *String) String() string {
	return s.Value
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
func (s *String) Negative(posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	return nil, &OperationError{
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
func (s *String) BitNot(posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (s *String) Not(posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	return nil, &OperationError{
		Frame:    frame,
		Message:  "invalid operation \"!\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// Add 实现字符串的加法运算
//
// 参数:
//
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	运算结果；若操作失败则返回运行时错误
//
// 支持的操作:
//
//   - 与*String类型相加：返回连接后的新字符串
//   - 与其他类型相加：返回操作错误
func (s *String) Add(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	// 字符串加法运算: 仅支持与另一个字符串拼接
	switch o := other.(type) {
	case *String:
		// 与字符串类型相加: 返回连接后的新字符串
		return &String{Value: s.Value + o.Value}, nil
	default:
		// 与非字符串类型相加: 返回操作错误
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
//	other - 另一个操作数
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	Object - 运算结果
//	error - 可能出现的错误
func (s *String) Subtract(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	return nil, &OperationError{
		Frame:    frame,
		Message:  "invalid operation \"-\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// Multiply 实现字符串的乘法运算
//
// 参数:
//
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	运算结果；若操作失败则返回运行时错误
//
// 支持的操作:
//
//   - 与*Int类型相乘：返回重复指定次数的新字符串
//   - 若整数为负数或超过math.MaxInt：返回操作错误
//   - 与其他类型相乘：返回操作错误
func (s *String) Multiply(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	// 字符串乘法运算: 仅支持与整数相乘，表示重复次数
	switch o := other.(type) {
	case *Int:
		// 检查乘数是否为负数
		if o.Value < 0 {
			return nil, &OperationError{
				Frame:    frame,
				Message:  "invalid operation \"*\".",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
		// 检查乘数是否超出最大整数限制
		if o.Value > math.MaxInt {
			return nil, &OperationError{
				Frame:    frame,
				Message:  "invalid operation \"*\".",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
		// 执行字符串重复操作
		return &String{Value: strings.Repeat(s.Value, int(o.Value))}, nil
	default:
		// 与非整数类型相乘: 返回操作错误
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
//	other - 另一个操作数
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	Object - 运算结果
//	error - 可能出现的错误
func (s *String) Divide(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	return nil, &OperationError{
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
func (s *String) Mod(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	return nil, &OperationError{
		Frame:    frame,
		Message:  "invalid operation \"%\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// Equal 判断当前字符串与另一个值是否相等
//
// 参数:
//
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	布尔值，表示比较结果；若操作失败则返回运行时错误
//
// 支持的比较:
//
//   - 与*String类型比较：比较字符串内容是否相同
//   - 与*Null类型比较：始终返回false
//   - 与其他类型比较：返回操作错误
func (s *String) Equal(other Object, _, _ *util.Pos, _ *frame.Frame) (Object, error) {
	// 字符串相等比较: 支持与字符串和空值比较
	switch o := other.(type) {
	case *String:
		// 与字符串类型比较: 比较内容是否相同
		return &Bool{Value: s.Value == o.Value}, nil
	default:
		// 与其他类型比较：返回false
		return &Bool{Value: false}, nil
	}
}

// NotEqual 判断当前字符串与另一个值是否不相等
//
// 参数:
//
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	布尔值，表示比较结果；若操作失败则返回运行时错误
//
// 支持的比较:
//
//   - 与*String类型比较：比较字符串内容是否不同
//   - 与*Null类型比较：始终返回true
//   - 与其他类型比较：返回操作错误
func (s *String) NotEqual(other Object, _, _ *util.Pos, _ *frame.Frame) (Object, error) {
	// 字符串不等比较: 支持与字符串和空值比较
	switch o := other.(type) {
	case *String:
		// 与字符串类型比较: 比较内容是否不同
		return &Bool{Value: s.Value != o.Value}, nil
	default:
		// 与其他类型比较：返回true
		return &Bool{Value: true}, nil
	}
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
func (s *String) LessThan(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	return nil, &OperationError{
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
func (s *String) GreaterThan(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	return nil, &OperationError{
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
func (s *String) LessThanOrEqual(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	return nil, &OperationError{
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
func (s *String) GreaterThanOrEqual(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	return nil, &OperationError{
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
func (s *String) BitAnd(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (s *String) BitOr(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (s *String) Xor(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (s *String) LeftShift(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (s *String) RightShift(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (s *String) And(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (s *String) Or(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (s *String) Index(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	// 以 rune 为单位的索引，支持 Unicode
	runes := []rune(s.Value)
	length := int64(len(runes))
	real := other.(*Int).Value
	if real < 0 {
		real = length + real
	}
	if real < 0 || real >= length {
		return nil, &IndexError{
			Frame:    frame,
			Message:  "index out of range.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	return &String{Value: string(runes[int(real)])}, nil
}

// Set 设置索引位置的值
//
// 参数:
//
//	index - 索引值
//	value - 要设置的值
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	error - 可能出现的错误
func (s *String) Set(index Object, value Object, posStart, posEnd *util.Pos, frame *frame.Frame) error {
	length := int64(len(s.Value))
	real := index.(*Int).Value
	if real < 0 {
		real = length + real
	}
	if real < 0 || real >= length {
		return &IndexError{
			Frame:    frame,
			Message:  "index out of range.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	if str, ok := value.(*String); ok {
		s.Value = s.Value[:int(real)] + str.Value + s.Value[int(real)+1:]
		return nil
	}
	return &TypeError{
		Frame:    frame,
		Message:  "invalid assignment.",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}
