// 定义了GoGhost语言词法分析器使用的令牌类型和相关操作
// 包含令牌结构、预定义令牌类型常量、关键字映射和操作符映射
// 提供令牌复制、字符串转换和标识符查找等功能，是语法分析的基础

package lexer

import "github.com/Ghost-Xiao/ghost-lang/internal/util"

// Token 表示词法分析器生成的令牌结构
// 包含令牌类型、字面量值和位置信息，用于语法分析和错误报告

type Token struct {
	Type     string    // 令牌类型，如INT、PLUS、IDENT等
	Literal  string    // 令牌的字面量值，如数字内容、标识符名称
	PosStart *util.Pos // 令牌在源代码中的起始位置
	PosEnd   *util.Pos // 令牌在源代码中的结束位置
}

// Copy 创建当前Token的深拷贝
// 返回值:
//
//	*Token - 与原Token内容完全相同的新实例
func (t *Token) Copy() *Token {
	return &Token{Type: t.Type, Literal: t.Literal, PosStart: t.PosStart, PosEnd: t.PosEnd}
}

// String 将Token转换为字符串表示形式
// 格式为"类型: 字面量"，用于调试和错误信息展示
// 返回值:
//
//	string - 格式化的令牌字符串
func (t *Token) String() string {
	return t.Type + ": " + t.Literal
}

// 以下为预定义的令牌类型常量
// 基础类型令牌
const (
	ILLEGAL = "ILLEGAL" // 非法令牌，表示无法识别的字符
	EOF     = "EOF"     // 结束符，表示源代码结束
	INT     = "INT"     // 整数类型令牌
	FLOAT   = "FLOAT"   // 浮点数类型令牌
	STRING  = "STRING"  // 字符串类型令牌
	IDENT   = "IDENT"   // 标识符令牌，如变量名、函数名

	// 关键字令牌
	VAR    = "VAR"    // var关键字，用于变量声明
	CONST  = "CONST"  // const关键字，用于常量声明
	FUNC   = "FUNC"   // func关键字，用于函数定义
	IF     = "IF"     // if关键字，条件语句
	ELSE   = "ELSE"   // else关键字，条件语句的分支
	FOR    = "FOR"    // for关键字，循环语句
	RETURN = "RETURN" // return关键字，函数返回
	TRUE   = "TRUE"   // true关键字，布尔值
	FALSE  = "FALSE"  // false关键字，布尔值
	NULL   = "NULL"   // null关键字，表示空值

	// 运算符令牌
	PLUS        = "PLUS"        // 加号运算符(+)
	MINUS       = "MINUS"       // 减号运算符(-)
	ASTERISK    = "ASTERISK"    // 乘号运算符(*)
	SLASH       = "SLASH"       // 除号运算符(/)
	PERCENT     = "PERCENT"     // 取模运算符(%)
	GT          = "GT"          // 大于运算符(>)
	LT          = "LT"          // 小于运算符(<)
	DOT         = "DOT"         // 点运算符(.)
	COMMA       = "COMMA"       // 逗号(,)
	EQUAL       = "EQUAL"       // 等号(=)
	LBRACKET    = "LBRACKET"    // 左中括号([)
	RBRACKET    = "RBRACKET"    // 右中括号(])
	LPAREN      = "LPAREN"      // 左圆括号(()
	RPAREN      = "RPAREN"      // 右圆括号())
	LBRACE      = "LBRACE"      // 左花括号({)
	RBRACE      = "RBRACE"      // 右花括号(})
	BANG        = "BANG"        // 感叹号(!)
	BITWISE_AND = "BITWISE_AND" // 按位与(&)
	BITWISE_OR  = "BITWISE_OR"  // 按位或(|)
	BITWISE_XOR = "BITWISE_XOR" // 按位异或(^)
	BITWISE_NOT = "BITWISE_NOT" // 按位非(~)
	LEFT_SHIFT  = "LEFT_SHIFT"  // 左移运算符(<<)
	RIGHT_SHIFT = "RIGHT_SHIFT" // 右移运算符(>>)
	EQUALS      = "EQUALS"      // 等于比较运算符(==)
	NOT_EQUALS  = "NOT_EQUALS"  // 不等于比较运算符(!=)
	LTE         = "LTE"         // 小于等于运算符(<=)
	GTE         = "GTE"         // 大于等于运算符(>=)
	LOGICAL_AND = "LOGICAL_AND" // 逻辑与(&&)
	LOGICAL_OR  = "LOGICAL_OR"  // 逻辑或(||)
	INCREMENT   = "INCREMENT"   // 自增运算符(++)
	DECREMENT   = "DECREMENT"   // 自减运算符(--)
	ARROW       = "ARROW"       // 箭头运算符(->)，用于函数返回类型
	SEMICOLON   = "SEMICOLON"   // 分号(;)

	// 复合赋值运算符令牌
	PLUS_EQUAL        = "PLUS_EQUAL"        // 加法赋值运算符(+=)
	MINUS_EQUAL       = "MINUS_EQUAL"       // 减法赋值运算符(-=)
	ASTERISK_EQUAL    = "ASTERISK_EQUAL"    // 乘法赋值运算符(*=)
	SLASH_EQUAL       = "SLASH_EQUAL"       // 除法赋值运算符(/=)
	PERCENT_EQUAL     = "PERCENT_EQUAL"     // 取模赋值运算符(%=)
	BITWISE_AND_EQUAL = "BITWISE_AND_EQUAL" // 按位与赋值运算符(&=)
	BITWISE_OR_EQUAL  = "BITWISE_OR_EQUAL"  // 按位或赋值运算符(|=)
	BITWISE_XOR_EQUAL = "BITWISE_XOR_EQUAL" // 按位异或赋值运算符(^=)
	LEFT_SHIFT_EQUAL  = "LEFT_SHIFT_EQUAL"  // 左移赋值运算符(<<=)
	RIGHT_SHIFT_EQUAL = "RIGHT_SHIFT_EQUAL" // 右移赋值运算符(>>=)
)

