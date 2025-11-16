package object

import (
	"github.com/Ghost-Xiao/ghost-lang/internal/frame"
	"github.com/Ghost-Xiao/ghost-lang/internal/util"
)

// Value 运行时值接口，定义所有可计算值的通用操作
// 实现此接口的类型包括整数、浮点数、字符串、函数等

type Object interface {
	// Type 返回值的类型
	//
	// 返回值:
	//
	//  string - 值的类型
	Type() string

	// String 返回值的字符串表示
	//
	// 返回值:
	//
	//  string - 格式化的字符串表示
	String() string

	// Negative 对值进行负运算
	//
	// 参数:
	//
	//  posStart - 表达式起始位置
	//  posEnd - 表达式结束位置
	//  frame - 当前调用栈
	//
	// 返回值:
	//
	//  Object - 运算结果
	//  error - 可能出现的错误
	Negative(posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error)

	// BitNot 对值进行按位非运算
	//
	// 参数:
	//
	//  posStart - 表达式起始位置
	//  posEnd - 表达式结束位置
	//  frame - 当前调用栈
	//
	// 返回值:
	//
	//  Object - 运算结果
	//  error - 可能出现的错误
	BitNot(posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error)

	// Not 对值进行逻辑非运算
	//
	// 参数:
	//
	//  posStart - 表达式起始位置
	//  posEnd - 表达式结束位置
	//  frame - 当前调用栈
	//
	// 返回值:
	//
	//  Object - 运算结果
	//  error - 可能出现的错误
	Not(posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error)

	// Add 对值进行加法运算
	//
	// 参数:
	//
	//  other - 另一个操作数
	//  posStart - 表达式起始位置
	//  posEnd - 表达式结束位置
	//  frame - 当前调用栈
	//
	// 返回值:
	//
	//  Object - 运算结果
	//  error - 可能出现的错误
	Add(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error)

	// Subtract 对值进行减法运算
	//
	// 参数:
	//
	//  other - 另一个操作数
	//  posStart - 表达式起始位置
	//  posEnd - 表达式结束位置
	//  frame - 当前调用栈
	//
	// 返回值:
	//
	//  Object - 运算结果
	//  error - 可能出现的错误
	Subtract(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error)

	// Multiply 对值进行乘法运算
	//
	// 参数:
	//
	//  other - 另一个操作数
	//  posStart - 表达式起始位置
	//  posEnd - 表达式结束位置
	//  frame - 当前调用栈
	//
	// 返回值:
	//
	//  Object - 运算结果
	//  error - 可能出现的错误
	Multiply(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error)

	// Divide 对值进行除法运算
	//
	// 参数:
	//
	//  other - 另一个操作数
	//  posStart - 表达式起始位置
	//  posEnd - 表达式结束位置
	//  frame - 当前调用栈
	//
	// 返回值:
	//
	//  Object - 运算结果
	//  error - 可能出现的错误
	Divide(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error)

	// Mod 对值进行取模运算
	//
	// 参数:
	//
	//  other - 另一个操作数
	//  posStart - 表达式起始位置
	//  posEnd - 表达式结束位置
	//  frame - 当前调用栈
	//
	// 返回值:
	//
	//  Object - 运算结果
	//  error - 可能出现的错误
	Mod(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error)

	// Equal 对值进行等于比较
	//
	// 参数:
	//
	//  other - 另一个操作数
	//  posStart - 表达式起始位置
	//  posEnd - 表达式结束位置
	//  frame - 当前调用栈
	//
	// 返回值:
	//
	//  Object - 比较结果
	Equal(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error)

	// NotEqual 对值进行不等于比较
	//
	// 参数:
	//
	//  other - 另一个操作数
	//  posStart - 表达式起始位置
	//  posEnd - 表达式结束位置
	//  frame - 当前调用栈
	//
	// 返回值:
	//
	//  Object - 比较结果
	NotEqual(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error)

	// LessThan 对值进行小于比较
	//
	// 参数:
	//
	//  other - 另一个操作数
	//  posStart - 表达式起始位置
	//  posEnd - 表达式结束位置
	//  frame - 当前调用栈
	//
	// 返回值:
	//
	//  Object - 比较结果
	LessThan(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error)

	// GreaterThan 对值进行大于比较
	//
	// 参数:
	//
	//  other - 另一个操作数
	//  posStart - 表达式起始位置
	//  posEnd - 表达式结束位置
	//  frame - 当前调用栈
	//
	// 返回值:
	//
	//  Object - 比较结果
	GreaterThan(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error)

	// LessThanOrEqual 对值进行小于等于比较
	//
	// 参数:
	//
	//  other - 另一个操作数
	//  posStart - 表达式起始位置
	//  posEnd - 表达式结束位置
	//  frame - 当前调用栈
	//
	// 返回值:
	//
	//  Object - 比较结果
	LessThanOrEqual(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error)

	// GreaterThanOrEqual 对值进行大于等于比较
	//
	// 参数:
	//
	//  other - 另一个操作数
	//  posStart - 表达式起始位置
	//  posEnd - 表达式结束位置
	//  frame - 当前调用栈
	//
	// 返回值:
	//
	//  Object - 比较结果
	GreaterThanOrEqual(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error)

	// BitAnd 对值进行按位与运算
	//
	// 参数:
	//
	//  other - 另一个操作数
	//  posStart - 表达式起始位置
	//  posEnd - 表达式结束位置
	//  frame - 当前调用栈
	//
	// 返回值:
	//
	//  Object - 运算结果
	//  error - 可能出现的错误
	BitAnd(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error)

	// BitOr 对值进行按位或运算
	//
	// 参数:
	//
	//  other - 另一个操作数
	//  posStart - 表达式起始位置
	//  posEnd - 表达式结束位置
	//  frame - 当前调用栈
	//
	// 返回值:
	//
	//  Object - 运算结果
	//  error - 可能出现的错误
	BitOr(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error)

	// Xor 对值进行异或运算
	//
	// 参数:
	//
	//  other - 另一个操作数
	//  posStart - 表达式起始位置
	//  posEnd - 表达式结束位置
	//  frame - 当前调用栈
	//
	// 返回值:
	//
	//  Object - 运算结果
	//  error - 可能出现的错误
	Xor(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error)

	// LeftShift 对值进行左移运算
	//
	// 参数:
	//
	//  other - 另一个操作数
	//  posStart - 表达式起始位置
	//  posEnd - 表达式结束位置
	//  frame - 当前调用栈
	//
	// 返回值:
	//
	//  Object - 运算结果
	//  error - 可能出现的错误
	LeftShift(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error)

	// RightShift 对值进行右移运算
	//
	// 参数:
	//
	//  other - 另一个操作数
	//  posStart - 表达式起始位置
	//  posEnd - 表达式结束位置
	//  frame - 当前调用栈
	//
	// 返回值:
	//
	//  Object - 运算结果
	//  error - 可能出现的错误
	RightShift(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error)

	// And 对值进行逻辑与运算
	//
	// 参数:
	//
	//  other - 另一个操作数
	//  posStart - 表达式起始位置
	//  posEnd - 表达式结束位置
	//  frame - 当前调用栈
	//
	// 返回值:
	//
	//  Object - 运算结果
	//  error - 可能出现的错误
	And(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error)

	// Or 对值进行逻辑或运算
	//
	// 参数:
	//
	//  other - 另一个操作数
	//  posStart - 表达式起始位置
	//  posEnd - 表达式结束位置
	//  frame - 当前调用栈
	//
	// 返回值:
	//
	//  Object - 运算结果
	//  error - 可能出现的错误
	Or(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error)

	// Index 对值进行索引运算
	//
	// 参数:
	//
	//  other - 另一个操作数
	//  posStart - 表达式起始位置
	//  posEnd - 表达式结束位置
	//  frame - 当前调用栈
	//
	// 返回值:
	//
	//  Object - 运算结果
	//  error - 可能出现的错误
	Index(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error)
}
