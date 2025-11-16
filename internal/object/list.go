package object

import (
	"strings"

	"github.com/Ghost-Xiao/ghost-lang/internal/frame"
	"github.com/Ghost-Xiao/ghost-lang/internal/util"
)

// List 列表类型结构体，表示运行时的列表
// 实现Object接口

type List struct {
	Elements []Object // 列表元素
}

// Type 返回值的类型
//
// 返回值:
//
//	string - 值的类型
func (l *List) Type() string {
	return "LIST"
}

// String 返回值的字符串表示
//
// 返回值:
//
//	string - 格式化的字符串表示
func (l *List) String() string {
	var elements []string
	for _, elem := range l.Elements {
		elements = append(elements, elem.String())
	}
	return "[" + strings.Join(elements, ", ") + "]"
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
func (l *List) Negative(posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (l *List) BitNot(posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (l *List) Not(posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (l *List) Add(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	if otherList, ok := other.(*List); ok {
		// 检查列表元素类型一致性
		if len(l.Elements) > 0 && len(otherList.Elements) > 0 {
			if l.Elements[0].Type() != otherList.Elements[0].Type() {
				return nil, &OperationError{
					Frame:    frame,
					Message:  "cannot concatenate lists with different element types.",
					PosStart: posStart,
					PosEnd:   posEnd,
				}
			}
		}
		// 创建新列表
		newElements := make([]Object, 0, len(l.Elements)+len(otherList.Elements))
		newElements = append(newElements, l.Elements...)
		newElements = append(newElements, otherList.Elements...)
		return &List{Elements: newElements}, nil
	}
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
func (l *List) Subtract(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
//	Object - 运算结果（重复后的新列表）
//	error - 可能出现的错误
func (l *List) Multiply(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	if intObj, ok := other.(*Int); ok {
		times := intObj.Value
		// 负数或零次重复返回错误
		if times <= 0 {
			return nil, &OperationError{
				Frame:    frame,
				Message:  "invalid operation \"*\".",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
		// 空列表重复任意次都是空列表
		if len(l.Elements) == 0 {
			return &List{Elements: make([]Object, 0)}, nil
		}
		// 创建新的元素切片
		newElements := make([]Object, 0, len(l.Elements)*int(times))
		// 重复添加原列表元素
		for i := int64(0); i < times; i++ {
			newElements = append(newElements, l.Elements...)
		}
		return &List{Elements: newElements}, nil
	}
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
func (l *List) Divide(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (l *List) Mod(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (l *List) Equal(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	if otherList, ok := other.(*List); ok {
		if len(l.Elements) != len(otherList.Elements) {
			return &Bool{Value: false}, nil
		}
		for i := range l.Elements {
			equal, err := l.Elements[i].Equal(otherList.Elements[i], posStart, posEnd, frame)
			if err != nil {
				return nil, err
			}
			if !equal.(*Bool).Value {
				return &Bool{Value: false}, nil
			}
		}
		return &Bool{Value: true}, nil
	}
	return &Bool{Value: false}, nil
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
func (l *List) NotEqual(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	equal, err := l.Equal(other, posStart, posEnd, frame)
	if err != nil {
		return nil, err
	}
	return &Bool{Value: !equal.(*Bool).Value}, nil
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
func (l *List) LessThan(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (l *List) GreaterThan(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (l *List) LessThanOrEqual(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (l *List) GreaterThanOrEqual(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (l *List) BitAnd(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (l *List) BitOr(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (l *List) Xor(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (l *List) LeftShift(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (l *List) RightShift(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (l *List) And(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (l *List) Or(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (l *List) Index(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	length := int64(len(l.Elements))
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
	return l.Elements[int(real)], nil
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
func (l *List) Set(index Object, value Object, posStart, posEnd *util.Pos, frame *frame.Frame) error {
	length := int64(len(l.Elements))
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
	if value.Type() != l.Elements[0].Type() {
		return &TypeError{
			Frame:    frame,
			Message:  "list elements must have consistent types.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	l.Elements[int(real)] = value
	return nil
}
