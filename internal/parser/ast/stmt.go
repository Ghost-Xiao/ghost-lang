package ast

import (
	"strings"

	"github.com/Ghost-Xiao/ghost-lang/internal/util"
)

// ForStatement 是for语句节点
// 用于执行for语句

type ForStatement struct {
	Initialization Statement  // 初始化语句
	Condition      Expression // 条件表达式
	Update         Statement  // 更新语句
	Body           Statement  // 循环体语句
	PosStart       *util.Pos  // 语句的起始位置
	PosEnd         *util.Pos  // 语句的结束位置
}

// String 返回for语句的字符串表示
// 格式为：for (<initialization>; <condition>; <update>) <body>
//
// 返回值:
//
//	for语句的字符串表示
func (fs *ForStatement) String() string {
	var sb strings.Builder
	sb.WriteString("for ")
	sb.WriteString(fs.Initialization.String())
	sb.WriteString("; ")
	sb.WriteString(fs.Condition.String())
	sb.WriteString("; ")
	sb.WriteString(fs.Update.String())
	sb.WriteString(" ")
	sb.WriteString(fs.Body.String())
	return sb.String()
}

// Statement 是标记方法，用于类型判断
// 实现Statement接口
func (fs *ForStatement) Statement() {}

// ExpressionStatement 是表达式语句节点
// 用于将表达式作为独立语句执行

type ExpressionStatement struct {
	Expr     Expression // 包裹的表达式
	PosStart *util.Pos  // 语句的起始位置
	PosEnd   *util.Pos  // 语句的结束位置
}

// String 返回表达式语句的字符串表示
// 直接返回内部表达式的字符串表示
//
// 返回值:
//
//	表达式的字符串表示
func (es *ExpressionStatement) String() string {
	return es.Expr.String()
}

// Statement 是标记方法，用于类型判断
// 实现Statement接口
func (es *ExpressionStatement) Statement() {}

// Parameter 是函数参数节点

type Parameter struct {
	Name         *IdentifierExpression // 参数名
	DefaultValue Expression            // 默认值
	PosStart     *util.Pos             // 参数名的起始位置
	PosEnd       *util.Pos             // 参数名的结束位置
}

// String 返回参数的字符串表示
// 格式为：a或a=<default-value>
//
// 返回值:
//
//	参数表达式的字符串表示
func (p *Parameter) String() string {
	var sb strings.Builder
	sb.WriteString(p.Name.String())
	if p.DefaultValue != nil {
		sb.WriteString("=")
		sb.WriteString(p.DefaultValue.String())
	}
	return sb.String()
}

// FunctionDeclarationStatement 是函数声明节点

type FunctionDeclarationStatement struct {
	Name      Expression   // 函数名
	Parameter []*Parameter // 参数
	Body      Statement    // 函数体
	PosStart  *util.Pos    // 语句的起始位置
	PosEnd    *util.Pos    // 语句的结束位置
}

// String 返回函数声明语句的字符串表示
// 格式为：func <name>(<para>) <body>
//
// 返回值:
//
//	函数声明语句的字符串表示
func (fs *FunctionDeclarationStatement) String() string {
	var sb strings.Builder
	sb.WriteString("func ")
	sb.WriteString(fs.Name.String())
	sb.WriteString("(")
	for i, p := range fs.Parameter {
		sb.WriteString(p.String())
		if i != len(fs.Parameter)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString(") ")
	sb.WriteString(fs.Body.String())
	return sb.String()
}

// Statement 是标记方法，用于类型判断
// 实现Statement接口
func (fs *FunctionDeclarationStatement) Statement() {}

// IsLvalue 方法，返回是否为左值
func (fs *FunctionDeclarationStatement) IsLvalue() bool {
	return false
}

// ReturnStatement 是返回语句节点
// 用于返回值

type ReturnStatement struct {
	ReturnValue Expression // 返回的表达式
	PosStart    *util.Pos  // 语句的起始位置
	PosEnd      *util.Pos  // 语句的结束位置
}

// String 返回返回语句的字符串表示
// 格式为：return <expr>
//
// 返回值:
//
//	表达式的字符串表示
func (rs *ReturnStatement) String() string {
	return "return " + rs.ReturnValue.String()
}

// Statement 是标记方法，用于类型判断
// 实现Statement接口
func (rs *ReturnStatement) Statement() {}
