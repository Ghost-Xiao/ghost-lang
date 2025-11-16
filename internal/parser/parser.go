// 实现GoGhost语言的语法分析器，负责将词法分析产生的token流转换为抽象语法树(AST)

package parser

import (
	"fmt"
	"strconv"

	"github.com/Ghost-Xiao/ghost-lang/internal/lexer"
	"github.com/Ghost-Xiao/ghost-lang/internal/parser/ast"
	"github.com/Ghost-Xiao/ghost-lang/internal/util"
)

// 运算符优先级常量定义，数值越大优先级越高
const (
	LOWEST  = iota // 最低优先级
	ASSIGN         // 赋值运算符优先级(=, +=, -=, *=, /= 等)
	LOGIC          // 逻辑运算符优先级(&&, ||)
	BIT            // 位运算符优先级(^, &, |, <<, >>)
	EQUALS         // 相等性运算符优先级(==, !=)
	COMPARE        // 比较运算符优先级(<, <=, >, >=)
	SUM            // 加减运算符优先级(+, -)
	MUL            // 乘除运算符优先级(*, /, %)
	PREFIX         // 前缀运算符优先级(!, -, ~, +)
	POSTFIX        // 后缀运算符优先级(++, --)
	CALL           // 函数调用优先级(fn())
)

// precedences 运算符优先级映射表，将token类型映射到对应的优先级常量
var precedences = map[string]int{
	lexer.EQUAL:             ASSIGN,
	lexer.PLUS_EQUAL:        ASSIGN,
	lexer.MINUS_EQUAL:       ASSIGN,
	lexer.ASTERISK_EQUAL:    ASSIGN,
	lexer.SLASH_EQUAL:       ASSIGN,
	lexer.PERCENT_EQUAL:     ASSIGN,
	lexer.BITWISE_AND_EQUAL: ASSIGN,
	lexer.BITWISE_OR_EQUAL:  ASSIGN,
	lexer.BITWISE_XOR_EQUAL: ASSIGN,
	lexer.LEFT_SHIFT_EQUAL:  ASSIGN,
	lexer.RIGHT_SHIFT_EQUAL: ASSIGN,
	lexer.LOGICAL_AND:       LOGIC,
	lexer.LOGICAL_OR:        LOGIC,
	lexer.BITWISE_XOR:       BIT,
	lexer.BITWISE_AND:       BIT,
	lexer.BITWISE_OR:        BIT,
	lexer.LEFT_SHIFT:        BIT,
	lexer.RIGHT_SHIFT:       BIT,
	lexer.EQUALS:            EQUALS,
	lexer.NOT_EQUALS:        EQUALS,
	lexer.LT:                COMPARE,
	lexer.LTE:               COMPARE,
	lexer.GT:                COMPARE,
	lexer.GTE:               COMPARE,
	lexer.PLUS:              SUM,
	lexer.MINUS:             SUM,
	lexer.ASTERISK:          MUL,
	lexer.SLASH:             MUL,
	lexer.PERCENT:           MUL,
	lexer.INCREMENT:         POSTFIX,
	lexer.DECREMENT:         POSTFIX,
	lexer.LPAREN:            CALL,
	lexer.LBRACKET:          CALL,
}

// Parser 语法解析器结构体，负责将词法分析器产生的token流解析为AST

type Parser struct {
	L              *lexer.Lexer                                              // 词法分析器实例
	CurrToken      *lexer.Token                                              // 当前正在处理的token
	NextToken      *lexer.Token                                              // 下一个待处理的token
	Err            error                                                     // 解析过程中产生的错误
	PrefixParseFns map[string]func(*util.Pos) ast.Expression                 // 前缀表达式解析函数映射表
	InfixParseFns  map[string]func(ast.Expression, *util.Pos) ast.Expression // 中缀表达式解析函数映射表
}

