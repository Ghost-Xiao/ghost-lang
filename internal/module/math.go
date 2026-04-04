package module

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/Ghost-Xiao/ghost-lang/internal/errors"
	"github.com/Ghost-Xiao/ghost-lang/internal/frame"
	"github.com/Ghost-Xiao/ghost-lang/internal/object"
	"github.com/Ghost-Xiao/ghost-lang/internal/util"
)

// Math 数学模块
type Math struct{}

// Name 获取模块名称
//
// 返回值:
//
// string - 模块名称
func (m *Math) Name() string {
	return "math"
}

// Load 加载模块
//
// 返回值:
//
// *Environment - 模块环境
func (m *Math) Load() *object.Environment {
	env := &object.Environment{
		Name:  m.Name(),
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
	env.Set("seed", &object.Symbol{Name: "seed", Value: &SEED, IsConst: true})

	return env
}

// Math模块成员
var (
	PI  = object.Float{Value: 3.141592653589793} // 圆周率
	E   = object.Float{Value: 2.718281828459045} // 自然对数的底
	TAU = object.Float{Value: 6.283185307179586} // 圆周率的两倍，即圆周长
	ABS = object.BuiltinFunction{
		Name:         "abs",
		Parameter:    []string{"x"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           MathAbs,
	} // 绝对值函数
	SQRT = object.BuiltinFunction{
		Name:         "sqrt",
		Parameter:    []string{"x"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           MathSqrt,
	} // 平方根函数
	SIN = object.BuiltinFunction{
		Name:         "sin",
		Parameter:    []string{"x"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           MathSin,
	} // 正弦函数
	COS = object.BuiltinFunction{
		Name:         "cos",
		Parameter:    []string{"x"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           MathCos,
	} // 余弦函数
	TAN = object.BuiltinFunction{
		Name:         "tan",
		Parameter:    []string{"x"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           MathTan,
	} // 正切函数
	ASIN = object.BuiltinFunction{
		Name:         "asin",
		Parameter:    []string{"x"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           MathAsin,
	} // 反正弦函数
	ACOS = object.BuiltinFunction{
		Name:         "acos",
		Parameter:    []string{"x"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           MathAcos,
	} // 反余弦函数
	ATAN = object.BuiltinFunction{
		Name:         "atan",
		Parameter:    []string{"x"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           MathAtan,
	} // 反正切函数
	LOG = object.BuiltinFunction{
		Name:         "log",
		Parameter:    []string{"a", "x"},
		DefaultValue: []object.Object{nil, nil},
		HaveVariadic: false,
		Fn:           MathLog,
	} // 以a为底x的对数函数
	LG = object.BuiltinFunction{
		Name:         "lg",
		Parameter:    []string{"x"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           MathLg,
	} // 10为底对数函数
	LN = object.BuiltinFunction{
		Name:         "ln",
		Parameter:    []string{"x"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           MathLn,
	} // 自然对数函数
	FLOOR = object.BuiltinFunction{
		Name:         "floor",
		Parameter:    []string{"x", "decimalPlaces"},
		DefaultValue: []object.Object{nil, &object.Int{Value: 0}},
		HaveVariadic: false,
		Fn:           MathFloor,
	} // 向下取整函数
	CEIL = object.BuiltinFunction{
		Name:         "ceil",
		Parameter:    []string{"x", "decimalPlaces"},
		DefaultValue: []object.Object{nil, &object.Int{Value: 0}},
		HaveVariadic: false,
		Fn:           MathCeil,
	} // 向上取整函数
	ROUND = object.BuiltinFunction{
		Name:         "round",
		Parameter:    []string{"x", "decimalPlaces"},
		DefaultValue: []object.Object{nil, &object.Int{Value: 0}},
		HaveVariadic: false,
		Fn:           MathRound,
	} // 四舍五入函数
	MIN = object.BuiltinFunction{
		Name:         "min",
		Parameter:    []string{"a"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: true,
		Fn:           MathMin,
	} // 最小值函数
	MAX = object.BuiltinFunction{
		Name:         "max",
		Parameter:    []string{"a"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: true,
		Fn:           MathMax,
	} // 最大值函数
	SUM = object.BuiltinFunction{
		Name:         "sum",
		Parameter:    []string{"a"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: true,
		Fn:           MathSum,
	} // 求和函数
	PRODUCT = object.BuiltinFunction{
		Name:         "product",
		Parameter:    []string{"a"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: true,
		Fn:           MathProduct,
	} // 求积函数
	MEAN = object.BuiltinFunction{
		Name:         "mean",
		Parameter:    []string{"a"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: true,
		Fn:           MathMean,
	} // 平均值函数
	MEDIAN = object.BuiltinFunction{
		Name:         "median",
		Parameter:    []string{"a"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: true,
		Fn:           MathMedian,
	} // 中位数函数
	VARIANCE = object.BuiltinFunction{
		Name:         "variance",
		Parameter:    []string{"a"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: true,
		Fn:           MathVariance,
	} // 方差函数
	STDDEV = object.BuiltinFunction{
		Name:         "stdDev",
		Parameter:    []string{"a"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: true,
		Fn:           MathStdDev,
	} // 标准差函数
	RAND = object.BuiltinFunction{
		Name:         "rand",
		Parameter:    []string{},
		DefaultValue: []object.Object{},
		HaveVariadic: false,
		Fn:           MathRand,
	} // 随机数函数（0-1之间）
	RANDINT = object.BuiltinFunction{
		Name:         "randInt",
		Parameter:    []string{"min", "max"},
		DefaultValue: []object.Object{nil, nil},
		HaveVariadic: false,
		Fn:           MathRandInt,
	} // 随机整数函数
	SEED = object.BuiltinFunction{
		Name:         "seed",
		Parameter:    []string{"s"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           MathSeed,
	} // 随机数种子函数
)

func MathAbs(f *frame.Frame, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
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

func MathSqrt(f *frame.Frame, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
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

func MathSin(f *frame.Frame, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
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

func MathCos(f *frame.Frame, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
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

func MathTan(f *frame.Frame, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
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

func MathAsin(f *frame.Frame, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
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

func MathAcos(f *frame.Frame, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
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

func MathAtan(f *frame.Frame, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
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

func MathLog(f *frame.Frame, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
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

func MathLg(f *frame.Frame, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
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

func MathLn(f *frame.Frame, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
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

func MathFloor(f *frame.Frame, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
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

func MathCeil(f *frame.Frame, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
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

func MathRound(f *frame.Frame, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
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

func MathMin(f *frame.Frame, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	if len(args) == 0 {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "min requires at least one argument.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	var minVal float64
	var isInt bool
	for i, arg := range args {
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

func MathMax(f *frame.Frame, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	if len(args) == 0 {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "max requires at least one argument.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	var maxVal float64
	var isInt bool
	for i, arg := range args {
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

func MathSum(f *frame.Frame, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	if len(args) == 0 {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "sum requires at least one argument.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	var sum float64
	var isInt bool
	for i, arg := range args {
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

func MathProduct(f *frame.Frame, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	if len(args) == 0 {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "product requires at least one argument.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	var product float64 = 1
	var isInt bool
	for i, arg := range args {
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

func MathMean(f *frame.Frame, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	if len(args) == 0 {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "mean requires at least one argument.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	var sum float64
	for _, arg := range args {
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

	return &object.Float{Value: sum / float64(len(args))}, nil
}

func MathMedian(f *frame.Frame, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	if len(args) == 0 {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "median requires at least one argument.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	var vals []float64
	for _, arg := range args {
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

func MathVariance(f *frame.Frame, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	if len(args) == 0 {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "variance requires at least one argument.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	var sum float64
	var vals []float64
	for _, arg := range args {
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

func MathStdDev(f *frame.Frame, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	varianceObj, err := MathVariance(f, posStart, posEnd, args...)
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

func MathRand(f *frame.Frame, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	return &object.Float{Value: rand.Float64()}, nil
}

func MathRandInt(f *frame.Frame, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
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

func MathSeed(f *frame.Frame, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "seed requires one argument.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	var seedVal int64
	switch seedArg := args[0].(type) {
	case *object.Int:
		seedVal = seedArg.Value
	case *object.Float:
		seedVal = int64(seedArg.Value)
	default:
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "seed must be a number.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	rand.Seed(seedVal)
	return &object.Null{}, nil
}

// Type 返回值的类型
//
// 返回值:
//
//	string - 值的类型
func (m *Math) Type() string {
	return "Module"
}

// String 返回值的字符串表示
//
// 返回值:
//
//	string - 格式化的字符串表示
func (m *Math) String() string {
	return fmt.Sprintf("module \"%s\"", m.Name())
}

// Negative 对值进行负运算
//
// 参数:
//
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	object.Object - 运算结果
//	error - 可能出现的错误
func (m *Math) Negative(posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"-\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// BitNot 对值进行按位非运算
//
// 参数:
//
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	object.Object - 运算结果
//	error - 可能出现的错误
func (m *Math) BitNot(posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"~\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// Not 对值进行逻辑非运算
//
// 参数:
//
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	object.Object - 运算结果
//	error - 可能出现的错误
func (m *Math) Not(posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"!\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// Add 对值进行加法运算
//
// 参数:
//
//	other - 另一个操作数
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	object.Object - 运算结果
//	error - 可能出现的错误
func (m *Math) Add(_ object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"+\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// Subtract 对值进行减法运算
//
// 参数:
//
//	other - 另一个操作数
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	object.Object - 运算结果
//	error - 可能出现的错误
func (m *Math) Subtract(_ object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"-\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// Multiply 对值进行乘法运算
//
// 参数:
//
//	other - 另一个操作数
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	object.Object - 运算结果
//	error - 可能出现的错误
func (m *Math) Multiply(_ object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"*\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// Divide 对值进行除法运算
//
// 参数:
//
//	other - 另一个操作数
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	object.Object - 运算结果
//	error - 可能出现的错误
func (m *Math) Divide(_ object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"/\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// Mod 对值进行取模运算
//
// 参数:
//
//	other - 另一个操作数
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	object.Object - 运算结果
//	error - 可能出现的错误
func (m *Math) Mod(_ object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"%\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// Equal 判断当前值与另一个值是否相等
//
// 参数:
//
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	object.Object - 运算结果
//	error - 可能出现的错误
func (m *Math) Equal(other object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"==\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// NotEqual 判断当前值与另一个值是否不相等
//
// 参数:
//
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	object.Object - 运算结果
//	error - 可能出现的错误
func (m *Math) NotEqual(other object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"!=\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// LessThan 对值进行小于比较
//
// 参数:
//
//	other - 另一个操作数
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	object.Object - 比较结果
func (m *Math) LessThan(_ object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"<\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// GreaterThan 对值进行大于比较
//
// 参数:
//
//	other - 另一个操作数
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	object.Object - 比较结果
func (m *Math) GreaterThan(_ object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \">\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// LessThanOrEqual 对值进行小于等于比较
//
// 参数:
//
//	other - 另一个操作数
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	object.Object - 比较结果
func (m *Math) LessThanOrEqual(_ object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"<=\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// GreaterThanOrEqual 对值进行大于等于比较
//
// 参数:
//
//	other - 另一个操作数
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	object.Object - 比较结果
func (m *Math) GreaterThanOrEqual(_ object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \">=\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// BitAnd 对值进行按位与运算
//
// 参数:
//
//	other - 另一个操作数
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	object.Object - 运算结果
//	error - 可能出现的错误
func (m *Math) BitAnd(_ object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"&\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// BitOr 对值进行按位或运算
//
// 参数:
//
//	other - 另一个操作数
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	object.Object - 运算结果
//	error - 可能出现的错误
func (m *Math) BitOr(_ object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"|\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// Xor 对值进行异或运算
//
// 参数:
//
//	other - 另一个操作数
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	object.Object - 运算结果
//	error - 可能出现的错误
func (m *Math) Xor(_ object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"^\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// LeftShift 对值进行左移运算
//
// 参数:
//
//	other - 另一个操作数
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	object.Object - 运算结果
//	error - 可能出现的错误
func (m *Math) LeftShift(_ object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"<<\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// RightShift 对值进行右移运算
//
// 参数:
//
//	other - 另一个操作数
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	object.Object - 运算结果
//	error - 可能出现的错误
func (m *Math) RightShift(_ object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \">>\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// And 对值进行逻辑与运算
//
// 参数:
//
//	other - 另一个操作数
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	object.Object - 运算结果
//	error - 可能出现的错误
func (m *Math) And(_ object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"&&\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// Or 对值进行逻辑或运算
//
// 参数:
//
//	other - 另一个操作数
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	object.Object - 运算结果
//	error - 可能出现的错误
func (m *Math) Or(_ object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"||\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// Index 执行索引运算
//
// 参数:
//
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	frame - 当前调用栈
//
// 返回值:
//
//	object.Object - 运算结果
//	error - 可能出现的错误
func (m *Math) Index(other object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.TypeError{
		Frame:    frame,
		Message:  "index expression not supported for this type.",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}
