// 实现GoGhost语言的词法分析器，负责将源代码转换为标记流(token stream)

package lexer

import (
	"fmt"
	"strings"

	"github.com/Ghost-Xiao/ghost-lang/internal/util"
)

// Lexer 词法分析器结构体，维护词法分析过程中的状态信息
// 负责读取源代码字符并生成对应的标记(token)
type Lexer struct {
	File    string    // 当前处理的文件名
	Input   string    // 待分析的源代码字符串
	CurrPos *util.Pos // 当前字符的位置信息
	NextPos *util.Pos // 下一个字符的位置信息
}

// NewLexer 创建一个新的词法分析器实例
//
// 参数:
//
//	file - 源代码文件名，用于错误报告
//	input - 要分析的源代码字符串
//
// 返回值:
//
//	初始化后的Lexer指针
func NewLexer(file string, input string) *Lexer {
	l := &Lexer{
		File:    file,
		Input:   input,
		CurrPos: util.NewPos(1, 0, -1, file, input),
		NextPos: util.NewPos(1, 1, 0, file, input),
	}
	l.NextChar() // 初始化时移动到第一个字符
	return l
}

// NextChar 移动到下一个字符位置
// 更新CurrPos和NextPos指针，实现字符流的顺序读取
func (l *Lexer) NextChar() {
	l.CurrPos = l.NextPos.Copy()
	l.NextPos.Advance()
}

// Backup 回退一个字符位置
// 在读取到不需要的字符时使用，将位置指针向后移动一位
func (l *Lexer) Backup() {
	l.NextPos = l.CurrPos.Copy()
	l.CurrPos.Backup()
}

// NextToken 获取下一个标记
// 这是词法分析器的核心方法，根据当前字符类型生成相应的token
//
// 返回值:
//
//	解析出的Token实例和可能的静态错误
func (l *Lexer) NextToken() (*Token, error) {
	for {
		// 根据当前字符类型进行不同处理
		switch l.CurrPos.Char {
		case 0:
			// 到达文件末尾，返回EOF标记
			return &Token{Type: EOF, Literal: "EOF", PosStart: l.CurrPos.Copy(), PosEnd: l.NextPos.Copy()}, nil
		case ' ', '\t', '\r', '\n':
			// 跳过空白字符（空格、制表符、回车、换行）
			l.eatWhitespace()
		default:
			// 处理数字字面量（整数或浮点数）
			if isNumber(l.CurrPos.Char) {
				posStart := l.CurrPos.Copy()
				num, err := l.scanNumber()
				// 单独的点号作为DOT标记处理
				if num == "." {
					return &Token{Type: DOT, Literal: ".", PosStart: posStart, PosEnd: l.NextPos.Copy()}, nil
				}
				if err != nil {
					return &Token{Type: ILLEGAL, Literal: "ILLEGAL", PosStart: posStart, PosEnd: l.NextPos.Copy()}, err
				}
				// 根据是否包含小数点判断是整数还是浮点数
				if strings.Contains(num, ".") {
					return &Token{Type: FLOAT, Literal: num, PosStart: posStart, PosEnd: l.NextPos.Copy()}, nil
				}
				return &Token{Type: INT, Literal: num, PosStart: posStart, PosEnd: l.NextPos.Copy()}, nil
				// 处理标识符或关键字（变量名、函数名等）
			} else if isLetter(l.CurrPos.Char) {
				posStart := l.CurrPos.Copy()
				id := l.scanIdentifier()
				return &Token{Type: LookupIdent(id), Literal: id, PosStart: posStart, PosEnd: l.NextPos.Copy()}, nil
				// 处理运算符
			} else if isOperator(l.CurrPos.Char) {
				posStart := l.CurrPos.Copy()
				// 如果是'/'
				if l.CurrPos.Char == '/' {
					// 如果下一个字符是'/'，说明是单行注释
					if l.NextPos.Char == '/' {
						l.skipComment()
						continue
						// 如果下一个字符是'*'，说明是多行注释
					} else if l.NextPos.Char == '*' {
						err := l.skipMultilineComment()
						if err != nil {
							return &Token{Type: ILLEGAL, Literal: "ILLEGAL", PosStart: posStart, PosEnd: l.NextPos.Copy()}, err
						}
						continue
					}
				}
				op := l.scanOperator()
				return &Token{Type: Operators[op], Literal: op, PosStart: posStart, PosEnd: l.NextPos.Copy()}, nil
				// 处理字符串字面量（支持单引号、双引号和反引号）
			} else if l.CurrPos.Char == '"' || l.CurrPos.Char == '\'' || l.CurrPos.Char == '`' {
				posStart := l.CurrPos.Copy()
				str, err := l.scanString()
				if err != nil {
					return &Token{Type: ILLEGAL, Literal: "ILLEGAL", PosStart: posStart, PosEnd: l.NextPos.Copy()}, err
				}
				return &Token{Type: STRING, Literal: str, PosStart: posStart, PosEnd: l.NextPos.Copy()}, nil
				// 非法字符处理
			} else {
				return &Token{Type: ILLEGAL, Literal: "ILLEGAL"}, &IllegalTokenError{
					Message:  fmt.Sprintf("illegal token \"%c\".", l.CurrPos.Char),
					PosStart: l.CurrPos.Copy(),
					PosEnd:   l.NextPos.Copy(),
				}
			}
		}
		l.NextChar()
	}
}

