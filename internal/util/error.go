// 定义GoGhost语言静态分析阶段的错误类型和错误格式化工具
// 提供错误位置标记、错误信息生成等功能，用于解释器前端的语法和语义错误处理

package util

import (
	"strings"
	"unicode"
)

// getDisplayWidth 计算字符串的显示宽度，考虑东亚字符（中文、日文、韩文）占2个字符宽度
//
// 参数:
//
//	s - 输入字符串
//
// 返回值:
//
//	int - 计算后的显示宽度
//
// 注意:
//
//	普通ASCII字符和符号计为1个宽度，东亚字符计为2个宽度
func getDisplayWidth(s string) int {
	width := 0
	count := 0
	for _, r := range s {
		// 判断是否为东亚字符（中文、日文、韩文）
		if unicode.Is(unicode.Han, r) || unicode.Is(unicode.Hangul, r) || unicode.In(r, unicode.Hiragana, unicode.Katakana) {
			width += 2
		} else {
			width += 1
		}
		count++
	}
	return width
}

// StringsWithArrows 生成带有箭头标记的错误位置可视化字符串
// 用于在源代码中标记错误发生的位置范围，帮助开发者定位问题
//
// 参数:
//
//	text - 源代码文本
//	posStart - 错误起始位置
//	posEnd - 错误结束位置
//	special - 是否使用特殊格式（额外缩进）
//
// 返回值:
//
//	string - 包含源代码片段和箭头标记的格式化字符串
func StringsWithArrows(text string, posStart *Pos, posEnd *Pos, special bool) string {
	var res strings.Builder
	// 查找错误起始位置所在行的开始索引
	var lineStart int
	if posStart.Idx > len(text) {
		lineStart = max(strings.LastIndex(text, "\n")+1, 0)
	} else {
		lineStart = max(strings.LastIndex(text[:posStart.Idx], "\n")+1, 0)
	}
	var lineEnd int
	startLine := lineStart + 1
	if startLine > len(text) {
		lineEnd = -1
	} else {
		// 查找错误起始位置所在行的结束索引
		lineEnd = strings.Index(text[startLine:], "\n")
	}
	if lineEnd < 0 {
		lineEnd = len(text)
	} else {
		lineEnd += lineStart + 1
	}
	// 计算需要显示的错误行数
	lineCount := posEnd.Row - posStart.Row + 1
	if lineCount == 1 {
		lineWithSpace := text[lineStart:lineEnd]
		// 去除左侧空格
		line := strings.TrimLeft(lineWithSpace, " ")
		spaceCount := len(lineWithSpace) - len(line)
		idxStart := posStart.Idx - lineStart - spaceCount
		idxEnd := posEnd.Idx - lineStart - spaceCount
		// 写入缩进空格
		if special {
			res.WriteString("        ")
		} else {
			res.WriteString("    ")
		}
		// 写入源代码行
		res.WriteString(line)
		res.WriteString("\n")
		// 写入箭头前的空格（考虑字符显示宽度）
		if idxStart > len(line) {
			res.WriteString(strings.Repeat(" ", getDisplayWidth(line)+idxStart-len(line)+4))
			if special {
				res.WriteString("    ")
			}
			// 根据错误范围长度写入箭头
			res.WriteString(strings.Repeat("^", idxEnd-idxStart))
		} else {
			res.WriteString(strings.Repeat(" ", getDisplayWidth(line[:idxStart])+4))
			if special {
				res.WriteString("    ")
			}
			// 根据错误范围长度写入箭头
			if idxEnd > len(line) {
				res.WriteString(strings.Repeat("^", getDisplayWidth(line[idxStart:])+idxEnd-len(line)))
			} else {
				res.WriteString(strings.Repeat("^", getDisplayWidth(line[idxStart:idxEnd])))
			}
		}
	} else {
		// 计算左侧空格数量最小值
		var minSpaceCount int
		_lineStart := lineStart
		_lineEnd := lineEnd
		_startLine := startLine
		for j := range lineCount {
			// 计算当前行的空格数量
			_lineWithSpace := text[_lineStart:_lineEnd]
			spaceCount := len(_lineWithSpace) - len(strings.TrimLeft(_lineWithSpace, " "))
			// 更新最小空格数量
			if j == 0 {
				minSpaceCount = spaceCount
			} else if spaceCount < minSpaceCount {
				minSpaceCount = spaceCount
			}
			// 计算下一行的索引
			_lineStart = _lineEnd + 1
			_startLine = _lineStart + 1
			if _startLine > len(text) {
				_lineEnd = -1
			} else {
				_lineEnd = strings.Index(text[_startLine:], "\n")
			}
			if _lineEnd < 0 {
				_lineEnd = len(text)
			} else {
				_lineEnd += _lineStart + 1
			}
		}
		for i := range lineCount {
			// 去除左侧空格
			lineWithSpace := text[lineStart:lineEnd]
			line := lineWithSpace[minSpaceCount:]
			// 在每行的前面加上“|”
			if special {
				res.WriteString("      | ")
			} else {
				res.WriteString("  | ")
			}
			// 写入源代码行
			res.WriteString(line)
			// 非最后一行换行
			if i != lineCount-1 {
				res.WriteString("\n")
			}
			// 更新下一行的索引
			lineStart = lineEnd + 1
			startLine = lineStart + 1
			if startLine > len(text) {
				lineEnd = -1
			} else {
				lineEnd = strings.Index(text[startLine:], "\n")
			}
			if lineEnd < 0 {
				lineEnd = len(text)
			} else {
				lineEnd += lineStart + 1
			}
		}
	}
	// 移除多余的制表符和前导换行
	return strings.TrimLeft(res.String(), "\n")
}
