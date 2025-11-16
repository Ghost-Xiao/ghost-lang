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
	case *ast.ExpressionStatement:
		return e.evalExpressionStatement(n, env)
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
		ret := e.Eval(forStatement.Body, forEnv)
		if e.Err != nil {
			return nil
		}
		if returnValue, ok := ret.(*object.ReturnValue); ok {
			return returnValue
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
	switch varAssignment.Name.(type) {
	case *ast.IdentifierExpression:
		varName := varAssignment.Name.(*ast.IdentifierExpression).Name
		// 检查变量是否已定义
		sym, ok := env.Get(varName)
		if !ok {
			e.Err = &VariableError{
				Frame:    e.Frame,
				Message:  fmt.Sprintf("undefined variable \"%s\".", varName),
				PosStart: varAssignment.PosStart,
				PosEnd:   varAssignment.PosEnd,
			}
			return nil
		}
		// 检查是否是常量
		if sym.IsConst {
			e.Err = &VariableError{
				Frame:    e.Frame,
				Message:  fmt.Sprintf("cannot redefine constant \"%s\".", varName),
				PosStart: varAssignment.PosStart,
				PosEnd:   varAssignment.PosEnd,
			}
			return nil
		}
		value := e.Eval(varAssignment.Value, env)
		if e.Err != nil {
			return nil
		}
		newSym := &object.Symbol{
			Name:    varName,
			Value:   value,
			IsConst: false,
		}
		env.Assign(varName, newSym)
		return value
	case *ast.IndexExpression:
		indexExpr := varAssignment.Name.(*ast.IndexExpression)
		err := e.checkIndexTargetConst(indexExpr.Target, env, indexExpr.PosStart, indexExpr.PosEnd)
		if err != nil {
			e.Err = err
			return nil
		}
		target := e.Eval(indexExpr.Target, env)
		if e.Err != nil {
			return nil
		}
		index := e.Eval(indexExpr.Index, env)
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
		value := e.Eval(varAssignment.Value, env)
		if e.Err != nil {
			return nil
		}
		err2 := idxable.Set(index, value, varAssignment.PosStart, varAssignment.PosEnd, e.Frame)
		if err2 != nil {
			e.Err = err2
			return nil
		}
		// 返回新值
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
	switch compoundAssignmentExpression.Name.(type) {
	case *ast.IdentifierExpression:
		// 获取变量名
		varName := compoundAssignmentExpression.Name.(*ast.IdentifierExpression).Name
		// 检查变量是否已定义
		sym, ok := env.Get(varName)
		if !ok {
			e.Err = &VariableError{
				Frame:    e.Frame,
				Message:  fmt.Sprintf("undefined variable \"%s\".", varName),
				PosStart: compoundAssignmentExpression.PosStart,
				PosEnd:   compoundAssignmentExpression.PosEnd,
			}
			return nil
		}
		// 检查是否是常量
		if sym.IsConst {
			e.Err = &VariableError{
				Frame:    e.Frame,
				Message:  fmt.Sprintf("cannot redefine constant \"%s\".", varName),
				PosStart: compoundAssignmentExpression.PosStart,
				PosEnd:   compoundAssignmentExpression.PosEnd,
			}
			return nil
		}
		// 计算右侧表达式
		right := e.Eval(compoundAssignmentExpression.Right, env)
		if e.Err != nil {
			return nil
		}
		// 获取运算符字面量
		literal := compoundAssignmentExpression.Operator.Literal[:len(compoundAssignmentExpression.Operator.Literal)-1]
		// 获取并创建基础运算符令牌
		baseOperator := &lexer.Token{
			Type:    lexer.CompoundAssignmentOperators[compoundAssignmentExpression.Operator.Type],
			Literal: literal,
		}
		// 执行复合赋值
		value := e.evalInfixOperator(&ast.InfixExpression{
			Left:     compoundAssignmentExpression.Name,
			Operator: baseOperator,
			Right:    compoundAssignmentExpression.Right,
			PosStart: compoundAssignmentExpression.PosStart,
			PosEnd:   compoundAssignmentExpression.PosEnd,
		}, sym.Value, right)
		if e.Err != nil {
			return nil
		}
		// 构建新符号
		newSym := &object.Symbol{
			Name:    varName,
			Value:   value,
			IsConst: false,
		}
		env.Assign(varName, newSym)
		return value
	case *ast.IndexExpression:
		indexExpr := compoundAssignmentExpression.Name.(*ast.IndexExpression)
		err := e.checkIndexTargetConst(indexExpr.Target, env, indexExpr.PosStart, indexExpr.PosEnd)
		if err != nil {
			e.Err = err
			return nil
		}
		target := e.Eval(indexExpr.Target, env)
		if e.Err != nil {
			return nil
		}
		index := e.Eval(indexExpr.Index, env)
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
		// 计算右侧表达式
		right := e.Eval(compoundAssignmentExpression.Right, env)
		if e.Err != nil {
			return nil
		}
		// 获取运算符字面量
		literal := compoundAssignmentExpression.Operator.Literal[:len(compoundAssignmentExpression.Operator.Literal)-1]
		// 获取并创建基础运算符令牌
		baseOperator := &lexer.Token{
			Type:    lexer.CompoundAssignmentOperators[compoundAssignmentExpression.Operator.Type],
			Literal: literal,
		}
		// 获取目标索引的值
		idxValue := e.Eval(compoundAssignmentExpression.Name, env)
		if e.Err != nil {
			return nil
		}
		// 执行复合赋值
		value := e.evalInfixOperator(&ast.InfixExpression{
			Left:     compoundAssignmentExpression.Name,
			Operator: baseOperator,
			Right:    compoundAssignmentExpression.Right,
			PosStart: compoundAssignmentExpression.PosStart,
			PosEnd:   compoundAssignmentExpression.PosEnd,
		}, idxValue, right)
		if e.Err != nil {
			return nil
		}
		err2 := idxable.Set(index, value, compoundAssignmentExpression.PosStart, compoundAssignmentExpression.PosEnd, e.Frame)
		if err2 != nil {
			e.Err = err2
			return nil
		}
		// 返回新值
		return value
	default:
		e.Err = &TypeError{
			Frame:    e.Frame,
			Message:  "invalid variable name type.",
			PosStart: compoundAssignmentExpression.PosStart,
			PosEnd:   compoundAssignmentExpression.PosEnd,
		}
		return nil
	}
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
	switch prefixUnaryIncDecExpression.Right.(type) {
	case *ast.IdentifierExpression:
		name := prefixUnaryIncDecExpression.Right.(*ast.IdentifierExpression).Name
		sym, ok := env.Get(name)
		if !ok {
			e.Err = &VariableError{
				Frame:    e.Frame,
				Message:  fmt.Sprintf("undefined variable \"%s\".", name),
				PosStart: prefixUnaryIncDecExpression.PosStart,
				PosEnd:   prefixUnaryIncDecExpression.PosEnd,
			}
			return nil
		}
		// 检查是否是常量
		if sym.IsConst {
			e.Err = &VariableError{
				Frame:    e.Frame,
				Message:  fmt.Sprintf("cannot redefine constant \"%s\".", name),
				PosStart: prefixUnaryIncDecExpression.PosStart,
				PosEnd:   prefixUnaryIncDecExpression.PosEnd,
			}
			return nil
		}
		right := e.Eval(prefixUnaryIncDecExpression.Right, env)
		if e.Err != nil {
			return nil
		}
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
		}, right, &object.Int{Value: 1})
		if e.Err != nil {
			return nil
		}
		// 构建新符号
		newSym := &object.Symbol{
			Name:    name,
			Value:   val,
			IsConst: false,
		}
		// 更新变量值
		env.Set(name, newSym)
		return val
	case *ast.IndexExpression:
		indexExpr := prefixUnaryIncDecExpression.Right.(*ast.IndexExpression)
		err := e.checkIndexTargetConst(indexExpr.Target, env, indexExpr.PosStart, indexExpr.PosEnd)
		if err != nil {
			e.Err = err
			return nil
		}
		target := e.Eval(indexExpr.Target, env)
		if e.Err != nil {
			return nil
		}
		index := e.Eval(indexExpr.Index, env)
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
		right := e.Eval(indexExpr, env)
		if e.Err != nil {
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
		}, right, &object.Int{Value: 1})
		if e.Err != nil {
			return nil
		}
		err2 := idxable.Set(index, val, prefixUnaryIncDecExpression.PosStart, prefixUnaryIncDecExpression.PosEnd, e.Frame)
		if err2 != nil {
			e.Err = err2
			return nil
		}
		// 返回新值
		return val
	default:
		e.Err = &TypeError{
			Frame:    e.Frame,
			Message:  "invalid variable name type.",
			PosStart: prefixUnaryIncDecExpression.PosStart,
			PosEnd:   prefixUnaryIncDecExpression.PosEnd,
		}
		return nil
	}
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
	switch postfixUnaryIncDecExpression.Left.(type) {
	case *ast.IdentifierExpression:
		name := postfixUnaryIncDecExpression.Left.(*ast.IdentifierExpression).Name
		sym, ok := env.Get(name)
		if !ok {
			e.Err = &VariableError{
				Frame:    e.Frame,
				Message:  fmt.Sprintf("undefined variable \"%s\".", name),
				PosStart: postfixUnaryIncDecExpression.PosStart,
				PosEnd:   postfixUnaryIncDecExpression.PosEnd,
			}
			return nil
		}
		// 检查是否是常量
		if sym.IsConst {
			e.Err = &VariableError{
				Frame:    e.Frame,
				Message:  fmt.Sprintf("cannot redefine constant \"%s\".", name),
				PosStart: postfixUnaryIncDecExpression.PosStart,
				PosEnd:   postfixUnaryIncDecExpression.PosEnd,
			}
			return nil
		}
		left := e.Eval(postfixUnaryIncDecExpression.Left, env)
		if e.Err != nil {
			return nil
		}
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
		}, left, &object.Int{Value: 1})
		if e.Err != nil {
			return nil
		}
		// 构建新符号
		newSym := &object.Symbol{
			Name:    name,
			Value:   val,
			IsConst: false,
		}
		// 更新变量值
		env.Set(name, newSym)
		return left
	case *ast.IndexExpression:
		indexExpr := postfixUnaryIncDecExpression.Left.(*ast.IndexExpression)
		err := e.checkIndexTargetConst(indexExpr.Target, env, indexExpr.PosStart, indexExpr.PosEnd)
		if err != nil {
			e.Err = err
			return nil
		}
		target := e.Eval(indexExpr.Target, env)
		if e.Err != nil {
			return nil
		}
		index := e.Eval(indexExpr.Index, env)
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
		// 获取索引值
		left := e.evalIndexExpression(indexExpr, env)
		if e.Err != nil {
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
		}, left, &object.Int{Value: 1})
		if e.Err != nil {
			return nil
		}
		err2 := idxable.Set(index, val, postfixUnaryIncDecExpression.PosStart, postfixUnaryIncDecExpression.PosEnd, e.Frame)
		if err2 != nil {
			e.Err = err2
			return nil
		}
		return left
	default:
		e.Err = &TypeError{
			Frame:    e.Frame,
			Message:  "invalid variable name type.",
			PosStart: postfixUnaryIncDecExpression.PosStart,
			PosEnd:   postfixUnaryIncDecExpression.PosEnd,
		}
		return nil
	}
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

