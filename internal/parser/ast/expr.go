package ast

import (
	"strconv"
	"strings"

	"github.com/Ghost-Xiao/ghost-lang/internal/lexer"
	"github.com/Ghost-Xiao/ghost-lang/internal/util"
)

// PrefixExpression 是前缀表达式节点
// 表示带前缀运算符的表达式，如!5、-3等
type PrefixExpression struct {
	Operator *lexer.Token // 前缀运算符token
	Value    Expression   // 右侧表达式
	PosStart *util.Pos    // 表达式的起始位置
	PosEnd   *util.Pos    // 表达式的结束位置
}

// String 返回前缀表达式的字符串表示
// 格式为：<operator><value>
//
// 返回值:
//
//	前缀表达式的字符串表示
func (pe *PrefixExpression) String() string {
	var sb strings.Builder
	sb.WriteString(pe.Operator.Literal)
	sb.WriteString(pe.Value.String())
	return sb.String()
}

// Expression 是标记方法，用于类型判断
// 实现Expression接口
func (pe *PrefixExpression) Expression() {}

// IsLvalue 方法，返回是否为左值
func (pe *PrefixExpression) IsLvalue() bool {
	return false
}

// IntExpression 是整数表达式节点
// 表示源代码中的整数常量
type IntExpression struct {
	Value    int64     // 整数值
	PosStart *util.Pos // 表达式的起始位置
	PosEnd   *util.Pos // 表达式的结束位置
}

// String 返回整数表达式的字符串表示
// 将整数值转换为十进制字符串
//
// 返回值:
//
//	整数的字符串表示
func (ie *IntExpression) String() string {
	return strconv.FormatInt(ie.Value, 10)
}

// Expression 是标记方法，用于类型判断
// 实现Expression接口
func (ie *IntExpression) Expression() {}

// IsLvalue 方法，返回是否为左值
func (ie *IntExpression) IsLvalue() bool {
	return false
}

// FloatExpression 是浮点数表达式节点
// 表示源代码中的浮点数常量
type FloatExpression struct {
	Value    float64   // 浮点数值
	PosStart *util.Pos // 表达式的起始位置
	PosEnd   *util.Pos // 表达式的结束位置
}

// String 返回浮点数表达式的字符串表示
// 将浮点数值转换为字符串
//
// 返回值:
//
//	浮点数的字符串表示
func (fe *FloatExpression) String() string {
	return strconv.FormatFloat(fe.Value, 'f', -1, 64)
}

// Expression 是标记方法，用于类型判断
// 实现Expression接口
func (fe *FloatExpression) Expression() {}

// IsLvalue 方法，返回是否为左值
func (fe *FloatExpression) IsLvalue() bool {
	return false
}

// IdentifierExpression 是标识符表达式节点
// 表示变量名、函数名等标识符
type IdentifierExpression struct {
	Name     string    // 标识符名称
	PosStart *util.Pos // 表达式的起始位置
	PosEnd   *util.Pos // 表达式的结束位置
}

// String 返回标识符的字符串表示
// 直接返回标识符名称
//
// 返回值:
//
//	标识符的名称字符串
func (ie *IdentifierExpression) String() string {
	return ie.Name
}

// Expression 是标记方法，用于类型判断
// 实现Expression接口
func (ie *IdentifierExpression) Expression() {}

// IsLvalue 方法，返回是否为左值
func (ie *IdentifierExpression) IsLvalue() bool {
	return true
}

// BoolExpression 是布尔表达式节点
// 表示布尔值(true/false)
type BoolExpression struct {
	Value    bool      // 布尔值
	PosStart *util.Pos // 表达式的起始位置
	PosEnd   *util.Pos // 表达式的结束位置
}

// String 返回布尔表达式的字符串表示
// 返回"true"或"false"字符串
//
// 返回值:
//
//	布尔值的字符串表示
func (be *BoolExpression) String() string {
	if be.Value {
		return "true"
	}
	return "false"
}