// isNumber 判断字符是否为数字(0-9)
//
// 参数:
//
//	ch - 要检查的字符
//
// 返回值:
//
//	如果是数字则返回true，否则返回false
func isNumber(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

// isLetter 判断字符是否为字母或下划线
// 支持ASCII字母、下划线和扩展ASCII字符(>=0x80)
//
// 参数:
//
//	ch - 要检查的字符
//
// 返回值:
//
//	如果是字母或下划线则返回true，否则返回false
func isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch >= 0x80
}

// isOperator 判断字符是否为运算符
//
// 参数:
//
//	ch - 要检查的字符
//
// 返回值:
//
//	如果是运算符则返回true，否则返回false
func isOperator(ch rune) bool {
	_, ok := Operators[string(ch)]
	return ok
}

// eatWhitespace 跳过连续的空白字符
// 包括空格、制表符、回车和换行
func (l *Lexer) eatWhitespace() {
	for l.CurrPos.Char == ' ' || l.CurrPos.Char == '\t' || l.CurrPos.Char == '\n' || l.CurrPos.Char == '\r' {
		l.NextChar()
	}
	l.Backup()
}

// skipComment 跳过单行注释
// 从当前'/'字符开始，直到行尾
func (l *Lexer) skipComment() {
	l.NextChar()
	l.NextChar()
	for l.CurrPos.Char != '\n' && l.CurrPos.Char != 0 {
		l.NextChar()
	}
	if l.CurrPos.Char == '\n' {
		l.NextChar()
	}
}

// skipMultilineComment 跳过多行注释
// 从当前'/*'字符开始，直到找到闭合的'*/'
//
// 返回值:
//
//	如果注释未正确闭合则返回语法错误
func (l *Lexer) skipMultilineComment() error {
	// 跳过起始的/*
	l.NextChar()
	l.NextChar()
	// 寻找结束的*/
	for (l.CurrPos.Char != '*' || l.NextPos.Char != '/') && l.CurrPos.Char != 0 {
		l.NextChar()
	}
	// 如果没有找到结束的*/，返回错误
	if l.CurrPos.Char != '*' || l.NextPos.Char != '/' {
		posEnd := l.NextPos.Copy()
		posEnd.Advance()
		return &SyntaxError{
			Message:  "\"*/\" is expected.",
			PosStart: l.CurrPos.Copy(),
			PosEnd:   posEnd,
		}
	}
	l.NextChar()
	l.NextChar()
	return nil
}

