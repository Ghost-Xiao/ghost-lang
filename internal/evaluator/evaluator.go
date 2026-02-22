package evaluator

import (
	"fmt"

	"github.com/Ghost-Xiao/ghost-lang/internal/frame"
	"github.com/Ghost-Xiao/ghost-lang/internal/lexer"
	"github.com/Ghost-Xiao/ghost-lang/internal/object"
	"github.com/Ghost-Xiao/ghost-lang/internal/parser/ast"
	"github.com/Ghost-Xiao/ghost-lang/internal/util"
)

// indexable 表示可索引接口

type indexable interface {
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
	Set(index object.Object, value object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) error
}

// Evaluator 解释器结构体，负责执行AST节点并管理运行时状态
// 包含一个错误字段用于捕获和传递运行时错误

type Evaluator struct {
	Frame *frame.Frame // 调用栈帧
	Err   error        // 运行时错误信息
}

// NewEvaluator 创建一个新的解释器实例
//
// 参数：
//
//	frame - 调用栈帧
//
// 返回值:
//
//	*Evaluator - 初始化后的解释器指针
func NewEvaluator(frame *frame.Frame) *Evaluator {
	return &Evaluator{
		Frame: frame,
		Err:   nil,
	}
}

// Eval 根据节点类型调用相应的访问方法
//
// 参数:
//
//		nodes - 要访问的AST节点
//	 env - 执行环境
//
// 返回值:
//
//	object.Object - 节点执行结果值，发生错误时为nil
func (e *Evaluator) Eval(nodes ast.Node, env *object.Environment) object.Object {
	// 根据节点类型分发到对应的处理方法
	switch n := nodes.(type) {
	case *ast.Program:
		return e.evalProgram(n, env)
	case *ast.ForStatement:
		return e.evalForStatement(n, env)
	case *ast.FunctionDeclarationStatement:
		return e.evalFunctionDeclarationStatement(n, env)
	case *ast.ReturnStatement:
		return e.evalReturnStatement(n, env)
	case *ast.BreakStatement:
		return e.evalBreakStatement(n, env)
	case *ast.ContinueStatement:
		return e.evalContinueStatement(n, env)
	case *ast.ExpressionStatement:
		return e.evalExpressionStatement(n, env)
	case *ast.EllipsisStatement:
		return nil
	case *ast.NamespaceStatement:
		return e.evalNamespaceStatement(n, env)
	case *ast.PrefixExpression:
		return e.evalPrefixExpression(n, env)
	case *ast.InfixExpression:
		return e.evalInfixExpression(n, env)
	case *ast.IntExpression:
		return e.evalIntExpression(n, env)
	case *ast.FloatExpression:
		return e.evalFloatExpression(n, env)
	case *ast.BoolExpression:
		return e.evalBooleanExpression(n, env)
	case *ast.NullExpression:
		return e.evalNullExpression(n, env)
	case *ast.StringExpression:
		return e.evalStringExpression(n, env)
	case *ast.ListExpression:
		return e.evalListExpression(n, env)
	case *ast.IdentifierExpression:
		return e.evalIdentifierExpression(n, env)
	case *ast.GroupedExpression:
		return e.Eval(n.Expr, env)
	case *ast.VarInitializationExpression:
		return e.evalVarInitializationExpression(n, env)
	case *ast.VarAssignmentExpression:
		return e.evalVarAssignmentExpression(n, env)
	case *ast.CompoundAssignmentExpression:
		return e.evalCompoundAssignmentExpression(n, env)
	case *ast.PrefixUnaryIncDecExpression:
		return e.evalPrefixUnaryIncDecExpression(n, env)
	case *ast.PostfixUnaryIncDecExpression:
		return e.evalPostfixUnaryIncDecExpression(n, env)
	case *ast.BlockExpression:
		return e.evalBlockExpression(n, env)
	case *ast.IfExpression:
		return e.evalIfExpression(n, env)
	case *ast.CallExpression:
		return e.evalCallExpression(n, env)
	case *ast.IndexExpression:
		return e.evalIndexExpression(n, env)
	case *ast.NamespaceAccessExpression:
		return e.evalNamespaceAccessExpression(n, env)
	default:
		panic(fmt.Sprintf("unknown node type: %T", n))
	}
}

// evalProgram 处理程序节点，依次执行所有语句
//
// 参数:
//
//	program - 程序节点，包含一系列语句
//	env - 执行环境
//
// 返回值:
//
//	object.Object - 程序执行结果(通常为nil)
//
// 错误处理:
//
//	若执行过程中发生错误，立即返回nil并设置e.Err
func (e *Evaluator) evalProgram(program *ast.Program, env *object.Environment) object.Object {
	for _, statement := range program.Statements {
		e.Eval(statement, env)
		if e.Err != nil {
			return nil
		}
	}
	return nil
}

// evalForStatement 处理for语句节点
// 执行for循环
//
// 参数:
//
//	forStatement - for语句节点
//	env - 执行环境
//
// 返回值:
//
//	object.Object - 始终返回nil
func (e *Evaluator) evalForStatement(forStatement *ast.ForStatement, env *object.Environment) object.Object {
	// 创建新环境
	forEnv := &object.Environment{
		Name:  "for",
		Store: make(map[string]*object.Symbol),
		Outer: env,
	}
	// 执行初始化语句
	e.Eval(forStatement.Initialization, forEnv)
	if e.Err != nil {
		return nil
	}
	// 执行条件表达式
	condition := e.Eval(forStatement.Condition, forEnv)
	if e.Err != nil {
		return nil
	}
	// 判断是不是布尔值
	if _, ok := condition.(*object.Bool); !ok {
		e.Err = &TypeError{
			Frame:    e.Frame,
			Message:  "non-bool condition in for loop.",
			PosStart: forStatement.PosStart,
			PosEnd:   forStatement.PosEnd,
		}
		return nil
	}
	// 执行循环体
	for condition.(*object.Bool).Value {
		// 执行循环体
		ret := e.evalWithSpecialValue(forStatement.Body, forEnv)
		if e.Err != nil {
			return nil
		}
		if returnValue, ok := ret.(*object.ReturnValue); ok {
			return returnValue
		}
		if _, ok := ret.(*object.BreakValue); ok {
			break
		}
		// 执行更新语句
		e.Eval(forStatement.Update, forEnv)
		if e.Err != nil {
			return nil
		}
		// 重新评估条件表达式
		condition = e.Eval(forStatement.Condition, forEnv)
		if e.Err != nil {
			return nil
		}
		// 判断是不是布尔值
		if _, ok := condition.(*object.Bool); !ok {
			e.Err = &TypeError{
				Frame:    e.Frame,
				Message:  "non-bool condition in for loop.",
				PosStart: forStatement.PosStart,
				PosEnd:   forStatement.PosEnd,
			}
			return nil
		}
	}
	return nil
}

// evalFunctionDeclarationStatement 处理函数声明语句节点
// 解释函数表达式
//
// 参数:
//
//	functionDeclarationStatement - 函数声明语句节点
//	env - 执行环境
//
// 返回值:
//
//	object.Object - 函数表达式的结果，发生错误时返回nil
func (e *Evaluator) evalFunctionDeclarationStatement(functionDeclarationStatement *ast.FunctionDeclarationStatement, env *object.Environment) object.Object {
	// 函数名字
	funcName := functionDeclarationStatement.Name.(*ast.IdentifierExpression).Name
	// 是否已定义过函数
	if _, ok := env.Get(funcName); ok {
		e.Err = &VariableError{
			Frame:    e.Frame,
			Message:  fmt.Sprintf("function \"%s\" already defined.", funcName),
			PosStart: functionDeclarationStatement.PosStart,
			PosEnd:   functionDeclarationStatement.PosEnd,
		}
		return nil
	}
	// 创建函数
	fn := &object.Function{
		Name:      funcName,
		Parameter: functionDeclarationStatement.Parameter,
		Body:      functionDeclarationStatement.Body,
		Env:       env,
	}
	// 绑定函数
	env.Set(funcName, &object.Symbol{
		Name:    funcName,
		Value:   fn,
		IsConst: true,
	})
	return nil
}

