package object

import (
	"strings"

	"github.com/Ghost-Xiao/ghost-lang/internal/errors"
	"github.com/Ghost-Xiao/ghost-lang/internal/frame"
	"github.com/Ghost-Xiao/ghost-lang/internal/util"
)

// HashKey 映射键类型结构体，表示运行时的映射键的哈希值
type HashKey struct {
	Type  string // 映射键类型
	Value uint64 // 映射键哈希值
}

// Pair 映射键值对类型结构体，表示运行时的映射键值对
type Pair struct {
	Key   Object // 映射键
	Value Object // 映射值
}

// Map 映射类型结构体，表示运行时的映射
// 实现Object接口
type Map struct {
	Pairs map[HashKey]Pair // 映射键值对映射
}

// Type 返回值的类型
//
// 返回值:
//
//	string - 值的类型
func (m *Map) Type() string {
	return "Map"
}

// String 返回值的字符串表示
//
// 返回值:
//
//	string - 格式化的字符串表示
func (m *Map) String() string {
	var pairs []string
	for _, pair := range m.Pairs {
		pairs = append(pairs, pair.Key.String()+": "+pair.Value.String())
	}
	return "{" + strings.Join(pairs, ", ") + "}"
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
func (m *Map) Negative(posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (m *Map) BitNot(posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (m *Map) Not(posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (m *Map) Add(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	if otherMap, ok := other.(*Map); ok {
		// 创建新映射
		newPairs := make(map[HashKey]Pair)
		// 合并键值对
		for hashKey, pair := range m.Pairs {
			newPairs[hashKey] = pair
		}
		for hashKey, pair := range otherMap.Pairs {
			newPairs[hashKey] = pair
		}
		return &Map{Pairs: newPairs}, nil
	}
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
func (m *Map) Subtract(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
//	Object - 运算结果（重复后的新列表）
//	error - 可能出现的错误
func (m *Map) Multiply(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (m *Map) Divide(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (m *Map) Mod(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	return nil, &errors.OperationError{
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
func (m *Map) Equal(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	if otherMap, ok := other.(*Map); ok {
		// 检查键值对数量是否相同
		if len(m.Pairs) != len(otherMap.Pairs) {
			return &Bool{Value: false}, nil
		}
		for hashKey, pair := range m.Pairs {
			otherPair, ok := otherMap.Pairs[hashKey]
			if !ok {
				return &Bool{Value: false}, nil
			}
			// 检查键是否相等
			keyEqual, err := pair.Key.Equal(otherPair.Key, posStart, posEnd, frame)
			if err != nil {
				return &Bool{Value: false}, nil
			}
			// 检查值是否相等
			valueEqual, err := pair.Value.Equal(otherPair.Value, posStart, posEnd, frame)
			if err != nil {
				return &Bool{Value: false}, nil
			}
			if !(keyEqual.(*Bool).Value && valueEqual.(*Bool).Value) {
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
func (m *Map) NotEqual(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	equal, err := m.Equal(other, posStart, posEnd, frame)
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
func (m *Map) LessThan(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (m *Map) GreaterThan(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (m *Map) LessThanOrEqual(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (m *Map) GreaterThanOrEqual(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (m *Map) BitAnd(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (m *Map) BitOr(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (m *Map) Xor(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (m *Map) LeftShift(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (m *Map) RightShift(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (m *Map) And(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (m *Map) Or(_ Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
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
func (m *Map) Index(other Object, posStart, posEnd *util.Pos, frame *frame.Frame) (Object, error) {
	hashable, ok := other.(Hashable)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    frame,
			Message:  "invalid type for map index.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	hashKey := HashKey{
		Type:  hashable.Type(),
		Value: hashable.Hash(),
	}
	pair, ok := m.Pairs[hashKey]
	if !ok {
		return nil, &errors.IndexError{
			Frame:    frame,
			Message:  "key not found.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	return pair.Value, nil
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
func (m *Map) Set(index Object, value Object, posStart, posEnd *util.Pos, frame *frame.Frame) error {
	hashable, ok := index.(Hashable)
	if !ok {
		return &errors.TypeError{
			Frame:    frame,
			Message:  "invalid type for map index.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	hashKey := HashKey{
		Type:  hashable.Type(),
		Value: hashable.Hash(),
	}
	m.Pairs[hashKey] = Pair{Key: index, Value: value}
	return nil
}

// Length 返回列表的长度
//
// 返回值:
//
//	int64 - 列表的长度
func (m *Map) Length() int64 {
	return int64(len(m.Pairs))
}