// NewParser 创建一个新的语法解析器实例
//
// 参数:
//
//	l - 词法分析器实例
//
// 返回值:
//
//	新的Parser实例和可能的初始化错误
func NewParser(l *lexer.Lexer) (*Parser, error) {
	p := &Parser{L: l}
	// 初始化当前token
	p.CurrToken, p.Err = p.L.NextToken()
	if p.Err != nil {
		return nil, p.Err
	}
	p.L.NextChar()
	// 初始化下一个token
	p.NextToken, p.Err = p.L.NextToken()
	if p.Err != nil {
		return nil, p.Err
	}
	p.L.NextChar()
	// 初始化前缀解析函数映射
	p.PrefixParseFns = map[string]func(*util.Pos) ast.Expression{
		lexer.INT:         p.parseIntegerExpression,
		lexer.FLOAT:       p.parseFloatExpression,
		lexer.IDENT:       p.parseIdentifierExpression,
		lexer.TRUE:        p.parseBoolExpression,
		lexer.FALSE:       p.parseBoolExpression,
		lexer.NULL:        p.parseNullExpression,
		lexer.STRING:      p.parseStringExpression,
		lexer.PLUS:        p.parsePrefixExpression,
		lexer.MINUS:       p.parsePrefixExpression,
		lexer.BANG:        p.parsePrefixExpression,
		lexer.BITWISE_NOT: p.parsePrefixExpression,
		lexer.LPAREN:      p.parseGroupedExpression,
		lexer.VAR:         p.parseVarInitializationExpression,
		lexer.CONST:       p.parseVarInitializationExpression,
		lexer.INCREMENT:   p.parsePrefixUnaryIncDecExpression,
		lexer.DECREMENT:   p.parsePrefixUnaryIncDecExpression,
		lexer.LBRACE:      p.parseBlockExpression,
		lexer.IF:          p.parseIfExpression,
		lexer.LBRACKET:    p.parseListExpression,
	}
	// 初始化中缀解析函数映射
	p.InfixParseFns = map[string]func(ast.Expression, *util.Pos) ast.Expression{
		lexer.LOGICAL_AND:       p.parseInfixExpression,
		lexer.LOGICAL_OR:        p.parseInfixExpression,
		lexer.BITWISE_XOR:       p.parseInfixExpression,
		lexer.BITWISE_AND:       p.parseInfixExpression,
		lexer.BITWISE_OR:        p.parseInfixExpression,
		lexer.LEFT_SHIFT:        p.parseInfixExpression,
		lexer.RIGHT_SHIFT:       p.parseInfixExpression,
		lexer.EQUALS:            p.parseInfixExpression,
		lexer.NOT_EQUALS:        p.parseInfixExpression,
		lexer.LT:                p.parseInfixExpression,
		lexer.LTE:               p.parseInfixExpression,
		lexer.GT:                p.parseInfixExpression,
		lexer.GTE:               p.parseInfixExpression,
		lexer.PLUS:              p.parseInfixExpression,
		lexer.MINUS:             p.parseInfixExpression,
		lexer.ASTERISK:          p.parseInfixExpression,
		lexer.SLASH:             p.parseInfixExpression,
		lexer.PERCENT:           p.parseInfixExpression,
		lexer.EQUAL:             p.parseVarAssignmentExpression,
		lexer.PLUS_EQUAL:        p.parseCompoundAssignmentExpression,
		lexer.MINUS_EQUAL:       p.parseCompoundAssignmentExpression,
		lexer.ASTERISK_EQUAL:    p.parseCompoundAssignmentExpression,
		lexer.SLASH_EQUAL:       p.parseCompoundAssignmentExpression,
		lexer.PERCENT_EQUAL:     p.parseCompoundAssignmentExpression,
		lexer.BITWISE_AND_EQUAL: p.parseCompoundAssignmentExpression,
		lexer.BITWISE_OR_EQUAL:  p.parseCompoundAssignmentExpression,
		lexer.BITWISE_XOR_EQUAL: p.parseCompoundAssignmentExpression,
		lexer.LEFT_SHIFT_EQUAL:  p.parseCompoundAssignmentExpression,
		lexer.RIGHT_SHIFT_EQUAL: p.parseCompoundAssignmentExpression,
		lexer.INCREMENT:         p.parsePostfixUnaryIncDecExpression,
		lexer.DECREMENT:         p.parsePostfixUnaryIncDecExpression,
		lexer.LPAREN:            p.parseCallExpression,
		lexer.LBRACKET:          p.parseIndexExpression,
	}
	return p, nil
}

// Advance 前进到下一个token，更新CurrToken和NextToken
func (p *Parser) Advance() {
	p.CurrToken = p.NextToken.Copy()
	p.NextToken, p.Err = p.L.NextToken()
	p.L.NextChar()
}

