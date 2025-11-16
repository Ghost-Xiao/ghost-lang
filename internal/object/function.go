package object

import (
	"fmt"
	"strings"

	"github.com/Ghost-Xiao/ghost-lang/internal/frame"
	"github.com/Ghost-Xiao/ghost-lang/internal/parser/ast"
	"github.com/Ghost-Xiao/ghost-lang/internal/util"
)

// Function 表示函数类型，实现了Object接口
// 支持的操作包括调用函数等

type Function struct {
	Name      string           // 函数名
	Parameter []*ast.Parameter // 参数
	Body      ast.Statement    // 函数体
	Env       *Environment     // 环境
}

// Type 返回值的类型
//
// 返回值:
//
//	string - 值的类型
func (f *Function) Type() string {
	return "Function"
}

// String 返回值的字符串表示
//
// 返回值:
//
//	string - 格式化的字符串表示
func (f *Function) String() string {
	var params []string
	for _, param := range f.Parameter {
		params = append(params, param.String())
	}
	return fmt.Sprintf("func %s(%s) {...}", f.Name, strings.Join(params, ", "))
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
func (f *Function) Negative(posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (f *Function) BitNot(posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (f *Function) Not(posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (f *Function) Add(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (f *Function) Subtract(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (f *Function) Multiply(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (f *Function) Divide(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (f *Function) Mod(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	return nil, &OperationError{
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
func (f *Function) Equal(other Object, _, _ *util.Pos, _ *frame.Frame) (Object, error) {
	// 函数相等比较规则: 比较引用是否相等
	otherFunc, ok := other.(*Function)
	if !ok {
		return &Bool{Value: false}, nil
	}
	return &Bool{Value: f == otherFunc}, nil
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
func (f *Function) NotEqual(other Object, _, _ *util.Pos, _ *frame.Frame) (Object, error) {
	// 函数不等比较规则: 比较引用是否不等
	otherFunc, ok := other.(*Function)
	if !ok {
		return &Bool{Value: true}, nil
	}
	return &Bool{Value: f != otherFunc}, nil
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
func (f *Function) LessThan(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (f *Function) GreaterThan(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (f *Function) LessThanOrEqual(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (f *Function) GreaterThanOrEqual(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (f *Function) BitAnd(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (f *Function) BitOr(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (f *Function) Xor(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (f *Function) LeftShift(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (f *Function) RightShift(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (f *Function) And(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (f *Function) Or(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (f *Function) Index(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	return nil, &TypeError{
		Frame:    frame,
		Message:  "index expression not supported for this type.",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}
