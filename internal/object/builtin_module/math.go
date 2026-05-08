package builtin_module

import (
	"math"
	"math/rand"

	"github.com/Ghost-Xiao/ghost-lang/internal/errors"
	"github.com/Ghost-Xiao/ghost-lang/internal/frame"
	"github.com/Ghost-Xiao/ghost-lang/internal/object"
	"github.com/Ghost-Xiao/ghost-lang/internal/util"
)

// MathModule 表示 math 内置模块
var MathModule = initMathModule()

// initMathModule 初始化 math 模块
//
// 返回值:
//
//	*object.Module - 初始化后的 math 模块
func initMathModule() *object.Module {
	env := &object.Environment{
		Name:  "math",
		Store: map[string]*object.Symbol{},
		Outer: nil,
	}

	env.Set("PI", &object.Symbol{Name: "PI", Value: &PI, IsConst: true})
	env.Set("E", &object.Symbol{Name: "E", Value: &E, IsConst: true})
	env.Set("TAU", &object.Symbol{Name: "TAU", Value: &TAU, IsConst: true})
	env.Set("abs", &object.Symbol{Name: "abs", Value: &ABS, IsConst: true})
	env.Set("sqrt", &object.Symbol{Name: "sqrt", Value: &SQRT, IsConst: true})
	env.Set("sin", &object.Symbol{Name: "sin", Value: &SIN, IsConst: true})
	env.Set("cos", &object.Symbol{Name: "cos", Value: &COS, IsConst: true})
	env.Set("tan", &object.Symbol{Name: "tan", Value: &TAN, IsConst: true})
	env.Set("asin", &object.Symbol{Name: "asin", Value: &ASIN, IsConst: true})
	env.Set("acos", &object.Symbol{Name: "acos", Value: &ACOS, IsConst: true})
	env.Set("atan", &object.Symbol{Name: "atan", Value: &ATAN, IsConst: true})
	env.Set("log", &object.Symbol{Name: "log", Value: &LOG, IsConst: true})
	env.Set("lg", &object.Symbol{Name: "lg", Value: &LG, IsConst: true})
	env.Set("ln", &object.Symbol{Name: "ln", Value: &LN, IsConst: true})
	env.Set("floor", &object.Symbol{Name: "floor", Value: &FLOOR, IsConst: true})
	env.Set("ceil", &object.Symbol{Name: "ceil", Value: &CEIL, IsConst: true})
	env.Set("round", &object.Symbol{Name: "round", Value: &ROUND, IsConst: true})
	env.Set("min", &object.Symbol{Name: "min", Value: &MIN, IsConst: true})
	env.Set("max", &object.Symbol{Name: "max", Value: &MAX, IsConst: true})
	env.Set("sum", &object.Symbol{Name: "sum", Value: &SUM, IsConst: true})
	env.Set("product", &object.Symbol{Name: "product", Value: &PRODUCT, IsConst: true})
	env.Set("mean", &object.Symbol{Name: "mean", Value: &MEAN, IsConst: true})
	env.Set("median", &object.Symbol{Name: "median", Value: &MEDIAN, IsConst: true})
	env.Set("variance", &object.Symbol{Name: "variance", Value: &VARIANCE, IsConst: true})
	env.Set("stdDev", &object.Symbol{Name: "stdDev", Value: &STDDEV, IsConst: true})
	env.Set("rand", &object.Symbol{Name: "rand", Value: &RAND, IsConst: true})
	env.Set("randInt", &object.Symbol{Name: "randInt", Value: &RANDINT, IsConst: true})

	return &object.Module{
		Name: "math",
		Env:  env,
	}
}

