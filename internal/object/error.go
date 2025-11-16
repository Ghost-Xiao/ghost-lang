package object

import (
	"strconv"

	"github.com/Ghost-Xiao/ghost-lang/internal/frame"
	"github.com/Ghost-Xiao/ghost-lang/internal/util"
)

// OperationError 操作错误类型，表示执行不支持的操作时发生的错误
// 例如对不兼容类型执行运算、调用未定义的方法等
// 拥有完整的错误跟踪和格式化能力

type OperationError struct {
	Frame    *frame.Frame // 错误发生时的调用栈
	Message  string       // 错误描述文本
	PosStart *util.Pos    // 错误起始位置
	PosEnd   *util.Pos    // 错误结束位置
}

// Error 生成格式化的操作错误信息字符串
// 前缀为"Operation Error"
//
// 返回值:
//
//	string - 格式化的操作错误信息，格式同基础Error但错误类型为"Operation Error"
func (e *OperationError) Error() string {
	res := ""
	posStart := e.PosStart
	posEnd := e.PosEnd
	currFrame := e.Frame
	// 构建调用栈跟踪信息
	for currFrame != nil {
		var linePos string
		if posStart.Row == posEnd.Row {
			linePos = "line " + strconv.Itoa(posStart.Row)
		} else {
			linePos = "lines " + strconv.Itoa(posStart.Row) + "-" + strconv.Itoa(posEnd.Row)
		}
		str := "    File " + posStart.File + ", " + linePos + ", in " + currFrame.FuncName + "\n"
		// 添加代码位置指示箭头
		str += util.StringsWithArrows(e.PosStart.Text, posStart, posEnd, true)
		res = str + "\n" + res
		posStart = currFrame.PosStart
		posEnd = currFrame.PosEnd
		currFrame = currFrame.Parent
	}
	res = "Traceback:\n" + res
	res += "Operation Error"
	if e.Message != "" {
		res += ": " + e.Message
	}
	return res
}

// MathError 数学错误类型，表示数学运算相关的错误
// 例如除以零、数值溢出、无效的数学函数参数等
// 拥有完整的错误跟踪和格式化能力

type MathError struct {
	Frame    *frame.Frame // 错误发生时的调用栈
	Message  string       // 错误描述文本
	PosStart *util.Pos    // 错误起始位置
	PosEnd   *util.Pos    // 错误结束位置
}

// Error 生成格式化的数学错误信息字符串
// 前缀为"Math Error"
//
// 返回值:
//
//	string - 格式化的数学错误信息，格式同基础Error但错误类型为"Math Error"
func (e *MathError) Error() string {
	res := ""
	posStart := e.PosStart
	posEnd := e.PosEnd
	currFrame := e.Frame
	// 构建调用栈跟踪信息
	for currFrame != nil {
		var linePos string
		if posStart.Row == posEnd.Row {
			linePos = "line " + strconv.Itoa(posStart.Row)
		} else {
			linePos = "lines " + strconv.Itoa(posStart.Row) + "-" + strconv.Itoa(posEnd.Row)
		}
		str := "    File " + posStart.File + ", " + linePos + ", in " + currFrame.FuncName + "\n"
		// 添加代码位置指示箭头
		str += util.StringsWithArrows(e.PosStart.Text, posStart, posEnd, true)
		res = str + "\n" + res
		posStart = currFrame.PosStart
		posEnd = currFrame.PosEnd
		currFrame = currFrame.Parent
	}
	res = "Traceback:\n" + res
	res += "Math Error"
	if e.Message != "" {
		res += ": " + e.Message
	}
	return res
}

// TypeError 类型错误类型，表示类型相关的运行时错误
// 例如访问类型不匹配等
// 拥有完整的错误跟踪和格式化能力

type TypeError struct {
	Frame    *frame.Frame // 错误发生时的调用栈
	Message  string       // 错误描述文本
	PosStart *util.Pos    // 错误起始位置
	PosEnd   *util.Pos    // 错误结束位置
}

// Error 生成格式化的类型错误信息字符串
// 前缀为"Type Error"
//
// 返回值:
//
//	string - 格式化的变量错误信息，格式同基础Error但错误类型为"Type Error"
func (e *TypeError) Error() string {
	res := ""
	posStart := e.PosStart
	posEnd := e.PosEnd
	currFrame := e.Frame
	// 构建调用栈跟踪信息
	for currFrame != nil {
		var linePos string
		if posStart.Row == posEnd.Row {
			linePos = "line " + strconv.Itoa(posStart.Row)
		} else {
			linePos = "lines " + strconv.Itoa(posStart.Row) + "-" + strconv.Itoa(posEnd.Row)
		}
		str := "    File " + posStart.File + ", " + linePos + ", in " + currFrame.FuncName + "\n"
		// 添加代码位置指示箭头
		str += util.StringsWithArrows(e.PosStart.Text, posStart, posEnd, true)
		res = str + "\n" + res
		posStart = currFrame.PosStart
		posEnd = currFrame.PosEnd
		currFrame = currFrame.Parent
	}
	res = "Traceback:\n" + res
	res += "Type Error"
	if e.Message != "" {
		res += ": " + e.Message
	}
	return res
}

// IndexError 索引错误类型，表示索引越界等相关的运行时错误
// 拥有完整的错误跟踪和格式化能力

type IndexError struct {
	Frame    *frame.Frame // 错误发生时的调用栈
	Message  string       // 错误描述文本
	PosStart *util.Pos    // 错误起始位置
	PosEnd   *util.Pos    // 错误结束位置
}

// Error 生成格式化的索引错误信息字符串
// 前缀为"Index Error"
func (e *IndexError) Error() string {
	res := ""
	posStart := e.PosStart
	posEnd := e.PosEnd
	currFrame := e.Frame
	// 构建调用栈跟踪信息
	for currFrame != nil {
		var linePos string
		if posStart.Row == posEnd.Row {
			linePos = "line " + strconv.Itoa(posStart.Row)
		} else {
			linePos = "lines " + strconv.Itoa(posStart.Row) + "-" + strconv.Itoa(posEnd.Row)
		}
		str := "    File " + posStart.File + ", " + linePos + ", in " + currFrame.FuncName + "\n"
		// 添加代码位置指示箭头
		str += util.StringsWithArrows(e.PosStart.Text, posStart, posEnd, true)
		res = str + "\n" + res
		posStart = currFrame.PosStart
		posEnd = currFrame.PosEnd
		currFrame = currFrame.Parent
	}
	res = "Traceback:\n" + res
	res += "Index Error"
	if e.Message != "" {
		res += ": " + e.Message
	}
	return res
}