// Expression 是标记方法，用于类型判断
// 实现Expression接口
func (be *BoolExpression) Expression() {}

// IsLvalue 方法，返回是否为左值
func (be *BoolExpression) IsLvalue() bool {
	return false
}

// NullExpression 是空值表达式节点
// 表示null值
type NullExpression struct {
	PosStart *util.Pos // 表达式的起始位置
	PosEnd   *util.Pos // 表达式的结束位置
}

// String 返回空值表达式的字符串表示
// 始终返回"null"
//
// 返回值:
//
//	"null"字符串
func (ne *NullExpression) String() string {
	return "null"
}

// Expression 是标记方法，用于类型判断
// 实现Expression接口
func (ne *NullExpression) Expression() {}

// IsLvalue 方法，返回是否为左值
func (ne *NullExpression) IsLvalue() bool {
	return false
}

// StringExpression 是字符串表达式节点
// 表示源代码中的字符串常量
type StringExpression struct {
	Value    string    // 字符串值
	PosStart *util.Pos // 表达式的起始位置
	PosEnd   *util.Pos // 表达式的结束位置
}

// String 返回字符串表达式的字符串表示
// 返回带引号的字符串
//
// 返回值:
//
//	带引号的字符串表示
func (se *StringExpression) String() string {
	return strconv.Quote(se.Value)
}

// Expression 是标记方法，用于类型判断
// 实现Expression接口
func (se *StringExpression) Expression() {}

// IsLvalue 方法，返回是否为左值
func (se *StringExpression) IsLvalue() bool {
	return false
}

// ListExpression 是列表表达式节点
// 表示源代码中的列表
type ListExpression struct {
	Value    []Expression // 列表值
	PosStart *util.Pos    // 表达式的起始位置
	PosEnd   *util.Pos    // 表达式的结束位置
}

// String 返回列表表达式的字符串表示
// 返回带中括号的列表
//
// 返回值:
//
//	带中括号的列表表示
func (le *ListExpression) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	for i, expr := range le.Value {
		sb.WriteString(expr.String())
		if i < len(le.Value)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("]")
	return sb.String()
}

// Expression 是标记方法，用于类型判断
// 实现Expression接口
func (le *ListExpression) Expression() {}

// IsLvalue 方法，返回是否为左值
func (le *ListExpression) IsLvalue() bool {
	for _, expr := range le.Value {
		if !expr.IsLvalue() {
			return false
		}
	}
	return true
}

// GroupedExpression 是分组表达式节点
// 表示带括号的表达式，用于改变运算优先级
type GroupedExpression struct {
	Expr     Expression // 包裹的表达式
	PosStart *util.Pos  // 表达式的起始位置
	PosEnd   *util.Pos  // 表达式的结束位置
}

// String 返回分组表达式的字符串表示
// 格式为：(expr)
//
// 返回值:
//
//	带括号的表达式字符串
func (ge *GroupedExpression) String() string {
	return "(" + ge.Expr.String() + ")"
}

// Expression 是标记方法，用于类型判断
// 实现Expression接口
func (ge *GroupedExpression) Expression() {}

// IsLvalue 方法，返回是否为左值
func (ge *GroupedExpression) IsLvalue() bool {
	return ge.Expr.IsLvalue()
}

// VarInitializationExpression 是变量初始化表达式节点
// 表示变量的声明和初始化
type VarInitializationExpression struct {
	IsConst  bool
	Name     Expression
	Value    Expression
	PosStart *util.Pos
	PosEnd   *util.Pos
}

// String 返回变量初始化表达式的字符串表示
// 格式为：var <name> = <value>或const <name> = <value>
//
// 返回值:
//
//	变量初始化表达式的字符串表示
func (vi *VarInitializationExpression) String() string {
	var sb strings.Builder
	if vi.IsConst {
		sb.WriteString("const ")
	} else {
		sb.WriteString("var ")
	}
	sb.WriteString(vi.Name.String())
	if vi.Value != nil {
		sb.WriteString(" = ")
		sb.WriteString(vi.Value.String())
	}
	return sb.String()
}

