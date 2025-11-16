package frame

import "github.com/Ghost-Xiao/ghost-lang/internal/util"

type Frame struct {
	FuncName string    // 函数名
	Parent   *Frame    // 父级
	PosStart *util.Pos // 函数调用开始位置
	PosEnd   *util.Pos // 函数调用结束位置
}