// CheckNextAndAdvance 检查下一个token是否为预期类型，如果是则前进，否则设置错误
//
// 参数:
//
//	excepted - 预期的token类型
func (p *Parser) CheckNextAndAdvance(excepted string) {
	if p.NextToken.Type != excepted {
		// 创建语法错误，包含预期和实际token类型信息
		p.Err = &SyntaxError{
			Message:  fmt.Sprintf("expected \"%s\", but got \"%s\".", excepted, p.NextToken.Type),
			PosStart: p.NextToken.PosStart.Copy(),
			PosEnd:   p.NextToken.PosEnd.Copy(),
		}
	} else {
		p.Advance()
	}
}

// ParseProgram 解析整个程序，生成AST的根节点Program
//
// 返回值:
//
//	包含所有语句的Program节点
func (p *Parser) ParseProgram() *ast.Program {
	posStart := p.CurrToken.PosStart.Copy()
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	// 循环解析所有语句直到文件结束
	for p.CurrToken.Type != lexer.EOF {
		if p.Err != nil {
			return nil
		}
		// 跳过空分号
		for p.CurrToken.Type == lexer.SEMICOLON {
			p.Advance()
		}
		if p.CurrToken.Type == lexer.EOF {
			break
		}
		// 解析单个语句
		statPosStart := p.CurrToken.PosStart.Copy()
		stat := p.parseStatement(statPosStart)
		if p.Err != nil {
			return nil
		}
		// 检查语句后的分号
		p.CheckNextAndAdvance(lexer.SEMICOLON)
		if p.Err != nil {
			return nil
		}
		// 添加语句到程序节点
		program.Statements = append(program.Statements, stat)
		p.Advance()
	}
	program.PosStart = posStart
	program.PosEnd = p.CurrToken.PosEnd.Copy()
	return program
}

// parseStatement 解析单个语句
//
// 参数:
//
//	posStart - 语句的起始位置
//
// 返回值:
//
//	解析得到的语句节点
func (p *Parser) parseStatement(posStart *util.Pos) ast.Statement {
	switch p.CurrToken.Type {
	case lexer.FOR:
		// 解析为for语句
		return p.parseForStatement(posStart)
	case lexer.FUNC:
		// 解析为函数声明语句
		return p.parseFunctionDeclarationStatement(posStart)
	case lexer.RETURN:
		// 解析为return语句
		return p.parseReturnStatement(posStart)
	default:
		// 解析为表达式语句
		return p.parseExpressionStatement(posStart)
	}
}

// parseForStatement 解析for语句
//
// 参数:
//
//	posStart - 语句的起始位置
//
// 返回值:
//
//	for语句节点ForStatement
func (p *Parser) parseForStatement(posStart *util.Pos) *ast.ForStatement {
	fs := &ast.ForStatement{
		PosStart: posStart,
	}
	p.Advance()
	// 解析初始化语句
	fs.Initialization = p.parseStatement(p.CurrToken.PosStart.Copy())
	if p.Err != nil {
		return nil
	}
	p.CheckNextAndAdvance(lexer.SEMICOLON)
	if p.Err != nil {
		return nil
	}
	p.Advance()
	// 解析条件表达式
	fs.Condition = p.ParseExpression(LOWEST)
	if p.Err != nil {
		return nil
	}
	p.CheckNextAndAdvance(lexer.SEMICOLON)
	if p.Err != nil {
		return nil
	}
	p.Advance()
	// 解析更新语句
	fs.Update = p.parseStatement(p.CurrToken.PosStart.Copy())
	if p.Err != nil {
		return nil
	}
	p.Advance()
	// 解析循环体语句
	fs.Body = p.parseStatement(p.CurrToken.PosStart.Copy())
	if p.Err != nil {
		return nil
	}
	fs.PosEnd = p.CurrToken.PosEnd.Copy()
	return fs
}

