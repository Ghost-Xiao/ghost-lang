// 提供代码位置信息的表示和操作功能
// 定义了Pos结构体及相关方法，用于精确描述源代码中的位置，支持位置的前进、后退和格式化输出
// 在词法分析、语法分析和错误报告中广泛使用，帮助精确定位代码问题

package util

import (
	"strconv"
	"strings"
	"unicode/utf8"
)

// Pos 表示源代码中的位置信息，包含行列号、字节索引、文件路径及当前字符
// 用于在编译过程中精确定位代码元素，是错误提示和语法分析的基础数据结构

type Pos struct {
	Row  int    // 行号，从1开始计数
	Col  int    // 列号，从1开始计数
	Idx  int    // 字节索引，从0开始计数
	File string // 文件路径
	Text string // 源代码文本
	Char rune   // 当前位置的字符
}

// NewPos 创建一个新的Pos实例
//
// 参数:
//
//	row  - 行号
//	col  - 列号
//	idx  - 字节索引
//	File - 文件路径
//	text - 源代码文本
//
// 返回值:
//
//	*Pos - 新创建的Pos指针
//
// 注意:
//
//	如果idx超出text范围，Char将被设置为0
func NewPos(row, col int, idx int, File string, text string) *Pos {
	p := &Pos{Row: row, Col: col, Idx: idx, File: File, Text: text}
	if idx < len(text) && idx >= 0 {
		p.Char, _ = utf8.DecodeRuneInString(text[idx:])
		if p.Char == utf8.RuneError {
			p.Char = 0
		}
	} else {
		p.Char = 0
	}
	return p
}

// Copy 创建当前Pos实例的深拷贝
//
// 返回值:
//
//	*Pos - 包含相同位置信息的新Pos指针
//
// 用途:
//
//	用于在不修改原位置信息的情况下创建独立的位置副本
func (p *Pos) Copy() *Pos {
	return &Pos{Row: p.Row, Col: p.Col, Idx: p.Idx, File: p.File, Text: p.Text, Char: p.Char}
}

// Advance 将位置向前移动一个字符
// 更新行号、列号、字节索引和当前字符
//
// 特殊处理:
//
//   - 遇到换行符('\n')时行号加1，列号重置为1
//   - 如果当前位置已超出文本范围，仍会增加列号和索引
func (p *Pos) Advance() {
	if p.Idx < len(p.Text) {
		// 获取当前字符的字节长度
		size := utf8.RuneLen(p.Char)
		p.Idx += size
		// 如果是换行符，更新行号并重置列号
		if p.Char == '\n' {
			p.Row++
			p.Col = 1
		} else {
			p.Col++
		}
		// 更新当前字符
		if p.Idx >= len(p.Text) {
			p.Char = 0
		} else {
			p.Char, _ = utf8.DecodeRuneInString(p.Text[p.Idx:])
			if p.Char == utf8.RuneError {
				p.Char = 0
			}
		}
	} else {
		// 超出文本范围时的处理
		p.Col++
		p.Idx++
		p.Char = 0
	}
}

// Backup 将位置向后移动一个字符
// 更新行号、列号、字节索引和当前字符
//
// 特殊处理:
//
//   - 遇到换行符('\n')时行号减1，列号设置为上一行的字符数
//   - 如果当前位置在文本起始处，仍会减少列号和索引
func (p *Pos) Backup() {
	if p.Idx <= 0 {
		p.Idx--
		p.Col--
		p.Char = 0
	} else {
		// 获取前一个字符及其字节长度
		char, size := utf8.DecodeLastRuneInString(p.Text[:p.Idx])
		if char == utf8.RuneError {
			p.Char = 0
		}
		p.Idx -= size
		p.Char = char
		// 如果是换行符，更新行号并计算列号
		if p.Char == '\n' {
			p.Row--
			if p.Row < 0 {
				p.Col = 0
			} else {
				lines := strings.Split(p.Text, "\n")
				p.Col = utf8.RuneCountInString(lines[p.Row-1]) + 1
			}
		} else {
			p.Col--
		}
	}
}

// String 将位置信息转换为字符串表示
// 格式为：文件路径:字符:行号:列号
//
// 返回值:
//
//	string - 格式化的位置字符串
//
// 用途:
//
//	主要用于错误信息展示，帮助开发者快速定位问题代码位置
func (p *Pos) String() string {
	return p.File + ":" + string(p.Char) + ":" + strconv.Itoa(p.Row) + ":" + strconv.Itoa(p.Col)
}