// Expression 是标记方法，用于类型判断
// 实现Expression接口
func (vi *VarInitializationExpression) Expression() {}

// IsLvalue 方法，返回是否为左值
func (vi *VarInitializationExpression) IsLvalue() bool {
	return false
}

// Definition 是空方法，用于标识该接口为定义类型
func (vi *VarInitializationExpression) Definition() {}

// VarAssignmentExpression 是变量赋值表达式节点
// 表示对变量进行赋值操作
type VarAssignmentExpression struct {
	Name     Expression
	Value    Expression
	PosStart *util.Pos
	PosEnd   *util.Pos
}

// String 返回变量赋值表达式的字符串表示
// 格式为：<name> = <value>
//
// 返回值:
//
//	变量赋值表达式的字符串表示
func (va *VarAssignmentExpression) String() string {
	var sb strings.Builder
	sb.WriteString(va.Name.String())
	sb.WriteString(" = ")
	sb.WriteString(va.Value.String())
	return sb.String()
}

// Expression 是标记方法，用于类型判断
// 实现Expression接口
func (va *VarAssignmentExpression) Expression() {}

// IsLvalue 方法，返回是否为左值
func (va *VarAssignmentExpression) IsLvalue() bool {
	return false
}

// CompoundAssignmentExpression 是复合赋值表达式节点
// 表示对变量进行复合赋值
type CompoundAssignmentExpression struct {
	Name     Expression   // 变量名称
	Operator *lexer.Token // 基础运算符
	Right    Expression   // 右操作数
	PosStart *util.Pos    // 表达式的起始位置
	PosEnd   *util.Pos    // 表达式的结束位置
}

// String 返回复合赋值表达式的字符串表示
// 格式为：<name> <operator> <right>
//
// 返回值:
//
//	复合赋值表达式的字符串表示
func (ce *CompoundAssignmentExpression) String() string {
	var sb strings.Builder
	sb.WriteString("var ")
	sb.WriteString(ce.Name.String())
	sb.WriteString(" ")
	sb.WriteString(ce.Operator.Literal)
	sb.WriteString(" ")
	sb.WriteString(ce.Right.String())
	return sb.String()
}

// Expression 是标记方法，用于类型判断
// 实现Expression接口
func (ce *CompoundAssignmentExpression) Expression() {}

// IsLvalue 方法，返回是否为左值
func (ce *CompoundAssignmentExpression) IsLvalue() bool {
	return false
}

// InfixExpression 是中缀表达式节点
// 表示带中缀运算符的表达式，如a + b、x * y等
type InfixExpression struct {
	Left     Expression   // 左侧表达式
	Operator *lexer.Token // 中缀运算符
	Right    Expression   // 右侧表达式
	PosStart *util.Pos    // 表达式的起始位置
	PosEnd   *util.Pos    // 表达式的结束位置
}

// String 返回中缀表达式的字符串表示
// 格式为：<left> <operator> <right>
//
// 返回值:
//
//	中缀表达式的字符串表示
func (ie *InfixExpression) String() string {
	var sb strings.Builder
	sb.WriteString(ie.Left.String())
	sb.WriteString(" ")
	sb.WriteString(ie.Operator.Literal)
	sb.WriteString(" ")
	sb.WriteString(ie.Right.String())
	return sb.String()
}

// Expression 是标记方法，用于类型判断
// 实现Expression接口
func (ie *InfixExpression) Expression() {}

// IsLvalue 方法，返回是否为左值
func (ie *InfixExpression) IsLvalue() bool {
	return false
}

// PrefixUnaryIncDecExpression 是前缀自增 / 自减表达式节点
// 表示前缀自增 / 自减表达式，如++a、--b等
type PrefixUnaryIncDecExpression struct {
	Operator *lexer.Token // 运算符
	Right    Expression   // 右侧表达式
	PosStart *util.Pos    // 表达式的起始位置
	PosEnd   *util.Pos    // 表达式的结束位置
}