// parseFunctionDeclarationStatement 解析函数表达式
//
// 参数:
//
//	posStart - 表达式的起始位置
//
// 返回值:
//
//	函数表达式节点FunctionDeclarationStatement
func (p *Parser) parseFunctionDeclarationStatement(posStart *util.Pos) *ast.FunctionDeclarationStatement {
	fe := &ast.FunctionDeclarationStatement{
		PosStart:  posStart,
		Parameter: make([]*ast.Parameter, 0),
	}
	// 解析函数名
	p.CheckNextAndAdvance(lexer.IDENT)
	if p.Err != nil {
		return nil
	}
	fe.Name = p.parseIdentifierExpression(p.CurrToken.PosStart.Copy())
	p.CheckNextAndAdvance(lexer.LPAREN)
	if p.Err != nil {
		return nil
	}
	p.Advance()
	haveDefault := false
	// 解析函数参数
	for p.CurrToken.Type != lexer.RPAREN {
		paraPosStart := p.CurrToken.PosStart.Copy()
		// 解析参数
		expr := p.parseIdentifierExpression(paraPosStart)
		if p.Err != nil {
			return nil
		}
		para := expr.(*ast.IdentifierExpression)
		var defaultValue ast.Expression = nil
		if haveDefault && p.NextToken.Type != lexer.EQUAL {
			p.Err = &SyntaxError{
				Message:  "non-default parameter follows default parameter.",
				PosStart: paraPosStart,
				PosEnd:   p.CurrToken.PosEnd.Copy(),
			}
			return nil
		}
		// 解析默认值
		if p.NextToken.Type == lexer.EQUAL {
			p.Advance()
			p.Advance()
			defaultExpr := p.ParseExpression(LOWEST)
			if p.Err != nil {
				return nil
			}
			defaultValue = defaultExpr
			haveDefault = true
		}
		// 创建参数节点
		parameter := &ast.Parameter{
			Name:         para,
			DefaultValue: defaultValue,
			PosStart:     paraPosStart,
			PosEnd:       p.CurrToken.PosEnd.Copy(),
		}
		fe.Parameter = append(fe.Parameter, parameter)
		if p.Err != nil {
			return nil
		}
		// 检查参数后的逗号
		if p.NextToken.Type != lexer.RPAREN {
			p.CheckNextAndAdvance(lexer.COMMA)
			if p.Err != nil {
				return nil
			}
		}
		p.Advance()
	}
	p.Advance()
	// 解析函数体
	fe.Body = p.parseStatement(p.CurrToken.PosStart.Copy())
	if p.Err != nil {
		return nil
	}
	fe.PosEnd = p.CurrToken.PosEnd.Copy()
	return fe
}

// parseReturnStatement 解析return语句
//
// 参数:
//
//	posStart - 语句的起始位置
//
// 返回值:
//
//	return语句节点ReturnStatement
func (p *Parser) parseReturnStatement(posStart *util.Pos) *ast.ReturnStatement {
	rs := &ast.ReturnStatement{
		PosStart: posStart,
	}
	p.Advance()
	// 解析返回值表达式
	rs.ReturnValue = p.ParseExpression(LOWEST)
	if p.Err != nil {
		return nil
	}
	rs.PosEnd = p.CurrToken.PosEnd.Copy()
	return rs
}

// parseExpressionStatement 解析表达式语句(由单个表达式组成的语句)
//
// 参数:
//
//	posStart - 语句的起始位置
//
// 返回值:
//
//	包含表达式的ExpressionStatement节点
func (p *Parser) parseExpressionStatement(posStart *util.Pos) *ast.ExpressionStatement {
	expr := p.ParseExpression(LOWEST)
	if p.Err != nil {
		return nil
	}
	return &ast.ExpressionStatement{Expr: expr, PosStart: posStart, PosEnd: p.CurrToken.PosEnd.Copy()}
}

// ParseExpression 解析表达式，根据运算符优先级递归构建表达式节点
//
// 参数:
//
//	precedence - 当前表达式的最低优先级要求
//
// 返回值:
//
//	解析得到的表达式节点
func (p *Parser) ParseExpression(precedence int) ast.Expression {
	posStart := p.CurrToken.PosStart.Copy()
	// 根据当前token类型获取对应的前缀解析函数
	prefixFn := p.PrefixParseFns[p.CurrToken.Type]
	if prefixFn == nil {
		// 如果没有对应的前缀解析函数，返回语法错误
		p.Err = &SyntaxError{
			Message:  fmt.Sprintf("unexpected \"%s\".", p.CurrToken.Type),
			PosStart: posStart,
			PosEnd:   p.CurrToken.PosEnd.Copy(),
		}
		return nil
	}
	expr := prefixFn(posStart)
	if p.Err != nil {
		return nil
	}
	// 根据运算符优先级处理中缀表达式
	for precedence < precedences[p.NextToken.Type] {
		infixFn := p.InfixParseFns[p.NextToken.Type]
		if infixFn == nil {
			return expr
		}
		p.Advance()
		expr = infixFn(expr, posStart)
		if p.Err != nil {
			return nil
		}
	}
	return expr
}

