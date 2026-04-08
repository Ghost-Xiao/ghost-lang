package object

import (
	"github.com/Ghost-Xiao/ghost-lang/internal/frame"
	"github.com/Ghost-Xiao/ghost-lang/internal/util"
)

// Indexable 表示可索引接口
type Indexable interface {
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
	Set(index Object, value Object, posStart, posEnd *util.Pos, frame *frame.Frame) error

	// Length 返回可索引对象的长度
	//
	// 返回值:
	//
	//	int64 - 可索引对象的长度
	Length() int64

	// 嵌入 object.Object 接口
	Object
}
