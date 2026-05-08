package object

import (
	"strconv"

	"github.com/Ghost-Xiao/ghost-lang/internal/errors"
	"github.com/Ghost-Xiao/ghost-lang/internal/frame"
	"github.com/Ghost-Xiao/ghost-lang/internal/util"
)

// UserError 用户定义错误
// 用于在运行时抛出用户定义的错误
type UserError struct {
	Err      Object       // 错误信息
	Frame    *frame.Frame // 错误发生时的调用栈帧
	PosStart *util.Pos    // 错误起始位置
	PosEnd   *util.Pos    // 错误结束位置
}

// Error 格式化错误信息
// 返回格式化后的错误信息和可能的错误
//
// 返回值:
//
//	string - 格式化后的错误信息
func (ue *UserError) Error() string {
	res := ""
	posStart := ue.PosStart
	posEnd := ue.PosEnd
	currFrame := ue.Frame
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
		str += util.StringsWithArrows(ue.PosStart.Text, posStart, posEnd, true)
		res = str + "\n" + res
		posStart = currFrame.PosStart
		posEnd = currFrame.PosEnd
		currFrame = currFrame.Parent
	}
	res = "Traceback:\n" + res
	res += ue.Err.(*Instance).Class.Name
	message, ok := ue.Err.(*Instance).Member.Get("message")
	if !ok {
		err := &errors.VariableError{
			Frame:    ue.Frame,
			Message:  "undefined member \"message\".",
			PosStart: ue.PosStart,
			PosEnd:   ue.PosEnd,
		}
		return err.Error()
	}
	msg, ok := message.Value.(*String)
	if !ok {
		err := &errors.TypeError{
			Frame:    ue.Frame,
			Message:  "message must be a string.",
			PosStart: ue.PosStart,
			PosEnd:   ue.PosEnd,
		}
		return err.Error()
	}
	if msg.Value != "" {
		res += ": " + msg.Value
	}
	return res
}