// Keywords 关键字映射表，将字符串标识符映射到对应的令牌类型
// 用于词法分析时识别保留关键字
var Keywords = map[string]string{
	"var":    VAR,    // 变量声明关键字
	"const":  CONST,  // 常量声明关键字
	"func":   FUNC,   // 函数定义关键字
	"if":     IF,     // 条件语句关键字
	"else":   ELSE,   // 条件语句分支关键字
	"for":    FOR,    // 循环语句关键字
	"return": RETURN, // 函数返回关键字
	"true":   TRUE,   // 布尔值true
	"false":  FALSE,  // 布尔值false
	"null":   NULL,   // 空值关键字
}

// Operators 操作符映射表，将字符串操作符映射到对应的令牌类型
// 用于词法分析时识别各种运算符
var Operators = map[string]string{
	"+":   PLUS,              // 加法运算符
	"-":   MINUS,             // 减法运算符
	"*":   ASTERISK,          // 乘法运算符
	"/":   SLASH,             // 除法运算符
	"%":   PERCENT,           // 取模运算符
	">":   GT,                // 大于比较运算符
	"<":   LT,                // 小于比较运算符
	".":   DOT,               // 点运算符
	",":   COMMA,             // 逗号分隔符
	"=":   EQUAL,             // 赋值运算符
	"[":   LBRACKET,          // 左中括号
	"]":   RBRACKET,          // 右中括号
	"(":   LPAREN,            // 左圆括号
	")":   RPAREN,            // 右圆括号
	"{":   LBRACE,            // 左花括号
	"}":   RBRACE,            // 右花括号
	"!":   BANG,              // 逻辑非运算符
	"&":   BITWISE_AND,       // 按位与运算符
	"|":   BITWISE_OR,        // 按位或运算符
	"^":   BITWISE_XOR,       // 按位异或运算符
	"~":   BITWISE_NOT,       // 按位非运算符
	"<<":  LEFT_SHIFT,        // 左移运算符
	">>":  RIGHT_SHIFT,       // 右移运算符
	"==":  EQUALS,            // 等于比较运算符
	"!=":  NOT_EQUALS,        // 不等于比较运算符
	"<=":  LTE,               // 小于等于运算符
	">=":  GTE,               // 大于等于运算符
	"&&":  LOGICAL_AND,       // 逻辑与运算符
	"||":  LOGICAL_OR,        // 逻辑或运算符
	"++":  INCREMENT,         // 自增运算符
	"--":  DECREMENT,         // 自减运算符
	"->":  ARROW,             // 箭头运算符
	";":   SEMICOLON,         // 分号结束符
	"+=":  PLUS_EQUAL,        // 加法赋值运算符
	"-=":  MINUS_EQUAL,       // 减法赋值运算符
	"*=":  ASTERISK_EQUAL,    // 乘法赋值运算符
	"/=":  SLASH_EQUAL,       // 除法赋值运算符
	"%=":  PERCENT_EQUAL,     // 取模赋值运算符
	"&=":  BITWISE_AND_EQUAL, // 按位与赋值运算符
	"|=":  BITWISE_OR_EQUAL,  // 按位或赋值运算符
	"^=":  BITWISE_XOR_EQUAL, // 按位异或赋值运算符
	"<<=": LEFT_SHIFT_EQUAL,  // 左移赋值运算符
	">>=": RIGHT_SHIFT_EQUAL, // 右移赋值运算符
}

// LookupIdent 检查标识符是否为关键字，并返回对应的令牌类型
// 参数:
//
//	ident - 要检查的标识符字符串
//
// 返回值:
//
//	string - 如果是关键字则返回对应的令牌类型，否则返回IDENT
func LookupIdent(ident string) string {
	if keyword, ok := Keywords[ident]; ok {
		return keyword
	}
	return IDENT
}

// Escape 转义字符映射表，将转义序列字符映射到实际字符
// 用于字符串解析时处理转义序列，如\n对应换行符
var Escape = map[rune]rune{
	'n':  '\n', // 换行符
	't':  '\t', // 制表符
	'r':  '\r', // 回车符
	'b':  '\b', // 退格符
	'\\': '\\', // 反斜杠
	'\'': '\'', // 单引号
	'"':  '"',  // 双引号
	'`':  '`',  // 反引号
}

// CompoundAssignmentOperators 包含复合赋值运算符到基础运算符的映射关系
var CompoundAssignmentOperators = map[string]string{
	PLUS_EQUAL:        PLUS,        // 加法运算符，对应+=
	MINUS_EQUAL:       MINUS,       // 减法运算符，对应-=
	ASTERISK_EQUAL:    ASTERISK,    // 乘法运算符，对应*=
	SLASH_EQUAL:       SLASH,       // 除法运算符，对应/=
	PERCENT_EQUAL:     PERCENT,     // 取模运算符，对应%=
	BITWISE_AND_EQUAL: BITWISE_AND, // 按位与运算符，对应&=
	BITWISE_OR_EQUAL:  BITWISE_OR,  // 按位或运算符，对应|=
	BITWISE_XOR_EQUAL: BITWISE_XOR, // 按位异或运算符，对应^=
	LEFT_SHIFT_EQUAL:  LEFT_SHIFT,  // 左移赋值运算符，对应<<=
	RIGHT_SHIFT_EQUAL: RIGHT_SHIFT, // 右移赋值运算符，对应>>=
}
