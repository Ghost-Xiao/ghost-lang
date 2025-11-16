package object

// Environment 表示程序运行时的上下文环境，用于管理符号表和上下文嵌套关系
// 在函数调用、作用域切换等场景中使用，实现变量的作用域隔离和查找

type Environment struct {
	Store map[string]*Symbol // 变量名到值的映射
	Outer *Environment       // 外部环境
}

// Get 查找符号的值，支持作用域链向上查找
// 先在当前环境中查找，若不存在且存在父环境，则递归查找父环境
//
// 参数:
//
//	name - 要查找的符号名称
//
// 返回值:
//
//	Symbol - 符号，若未找到则为nil
//	bool - 查找结果，true表示找到，false表示未找到
func (e *Environment) Get(name string) (*Symbol, bool) {
	val, ok := e.Store[name]
	if ok {
		return val, ok
	}
	// 若当前表未找到，尝试在父环境中查找
	if e.Outer != nil {
		return e.Outer.Get(name)
	}
	return nil, false
}

// Set 设置符号的值到当前环境
// 仅在当前作用域中添加或修改变量，不影响父环境
//
// 参数:
//
//	name - 要设置的符号名称
//	sym - 符号
func (e *Environment) Set(name string, sym *Symbol) {
	e.Store[name] = sym
}

// Assign 设置符号的值到当前环境
// 沿作用域链查找并设置符号的值
// 若当前作用域未定义该符号，则递归查找父作用域，直到找到或到达全局作用域
//
// 参数:
//
//	name - 要设置的符号名称
//	sym - 符号
func (e *Environment) Assign(name string, sym *Symbol) {
	// 先在当前作用域查找
	if _, ok := e.Store[name]; ok {
		e.Store[name] = sym
	} else {
		// 若当前作用域未定义，递归查找父作用域
		if e.Outer != nil {
			e.Outer.Assign(name, sym)
		}
	}
}

// Exists 检查符号是否存在于当前环境（不包含父环境）
// 仅判断当前作用域中是否已定义该符号，不进行作用域链查找
//
// 参数:
//
//	name - 要检查的符号名称
//
// 返回值:
//
//	bool - 存在性结果，true表示存在，false表示不存在
func (e *Environment) Exists(name string) bool {
	_, ok := e.Store[name]
	return ok
}