// parsePrefixExpression 解析前缀表达式(如!5, -3, +2等)
//
// 参数:
//
//	posStart - 表达式的起始位置
//
// 返回值:
//
//	前缀表达式节点PrefixExpression
func (p *Parser) parsePrefixExpression(posStart *util.Pos) ast.Expression {
	pe := &ast.PrefixExpression{
		Operator: p.CurrToken.Copy(),
		PosStart: posStart,
	}
	p.Advance()
	// 递归解析前缀运算符后的表达式
	expr := p.ParseExpression(PREFIX)
	if p.Err != nil {
		return nil
	}
	pe.Value = expr
	pe.PosEnd = p.CurrToken.PosEnd.Copy()
	return pe
}

// parseIntegerExpression 解析整数表达式
//
// 参数:
//
//	posStart - 表达式的起始位置
//
// 返回值:
//
//	整数表达式节点IntExpression
func (p *Parser) parseIntegerExpression(posStart *util.Pos) ast.Expression {
	// 将token字面量转换为int64类型
	num, ok := strconv.ParseInt(p.CurrToken.Literal, 10, 64)
	if ok != nil {
		// 转换失败时返回非法token错误
		p.Err = &lexer.IllegalTokenError{
			Message:  "illegal integer.",
			PosStart: posStart,
			PosEnd:   p.CurrToken.PosEnd.Copy(),
		}
		return nil
	}
	return &ast.IntExpression{Value: num, PosStart: posStart, PosEnd: p.CurrToken.PosEnd.Copy()}
}

// parseFloatExpression 解析浮点数表达式
//
// 参数:
//
//	posStart - 表达式的起始位置
//
// 返回值:
//
//	浮点数表达式节点FloatExpression
func (p *Parser) parseFloatExpression(posStart *util.Pos) ast.Expression {
	// 将token字面量转换为float64类型
	num, ok := strconv.ParseFloat(p.CurrToken.Literal, 64)
	if ok != nil {
		// 转换失败时返回非法token错误
		p.Err = &lexer.IllegalTokenError{
			Message:  "illegal float.",
			PosStart: posStart,
			PosEnd:   p.CurrToken.PosEnd.Copy(),
		}
		return nil
	}
	return &ast.FloatExpression{Value: num, PosStart: posStart, PosEnd: p.CurrToken.PosEnd.Copy()}
}

// parseIdentifierExpression 解析标识符表达式(变量名、函数名等)
//
// 参数:
//
//	posStart - 表达式的起始位置
//
// 返回值:
//
//	标识符表达式节点IdentifierExpression
func (p *Parser) parseIdentifierExpression(posStart *util.Pos) ast.Expression {
	return &ast.IdentifierExpression{Name: p.CurrToken.Literal, PosStart: posStart, PosEnd: p.CurrToken.PosEnd.Copy()}
}

// parseBoolExpression 解析布尔表达式(true或false)
//
// 参数:
//
//	posStart - 表达式的起始位置
//
// 返回值:
//
//	布尔表达式节点BoolExpression
func (p *Parser) parseBoolExpression(posStart *util.Pos) ast.Expression {
	return &ast.BoolExpression{Value: p.CurrToken.Literal == "true", PosStart: posStart, PosEnd: p.CurrToken.PosEnd.Copy()}
}

// parseNullExpression 解析空值表达式(null)
//
// 参数:
//
//	posStart - 表达式的起始位置
//
// 返回值:
//
//	空值表达式节点NullExpression
func (p *Parser) parseNullExpression(posStart *util.Pos) ast.Expression {
	return &ast.NullExpression{PosStart: posStart, PosEnd: p.CurrToken.PosEnd.Copy()}
}