// String 返回前缀自增 / 自减表达式的字符串表示
// 格式为：<operator><right>
//
// 返回值:
//
//	前缀自增 / 自减表达式的字符串表示
func (pe *PrefixUnaryIncDecExpression) String() string {
	var sb strings.Builder
	sb.WriteString(pe.Operator.Literal)
	sb.WriteString(pe.Right.String())
	return sb.String()
}

// Expression 是标记方法，用于类型判断
// 实现Expression接口
func (pe *PrefixUnaryIncDecExpression) Expression() {}

// IsLvalue 方法，返回是否为左值
func (pe *PrefixUnaryIncDecExpression) IsLvalue() bool {
	return false
}

// PostfixUnaryIncDecExpression 是后缀自增 / 自减表达式节点
// 表示后缀自增 / 自减表达式，如a++、b--等
type PostfixUnaryIncDecExpression struct {
	Operator *lexer.Token // 运算符
	Left     Expression   // 左侧表达式
	PosStart *util.Pos    // 表达式的起始位置
	PosEnd   *util.Pos    // 表达式的结束位置
}

// String 返回后缀自增 / 自减表达式的字符串表示
// 格式为：<left><operator>
//
// 返回值:
//
//	后缀自增 / 自减表达式的字符串表示
func (pe *PostfixUnaryIncDecExpression) String() string {
	var sb strings.Builder
	sb.WriteString(pe.Left.String())
	sb.WriteString(pe.Operator.Literal)
	return sb.String()
}

// Expression 是标记方法，用于类型判断
// 实现Expression接口
func (pe *PostfixUnaryIncDecExpression) Expression() {}

// IsLvalue 方法，返回是否为左值
func (pe *PostfixUnaryIncDecExpression) IsLvalue() bool {
	return false
}

// BlockExpression 是块表达式节点
type BlockExpression struct {
	Statements []Statement // 语句块
	PosStart   *util.Pos   // 表达式的起始位置
	PosEnd     *util.Pos   // 表达式的结束位置
}

// String 返回块表达式的字符串表示
//
//	格式为：{
//			    <statement-1>;
//			    <statement-2>;
//			    ...
//			    <statement-n>;
//		   }
//
// 返回值:
//
//	块表达式的字符串表示
func (be *BlockExpression) String() string {
	var nodes []string
	for _, s := range be.Statements {
		nodes = append(nodes, s.String())
	}
	return "{\n    " + strings.Join(nodes, ";\n    ") + "\n}"
}

// Expression 是标记方法，用于类型判断
// 实现Expression接口
func (be *BlockExpression) Expression() {}

// IsLvalue 方法，返回是否为左值
func (be *BlockExpression) IsLvalue() bool {
	return false
}

// IfExpression 是if表达式节点
type IfExpression struct {
	Condition   Expression // 条件表达式
	Consequence Statement  // 条件为真时执行的分支体
	Alternative Statement  // else 分支体
	PosStart    *util.Pos  // 表达式的起始位置
	PosEnd      *util.Pos  // 表达式的结束位置
}

// String 返回if表达式的字符串表示
// 格式为：if <cond> <body> else <body>
//
// 返回值:
//
//	if表达式的字符串表示
func (ie *IfExpression) String() string {
	var sb strings.Builder
	sb.WriteString("if ")
	sb.WriteString(ie.Condition.String())
	sb.WriteString(" ")
	sb.WriteString(ie.Consequence.String())
	if ie.Alternative != nil {
		sb.WriteString(" else ")
		sb.WriteString(ie.Alternative.String())
	}
	return sb.String()
}

// Expression 是标记方法，用于类型判断
// 实现Expression接口
func (ie *IfExpression) Expression() {}

// IsLvalue 方法，返回是否为左值
func (ie *IfExpression) IsLvalue() bool {
	return false
}