// evalReturnStatement 处理return语句节点
// 执行return语句，返回值
//
// 参数:
//
//	returnStatement - return语句节点
//	env - 执行环境
//
// 返回值:
//
//	object.Object
func (e *Evaluator) evalReturnStatement(returnStatement *ast.ReturnStatement, env *object.Environment) object.Object {
	if e.Frame.Parent == nil {
		e.Err = &SyntaxError{
			Frame:    e.Frame,
			Message:  "return statement is only allowed inside functions.",
			PosStart: returnStatement.PosStart,
			PosEnd:   returnStatement.PosEnd,
		}
		return nil
	}
	returnValue := e.Eval(returnStatement.ReturnValue, env)
	if e.Err != nil {
		return nil
	}
	// 返回ReturnValue对象
	return &object.ReturnValue{
		Value: returnValue,
	}
}

// evalBreakStatement 处理break语句节点
// 执行break语句，跳出当前循环
//
// 参数:
//
//	breakStatement - break语句节点
//	env - 执行环境
//
// 返回值:
//
//	object.Object - BreakValue实例
func (e *Evaluator) evalBreakStatement(breakStatement *ast.BreakStatement, env *object.Environment) object.Object {
	inLoop := false
	for e := env.Outer; e != nil; e = e.Outer {
		if e.Name == "for" {
			inLoop = true
			break
		}
		if e.Name == "function" {
			break
		}
	}
	if !inLoop {
		e.Err = &SyntaxError{
			Frame:    e.Frame,
			Message:  "break statement is only allowed inside for loops.",
			PosStart: breakStatement.PosStart,
			PosEnd:   breakStatement.PosEnd,
		}
		return nil
	}
	return &object.BreakValue{}
}

// evalContinueStatement 处理continue语句节点
// 执行continue语句，跳出当前循环
//
// 参数:
//
//	continueStatement - continue语句节点
//	env - 执行环境
//
// 返回值:
//
//	object.Object - ContinueValue实例
func (e *Evaluator) evalContinueStatement(continueStatement *ast.ContinueStatement, env *object.Environment) object.Object {
	inLoop := false
	for e := env.Outer; e != nil; e = e.Outer {
		if e.Name == "for" {
			inLoop = true
			break
		}
		if e.Name == "function" {
			break
		}
	}
	if !inLoop {
		e.Err = &SyntaxError{
			Frame:    e.Frame,
			Message:  "continue statement is only allowed inside for loops.",
			PosStart: continueStatement.PosStart,
			PosEnd:   continueStatement.PosEnd,
		}
		return nil
	}
	return &object.ContinueValue{}
}

// evalIndexExpression 处理索引表达式节点
// 执行索引表达式
//
// 参数:
//
//	indexExpression - 索引表达式节点
//	env - 执行环境
//
// 返回值:
//
//	object.Object
func (e *Evaluator) evalIndexExpression(indexExpression *ast.IndexExpression, env *object.Environment) object.Object {
	target := e.Eval(indexExpression.Target, env)
	if e.Err != nil {
		return nil
	}
	idxObj := e.Eval(indexExpression.Index, env)
	if e.Err != nil {
		return nil
	}
	// 判断索引是否是整数
	intIdx, ok := idxObj.(*object.Int)
	if !ok {
		e.Err = &TypeError{
			Frame:    e.Frame,
			Message:  "index must be integer.",
			PosStart: indexExpression.PosStart,
			PosEnd:   indexExpression.PosEnd,
		}
		return nil
	}
	ret, err := target.Index(intIdx, indexExpression.PosStart, indexExpression.PosEnd, e.Frame)
	if err != nil {
		e.Err = err
		return nil
	}
	return ret
}

// evalNamespaceAccessExpression 处理命名空间访问表达式节点
// 执行命名空间访问表达式
//
// 参数:
//
//	namespaceAccessExpression - 命名空间访问表达式节点
//	env - 执行环境
//
// 返回值:
//
//	object.Object
func (e *Evaluator) evalNamespaceAccessExpression(namespaceAccessExpression *ast.NamespaceAccessExpression, env *object.Environment) object.Object {
	target := e.Eval(namespaceAccessExpression.Target, env)
	if e.Err != nil {
		return nil
	}
	if _, ok := namespaceAccessExpression.Member.(*ast.IdentifierExpression); !ok {
		e.Err = &TypeError{
			Frame:    e.Frame,
			Message:  "member must be identifier.",
			PosStart: namespaceAccessExpression.PosStart,
			PosEnd:   namespaceAccessExpression.PosEnd,
		}
		return nil
	}
	member := namespaceAccessExpression.Member.(*ast.IdentifierExpression).Name
	// 判断是否是命名空间
	if namespace, ok := target.(*object.Namespace); ok {
		// 获取成员
		ret, ok := namespace.Member.Get(member)
		if !ok {
			e.Err = &VariableError{
				Frame:    e.Frame,
				Message:  fmt.Sprintf("undefined member \"%s\".", member),
				PosStart: namespaceAccessExpression.PosStart,
				PosEnd:   namespaceAccessExpression.PosEnd,
			}
			return nil
		}
		return ret.Value
	} else {
		e.Err = &VariableError{
			Frame:    e.Frame,
			Message:  "target must be namespace.",
			PosStart: namespaceAccessExpression.PosStart,
			PosEnd:   namespaceAccessExpression.PosEnd,
		}
		return nil
	}
}

// evalExpressionStatement 处理表达式语句节点
// 执行表达式并忽略其返回值
//
// 参数:
//
//	expressionStatement - 表达式语句节点
//	env - 执行环境
//
// 返回值:
//
//	object.Object
func (e *Evaluator) evalExpressionStatement(expressionStatement *ast.ExpressionStatement, env *object.Environment) object.Object {
	ret := e.Eval(expressionStatement.Expr, env)
	if e.Err != nil {
		return nil
	}
	if returnValue, ok := ret.(*object.ReturnValue); ok {
		return returnValue
	}
	return nil
}

// evalNamespaceStatement 处理命名空间语句节点
// 创建一个新的命名空间环境
//
// 参数:
//
//	namespaceStatement - 命名空间语句节点
//	env - 执行环境
//
// 返回值:
//
//	object.Object
func (e *Evaluator) evalNamespaceStatement(namespaceStatement *ast.NamespaceStatement, env *object.Environment) object.Object {
	name := namespaceStatement.Name.(*ast.IdentifierExpression).Name
	// 是否已定义过命名空间
	if _, ok := env.Get(name); ok {
		e.Err = &VariableError{
			Frame:    e.Frame,
			Message:  fmt.Sprintf("namespace \"%s\" already defined.", name),
			PosStart: namespaceStatement.PosStart,
			PosEnd:   namespaceStatement.PosEnd,
		}
		return nil
	}
	// 创建新环境
	namespaceEnv := &object.Environment{
		Name:  name,
		Store: make(map[string]*object.Symbol),
		Outer: env,
	}
	// 执行命名空间体
	switch n := namespaceStatement.Body.(type) {
	case *ast.ExpressionStatement:
		switch expr := n.Expr.(type) {
		case *ast.BlockExpression:
			for _, stmt := range expr.Statements {
				switch s := stmt.(type) {
				case *ast.ExpressionStatement:
					if _, ok := s.Expr.(ast.Definition); !ok {
						e.Err = &SyntaxError{
							Frame:    e.Frame,
							Message:  "namespace body must be definitions.",
							PosStart: namespaceStatement.PosStart,
							PosEnd:   namespaceStatement.PosEnd,
						}
						return nil
					}
				default:
					if _, ok := s.(ast.Definition); !ok {
						e.Err = &SyntaxError{
							Frame:    e.Frame,
							Message:  "namespace body must be definitions.",
							PosStart: namespaceStatement.PosStart,
							PosEnd:   namespaceStatement.PosEnd,
						}
						return nil
					}
				}
				e.Eval(stmt, namespaceEnv)
				if e.Err != nil {
					return nil
				}
			}
		default:
			if _, ok := expr.(ast.Definition); !ok {
				e.Err = &SyntaxError{
					Frame:    e.Frame,
					Message:  "namespace body must be definitions.",
					PosStart: namespaceStatement.PosStart,
					PosEnd:   namespaceStatement.PosEnd,
				}
				return nil
			}
			e.Eval(expr, namespaceEnv)
			if e.Err != nil {
				return nil
			}
		}
	default:
		if _, ok := n.(ast.Definition); !ok {
			e.Err = &SyntaxError{
				Frame:    e.Frame,
				Message:  "namespace body must be definitions.",
				PosStart: namespaceStatement.PosStart,
				PosEnd:   namespaceStatement.PosEnd,
			}
			return nil
		}
		e.Eval(n, namespaceEnv)
		if e.Err != nil {
			return nil
		}
	}
	if e.Err != nil {
		return nil
	}
	// 创建命名空间对象
	namespace := &object.Namespace{
		Name:   name,
		Member: namespaceEnv,
	}
	sym := &object.Symbol{
		Name:    name,
		Value:   namespace,
		IsConst: true,
	}
	// 绑定命名空间
	env.Set(name, sym)
	return nil
}