// parseStringExpression 解析字符串表达式
//
// 参数:
//
//	posStart - 表达式的起始位置
//
// 返回值:
//
//	字符串表达式节点StringExpression
func (p *Parser) parseStringExpression(posStart *util.Pos) ast.Expression {
	return &ast.StringExpression{Value: p.CurrToken.Literal, PosStart: posStart, PosEnd: p.CurrToken.PosEnd.Copy()}
}

// parseGroupedExpression 解析分组表达式(括号内的表达式)
//
// 参数:
//
//	posStart - 表达式的起始位置
//
// 返回值:
//
//	分组表达式节点GroupedExpression
func (p *Parser) parseGroupedExpression(posStart *util.Pos) ast.Expression {
	p.Advance()
	// 解析括号内的表达式
	expr := p.ParseExpression(LOWEST)
	if p.Err != nil {
		return nil
	}
	// 确保括号匹配
	p.CheckNextAndAdvance(lexer.RPAREN)
	return &ast.GroupedExpression{Expr: expr, PosStart: posStart, PosEnd: p.CurrToken.PosEnd.Copy()}
}

// parseVarInitializationExpression 解析变量初始化表达式(var或const)
//
// 参数:
//
//	posStart - 表达式的起始位置
//
// 返回值:
//
//	变量初始化表达式节点VarInitialization
func (p *Parser) parseVarInitializationExpression(posStart *util.Pos) ast.Expression {
	// 区分const和var声明
	isConst := p.CurrToken.Type == lexer.CONST
	// 检查并消耗标识符
	p.CheckNextAndAdvance(lexer.IDENT)
	if p.Err != nil {
		return nil
	}
	// 解析变量名
	name := p.parseIdentifierExpression(p.CurrToken.PosStart.Copy())
	// 检查并消耗赋值运算符
	p.CheckNextAndAdvance(lexer.EQUAL)
	if p.Err != nil {
		return nil
	}
	p.Advance()
	// 解析变量值表达式
	value := p.ParseExpression(LOWEST)
	if p.Err != nil {
		return nil
	}
	return &ast.VarInitializationExpression{
		IsConst:  isConst,
		Name:     name,
		Value:    value,
		PosStart: posStart,
		PosEnd:   p.CurrToken.PosEnd.Copy(),
	}
}

// parseVarAssignmentExpression 解析变量赋值表达式
//
// 参数:
//
//	left - 左侧表达式节点
//	posStart - 表达式的起始位置
//
// 返回值:
//
//	变量赋值表达式节点VarAssignment
func (p *Parser) parseVarAssignmentExpression(left ast.Expression, posStart *util.Pos) ast.Expression {
	if !left.IsLvalue() {
		p.Err = &SyntaxError{
			Message:  "operation \"=\" requires an lvalue operand.",
			PosStart: posStart,
			PosEnd:   p.CurrToken.PosEnd.Copy(),
		}
		return nil
	}
	p.Advance()
	// 解析变量值表达式
	value := p.ParseExpression(LOWEST)
	if p.Err != nil {
		return nil
	}
	return &ast.VarAssignmentExpression{
		Name:     left,
		Value:    value,
		PosStart: posStart,
		PosEnd:   p.CurrToken.PosEnd.Copy(),
	}
}

func (p *Parser) parseCompoundAssignmentExpression(left ast.Expression, posStart *util.Pos) ast.Expression {
	// 检查左侧表达式是否为左值
	if !left.IsLvalue() {
		p.Err = &SyntaxError{
			Message:  fmt.Sprintf("operation \"%s\" requires an lvalue operand.", p.CurrToken.Literal),
			PosStart: posStart,
			PosEnd:   p.CurrToken.PosEnd.Copy(),
		}
		return nil
	}
	// 记录复合赋值运算符
	operator := p.CurrToken.Copy()
	p.Advance()
	// 解析右侧表达式
	right := p.ParseExpression(LOWEST)
	if p.Err != nil {
		return nil
	}
	return &ast.CompoundAssignmentExpression{
		Name:     left,
		Operator: operator,
		Right:    right,
		PosStart: posStart,
		PosEnd:   p.CurrToken.PosEnd.Copy(),
	}
}