// CallExpression 是函数调用表达式节点
type CallExpression struct {
	Function Expression   // 函数
	Argument []Expression // 参数
	IsUnpack []bool       // 是否解包
	PosStart *util.Pos    // 表达式的起始位置
	PosEnd   *util.Pos    // 表达式的结束位置
}

// String 返回函数调用表达式的字符串表示
// 格式为：<func>(<para>)
//
// 返回值:
//
//	函数调用表达式的字符串表示
func (ce *CallExpression) String() string {
	var sb strings.Builder
	sb.WriteString(ce.Function.String())
	sb.WriteString("(")
	for i, a := range ce.Argument {
		if a != nil {
			sb.WriteString(a.String())
		}
		if ce.IsUnpack[i] {
			sb.WriteString("...")
		}
		if i != len(ce.Argument)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString(")")
	return sb.String()
}

// Expression 是标记方法，用于类型判断
// 实现Expression接口
func (ce *CallExpression) Expression() {}

// IsLvalue 方法，返回是否为左值
func (ce *CallExpression) IsLvalue() bool {
	return false
}

// IndexExpression 是索引表达式节点
type IndexExpression struct {
	Target   Expression // 被索引的目标
	Index    Expression // 索引表达式
	PosStart *util.Pos  // 表达式的起始位置
	PosEnd   *util.Pos  // 表达式的结束位置
}

// String 返回索引表达式的字符串表示
// 格式为：<target>[<index>]
//
// 返回值:
//
//	索引表达式的字符串表示
func (ie *IndexExpression) String() string {
	var sb strings.Builder
	sb.WriteString(ie.Target.String())
	sb.WriteString("[")
	sb.WriteString(ie.Index.String())
	sb.WriteString("]")
	return sb.String()
}

// Expression 是标记方法，用于类型判断
// 实现Expression接口
func (ie *IndexExpression) Expression() {}

// IsLvalue 方法，返回是否为左值
func (ie *IndexExpression) IsLvalue() bool {
	return true
}

// NamespaceAccessExpression 是命名空间访问表达式节点
// 表示对命名空间成员的访问，如a::b、c::d等
type NamespaceAccessExpression struct {
	Target   Expression // 被访问的目标
	Member   Expression // 成员名
	PosStart *util.Pos  // 表达式的起始位置
	PosEnd   *util.Pos  // 表达式的结束位置
}

// String 返回命名空间访问表达式的字符串表示
// 格式为：<target>::<member>
//
// 返回值:
//
//	命名空间访问表达式的字符串表示
func (ne *NamespaceAccessExpression) String() string {
	var sb strings.Builder
	sb.WriteString(ne.Target.String())
	sb.WriteString("::")
	sb.WriteString(ne.Member.String())
	return sb.String()
}

// Expression 是标记方法，用于类型判断
// 实现Expression接口
func (ne *NamespaceAccessExpression) Expression() {}

// IsLvalue 方法，返回是否为左值
func (ne *NamespaceAccessExpression) IsLvalue() bool {
	return true
}

// ContainsExpression 是contains表达式节点
// 表示对contains操作的表达式，如a contains b等
type ContainsExpression struct {
	Target   Expression // 目标表达式
	Query    Expression // 查询表达式
	PosStart *util.Pos  // 表达式的起始位置
	PosEnd   *util.Pos  // 表达式的结束位置
}

// String 返回contains表达式的字符串表示
// 格式为：<target> contains <query>
//
// 返回值:
//
//	contains表达式的字符串表示
func (ie *ContainsExpression) String() string {
	var sb strings.Builder
	sb.WriteString(ie.Target.String())
	sb.WriteString(" contains ")
	sb.WriteString(ie.Query.String())
	return sb.String()
}

// Expression 是标记方法，用于类型判断
// 实现Expression接口
func (ie *ContainsExpression) Expression() {}

// IsLvalue 方法，返回是否为左值
func (ie *ContainsExpression) IsLvalue() bool {
	return false
}

// RangeExpression 是范围表达式节点
// 表示范围表达式，如a..b等
type RangeExpression struct {
	Start    Expression // 起始值
	End      Expression // 结束值
	PosStart *util.Pos  // 表达式的起始位置
	PosEnd   *util.Pos  // 表达式的结束位置
}

// String 返回范围表达式的字符串表示
// 格式为：<start>..<end>
//
// 返回值:
//
//	范围表达式的字符串表示
func (re *RangeExpression) String() string {
	var sb strings.Builder
	sb.WriteString(re.Start.String())
	sb.WriteString("..")
	sb.WriteString(re.End.String())
	return sb.String()
}

// Expression 是标记方法，用于类型判断
// 实现Expression接口
func (re *RangeExpression) Expression() {}

// IsLvalue 方法，返回是否为左值
func (re *RangeExpression) IsLvalue() bool {
	return false
}

// MemberAccessExpression 是成员访问表达式节点
// 表示对成员的访问，如a.b、c.d等
type MemberAccessExpression struct {
	Target   Expression // 被访问的目标
	Member   Expression // 成员名
	PosStart *util.Pos  // 表达式的起始位置
	PosEnd   *util.Pos  // 表达式的结束位置
}

// String 返回成员访问表达式的字符串表示
// 格式为：<target>.<member>
//
// 返回值:
//
//	成员访问表达式的字符串表示
func (re *MemberAccessExpression) String() string {
	var sb strings.Builder
	sb.WriteString(re.Target.String())
	sb.WriteString(".")
	sb.WriteString(re.Member.String())
	return sb.String()
}

// Expression 是标记方法，用于类型判断
// 实现Expression接口
func (re *MemberAccessExpression) Expression() {}

// IsLvalue 方法，返回是否为左值
func (re *MemberAccessExpression) IsLvalue() bool {
	return true
}

// MapExpression 是映射表达式节点
type MapExpression struct {
	Pairs    [][]Expression // 映射对
	PosStart *util.Pos      // 表达式的起始位置
	PosEnd   *util.Pos      // 表达式的结束位置
}

// String 返回映射表达式的字符串表示
// 格式为：{<key1>: <value1>, <key2>: <value2>, ...}
//
// 返回值:
//
//	映射表达式的字符串表示
func (re *MapExpression) String() string {
	var sb strings.Builder
	sb.WriteString("{")
	for i, pair := range re.Pairs {
		sb.WriteString(pair[0].String())
		sb.WriteString(": ")
		sb.WriteString(pair[1].String())
		if i != len(re.Pairs)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("}")
	return sb.String()
}

// Expression 是标记方法，用于类型判断
// 实现Expression接口
func (re *MapExpression) Expression() {}

// IsLvalue 方法，返回是否为左值
func (re *MapExpression) IsLvalue() bool {
	return false
}

// ThisExpression 是this表达式节点
type ThisExpression struct {
	PosStart *util.Pos // 表达式的起始位置
	PosEnd   *util.Pos // 表达式的结束位置
}

// String 返回this表达式的字符串表示
// 格式为：this
//
// 返回值:
//
//	this表达式的字符串表示
func (re *ThisExpression) String() string {
	return "this"
}

// Expression 是标记方法，用于类型判断
// 实现Expression接口
func (re *ThisExpression) Expression() {}

// IsLvalue 方法，返回是否为左值
func (re *ThisExpression) IsLvalue() bool {
	return false
}

// SuperExpression 是super表达式节点
type SuperExpression struct {
	PosStart *util.Pos // 表达式的起始位置
	PosEnd   *util.Pos // 表达式的结束位置
}

// String 返回super表达式的字符串表示
// 格式为：super
//
// 返回值:
//
//	super表达式的字符串表示
func (re *SuperExpression) String() string {
	return "super"
}

// Expression 是标记方法，用于类型判断
// 实现Expression接口
func (re *SuperExpression) Expression() {}

// IsLvalue 方法，返回是否为左值
func (re *SuperExpression) IsLvalue() bool {
	return false
}