// evalIntExpression 处理整数表达式节点
// 将AST整数节点转换为运行时整数值
//
// 参数:
//
//	numberExpression - 整数表达式节点
//	env - 执行环境
//
// 返回值:
//
//	object.Object - 包含整数值的value.Int实例
func (e *Evaluator) evalIntExpression(numberExpression *ast.IntExpression, _ *object.Environment) object.Object {
	return &object.Int{Value: numberExpression.Value}
}

// evalFloatExpression 处理浮点数表达式节点
// 将AST浮点数节点转换为运行时浮点值
//
// 参数:
//
//	numberExpression - 浮点数表达式节点
//	env - 执行环境
//
// 返回值:
//
//	object.Object - 包含浮点值的value.Float实例
func (e *Evaluator) evalFloatExpression(numberExpression *ast.FloatExpression, _ *object.Environment) object.Object {
	return &object.Float{Value: numberExpression.Value}
}

// evalBooleanExpression 处理布尔表达式节点
// 将AST布尔节点转换为运行时布尔值
//
// 参数:
//
//	booleanExpression - 布尔表达式节点
//	env - 执行环境
//
// 返回值:
//
//	object.Object - 包含布尔值的value.Bool实例
func (e *Evaluator) evalBooleanExpression(booleanExpression *ast.BoolExpression, _ *object.Environment) object.Object {
	return &object.Bool{Value: booleanExpression.Value}
}

// evalNullExpression 处理空值表达式节点
// 返回运行时空值
//
// 参数:
//
//	_ - 空值表达式节点(未使用)
//	env - 执行环境
//
// 返回值:
//
//	object.Object - 空值value.Null实例
func (e *Evaluator) evalNullExpression(_ *ast.NullExpression, _ *object.Environment) object.Object {
	return &object.Null{}
}

// evalStringExpression 处理字符串表达式节点
// 将AST字符串节点转换为运行时字符串值
//
// 参数:
//
//	stringExpression - 字符串表达式节点
//	env - 执行环境
//
// 返回值:
//
//	object.Object - 包含字符串值的value.String实例
func (e *Evaluator) evalStringExpression(stringExpression *ast.StringExpression, _ *object.Environment) object.Object {
	return &object.String{Value: stringExpression.Value}
}

// evalListExpression 处理列表表达式节点
// 将AST列表节点转换为运行时列表值，并验证元素类型一致性
//
// 参数:
//
//	listExpression - 列表表达式节点
//	env - 执行环境
//
// 返回值:
//
//	object.Object - 包含列表值的object.List实例，错误时返回nil
//
// 错误处理:
//
//	若列表元素类型不一致，设置TypeError并返回nil
func (e *Evaluator) evalListExpression(listExpression *ast.ListExpression, env *object.Environment) object.Object {
	elements := make([]object.Object, 0, len(listExpression.Value))
	var firstType string
	// 解释每个列表元素
	for i, elementExpr := range listExpression.Value {
		element := e.Eval(elementExpr, env)
		if e.Err != nil {
			return nil
		}
		// 第一个元素确定列表的类型
		if i == 0 {
			firstType = element.Type()
		} else {
			// 检查后续元素类型是否与第一个元素一致
			if element.Type() != firstType {
				e.Err = &TypeError{
					Frame:    e.Frame,
					Message:  "list elements must have consistent types.",
					PosStart: listExpression.PosStart,
					PosEnd:   listExpression.PosEnd,
				}
				return nil
			}
		}
		elements = append(elements, element)
	}
	return &object.List{Elements: elements}
}

// evalIdentifierExpression 处理标识符表达式节点
// 在符号表中查找标识符并返回对应的值
//
// 参数:
//
//	identifierExpression - 标识符表达式节点
//	env - 执行环境
//
// 返回值:
//
//	object.Object - 标识符对应的值，若未找到则返回nil并设置e.Err
//
// 错误处理:
//
//	若标识符未定义，设置VariableError并返回nil
func (e *Evaluator) evalIdentifierExpression(identifierExpression *ast.IdentifierExpression, env *object.Environment) object.Object {
	varName := identifierExpression.Name
	val, ok := env.Get(varName)
	if !ok {
		e.Err = &VariableError{
			Frame:    e.Frame,
			Message:  fmt.Sprintf("undefined variable \"%s\".", varName),
			PosStart: identifierExpression.PosStart,
			PosEnd:   identifierExpression.PosEnd,
		}
		return nil
	}
	return val.Value
}

// evalVarInitializationExpression 处理变量初始化节点
// 在当前上下文中声明并初始化变量或常量
//
// 参数:
//
//	varInitialization - 变量初始化节点
//	env - 执行环境
//
// 返回值:
//
//	object.Object - 已声明的变量值，发生错误时返回nil
//
// 错误处理:
//
//   - 尝试重定义常量时返回错误
//   - 尝试将变量重新声明为常量时返回错误
func (e *Evaluator) evalVarInitializationExpression(varInitialization *ast.VarInitializationExpression, env *object.Environment) object.Object {
	varName := varInitialization.Name.(*ast.IdentifierExpression).Name
	// 检查变量是否已定义
	if env.Exists(varName) {
		e.Err = &VariableError{
			Frame:    e.Frame,
			Message:  fmt.Sprintf("variable \"%s\" already defined.", varName),
			PosStart: varInitialization.PosStart,
			PosEnd:   varInitialization.PosEnd,
		}
		return nil
	}
	// 计算并赋值
	val := e.Eval(varInitialization.Value, env)
	if e.Err != nil {
		return nil
	}
	// 创建符号
	var sym = &object.Symbol{
		Name:    varName,
		Value:   val,
		IsConst: varInitialization.IsConst,
	}
	env.Set(varName, sym)
	return val
}