// parsePrefixUnaryIncDecExpression 解析前缀自增 / 自减表达式
//
// 参数:
//
//	posStart - 表达式的起始位置
//
// 返回值:
//
//	前缀自增 / 自减表达式节点PrefixUnaryIncDecExpression
func (p *Parser) parsePrefixUnaryIncDecExpression(posStart *util.Pos) ast.Expression {
	// 记录运算符
	operator := p.CurrToken.Copy()
	p.Advance()
	// 解析右侧表达式
	right := p.ParseExpression(LOWEST)
	if p.Err != nil {
		return nil
	}
	// 检查右侧表达式是否为左值
	if !right.IsLvalue() {
		p.Err = &SyntaxError{
			Message:  "operation \"++\" or \"--\" requires an lvalue operand.",
			PosStart: posStart,
			PosEnd:   p.CurrToken.PosEnd.Copy(),
		}
		return nil
	}
	return &ast.PrefixUnaryIncDecExpression{
		Operator: operator,
		Right:    right,
		PosStart: posStart,
		PosEnd:   p.CurrToken.PosEnd.Copy(),
	}
}

func (p *Parser) parsePostfixUnaryIncDecExpression(left ast.Expression, posStart *util.Pos) ast.Expression {
	// 记录运算符
	operator := p.CurrToken.Copy()
	// 检查左侧表达式是否为左值
	if !left.IsLvalue() {
		p.Err = &SyntaxError{
			Message:  "operation \"++\" or \"--\" requires an lvalue operand.",
			PosStart: posStart,
			PosEnd:   p.CurrToken.PosEnd.Copy(),
		}
		return nil
	}
	return &ast.PostfixUnaryIncDecExpression{
		Operator: operator,
		Left:     left,
		PosStart: posStart,
		PosEnd:   p.CurrToken.PosEnd.Copy(),
	}
}

// parseInfixExpression 解析中缀表达式(如a + b, x * y等)
//
// 参数:
//
//	left - 左侧表达式节点
//	posStart - 表达式的起始位置
//
// 返回值:
//
//	中缀表达式节点InfixExpression
func (p *Parser) parseInfixExpression(left ast.Expression, posStart *util.Pos) ast.Expression {
	ie := &ast.InfixExpression{
		Left:     left,
		Operator: p.CurrToken.Copy(),
		PosStart: posStart,
	}
	// 获取当前运算符优先级
	precedence := precedences[p.CurrToken.Type]
	p.Advance()
	// 根据优先级解析右侧表达式
	right := p.ParseExpression(precedence)
	if p.Err != nil {
		return nil
	}
	ie.Right = right
	ie.PosEnd = p.CurrToken.PosEnd.Copy()
	return ie
}

// parseBlockExpression 解析块表达式
//
// 参数:
//
//	posStart - 表达式的起始位置
//
// 返回值:
//
//	块表达式节点BlockExpression
func (p *Parser) parseBlockExpression(posStart *util.Pos) ast.Expression {
	expr := &ast.BlockExpression{
		PosStart: posStart,
	}
	p.Advance()
	// 循环解析所有语句直到遇到右大括号
	for p.CurrToken.Type != lexer.RBRACE {
		if p.Err != nil {
			return nil
		}
		// 跳过空分号
		for p.CurrToken.Type == lexer.SEMICOLON {
			p.Advance()
		}
		if p.CurrToken.Type == lexer.RBRACE {
			break
		}
		// 解析单个语句
		statPosStart := p.CurrToken.PosStart.Copy()
		stat := p.parseStatement(statPosStart)
		if p.Err != nil {
			return nil
		}
		// 如果块表达式还未结束
		if p.NextToken.Type != lexer.RBRACE {
			// 检查语句后的分号
			p.CheckNextAndAdvance(lexer.SEMICOLON)
			if p.Err != nil {
				return nil
			}
		}
		// 添加语句到块表达式节点
		expr.Statements = append(expr.Statements, stat)
		p.Advance()
	}
	expr.PosEnd = p.CurrToken.PosEnd.Copy()
	return expr
}

