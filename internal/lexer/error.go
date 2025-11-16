// 用于标识词法分析阶段的令牌错误

package lexer

import (
	"strconv"

	"github.com/Ghost-Xiao/ghost-lang/internal/util"
)

// IllegalTokenError 非法令牌错误，表示词法分析时遇到无效的令牌
// 实现 error 接口

type IllegalTokenError struct {
	Message  string    // 错误描述文本
	PosStart *util.Pos // 错误起始位置
	PosEnd   *util.Pos // 错误结束位置
}

// Error 生成格式化的非法令牌错误信息
// 包含错误位置、源代码片段和错误类型标识
//
// 返回值:
//
//	string - 格式化的非法令牌错误信息
func (e *IllegalTokenError) Error() string {
	var linePos string
	if e.PosStart.Row == e.PosEnd.Row {
		linePos = "line " + strconv.Itoa(e.PosStart.Row)
	} else {
		linePos = "lines " + strconv.Itoa(e.PosStart.Row) + "-" + strconv.Itoa(e.PosEnd.Row)
	}
	result := "File " + e.PosStart.File + ", " + linePos + "\n"
	result += util.StringsWithArrows(e.PosStart.Text, e.PosStart, e.PosEnd, false)
	result += "\nIllegal Token Error"
	if e.Message != "" {
		result += ": " + e.Message
	}
	return result
}

// SyntaxError 语法错误，表示遇到非法语法
// 实现 error 接口

type SyntaxError struct {
	Message  string    // 错误描述文本
	PosStart *util.Pos // 错误起始位置
	PosEnd   *util.Pos // 错误结束位置
}

// Error 生成格式化的非法令牌错误信息
// 包含错误位置、源代码片段和错误类型标识
//
// 返回值:
//
//	string - 格式化的非法令牌错误信息
func (e *SyntaxError) Error() string {
	var linePos string
	if e.PosStart.Row == e.PosEnd.Row {
		linePos = "line " + strconv.Itoa(e.PosStart.Row)
	} else {
		linePos = "lines " + strconv.Itoa(e.PosStart.Row) + "-" + strconv.Itoa(e.PosEnd.Row)
	}
	result := "File " + e.PosStart.File + ", " + linePos + "\n"
	result += util.StringsWithArrows(e.PosStart.Text, e.PosStart, e.PosEnd, false)
	result += "\nSyntax Error"
	if e.Message != "" {
		result += ": " + e.Message
	}
	return result
}