// getSymbol 根据表达式获取对应的符号对象
// 支持标识符表达式和命名空间访问表达式
//
// 参数:
//
//	expr - AST表达式节点（IdentifierExpression 或 NamespaceAccessExpression）
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//
// 返回值:
//
//	*object.Symbol - 符号对象
//	error - 错误信息（变量未定义、类型错误等）
func (e *Evaluator) getSymbol(expr ast.Expression, env *object.Environment, posStart, posEnd *util.Pos) (*object.Symbol, error) {
	switch ex := expr.(type) {
	case *ast.IdentifierExpression:
		sym, ok := env.Get(ex.Name)
		if !ok {
			return nil, &VariableError{
				Frame:    e.Frame,
				Message:  fmt.Sprintf("undefined variable \"%s\".", ex.Name),
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
		return sym, nil
	case *ast.NamespaceAccessExpression:
		tar := e.Eval(ex.Target, env)
		if e.Err != nil {
			return nil, e.Err
		}
		if _, ok := ex.Member.(*ast.IdentifierExpression); !ok {
			return nil, &TypeError{
				Frame:    e.Frame,
				Message:  "member must be identifier.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
		member := ex.Member.(*ast.IdentifierExpression).Name
		if namespace, ok := tar.(*object.Namespace); ok {
			ret, ok := namespace.Member.Get(member)
			if !ok {
				return nil, &VariableError{
					Frame:    e.Frame,
					Message:  fmt.Sprintf("undefined member \"%s\".", member),
					PosStart: posStart,
					PosEnd:   posEnd,
				}
			}
			return ret, nil
		}
		return nil, &VariableError{
			Frame:    e.Frame,
			Message:  "target must be namespace.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	default:
		return nil, &TypeError{
			Frame:    e.Frame,
			Message:  "invalid variable name type.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
}

// checkIndexTargetConst 检查索引表达式的目标是否为常量
//
// 参数:
//
//	target - 索引表达式的目标
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//
// 返回值:
//
//	error - 如果目标是常量则返回错误，否则返回nil
func (e *Evaluator) checkIndexTargetConst(target ast.Expression, env *object.Environment, posStart, posEnd *util.Pos) error {
	switch t := target.(type) {
	case *ast.IdentifierExpression:
		// 检查标识符是否为常量
		sym, ok := env.Get(t.Name)
		if !ok {
			return &VariableError{
				Frame:    e.Frame,
				Message:  fmt.Sprintf("undefined variable \"%s\".", t.Name),
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
		if sym.IsConst {
			return &VariableError{
				Frame:    e.Frame,
				Message:  fmt.Sprintf("cannot redefine constant \"%s\".", t.Name),
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
	case *ast.IndexExpression:
		// 递归检查嵌套索引表达式的目标
		return e.checkIndexTargetConst(t.Target, env, posStart, posEnd)
	case *ast.NamespaceAccessExpression:
		tar := e.Eval(t.Target, env)
		if e.Err != nil {
			return e.Err
		}
		if _, ok := t.Member.(*ast.IdentifierExpression); !ok {
			return &TypeError{
				Frame:    e.Frame,
				Message:  "member must be identifier.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
		member := t.Member.(*ast.IdentifierExpression).Name
		// 判断是否是命名空间
		if namespace, ok := tar.(*object.Namespace); ok {
			// 获取成员
			ret, ok := namespace.Member.Get(member)
			if !ok {
				return &VariableError{
					Frame:    e.Frame,
					Message:  fmt.Sprintf("undefined member \"%s\".", member),
					PosStart: posStart,
					PosEnd:   posEnd,
				}
			}
			if ret.IsConst {
				return &VariableError{
					Frame:    e.Frame,
					Message:  fmt.Sprintf("cannot redefine constant \"%s\".", member),
					PosStart: posStart,
					PosEnd:   posEnd,
				}
			}
			return nil
		} else {
			return &VariableError{
				Frame:    e.Frame,
				Message:  "target must be namespace.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
	}
	return nil
}

// evalVarAssignmentExpression 处理变量赋值节点
// 在当前上下文中对变量进行赋值
//
// 参数:
//
//	VarAssignmentExpression - 变量赋值节点
//	env - 执行环境
//
// 返回值:
//
//	object.Object - 已声明的变量值，发生错误时返回nil
//
// 错误处理:
//
//   - 尝试重定义常量时返回错误
//   - 尝试将变量重新声明为常量时返回错误
func (e *Evaluator) evalVarAssignmentExpression(varAssignment *ast.VarAssignmentExpression, env *object.Environment) object.Object {
	// 计算右侧表达式的值
	value := e.Eval(varAssignment.Value, env)
	if e.Err != nil {
		return nil
	}

	// 根据左值类型进行赋值
	switch name := varAssignment.Name.(type) {
	case *ast.IdentifierExpression, *ast.NamespaceAccessExpression:
		// 使用 getSymbol 获取符号
		sym, err := e.getSymbol(name, env, varAssignment.PosStart, varAssignment.PosEnd)
		if err != nil {
			e.Err = err
			return nil
		}
		// 检查是否是常量
		if sym.IsConst {
			var varName string
			if id, ok := name.(*ast.IdentifierExpression); ok {
				varName = id.Name
			} else if ns, ok := name.(*ast.NamespaceAccessExpression); ok {
				varName = ns.Member.(*ast.IdentifierExpression).Name
			}
			e.Err = &VariableError{
				Frame:    e.Frame,
				Message:  fmt.Sprintf("cannot redefine constant \"%s\".", varName),
				PosStart: varAssignment.PosStart,
				PosEnd:   varAssignment.PosEnd,
			}
			return nil
		}
		// 更新符号值
		newSym := &object.Symbol{
			Name:    sym.Name,
			Value:   value,
			IsConst: false,
		}
		if id, ok := name.(*ast.IdentifierExpression); ok {
			env.Assign(id.Name, newSym)
		} else if ns, ok := name.(*ast.NamespaceAccessExpression); ok {
			tar := e.Eval(ns.Target, env)
			if e.Err != nil {
				return nil
			}
			if namespace, ok := tar.(*object.Namespace); ok {
				namespace.Member.Set(sym.Name, newSym)
			}
		}
		return value
	case *ast.IndexExpression:
		// 检查索引目标是否为常量
		err := e.checkIndexTargetConst(name.Target, env, name.PosStart, name.PosEnd)
		if err != nil {
			e.Err = err
			return nil
		}
		target := e.Eval(name.Target, env)
		if e.Err != nil {
			return nil
		}
		index := e.Eval(name.Index, env)
		if e.Err != nil {
			return nil
		}
		// 判断索引是否是整数
		if _, ok := index.(*object.Int); !ok {
			e.Err = &TypeError{
				Frame:    e.Frame,
				Message:  "index must be integer.",
				PosStart: varAssignment.PosStart,
				PosEnd:   varAssignment.PosEnd,
			}
			return nil
		}
		// 检查目标是否可索引
		idxable, ok := target.(indexable)
		if !ok {
			e.Err = &TypeError{
				Frame:    e.Frame,
				Message:  "index expression not supported for this type.",
				PosStart: varAssignment.PosStart,
				PosEnd:   varAssignment.PosEnd,
			}
			return nil
		}
		// 设置值
		err2 := idxable.Set(index, value, varAssignment.PosStart, varAssignment.PosEnd, e.Frame)
		if err2 != nil {
			e.Err = err2
			return nil
		}
		return value
	default:
		e.Err = &TypeError{
			Frame:    e.Frame,
			Message:  "invalid variable name type.",
			PosStart: varAssignment.PosStart,
			PosEnd:   varAssignment.PosEnd,
		}
		return nil
	}
}

// evalCompoundAssignmentExpression 处理变量复合赋值节点
// 在当前上下文中对变量进行复合赋值
//
// 参数:
//
//	compoundAssignmentExpression - 变量复合赋值节点
//	env - 执行环境
//
// 返回值:
//
//	object.Object - 已声明的变量值，发生错误时返回nil
//
// 错误处理:
//
//   - 尝试重定义常量时返回错误
//   - 尝试将变量重新声明为常量时返回错误
func (e *Evaluator) evalCompoundAssignmentExpression(compoundAssignmentExpression *ast.CompoundAssignmentExpression, env *object.Environment) object.Object {
	// 计算右侧表达式
	right := e.Eval(compoundAssignmentExpression.Right, env)
	if e.Err != nil {
		return nil
	}

	// 获取运算符字面量（去掉最后的等号）
	literal := compoundAssignmentExpression.Operator.Literal[:len(compoundAssignmentExpression.Operator.Literal)-1]
	// 获取并创建基础运算符令牌
	baseOperator := &lexer.Token{
		Type:    lexer.CompoundAssignmentOperators[compoundAssignmentExpression.Operator.Type],
		Literal: literal,
	}

	// 用于存储原始值和设置新值的函数
	var oldValue object.Object
	var setValue func(newVal object.Object)

	// 根据左值类型获取原始值和设置函数
	switch name := compoundAssignmentExpression.Name.(type) {
	case *ast.IdentifierExpression, *ast.NamespaceAccessExpression:
		// 使用 getSymbol 获取符号
		sym, err := e.getSymbol(name, env, compoundAssignmentExpression.PosStart, compoundAssignmentExpression.PosEnd)
		if err != nil {
			e.Err = err
			return nil
		}
		// 检查是否是常量
		if sym.IsConst {
			var varName string
			if id, ok := name.(*ast.IdentifierExpression); ok {
				varName = id.Name
			} else if ns, ok := name.(*ast.NamespaceAccessExpression); ok {
				varName = ns.Member.(*ast.IdentifierExpression).Name
			}
			e.Err = &VariableError{
				Frame:    e.Frame,
				Message:  fmt.Sprintf("cannot redefine constant \"%s\".", varName),
				PosStart: compoundAssignmentExpression.PosStart,
				PosEnd:   compoundAssignmentExpression.PosEnd,
			}
			return nil
		}
		oldValue = sym.Value
		setValue = func(newVal object.Object) {
			newSym := &object.Symbol{
				Name:    sym.Name,
				Value:   newVal,
				IsConst: false,
			}
			if id, ok := name.(*ast.IdentifierExpression); ok {
				env.Assign(id.Name, newSym)
			} else if ns, ok := name.(*ast.NamespaceAccessExpression); ok {
				tar := e.Eval(ns.Target, env)
				if e.Err == nil {
					if namespace, ok := tar.(*object.Namespace); ok {
						namespace.Member.Set(sym.Name, newSym)
					}
				}
			}
		}
	case *ast.IndexExpression:
		// 检查索引目标是否为常量
		err := e.checkIndexTargetConst(name.Target, env, name.PosStart, name.PosEnd)
		if err != nil {
			e.Err = err
			return nil
		}
		target := e.Eval(name.Target, env)
		if e.Err != nil {
			return nil
		}
		index := e.Eval(name.Index, env)
		if e.Err != nil {
			return nil
		}
		// 判断索引是否是整数
		if _, ok := index.(*object.Int); !ok {
			e.Err = &TypeError{
				Frame:    e.Frame,
				Message:  "index must be integer.",
				PosStart: compoundAssignmentExpression.PosStart,
				PosEnd:   compoundAssignmentExpression.PosEnd,
			}
			return nil
		}
		// 检查目标是否可索引
		idxable, ok := target.(indexable)
		if !ok {
			e.Err = &TypeError{
				Frame:    e.Frame,
				Message:  "index expression not supported for this type.",
				PosStart: compoundAssignmentExpression.PosStart,
				PosEnd:   compoundAssignmentExpression.PosEnd,
			}
			return nil
		}
		// 获取目标索引的值
		oldValue = e.Eval(name, env)
		if e.Err != nil {
			return nil
		}
		setValue = func(newVal object.Object) {
			err2 := idxable.Set(index, newVal, compoundAssignmentExpression.PosStart, compoundAssignmentExpression.PosEnd, e.Frame)
			if err2 != nil {
				e.Err = err2
			}
		}
	default:
		e.Err = &TypeError{
			Frame:    e.Frame,
			Message:  "invalid variable name type.",
			PosStart: compoundAssignmentExpression.PosStart,
			PosEnd:   compoundAssignmentExpression.PosEnd,
		}
		return nil
	}

	// 执行复合赋值运算
	value := e.evalInfixOperator(&ast.InfixExpression{
		Left:     compoundAssignmentExpression.Name,
		Operator: baseOperator,
		Right:    compoundAssignmentExpression.Right,
		PosStart: compoundAssignmentExpression.PosStart,
		PosEnd:   compoundAssignmentExpression.PosEnd,
	}, oldValue, right)
	if e.Err != nil {
		return nil
	}

	// 设置新值
	setValue(value)
	if e.Err != nil {
		return nil
	}

	return value
}

// evalPrefixExpression 处理前缀表达式节点
// 执行前缀运算符(如!、-)运算
//
// 参数:
//
//	prefixExpression - 前缀表达式节点
//	env - 执行环境
//
// 返回值:
//
//	object.Object - 运算结果，发生错误时返回nil
//
// 错误处理:
//
//	若运算符不支持，设置OperationError并返回nil
func (e *Evaluator) evalPrefixExpression(prefixExpression *ast.PrefixExpression, env *object.Environment) object.Object {
	right := e.Eval(prefixExpression.Value, env)
	if e.Err != nil {
		return nil
	}
	val := e.evalPrefixOperator(prefixExpression, right)
	if e.Err != nil {
		return nil
	}
	return val
}

func (e *Evaluator) evalPrefixOperator(prefixExpression *ast.PrefixExpression, right object.Object) object.Object {
	switch prefixExpression.Operator.Type {
	case lexer.MINUS:
		val, err := right.Negative(prefixExpression.PosStart, prefixExpression.PosEnd, e.Frame)
		if err != nil {
			e.Err = err
			return nil
		}
		return val
	case lexer.BANG:
		val, err := right.Not(prefixExpression.PosStart, prefixExpression.PosEnd, e.Frame)
		if err != nil {
			e.Err = err
			return nil
		}
		return val
	case lexer.BITWISE_NOT:
		val, err := right.BitNot(prefixExpression.PosStart, prefixExpression.PosEnd, e.Frame)
		if err != nil {
			e.Err = err
			return nil
		}
		return val
	default:
		e.Err = &object.OperationError{
			Message:  fmt.Sprintf("invalid operation \"%s\".", prefixExpression.Operator.Type),
			PosStart: prefixExpression.PosStart,
			PosEnd:   prefixExpression.PosEnd,
		}
		return nil
	}
}

// evalPrefixUnaryIncDecExpression 处理前缀自增 / 自减表达式节点
// 执行前缀自增 / 自减表达式(如++a、--b)运算
//
// 参数:
//
//	prefixUnaryIncDecExpression - 前缀自增 / 自减表达式节点
//	env - 执行环境
//
// 返回值:
//
//	object.Object - 运算结果，发生错误时返回nil
//
// 错误处理:
//
//	若变量是常量，设置VariableError并返回nil
func (e *Evaluator) evalPrefixUnaryIncDecExpression(prefixUnaryIncDecExpression *ast.PrefixUnaryIncDecExpression, env *object.Environment) object.Object {
	// 构建运算符
	var operator *lexer.Token
	if prefixUnaryIncDecExpression.Operator.Type == lexer.INCREMENT {
		operator = &lexer.Token{
			Type:    lexer.PLUS,
			Literal: "+",
		}
	} else {
		operator = &lexer.Token{
			Type:    lexer.MINUS,
			Literal: "-",
		}
	}

	// 用于存储原始值和设置新值的函数
	var oldValue object.Object
	var setValue func(newVal object.Object)

	// 根据右值类型获取原始值和设置函数
	switch right := prefixUnaryIncDecExpression.Right.(type) {
	case *ast.IdentifierExpression, *ast.NamespaceAccessExpression:
		// 使用 getSymbol 获取符号
		sym, err := e.getSymbol(right, env, prefixUnaryIncDecExpression.PosStart, prefixUnaryIncDecExpression.PosEnd)
		if err != nil {
			e.Err = err
			return nil
		}
		// 检查是否是常量
		if sym.IsConst {
			var varName string
			if id, ok := right.(*ast.IdentifierExpression); ok {
				varName = id.Name
			} else if ns, ok := right.(*ast.NamespaceAccessExpression); ok {
				varName = ns.Member.(*ast.IdentifierExpression).Name
			}
			e.Err = &VariableError{
				Frame:    e.Frame,
				Message:  fmt.Sprintf("cannot redefine constant \"%s\".", varName),
				PosStart: prefixUnaryIncDecExpression.PosStart,
				PosEnd:   prefixUnaryIncDecExpression.PosEnd,
			}
			return nil
		}
		oldValue = sym.Value
		setValue = func(newVal object.Object) {
			newSym := &object.Symbol{
				Name:    sym.Name,
				Value:   newVal,
				IsConst: false,
			}
			if id, ok := right.(*ast.IdentifierExpression); ok {
				env.Set(id.Name, newSym)
			} else if ns, ok := right.(*ast.NamespaceAccessExpression); ok {
				tar := e.Eval(ns.Target, env)
				if e.Err == nil {
					if namespace, ok := tar.(*object.Namespace); ok {
						namespace.Member.Set(sym.Name, newSym)
					}
				}
			}
		}
	case *ast.IndexExpression:
		// 检查索引目标是否为常量
		err := e.checkIndexTargetConst(right.Target, env, right.PosStart, right.PosEnd)
		if err != nil {
			e.Err = err
			return nil
		}
		target := e.Eval(right.Target, env)
		if e.Err != nil {
			return nil
		}
		index := e.Eval(right.Index, env)
		if e.Err != nil {
			return nil
		}
		// 判断索引是否是整数
		if _, ok := index.(*object.Int); !ok {
			e.Err = &TypeError{
				Frame:    e.Frame,
				Message:  "index must be integer.",
				PosStart: prefixUnaryIncDecExpression.PosStart,
				PosEnd:   prefixUnaryIncDecExpression.PosEnd,
			}
			return nil
		}
		// 检查目标是否可索引
		idxable, ok := target.(indexable)
		if !ok {
			e.Err = &TypeError{
				Frame:    e.Frame,
				Message:  "index expression not supported for this type.",
				PosStart: prefixUnaryIncDecExpression.PosStart,
				PosEnd:   prefixUnaryIncDecExpression.PosEnd,
			}
			return nil
		}
		// 获取索引值
		oldValue = e.Eval(right, env)
		if e.Err != nil {
			return nil
		}
		setValue = func(newVal object.Object) {
			err2 := idxable.Set(index, newVal, prefixUnaryIncDecExpression.PosStart, prefixUnaryIncDecExpression.PosEnd, e.Frame)
			if err2 != nil {
				e.Err = err2
			}
		}
	default:
		e.Err = &TypeError{
			Frame:    e.Frame,
			Message:  "invalid variable name type.",
			PosStart: prefixUnaryIncDecExpression.PosStart,
			PosEnd:   prefixUnaryIncDecExpression.PosEnd,
		}
		return nil
	}

	// 执行运算符
	val := e.evalInfixOperator(&ast.InfixExpression{
		Left:     prefixUnaryIncDecExpression.Right,
		Operator: operator,
		Right: &ast.IntExpression{
			Value:    1,
			PosStart: prefixUnaryIncDecExpression.PosStart,
			PosEnd:   prefixUnaryIncDecExpression.PosEnd,
		},
		PosStart: prefixUnaryIncDecExpression.PosStart,
		PosEnd:   prefixUnaryIncDecExpression.PosEnd,
	}, oldValue, &object.Int{Value: 1})
	if e.Err != nil {
		return nil
	}

	// 设置新值
	setValue(val)
	if e.Err != nil {
		return nil
	}

	// 前缀自增/自减返回新值
	return val
}

// postfixUnaryIncDecExpression 处理后缀自增 / 自减表达式节点
// 执行后缀自增 / 自减表达式(如a++、b--)运算
//
// 参数:
//
//	postfixUnaryIncDecExpression - 后缀自增 / 自减表达式节点
//	env - 执行环境
//
// 返回值:
//
//	object.Object - 运算结果，发生错误时返回nil
//
// 错误处理:
//
//	若变量是常量，设置VariableError并返回nil
func (e *Evaluator) evalPostfixUnaryIncDecExpression(postfixUnaryIncDecExpression *ast.PostfixUnaryIncDecExpression, env *object.Environment) object.Object {
	// 构建运算符
	var operator *lexer.Token
	if postfixUnaryIncDecExpression.Operator.Type == lexer.INCREMENT {
		operator = &lexer.Token{
			Type:    lexer.PLUS,
			Literal: "+",
		}
	} else {
		operator = &lexer.Token{
			Type:    lexer.MINUS,
			Literal: "-",
		}
	}

	// 用于存储原始值和设置新值的函数
	var oldValue object.Object
	var setValue func(newVal object.Object)

	// 根据左值类型获取原始值和设置函数
	switch left := postfixUnaryIncDecExpression.Left.(type) {
	case *ast.IdentifierExpression, *ast.NamespaceAccessExpression:
		// 使用 getSymbol 获取符号
		sym, err := e.getSymbol(left, env, postfixUnaryIncDecExpression.PosStart, postfixUnaryIncDecExpression.PosEnd)
		if err != nil {
			e.Err = err
			return nil
		}
		// 检查是否是常量
		if sym.IsConst {
			var varName string
			if id, ok := left.(*ast.IdentifierExpression); ok {
				varName = id.Name
			} else if ns, ok := left.(*ast.NamespaceAccessExpression); ok {
				varName = ns.Member.(*ast.IdentifierExpression).Name
			}
			e.Err = &VariableError{
				Frame:    e.Frame,
				Message:  fmt.Sprintf("cannot redefine constant \"%s\".", varName),
				PosStart: postfixUnaryIncDecExpression.PosStart,
				PosEnd:   postfixUnaryIncDecExpression.PosEnd,
			}
			return nil
		}
		oldValue = sym.Value
		setValue = func(newVal object.Object) {
			newSym := &object.Symbol{
				Name:    sym.Name,
				Value:   newVal,
				IsConst: false,
			}
			if id, ok := left.(*ast.IdentifierExpression); ok {
				env.Set(id.Name, newSym)
			} else if ns, ok := left.(*ast.NamespaceAccessExpression); ok {
				tar := e.Eval(ns.Target, env)
				if e.Err == nil {
					if namespace, ok := tar.(*object.Namespace); ok {
						namespace.Member.Set(sym.Name, newSym)
					}
				}
			}
		}
	case *ast.IndexExpression:
		// 检查索引目标是否为常量
		err := e.checkIndexTargetConst(left.Target, env, left.PosStart, left.PosEnd)
		if err != nil {
			e.Err = err
			return nil
		}
		target := e.Eval(left.Target, env)
		if e.Err != nil {
			return nil
		}
		index := e.Eval(left.Index, env)
		if e.Err != nil {
			return nil
		}
		// 判断索引是否是整数
		if _, ok := index.(*object.Int); !ok {
			e.Err = &TypeError{
				Frame:    e.Frame,
				Message:  "index must be integer.",
				PosStart: postfixUnaryIncDecExpression.PosStart,
				PosEnd:   postfixUnaryIncDecExpression.PosEnd,
			}
			return nil
		}
		// 检查目标是否可索引
		idxable, ok := target.(indexable)
		if !ok {
			e.Err = &TypeError{
				Frame:    e.Frame,
				Message:  "index expression not supported for this type.",
				PosStart: postfixUnaryIncDecExpression.PosStart,
				PosEnd:   postfixUnaryIncDecExpression.PosEnd,
			}
			return nil
		}
		// 获取索引值
		oldValue = e.evalIndexExpression(left, env)
		if e.Err != nil {
			return nil
		}
		setValue = func(newVal object.Object) {
			err2 := idxable.Set(index, newVal, postfixUnaryIncDecExpression.PosStart, postfixUnaryIncDecExpression.PosEnd, e.Frame)
			if err2 != nil {
				e.Err = err2
			}
		}
	default:
		e.Err = &TypeError{
			Frame:    e.Frame,
			Message:  "invalid variable name type.",
			PosStart: postfixUnaryIncDecExpression.PosStart,
			PosEnd:   postfixUnaryIncDecExpression.PosEnd,
		}
		return nil
	}

	// 执行运算符
	val := e.evalInfixOperator(&ast.InfixExpression{
		Left:     postfixUnaryIncDecExpression.Left,
		Operator: operator,
		Right: &ast.IntExpression{
			Value:    1,
			PosStart: postfixUnaryIncDecExpression.PosStart,
			PosEnd:   postfixUnaryIncDecExpression.PosEnd,
		},
		PosStart: postfixUnaryIncDecExpression.PosStart,
		PosEnd:   postfixUnaryIncDecExpression.PosEnd,
	}, oldValue, &object.Int{Value: 1})
	if e.Err != nil {
		return nil
	}

	// 设置新值
	setValue(val)
	if e.Err != nil {
		return nil
	}

	// 后缀自增/自减返回原值
	return oldValue
}

// evalInfixExpression 处理中缀表达式节点
// 执行中缀运算符(如+、-、*、/、&&、||等)运算
//
// 参数:
//
//	infixExpression - 中缀表达式节点
//	env - 执行环境
//
// 返回值:
//
//	object.Object - 运算结果，发生错误时返回nil
//
// 特殊处理:
//
//   - 逻辑与(&&)和逻辑或(||)使用短路求值
//
// 错误处理:
//
//	若运算符不支持或操作数类型不匹配，设置OperationError并返回nil
func (e *Evaluator) evalInfixExpression(infixExpression *ast.InfixExpression, env *object.Environment) object.Object {
	left := e.Eval(infixExpression.Left, env)
	if e.Err != nil {
		return nil
	}
	// 逻辑与短路求值:若左操作数为false，直接返回false
	if infixExpression.Operator.Type == lexer.LOGICAL_AND {
		if leftValue, ok := left.(*object.Bool); ok {
			if !leftValue.Value {
				return &object.Bool{Value: false}
			}
		} else {
			e.Err = &object.OperationError{
				Frame:    e.Frame,
				Message:  "invalid operation \"&&\".",
				PosStart: infixExpression.PosStart,
				PosEnd:   infixExpression.PosEnd,
			}
			return nil
		}
	}
	// 逻辑或短路求值:若左操作数为true，直接返回true
	if infixExpression.Operator.Type == lexer.LOGICAL_OR {
		if leftValue, ok := left.(*object.Bool); ok {
			if leftValue.Value {
				return &object.Bool{Value: true}
			}
		} else {
			e.Err = &object.OperationError{
				Frame:    e.Frame,
				Message:  "invalid operation \"||\".",
				PosStart: infixExpression.PosStart,
				PosEnd:   infixExpression.PosEnd,
			}
			return nil
		}
	}
	// 计算右操作数并执行运算
	right := e.Eval(infixExpression.Right, env)
	if e.Err != nil {
		return nil
	}
	val := e.evalInfixOperator(infixExpression, left, right)
	if e.Err != nil {
		return nil
	}
	return val
}

func (e *Evaluator) evalInfixOperator(infixExpression *ast.InfixExpression, left, right object.Object) object.Object {
	switch infixExpression.Operator.Type {
	case lexer.PLUS:
		val, err := left.Add(right, infixExpression.PosStart, infixExpression.PosEnd, e.Frame)
		if err != nil {
			e.Err = err
			return nil
		}
		return val
	case lexer.MINUS:
		val, err := left.Subtract(right, infixExpression.PosStart, infixExpression.PosEnd, e.Frame)
		if err != nil {
			e.Err = err
			return nil
		}
		return val
	case lexer.ASTERISK:
		val, err := left.Multiply(right, infixExpression.PosStart, infixExpression.PosEnd, e.Frame)
		if err != nil {
			e.Err = err
			return nil
		}
		return val
	case lexer.SLASH:
		val, err := left.Divide(right, infixExpression.PosStart, infixExpression.PosEnd, e.Frame)
		if err != nil {
			e.Err = err
			return nil
		}
		return val
	case lexer.PERCENT:
		val, err := left.Mod(right, infixExpression.PosStart, infixExpression.PosEnd, e.Frame)
		if err != nil {
			e.Err = err
			return nil
		}
		return val
	case lexer.EQUALS:
		val, err := left.Equal(right, infixExpression.PosStart, infixExpression.PosEnd, e.Frame)
		if err != nil {
			e.Err = err
			return nil
		}
		return val
	case lexer.NOT_EQUALS:
		val, err := left.NotEqual(right, infixExpression.PosStart, infixExpression.PosEnd, e.Frame)
		if err != nil {
			e.Err = err
			return nil
		}
		return val
	case lexer.LT:
		val, err := left.LessThan(right, infixExpression.PosStart, infixExpression.PosEnd, e.Frame)
		if err != nil {
			e.Err = err
			return nil
		}
		return val
	case lexer.GT:
		val, err := left.GreaterThan(right, infixExpression.PosStart, infixExpression.PosEnd, e.Frame)
		if err != nil {
			e.Err = err
			return nil
		}
		return val
	case lexer.LTE:
		val, err := left.LessThanOrEqual(right, infixExpression.PosStart, infixExpression.PosEnd, e.Frame)
		if err != nil {
			e.Err = err
			return nil
		}
		return val
	case lexer.GTE:
		val, err := left.GreaterThanOrEqual(right, infixExpression.PosStart, infixExpression.PosEnd, e.Frame)
		if err != nil {
			e.Err = err
			return nil
		}
		return val
	case lexer.BITWISE_AND:
		val, err := left.BitAnd(right, infixExpression.PosStart, infixExpression.PosEnd, e.Frame)
		if err != nil {
			e.Err = err
			return nil
		}
		return val
	case lexer.BITWISE_OR:
		val, err := left.BitOr(right, infixExpression.PosStart, infixExpression.PosEnd, e.Frame)
		if err != nil {
			e.Err = err
			return nil
		}
		return val
	case lexer.BITWISE_XOR:
		val, err := left.Xor(right, infixExpression.PosStart, infixExpression.PosEnd, e.Frame)
		if err != nil {
			e.Err = err
			return nil
		}
		return val
	case lexer.LEFT_SHIFT:
		val, err := left.LeftShift(right, infixExpression.PosStart, infixExpression.PosEnd, e.Frame)
		if err != nil {
			e.Err = err
			return nil
		}
		return val
	case lexer.RIGHT_SHIFT:
		val, err := left.RightShift(right, infixExpression.PosStart, infixExpression.PosEnd, e.Frame)
		if err != nil {
			e.Err = err
			return nil
		}
		return val
	case lexer.LOGICAL_AND:
		val, err := left.And(right, infixExpression.PosStart, infixExpression.PosEnd, e.Frame)
		if err != nil {
			e.Err = err
			return nil
		}
		return val
	case lexer.LOGICAL_OR:
		val, err := left.Or(right, infixExpression.PosStart, infixExpression.PosEnd, e.Frame)
		if err != nil {
			e.Err = err
			return nil
		}
		return val
	default:
		e.Err = &object.OperationError{
			Message:  fmt.Sprintf("invalid operation \"%s\".", infixExpression.Operator.Type),
			PosStart: infixExpression.PosStart,
			PosEnd:   infixExpression.PosEnd,
		}
		return nil
	}
}

func (e *Evaluator) evalWithSpecialValue(node ast.Node, env *object.Environment) object.Object {
	var ret object.Object
	switch n := node.(type) {
	case *ast.ExpressionStatement:
		ret = e.Eval(n.Expr, env)
		if e.Err != nil {
			return nil
		}
	case *ast.ReturnStatement:
		ret = e.evalReturnStatement(n, env)
		if e.Err != nil {
			return nil
		}
		return ret
	case *ast.BreakStatement:
		ret = e.evalBreakStatement(n, env)
		if e.Err != nil {
			return nil
		}
		return ret
	case *ast.ContinueStatement:
		ret = e.evalContinueStatement(n, env)
		if e.Err != nil {
			return nil
		}
		return ret
	case ast.Statement:
		ret = e.Eval(n, env)
		if e.Err != nil {
			return nil
		}
		if returnValue, ok := ret.(*object.ReturnValue); ok {
			return returnValue
		}
		ret = &object.Null{}
	case ast.Expression:
		ret = e.Eval(n, env)
		if e.Err != nil {
			return nil
		}
	}
	return ret
}

// evalBlockExpression 处理块表达式节点
// 解释块表达式
//
// 参数:
//
//	blockExpression - 块表达式节点
//	env - 执行环境
//
// 返回值:
//
//	object.Object - 块表达式的结果，发生错误时返回nil
func (e *Evaluator) evalBlockExpression(blockExpression *ast.BlockExpression, env *object.Environment) object.Object {
	var ret object.Object
	// 创建新环境
	blockEnv := &object.Environment{
		Name:  "block",
		Store: make(map[string]*object.Symbol),
		Outer: env,
	}
	for _, statement := range blockExpression.Statements {
		// 获取返回值
		ret = e.evalWithSpecialValue(statement, blockEnv)
		if returnValue, ok := ret.(*object.ReturnValue); ok {
			return returnValue
		}
		if breakValue, ok := ret.(*object.BreakValue); ok {
			return breakValue
		}
		if continueValue, ok := ret.(*object.ContinueValue); ok {
			return continueValue
		}
	}
	return ret
}

// evalIfExpression 处理if表达式节点
// 解释if表达式
//
// 参数:
//
//	ifExpression - 块表达式节点
//	env - 执行环境
//
// 返回值:
//
//	object.Object - if表达式的结果，发生错误时返回nil
//
// 特殊处理：
//
// - 若条件为假且没有else分支，返回Null
func (e *Evaluator) evalIfExpression(ifExpression *ast.IfExpression, env *object.Environment) object.Object {
	condition := e.Eval(ifExpression.Condition, env)
	if e.Err != nil {
		return nil
	}
	if _, ok := condition.(*object.Bool); !ok {
		e.Err = &TypeError{
			Frame:    e.Frame,
			Message:  "non-bool condition in if expression.",
			PosStart: ifExpression.PosStart,
			PosEnd:   ifExpression.PosEnd,
		}
		return nil
	}
	// 创建新环境
	ifEnv := &object.Environment{
		Name:  "if",
		Store: make(map[string]*object.Symbol),
		Outer: env,
	}
	if condition.(*object.Bool).Value {
		return e.evalWithSpecialValue(ifExpression.Consequence, ifEnv)
	} else if ifExpression.Alternative != nil {
		return e.evalWithSpecialValue(ifExpression.Alternative, ifEnv)
	} else {
		return &object.Null{}
	}
}

// evalCallExpression 处理函数调用表达式节点
// 解释函数调用表达式
//
// 参数:
//
//	callExpression - 函数调用表达式节点
//	env - 执行环境
//
// 返回值:
//
//	object.Object - 函数表达式的结果，发生错误时返回nil
func (e *Evaluator) evalCallExpression(callExpression *ast.CallExpression, env *object.Environment) object.Object {
	function := e.Eval(callExpression.Function, env)
	if e.Err != nil {
		return nil
	}
	switch fn := function.(type) {
	// 函数
	case *object.Function:
		// 检查是否有可变参数
		hasVariadic := false
		for _, param := range fn.Parameter {
			if param.IsVariadic {
				hasVariadic = true
				break
			}
		}

		// 计算默认参数数量
		defaultLen := 0
		for _, param := range fn.Parameter {
			if param.DefaultValue != nil {
				defaultLen++
			}
		}
		// 计算传入参数数量
		argLen := 0
		for _, arg := range callExpression.Argument {
			if arg != nil {
				argLen++
			}
		}

		// 参数数量检查
		least := len(fn.Parameter) - defaultLen
		max := len(fn.Parameter)

		// 如果有可变参数，调整参数数量检查逻辑
		if hasVariadic {
			// 最小参数数 = 总参数数 - 默认参数数 - 1（减去可变参数）
			least--
			// 最大参数数没有限制
			max = -1
		}

		// 参数数量不匹配
		if max == -1 {
			// 有可变参数，只检查最小参数数
			if argLen < least {
				e.Err = &ArgumentError{
					Frame:    e.Frame,
					Message:  fmt.Sprintf("expected at least %d parameters, got %d.", least, argLen),
					PosStart: callExpression.PosStart,
					PosEnd:   callExpression.PosEnd,
				}
				return nil
			}
		} else {
			// 没有可变参数，检查参数数量是否在范围内
			if !(least <= argLen && argLen <= max) {
				if defaultLen == 0 {
					e.Err = &ArgumentError{
						Frame:    e.Frame,
						Message:  fmt.Sprintf("expected %d parameters, got %d.", len(fn.Parameter), argLen),
						PosStart: callExpression.PosStart,
						PosEnd:   callExpression.PosEnd,
					}
				} else if least == 1 {
					e.Err = &ArgumentError{
						Frame:    e.Frame,
						Message:  fmt.Sprintf("expected between 1 parameter and %d parameters, got %d.", len(fn.Parameter), argLen),
						PosStart: callExpression.PosStart,
						PosEnd:   callExpression.PosEnd,
					}
				} else {
					e.Err = &ArgumentError{
						Frame:    e.Frame,
						Message:  fmt.Sprintf("expected between %d and %d parameters, got %d.", least, len(fn.Parameter), argLen),
						PosStart: callExpression.PosStart,
						PosEnd:   callExpression.PosEnd,
					}
				}
				return nil
			}
		}

		// 计算实际参数值
		var argument []object.Object
		for _, arg := range callExpression.Argument {
			if arg == nil {
				// 如果参数为nil，用默认值填充
				if len(argument) < len(fn.Parameter) {
					defaultValue := e.Eval(fn.Parameter[len(argument)].DefaultValue, env)
					if e.Err != nil {
						return nil
					}
					argument = append(argument, defaultValue)
				}
				continue
			}
			a := e.Eval(arg, env)
			if e.Err != nil {
				return nil
			}
			argument = append(argument, a)
		}

		// 有默认参数未被赋值时，用默认值填充
		for i := len(argument); i < len(fn.Parameter); i++ {
			// 如果是可变参数，跳过默认值填充
			if fn.Parameter[i].IsVariadic {
				break
			}
			defaultValue := e.Eval(fn.Parameter[i].DefaultValue, env)
			if e.Err != nil {
				return nil
			}
			argument = append(argument, defaultValue)
		}

		// 创建函数环境
		funcEnv := &object.Environment{
			Name:  "function",
			Store: make(map[string]*object.Symbol),
			Outer: fn.Env,
		}
		e.Frame = &frame.Frame{
			FuncName: fmt.Sprintf("<function \"%s\">", fn.Name),
			Parent:   e.Frame,
			PosStart: callExpression.PosStart,
			PosEnd:   callExpression.PosEnd,
		}

		// 绑定参数到函数环境
		for i, param := range fn.Parameter {
			if param.IsVariadic {
				// 处理可变参数：收集剩余的所有参数到一个列表中
				variadicArgs := make([]object.Object, 0)
				for j := i; j < len(argument); j++ {
					if j != i && argument[j].Type() != variadicArgs[0].Type() {
						e.Err = &TypeError{
							Frame:    e.Frame,
							Message:  "all variadic arguments must be of the same type.",
							PosStart: callExpression.PosStart,
							PosEnd:   callExpression.PosEnd,
						}
						return nil
					}
					variadicArgs = append(variadicArgs, argument[j])
				}
				// 创建列表对象
				listObj := &object.List{
					Elements: variadicArgs,
				}
				// 绑定可变参数
				funcEnv.Set(param.Name.Name, &object.Symbol{
					Name:    param.Name.Name,
					Value:   listObj,
					IsConst: false,
				})
				break
			} else if i < len(argument) {
				// 普通参数
				funcEnv.Set(param.Name.Name, &object.Symbol{
					Name:    param.Name.Name,
					Value:   argument[i],
					IsConst: false,
				})
			}
		}

		// 执行函数体
		var returnValue = e.evalWithSpecialValue(fn.Body, funcEnv)
		if e.Err != nil {
			return nil
		}
		e.Frame = e.Frame.Parent
		if ret, ok := returnValue.(*object.ReturnValue); ok {
			return ret.Value
		} else {
			return returnValue
		}
	// 内置函数
	case *object.BuiltinFunction:
		// 检查是否有可变参数
		hasVariadic := fn.HaveVariadic

		// 计算默认参数数量
		defaultLen := 0
		for _, defaultValue := range fn.DefaultValue {
			if defaultValue != nil {
				defaultLen++
			}
		}
		// 计算传入参数数量
		argLen := 0
		for _, arg := range callExpression.Argument {
			if arg != nil {
				argLen++
			}
		}

		// 参数数量检查
		least := len(fn.Parameter) - defaultLen
		max := len(fn.Parameter)

		// 如果有可变参数，调整参数数量检查逻辑
		if hasVariadic {
			// 最小参数数 = 总参数数 - 默认参数数 - 1（减去可变参数）
			least--
			// 最大参数数没有限制
			max = -1
		}

		// 参数数量不匹配
		if max == -1 {
			// 有可变参数，只检查最小参数数
			if argLen < least {
				e.Err = &ArgumentError{
					Frame:    e.Frame,
					Message:  fmt.Sprintf("expected at least %d parameters, got %d.", least, argLen),
					PosStart: callExpression.PosStart,
					PosEnd:   callExpression.PosEnd,
				}
				return nil
			}
		} else {
			// 没有可变参数，检查参数数量是否在范围内
			if !(least <= argLen && argLen <= max) {
				if defaultLen == 0 {
					e.Err = &ArgumentError{
						Frame:    e.Frame,
						Message:  fmt.Sprintf("expected %d parameters, got %d.", len(fn.Parameter), argLen),
						PosStart: callExpression.PosStart,
						PosEnd:   callExpression.PosEnd,
					}
				} else if least == 1 {
					e.Err = &ArgumentError{
						Frame:    e.Frame,
						Message:  fmt.Sprintf("expected between 1 parameter and %d parameters, got %d.", len(fn.Parameter), argLen),
						PosStart: callExpression.PosStart,
						PosEnd:   callExpression.PosEnd,
					}
				} else {
					e.Err = &ArgumentError{
						Frame:    e.Frame,
						Message:  fmt.Sprintf("expected between %d and %d parameters, got %d.", least, len(fn.Parameter), argLen),
						PosStart: callExpression.PosStart,
						PosEnd:   callExpression.PosEnd,
					}
				}
				return nil
			}
		}

		// 评估参数表达式
		var argument []object.Object
		for _, arg := range callExpression.Argument {
			// 如果参数为nil，用默认值填充
			if arg == nil {
				if len(argument) < len(fn.Parameter) {
					defaultValue := e.Eval(fn.DefaultValue[len(argument)], env)
					if e.Err != nil {
						return nil
					}
					argument = append(argument, defaultValue)
				}
				continue
			}
			a := e.Eval(arg, env)
			if e.Err != nil {
				return nil
			}
			argument = append(argument, a)
		}

		// 有默认参数未被赋值时，用默认值填充
		for i := len(argument); i < len(fn.Parameter); i++ {
			// 如果是可变参数，跳过默认值填充
			if hasVariadic && i == len(fn.Parameter)-1 {
				break
			}
			defaultValue := e.Eval(fn.DefaultValue[i], env)
			if e.Err != nil {
				return nil
			}
			argument = append(argument, defaultValue)
		}
		// 判断可变参数传入的类型是否相同
		if hasVariadic {
			for i := len(fn.Parameter) - 1; i < len(argument); i++ {
				if i != len(fn.Parameter)-1 && argument[i].Type() != argument[i-1].Type() {
					e.Err = &TypeError{
						Frame:    e.Frame,
						Message:  "all variadic arguments must be of the same type.",
						PosStart: callExpression.PosStart,
						PosEnd:   callExpression.PosEnd,
					}
					return nil
				}
			}
		}

		// 调用内置函数
		e.Frame = &frame.Frame{
			FuncName: fmt.Sprintf("<builtin \"%s\">", fn.Name),
			Parent:   e.Frame,
			PosStart: callExpression.PosStart,
			PosEnd:   callExpression.PosEnd,
		}
		val, err := fn.Fn(e.Frame, callExpression.PosStart, callExpression.PosEnd, argument...)
		if err != nil {
			e.Err = err
			return nil
		}
		e.Frame = e.Frame.Parent
		return val
	default:
		// 调用非函数
		e.Err = &TypeError{
			Frame:    e.Frame,
			Message:  "the value is not a function and cannot be called.",
			PosStart: callExpression.PosStart,
			PosEnd:   callExpression.PosEnd,
		}
		return nil
	}
}