// parseIfExpression 解析if表达式
//
// 参数:
//
//	posStart - 表达式的起始位置
//
// 返回值:
//
//	if表达式节点IfExpression
func (p *Parser) parseIfExpression(posStart *util.Pos) ast.Expression {
	ie := &ast.IfExpression{
		PosStart: posStart,
	}
	p.Advance()
	// 解析if条件表达式
	ie.Condition = p.ParseExpression(LOWEST)
	if p.Err != nil {
		return nil
	}
	p.Advance()
	// 解析条件为真时执行的分支体
	ie.Consequence = p.parseStatement(p.CurrToken.PosStart)
	if p.Err != nil {
		return nil
	}
	// 检查是否存在else分支
	if p.NextToken.Type == lexer.ELSE {
		p.Advance()
		p.Advance()
		// 解析else分支
		ie.Alternative = p.parseStatement(p.CurrToken.PosStart)
		if p.Err != nil {
			return nil
		}
	} else {
		// 如果没有else分支，设置为nil
		ie.Alternative = nil
	}
	ie.PosEnd = p.CurrToken.PosEnd.Copy()
	return ie
}

// parseListExpression 解析列表表达式
//
// 参数:
//
//	posStart - 表达式的起始位置
//
// 返回值:
//
//	列表表达式节点ListExpression
func (p *Parser) parseListExpression(posStart *util.Pos) ast.Expression {
	le := &ast.ListExpression{
		Value:    make([]ast.Expression, 0),
		PosStart: posStart,
	}
	p.Advance()
	// 处理空列表的情况
	if p.CurrToken.Type == lexer.RBRACKET {
		le.PosEnd = p.CurrToken.PosEnd.Copy()
		return le
	}
	// 循环解析列表中的元素直到遇到右方括号
	for p.CurrToken.Type != lexer.RBRACKET {
		if p.Err != nil {
			return nil
		}
		// 解析列表元素表达式
		elem := p.ParseExpression(LOWEST)
		if p.Err != nil {
			return nil
		}
		// 添加元素到列表
		le.Value = append(le.Value, elem)
		// 检查是否还有更多元素(通过逗号分隔)
		if p.NextToken.Type != lexer.RBRACKET {
			// 检查并消耗逗号
			p.CheckNextAndAdvance(lexer.COMMA)
			if p.Err != nil {
				return nil
			}
		}
		p.Advance()
	}
	// 设置列表表达式的结束位置
	le.PosEnd = p.CurrToken.PosEnd.Copy()
	return le
}

// parseCallExpression 解析函数调用表达式
//
// 参数:
//
//	posStart - 表达式的起始位置
//
// 返回值:
//
//	函数表达式节点CallExpression
func (p *Parser) parseCallExpression(left ast.Expression, posStart *util.Pos) ast.Expression {
	ce := &ast.CallExpression{
		Function: left,
		Argument: make([]ast.Expression, 0),
		PosStart: posStart,
	}
	p.Advance()
	for p.CurrToken.Type != lexer.RPAREN {
		// 如果不是逗号
		if p.CurrToken.Type != lexer.COMMA {
			arg := p.ParseExpression(LOWEST)
			if p.Err != nil {
				return nil
			}
			ce.Argument = append(ce.Argument, arg)
			if p.NextToken.Type != lexer.RPAREN {
				p.CheckNextAndAdvance(lexer.COMMA)
				if p.Err != nil {
					return nil
				}
			}
		} else {
			ce.Argument = append(ce.Argument, nil)
		}
		p.Advance()
	}
	ce.PosEnd = p.CurrToken.PosEnd.Copy()
	return ce
}

// parseIndexExpression 解析索引表达式
//
// 参数:
//
//	left - 左侧目标表达式
//	posStart - 表达式的起始位置
//
// 返回值:
//
//	索引表达式节点 IndexExpression
func (p *Parser) parseIndexExpression(left ast.Expression, posStart *util.Pos) ast.Expression {
	// 当前 CurrToken 为 '['
	p.Advance()
	// 解析索引表达式
	indexExpr := p.ParseExpression(LOWEST)
	if p.Err != nil {
		return nil
	}
	// 期待并消耗 ']'
	p.CheckNextAndAdvance(lexer.RBRACKET)
	if p.Err != nil {
		return nil
	}
	ie := &ast.IndexExpression{
		Target:   left,
		Index:    indexExpr,
		PosStart: posStart,
		PosEnd:   p.CurrToken.PosEnd.Copy(),
	}
	return ie
}
