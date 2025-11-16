package ast

import (
	"strings"

	"github.com/Ghost-Xiao/ghost-lang/internal/util"
)

// Node 节点接口，是所有AST节点的基接口
// 定义了所有节点共有的基本操作

type Node interface {
	// String 返回节点的字符串表示，用于调试和AST可视化
	// 返回值:
	//   string - 格式化的节点描述
	String() string
}

// Expression 表达式接口，表示可计算值的AST节点
// 继承Node接口，添加Expression方法作为类型标记

type Expression interface {
	Node
	// Expression 空方法，用于标识该接口为表达式类型
	Expression()
	// IsLvalue 方法，返回是否为左值
	IsLvalue() bool
}

// Statement 语句接口，表示执行操作的AST节点
// 继承Node接口，添加Statement方法作为类型标记

type Statement interface {
	Node
	// Statement 空方法，用于标识该接口为语句类型
	Statement()
}

// Program 是AST的根节点，表示整个程序
// 包含一系列按顺序执行的语句节点及位置信息

type Program struct {
	Statements []Statement // 程序中的语句列表
	PosStart   *util.Pos   // 程序的起始位置
	PosEnd     *util.Pos   // 程序的结束位置
}

// String 返回程序节点的字符串表示
// 将所有语句的字符串表示按顺序拼接，用分号和换行分隔
//
// 返回值:
//
//	程序的完整字符串表示
func (p *Program) String() string {
	var nodes []string
	for _, s := range p.Statements {
		nodes = append(nodes, s.String())
	}
	return strings.Join(nodes, ";\n") + ";"
}
