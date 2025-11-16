package evaluator

import (
	"strconv"

	"github.com/Ghost-Xiao/ghost-lang/internal/frame"
	"github.com/Ghost-Xiao/ghost-lang/internal/util"
)

// VariableError 变量错误类型，表示变量相关的运行时错误
// 例如访问未定义变量、重复声明常量、类型不匹配等
// 拥有完整的错误跟踪和格式化能力

type VariableError struct {
	Frame    *frame.Frame // 错误发生时的调用栈
	Message  string       // 错误描述文本
	PosStart *util.Pos    // 错误起始位置
	PosEnd   *util.Pos    // 错误结束位置
}

// Error 生成格式化的变量错误信息字符串
// 前缀为"Variable Error"
//
// 返回值:
//
//	string - 格式化的变量错误信息，格式同基础Error但错误类型为"Variable Error"
func (e *VariableError) Error() string {
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
	res += "Variable Error"
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

// SyntaxError 语法错误类型，表示语法相关的运行时错误
// 例如缺少括号等
// 拥有完整的错误跟踪和格式化能力

type SyntaxError struct {
	Frame    *frame.Frame // 错误发生时的调用栈
	Message  string       // 错误描述文本
	PosStart *util.Pos    // 错误起始位置
	PosEnd   *util.Pos    // 错误结束位置
}

// Error 生成格式化的语法错误信息字符串
// 前缀为"Syntax Error"
//
// 返回值:
//
//	string - 格式化的变量错误信息，格式同基础Error但错误类型为"Syntax Error"
func (e *SyntaxError) Error() string {
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
	res += "Syntax Error"
	if e.Message != "" {
		res += ": " + e.Message
	}
	return res
}

// ArgumentError 参数错误类型，表示参数相关的运行时错误
// 例如参数数量不匹配等
// 拥有完整的错误跟踪和格式化能力

type ArgumentError struct {
	Frame    *frame.Frame // 错误发生时的调用栈
	Message  string       // 错误描述文本
	PosStart *util.Pos    // 错误起始位置
	PosEnd   *util.Pos    // 错误结束位置
}

// Error 生成格式化的语法错误信息字符串
// 前缀为"Argument Error"
//
// 返回值:
//
//	string - 格式化的变量错误信息，格式同基础Error但错误类型为"Argument Error"
func (e *ArgumentError) Error() string {
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
	res += "Argument Error"
	if e.Message != "" {
		res += ": " + e.Message
	}
	return res
}
