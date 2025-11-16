package object

import (
	"github.com/Ghost-Xiao/ghost-lang/internal/frame"
	"github.com/Ghost-Xiao/ghost-lang/internal/util"
)

// Null 表示空值类型，实现了Object接口
// 用于表示不存在的值或空值

type Null struct{}

// Type 返回值的类型
//
// 返回值:
//
//	string - 值的类型
func (n *Null) Type() string {
	return "Null"
}

// String 返回值的字符串表示
//
// 返回值:
//
//	string - 格式化的字符串表示
func (n *Null) String() string {
	return "null"
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
func (n *Null) Negative(posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (n *Null) BitNot(posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (n *Null) Not(posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
//	other - 另一个操作数
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	Object - 运算结果
//	error - 可能出现的错误
func (n *Null) Add(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	return nil, &OperationError{
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
func (n *Null) Subtract(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	return nil, &OperationError{
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
func (n *Null) Multiply(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	return nil, &OperationError{
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
func (n *Null) Divide(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (n *Null) Mod(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	return nil, &OperationError{
		Frame:    frame,
		Message:  "invalid operation \"%\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// Equal 判断当前值与另一个值是否相等
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
//   - 与*Null类型比较：返回true
//   - 与其他类型比较：返回false
func (n *Null) Equal(other Object, _, _ *util.Pos, _ *frame.Frame) (Object, error) {
	// 空值比较规则: 仅与另一个空值相等
	switch other.(type) {
	case *Null:
		// 与Null类型比较: 返回true
		return &Bool{Value: true}, nil
	default:
		// 与非Null类型比较: 返回false
		return &Bool{Value: false}, nil
	}
}

// NotEqual 判断当前值与另一个值是否不相等
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
//   - 与*Null类型比较：返回false
//   - 与其他类型比较：返回true
func (n *Null) NotEqual(other Object, _, _ *util.Pos, _ *frame.Frame) (Object, error) {
	// 空值不等比较规则: 仅与另一个空值不相等
	switch other.(type) {
	case *Null:
		// 与Null类型比较: 返回false
		return &Bool{Value: false}, nil
	default:
		// 与非Null类型比较: 返回true
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
func (n *Null) LessThan(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (n *Null) GreaterThan(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (n *Null) LessThanOrEqual(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (n *Null) GreaterThanOrEqual(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (n *Null) BitAnd(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (n *Null) BitOr(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (n *Null) Xor(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (n *Null) LeftShift(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (n *Null) RightShift(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (n *Null) And(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (n *Null) Or(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (n *Null) Index(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	return nil, &TypeError{
		Frame:    frame,
		Message:  "index expression not supported for this type.",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}
