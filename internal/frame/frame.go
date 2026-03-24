package frame

import "github.com/Ghost-Xiao/ghost-lang/internal/util"

// Frame 调用栈帧结构体，用于管理函数调用时的上下文信息
// 包含函数名、父级帧、调用位置等
type Frame struct {
	FuncName string    // 函数名
	Parent   *Frame    // 父级
	PosStart *util.Pos // 函数调用开始位置
	PosEnd   *util.Pos // 函数调用结束位置
}