// scanNumber 扫描数字字面量
// 支持整数(123)和浮点数(123.45, .45, 123.)
//
// 返回值:
//
//	解析出的数字字符串和可能的错误
func (l *Lexer) scanNumber() (string, error) {
	var num string
	var dotCount int // 小数点计数器，用于检查是否有多个小数点
	// 扫描数字字符和小数点
	for isNumber(l.CurrPos.Char) || l.CurrPos.Char == '.' {
		if l.CurrPos.Char == '.' {
			dotCount++
			// 检查是否有多个小数点，浮点数只能有一个小数点
			if dotCount > 1 {
				return "", &IllegalTokenError{
					Message:  "illegal float literal.",
					PosStart: l.CurrPos.Copy(),
					PosEnd:   l.NextPos.Copy(),
				}
			}
		}
		num += string(l.CurrPos.Char)
		l.NextChar()
	}
	l.Backup()
	return num, nil
}

// scanIdentifier 扫描标识符
// 标识符可以包含字母、数字和下划线，必须以字母或下划线开头
//
// 返回值:
//
//	解析出的标识符字符串
func (l *Lexer) scanIdentifier() string {
	var runes []rune
	for {
		runes = append(runes, l.CurrPos.Char)
		l.NextChar()
		// 标识符由字母、数字和下划线组成
		if !isLetter(l.CurrPos.Char) && !isNumber(l.CurrPos.Char) {
			break
		}
	}
	l.Backup()
	return string(runes)
}

// scanOperator 扫描运算符
// 支持多字符运算符(如==, !=, <=等)
//
// 返回值:
//
//	解析出的运算符字符串
func (l *Lexer) scanOperator() string {
	var op string
	for {
		op += string(l.CurrPos.Char)
		l.NextChar()
		// 检查当前运算符是否有效
		if _, ok := Operators[op]; !ok {
			op = op[:len(op)-1] // 回退最后一个字符
			l.Backup()
			break
		}
		// 如果下一个字符不是运算符，停止扫描
		if !isOperator(l.CurrPos.Char) {
			break
		}
	}
	l.Backup()
	return op
}

// scanString 扫描字符串字面量
// 支持单引号、双引号和反引号字符串，以及转义字符
//
// 返回值:
//
//	解析出的字符串内容和可能的错误
func (l *Lexer) scanString() (string, error) {
	posStart := l.CurrPos.Copy()
	var runes []rune
	quote := l.CurrPos.Char // 记录字符串开始的引号类型
	l.NextChar()
	// 扫描直到找到匹配的结束引号
	for l.CurrPos.Char != quote && l.CurrPos.Char != 0 {
		// 处理转义字符(仅在非反引号字符串中支持)
		if l.CurrPos.Char == '\\' && quote != '`' {
			slashPos := l.CurrPos.Copy()
			l.NextChar()
			// 检查转义字符后的字符是否存在
			if l.CurrPos.Char == 0 {
				return "", &IllegalTokenError{
					Message:  "trailing backslash.",
					PosStart: slashPos,
					PosEnd:   l.NextPos.Copy(),
				}
			}
			// 查找有效的转义字符
			escapeChar, ok := Escape[l.CurrPos.Char]
			if !ok {
				return "", &IllegalTokenError{
					Message:  "illegal escape character.",
					PosStart: slashPos,
					PosEnd:   l.NextPos.Copy(),
				}
			}
			runes = append(runes, escapeChar)
		} else {
			runes = append(runes, l.CurrPos.Char)
		}
		l.NextChar()
	}
	// 检查字符串是否正确闭合
	if l.CurrPos.Char != quote {
		return "", &IllegalTokenError{
			Message:  "unterminated string literal.",
			PosStart: posStart,
			PosEnd:   l.NextPos.Copy(),
		}
	}
	return string(runes), nil
}