func (e *Evaluator) evalWithReturnValue(node ast.Node, env *object.Environment) object.Object {
	var ret object.Object
	switch n := node.(type) {
	case *ast.ExpressionStatement:
		ret = e.Eval(n.Expr, env)
		if e.Err != nil {
			return nil
		}
	case *ast.ReturnStatement:
		ret = e.Eval(n.ReturnValue, env)
		if e.Err != nil {
			return nil
		}
		return &object.ReturnValue{Value: ret}
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
		Store: make(map[string]*object.Symbol),
		Outer: env,
	}
	for _, statement := range blockExpression.Statements {
		// 获取返回值
		ret = e.evalWithReturnValue(statement, blockEnv)
		if returnValue, ok := ret.(*object.ReturnValue); ok {
			return returnValue
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
		Store: make(map[string]*object.Symbol),
		Outer: env,
	}
	if condition.(*object.Bool).Value {
		return e.evalWithReturnValue(ifExpression.Consequence, ifEnv)
	} else if ifExpression.Alternative != nil {
		return e.evalWithReturnValue(ifExpression.Alternative, ifEnv)
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
		// 参数数量不匹配
		least := len(fn.Parameter) - defaultLen
		if !(least <= argLen && argLen <= len(fn.Parameter)) {
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
		var argument []object.Object
		for _, arg := range callExpression.Argument {
			// 如果参数为nil，用默认值填充
			if arg == nil {
				defaultValue := e.Eval(fn.Parameter[len(argument)].DefaultValue, env)
				if e.Err != nil {
					return nil
				}
				argument = append(argument, defaultValue)
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
			defaultValue := e.Eval(fn.Parameter[i].DefaultValue, env)
			if e.Err != nil {
				return nil
			}
			argument = append(argument, defaultValue)
		}
		// 创建函数环境
		funcEnv := &object.Environment{
			Store: make(map[string]*object.Symbol),
			Outer: fn.Env,
		}
		e.Frame = &frame.Frame{
			FuncName: fmt.Sprintf("<function \"%s\">", fn.Name),
			Parent:   e.Frame,
			PosStart: callExpression.PosStart,
			PosEnd:   callExpression.PosEnd,
		}
		// 创建参数
		for i, param := range fn.Parameter {
			funcEnv.Set(param.Name.Name, &object.Symbol{
				Name:    param.Name.Name,
				Value:   argument[i],
				IsConst: false,
			})
		}
		// 执行函数体
		var returnValue = e.evalWithReturnValue(fn.Body, funcEnv)
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
		// 参数数量不匹配
		least := len(fn.Parameter) - defaultLen
		if !(least <= argLen && argLen <= len(fn.Parameter)) {
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
		// 调用内置函数
		var argument []object.Object
		for _, arg := range callExpression.Argument {
			// 如果参数为nil，用默认值填充
			if arg == nil {
				defaultValue := e.Eval(fn.DefaultValue[len(argument)], env)
				if e.Err != nil {
					return nil
				}
				argument = append(argument, defaultValue)
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
			defaultValue := e.Eval(fn.DefaultValue[i], env)
			if e.Err != nil {
				return nil
			}
			argument = append(argument, defaultValue)
		}
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