var (
	// PI 圆周率常量
	PI = object.Float{Value: 3.141592653589793}
	// E 自然对数的底数
	E = object.Float{Value: 2.718281828459045}
	// TAU 2倍圆周率
	TAU = object.Float{Value: 6.283185307179586}
	// ABS 绝对值函数
	ABS = object.BuiltinFunction{
		Name:         "abs",
		Parameter:    []string{"x"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           MathAbs,
	}
	// SQRT 平方根函数
	SQRT = object.BuiltinFunction{
		Name:         "sqrt",
		Parameter:    []string{"x"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           MathSqrt,
	}
	// SIN 正弦函数
	SIN = object.BuiltinFunction{
		Name:         "sin",
		Parameter:    []string{"x"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           MathSin,
	}
	// COS 余弦函数
	COS = object.BuiltinFunction{
		Name:         "cos",
		Parameter:    []string{"x"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           MathCos,
	}
	// TAN 正切函数
	TAN = object.BuiltinFunction{
		Name:         "tan",
		Parameter:    []string{"x"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           MathTan,
	}
	// ASIN 反正弦函数
	ASIN = object.BuiltinFunction{
		Name:         "asin",
		Parameter:    []string{"x"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           MathAsin,
	}
	// ACOS 反余弦函数
	ACOS = object.BuiltinFunction{
		Name:         "acos",
		Parameter:    []string{"x"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           MathAcos,
	}
	// ATAN 反正切函数
	ATAN = object.BuiltinFunction{
		Name:         "atan",
		Parameter:    []string{"x"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           MathAtan,
	}
	// LOG 对数函数
	LOG = object.BuiltinFunction{
		Name:         "log",
		Parameter:    []string{"a", "x"},
		DefaultValue: []object.Object{nil, nil},
		HaveVariadic: false,
		Fn:           MathLog,
	}
	// LG 常用对数函数（以10为底）
	LG = object.BuiltinFunction{
		Name:         "lg",
		Parameter:    []string{"x"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           MathLg,
	}
	// LN 自然对数函数（以e为底）
	LN = object.BuiltinFunction{
		Name:         "ln",
		Parameter:    []string{"x"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           MathLn,
	}
	// FLOOR 向下取整函数
	FLOOR = object.BuiltinFunction{
		Name:         "floor",
		Parameter:    []string{"x", "decimalPlaces"},
		DefaultValue: []object.Object{nil, &object.Int{Value: 0}},
		HaveVariadic: false,
		Fn:           MathFloor,
	}
	// CEIL 向上取整函数
	CEIL = object.BuiltinFunction{
		Name:         "ceil",
		Parameter:    []string{"x", "decimalPlaces"},
		DefaultValue: []object.Object{nil, &object.Int{Value: 0}},
		HaveVariadic: false,
		Fn:           MathCeil,
	}
	// ROUND 四舍五入函数
	ROUND = object.BuiltinFunction{
		Name:         "round",
		Parameter:    []string{"x", "decimalPlaces"},
		DefaultValue: []object.Object{nil, &object.Int{Value: 0}},
		HaveVariadic: false,
		Fn:           MathRound,
	}
	// MIN 最小值函数
	MIN = object.BuiltinFunction{
		Name:         "min",
		Parameter:    []string{"a"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: true,
		Fn:           MathMin,
	}
	// MAX 最大值函数
	MAX = object.BuiltinFunction{
		Name:         "max",
		Parameter:    []string{"a"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: true,
		Fn:           MathMax,
	}
	// SUM 求和函数
	SUM = object.BuiltinFunction{
		Name:         "sum",
		Parameter:    []string{"a"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: true,
		Fn:           MathSum,
	}
	// PRODUCT 乘积函数
	PRODUCT = object.BuiltinFunction{
		Name:         "product",
		Parameter:    []string{"a"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: true,
		Fn:           MathProduct,
	}
	// MEAN 平均值函数
	MEAN = object.BuiltinFunction{
		Name:         "mean",
		Parameter:    []string{"a"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: true,
		Fn:           MathMean,
	}
	// MEDIAN 中位数函数
	MEDIAN = object.BuiltinFunction{
		Name:         "median",
		Parameter:    []string{"a"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: true,
		Fn:           MathMedian,
	}
	// VARIANCE 方差函数
	VARIANCE = object.BuiltinFunction{
		Name:         "variance",
		Parameter:    []string{"a"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: true,
		Fn:           MathVariance,
	}
	// STDDEV 标准差函数
	STDDEV = object.BuiltinFunction{
		Name:         "stdDev",
		Parameter:    []string{"a"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: true,
		Fn:           MathStdDev,
	}
	// RAND 随机数生成函数（0-1之间）
	RAND = object.BuiltinFunction{
		Name:         "rand",
		Parameter:    []string{},
		DefaultValue: []object.Object{},
		HaveVariadic: false,
		Fn:           MathRand,
	}
	// RANDINT 随机整数生成函数
	RANDINT = object.BuiltinFunction{
		Name:         "randInt",
		Parameter:    []string{"min", "max"},
		DefaultValue: []object.Object{nil, nil},
		HaveVariadic: false,
		Fn:           MathRandInt,
	}
)

// MathAbs 实现绝对值函数
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 函数参数
//
// 返回值:
//
//	object.Object - 绝对值结果
//	error - 可能出现的错误
func MathAbs(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	x := args[0]
	switch num := x.(type) {
	case *object.Int:
		if num.Value < 0 {
			return &object.Int{Value: -num.Value}, nil
		}
		return num, nil
	case *object.Float:
		if num.Value < 0 {
			return &object.Float{Value: -num.Value}, nil
		}
		return num, nil
	}
	return nil, &errors.TypeError{
		Frame:    f,
		Message:  "invalid type for abs function.",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// MathSqrt 实现平方根函数
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 函数参数
//
// 返回值:
//
//	object.Object - 平方根结果
//	error - 可能出现的错误
func MathSqrt(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	x := args[0]
	switch num := x.(type) {
	case *object.Int:
		if num.Value < 0 {
			return nil, &errors.OperationError{
				Frame:    f,
				Message:  "sqrt of negative number.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
		return &object.Float{Value: math.Sqrt(float64(num.Value))}, nil
	case *object.Float:
		if num.Value < 0 {
			return nil, &errors.OperationError{
				Frame:    f,
				Message:  "sqrt of negative number.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
		return &object.Float{Value: math.Sqrt(num.Value)}, nil
	}
	return nil, &errors.TypeError{
		Frame:    f,
		Message:  "invalid type for sqrt function.",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// MathSin 实现正弦函数
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 函数参数
//
// 返回值:
//
//	object.Object - 正弦结果
//	error - 可能出现的错误
func MathSin(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	x := args[0]
	switch num := x.(type) {
	case *object.Int:
		return &object.Float{Value: math.Sin(float64(num.Value))}, nil
	case *object.Float:
		return &object.Float{Value: math.Sin(num.Value)}, nil
	}
	return nil, &errors.TypeError{
		Frame:    f,
		Message:  "invalid type for sin function.",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// MathCos 实现余弦函数
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 函数参数
//
// 返回值:
//
//	object.Object - 余弦结果
//	error - 可能出现的错误
func MathCos(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	x := args[0]
	switch num := x.(type) {
	case *object.Int:
		return &object.Float{Value: math.Cos(float64(num.Value))}, nil
	case *object.Float:
		return &object.Float{Value: math.Cos(num.Value)}, nil
	}
	return nil, &errors.TypeError{
		Frame:    f,
		Message:  "invalid type for cos function.",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// MathTan 实现正切函数
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 函数参数
//
// 返回值:
//
//	object.Object - 正切结果
//	error - 可能出现的错误
func MathTan(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	x := args[0]
	switch num := x.(type) {
	case *object.Int:
		return &object.Float{Value: math.Tan(float64(num.Value))}, nil
	case *object.Float:
		return &object.Float{Value: math.Tan(num.Value)}, nil
	}
	return nil, &errors.TypeError{
		Frame:    f,
		Message:  "invalid type for tan function.",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// MathAsin 实现反正弦函数
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 函数参数
//
// 返回值:
//
//	object.Object - 反正弦结果
//	error - 可能出现的错误
func MathAsin(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	x := args[0]
	switch num := x.(type) {
	case *object.Int:
		return &object.Float{Value: math.Asin(float64(num.Value))}, nil
	case *object.Float:
		return &object.Float{Value: math.Asin(num.Value)}, nil
	}
	return nil, &errors.TypeError{
		Frame:    f,
		Message:  "invalid type for asin function.",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// MathAcos 实现反余弦函数
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 函数参数
//
// 返回值:
//
//	object.Object - 反余弦结果
//	error - 可能出现的错误
func MathAcos(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	x := args[0]
	switch num := x.(type) {
	case *object.Int:
		return &object.Float{Value: math.Acos(float64(num.Value))}, nil
	case *object.Float:
		return &object.Float{Value: math.Acos(num.Value)}, nil
	}
	return nil, &errors.TypeError{
		Frame:    f,
		Message:  "invalid type for acos function.",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// MathAtan 实现反正切函数
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 函数参数
//
// 返回值:
//
//	object.Object - 反正切结果
//	error - 可能出现的错误
func MathAtan(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	x := args[0]
	switch num := x.(type) {
	case *object.Int:
		return &object.Float{Value: math.Atan(float64(num.Value))}, nil
	case *object.Float:
		return &object.Float{Value: math.Atan(num.Value)}, nil
	}
	return nil, &errors.TypeError{
		Frame:    f,
		Message:  "invalid type for atan function.",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// MathLog 实现对数函数
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 函数参数
//
// 返回值:
//
//	object.Object - 对数结果
//	error - 可能出现的错误
func MathLog(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	a := args[0]
	x := args[1]
	var base, num float64
	switch v := a.(type) {
	case *object.Int:
		base = float64(v.Value)
	case *object.Float:
		base = v.Value
	default:
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "invalid type for log base.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	switch v := x.(type) {
	case *object.Int:
		num = float64(v.Value)
	case *object.Float:
		num = v.Value
	default:
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "invalid type for log argument.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	if base <= 0 || base == 1 {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "log base must be > 0 and != 1.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	if num <= 0 {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "log argument must be > 0.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	return &object.Float{Value: math.Log(num) / math.Log(base)}, nil
}

// MathLg 实现常用对数函数(以10为底)
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 函数参数
//
// 返回值:
//
//	object.Object - 常用对数结果
//	error - 可能出现的错误
func MathLg(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	x := args[0]
	switch num := x.(type) {
	case *object.Int:
		if num.Value <= 0 {
			return nil, &errors.OperationError{
				Frame:    f,
				Message:  "lg argument must be > 0.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
		return &object.Float{Value: math.Log10(float64(num.Value))}, nil
	case *object.Float:
		if num.Value <= 0 {
			return nil, &errors.OperationError{
				Frame:    f,
				Message:  "lg argument must be > 0.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
		return &object.Float{Value: math.Log10(num.Value)}, nil
	}
	return nil, &errors.TypeError{
		Frame:    f,
		Message:  "invalid type for lg function.",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// MathLn 实现自然对数函数(以e为底)
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 函数参数
//
// 返回值:
//
//	object.Object - 自然对数结果
//	error - 可能出现的错误
func MathLn(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	x := args[0]
	switch num := x.(type) {
	case *object.Int:
		if num.Value <= 0 {
			return nil, &errors.OperationError{
				Frame:    f,
				Message:  "ln argument must be > 0.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
		return &object.Float{Value: math.Log(float64(num.Value))}, nil
	case *object.Float:
		if num.Value <= 0 {
			return nil, &errors.OperationError{
				Frame:    f,
				Message:  "ln argument must be > 0.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
		return &object.Float{Value: math.Log(num.Value)}, nil
	}
	return nil, &errors.TypeError{
		Frame:    f,
		Message:  "invalid type for ln function.",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// MathFloor 实现向下取整函数
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 函数参数
//
// 返回值:
//
//	object.Object - 向下取整结果
//	error - 可能出现的错误
func MathFloor(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	x := args[0]
	var decimalPlaces int64 = 0

	if len(args) > 1 && args[1] != nil {
		switch dp := args[1].(type) {
		case *object.Int:
			decimalPlaces = dp.Value
		case *object.Float:
			decimalPlaces = int64(dp.Value)
		default:
			return nil, &errors.TypeError{
				Frame:    f,
				Message:  "decimalPlaces must be a number.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
	}

	var val float64
	switch num := x.(type) {
	case *object.Int:
		val = float64(num.Value)
	case *object.Float:
		val = num.Value
	default:
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "invalid type for floor function.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	if decimalPlaces == 0 {
		return &object.Int{Value: int64(math.Floor(val))}, nil
	}
	multiplier := math.Pow(10, float64(decimalPlaces))
	result := math.Floor(val*multiplier) / multiplier
	return &object.Float{Value: result}, nil
}

// MathCeil 实现向上取整函数
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 函数参数
//
// 返回值:
//
//	object.Object - 向上取整结果
//	error - 可能出现的错误
func MathCeil(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	x := args[0]
	var decimalPlaces int64 = 0

	if len(args) > 1 && args[1] != nil {
		switch dp := args[1].(type) {
		case *object.Int:
			decimalPlaces = dp.Value
		case *object.Float:
			decimalPlaces = int64(dp.Value)
		default:
			return nil, &errors.TypeError{
				Frame:    f,
				Message:  "decimalPlaces must be a number.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
	}

	var val float64
	switch num := x.(type) {
	case *object.Int:
		val = float64(num.Value)
	case *object.Float:
		val = num.Value
	default:
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "invalid type for ceil function.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	if decimalPlaces == 0 {
		return &object.Int{Value: int64(math.Ceil(val))}, nil
	}
	multiplier := math.Pow(10, float64(decimalPlaces))
	result := math.Ceil(val*multiplier) / multiplier
	return &object.Float{Value: result}, nil
}

// MathRound 实现四舍五入函数
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 函数参数
//
// 返回值:
//
//	object.Object - 四舍五入结果
//	error - 可能出现的错误
func MathRound(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	x := args[0]
	var decimalPlaces int64 = 0

	if len(args) > 1 && args[1] != nil {
		switch dp := args[1].(type) {
		case *object.Int:
			decimalPlaces = dp.Value
		case *object.Float:
			decimalPlaces = int64(dp.Value)
		default:
			return nil, &errors.TypeError{
				Frame:    f,
				Message:  "decimalPlaces must be a number.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
	}

	var val float64
	switch num := x.(type) {
	case *object.Int:
		val = float64(num.Value)
	case *object.Float:
		val = num.Value
	default:
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "invalid type for round function.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	if decimalPlaces == 0 {
		return &object.Int{Value: int64(math.Round(val))}, nil
	}
	multiplier := math.Pow(10, float64(decimalPlaces))
	result := math.Round(val*multiplier) / multiplier
	return &object.Float{Value: result}, nil
}

// MathMin 实现最小值函数
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 函数参数
//
// 返回值:
//
//	object.Object - 最小值结果
//	error - 可能出现的错误
func MathMin(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	if len(args[0].(*object.List).Elements) == 0 {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "min requires at least one argument.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	var minVal float64
	var isInt bool
	for i, arg := range args[0].(*object.List).Elements {
		var val float64
		switch num := arg.(type) {
		case *object.Int:
			val = float64(num.Value)
			if i == 0 {
				isInt = true
			}
		case *object.Float:
			val = num.Value
			isInt = false
		default:
			return nil, &errors.TypeError{
				Frame:    f,
				Message:  "invalid type for min function.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
		if i == 0 || val < minVal {
			minVal = val
		}
	}

	if isInt {
		return &object.Int{Value: int64(minVal)}, nil
	}
	return &object.Float{Value: minVal}, nil
}

// MathMax 实现最大值函数
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 函数参数
//
// 返回值:
//
//	object.Object - 最大值结果
//	error - 可能出现的错误
func MathMax(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	if len(args[0].(*object.List).Elements) == 0 {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "max requires at least one argument.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	var maxVal float64
	var isInt bool
	for i, arg := range args[0].(*object.List).Elements {
		var val float64
		switch num := arg.(type) {
		case *object.Int:
			val = float64(num.Value)
			if i == 0 {
				isInt = true
			}
		case *object.Float:
			val = num.Value
			isInt = false
		default:
			return nil, &errors.TypeError{
				Frame:    f,
				Message:  "invalid type for max function.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
		if i == 0 || val > maxVal {
			maxVal = val
		}
	}

	if isInt {
		return &object.Int{Value: int64(maxVal)}, nil
	}
	return &object.Float{Value: maxVal}, nil
}

// MathSum 实现求和函数
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 函数参数
//
// 返回值:
//
//	object.Object - 求和结果
//	error - 可能出现的错误
func MathSum(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	if len(args[0].(*object.List).Elements) == 0 {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "sum requires at least one argument.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	var sum float64
	var isInt bool
	for i, arg := range args[0].(*object.List).Elements {
		var val float64
		switch num := arg.(type) {
		case *object.Int:
			val = float64(num.Value)
			if i == 0 {
				isInt = true
			}
		case *object.Float:
			val = num.Value
			isInt = false
		default:
			return nil, &errors.TypeError{
				Frame:    f,
				Message:  "invalid type for sum function.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
		sum += val
	}

	if isInt {
		return &object.Int{Value: int64(sum)}, nil
	}
	return &object.Float{Value: sum}, nil
}

// MathProduct 实现乘积函数
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 函数参数
//
// 返回值:
//
//	object.Object - 乘积结果
//	error - 可能出现的错误
func MathProduct(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	if len(args[0].(*object.List).Elements) == 0 {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "product requires at least one argument.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	var product float64 = 1
	var isInt bool
	for i, arg := range args[0].(*object.List).Elements {
		var val float64
		switch num := arg.(type) {
		case *object.Int:
			val = float64(num.Value)
			if i == 0 {
				isInt = true
			}
		case *object.Float:
			val = num.Value
			isInt = false
		default:
			return nil, &errors.TypeError{
				Frame:    f,
				Message:  "invalid type for product function.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
		product *= val
	}

	if isInt {
		return &object.Int{Value: int64(product)}, nil
	}
	return &object.Float{Value: product}, nil
}

// MathMean 实现平均值函数
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 函数参数
//
// 返回值:
//
//	object.Object - 平均值结果
//	error - 可能出现的错误
func MathMean(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	if len(args[0].(*object.List).Elements) == 0 {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "mean requires at least one argument.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	var sum float64
	for _, arg := range args[0].(*object.List).Elements {
		var val float64
		switch num := arg.(type) {
		case *object.Int:
			val = float64(num.Value)
		case *object.Float:
			val = num.Value
		default:
			return nil, &errors.TypeError{
				Frame:    f,
				Message:  "invalid type for mean function.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
		sum += val
	}
	
	return &object.Float{Value: sum / float64(len(args[0].(*object.List).Elements))}, nil
}

// MathMedian 实现中位数函数
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 函数参数
//
// 返回值:
//
//	object.Object - 中位数结果
//	error - 可能出现的错误
func MathMedian(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	if len(args[0].(*object.List).Elements) == 0 {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "median requires at least one argument.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	var vals []float64
	for _, arg := range args[0].(*object.List).Elements {
		var val float64
		switch num := arg.(type) {
		case *object.Int:
			val = float64(num.Value)
		case *object.Float:
			val = num.Value
		default:
			return nil, &errors.TypeError{
				Frame:    f,
				Message:  "invalid type for median function.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
		vals = append(vals, val)
	}

	for i := 0; i < len(vals); i++ {
		for j := i + 1; j < len(vals); j++ {
			if vals[i] > vals[j] {
				vals[i], vals[j] = vals[j], vals[i]
			}
		}
	}

	n := len(vals)
	if n%2 == 1 {
		return &object.Float{Value: vals[n/2]}, nil
	}
	return &object.Float{Value: (vals[n/2-1] + vals[n/2]) / 2}, nil
}

// MathVariance 实现方差函数
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 函数参数
//
// 返回值:
//
//	object.Object - 方差结果
//	error - 可能出现的错误
func MathVariance(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	if len(args[0].(*object.List).Elements) == 0 {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "variance requires at least one argument.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	var sum float64
	var vals []float64
	for _, arg := range args[0].(*object.List).Elements {
		var val float64
		switch num := arg.(type) {
		case *object.Int:
			val = float64(num.Value)
		case *object.Float:
			val = num.Value
		default:
			return nil, &errors.TypeError{
				Frame:    f,
				Message:  "invalid type for variance function.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
		sum += val
		vals = append(vals, val)
	}

	mean := sum / float64(len(vals))
	var varianceSum float64
	for _, val := range vals {
		diff := val - mean
		varianceSum += diff * diff
	}

	return &object.Float{Value: varianceSum / float64(len(vals))}, nil
}

// MathStdDev 实现标准差函数
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 函数参数
//
// 返回值:
//
//	object.Object - 标准差结果
//	error - 可能出现的错误
func MathStdDev(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	varianceObj, err := MathVariance(f, env, posStart, posEnd, args...)
	if err != nil {
		return nil, err
	}

	variance, ok := varianceObj.(*object.Float)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "invalid variance result.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	return &object.Float{Value: math.Sqrt(variance.Value)}, nil
}

// MathRand 实现随机数生成函数(0-1之间)
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 函数参数
//
// 返回值:
//
//	object.Object - 随机数结果
//	error - 可能出现的错误
func MathRand(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	return &object.Float{Value: rand.Float64()}, nil
}

// MathRandInt 实现随机整数生成函数
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 函数参数
//
// 返回值:
//
//	object.Object - 随机整数结果
//	error - 可能出现的错误
func MathRandInt(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	if len(args) != 2 {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "randInt requires two arguments.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	var minVal, maxVal int64
	switch minArg := args[0].(type) {
	case *object.Int:
		minVal = minArg.Value
	case *object.Float:
		minVal = int64(minArg.Value)
	default:
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "min must be a number.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	switch maxArg := args[1].(type) {
	case *object.Int:
		maxVal = maxArg.Value
	case *object.Float:
		maxVal = int64(maxArg.Value)
	default:
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "max must be a number.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	if minVal > maxVal {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "min must be less than or equal to max.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	if minVal == maxVal {
		return &object.Int{Value: minVal}, nil
	}

	return &object.Int{Value: minVal + rand.Int63n(maxVal-minVal+1)}, nil
}
