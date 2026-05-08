package builtinclass

import (
	"fmt"
	"strings"

	"github.com/Ghost-Xiao/ghost-lang/internal/errors"
	"github.com/Ghost-Xiao/ghost-lang/internal/frame"
	"github.com/Ghost-Xiao/ghost-lang/internal/object"
	"github.com/Ghost-Xiao/ghost-lang/internal/util"
)

// StringClass 表示 String 内置类的类定义
var StringClass = initStringClass()

// initStringClass 初始化 String 类
//
// 返回值:
//
//	*object.Class - 初始化后的 String 类对象
func initStringClass() *object.Class {
	member := &object.Environment{
		Name:  "String",
		Store: map[string]*object.Symbol{},
		Outer: nil,
	}

	member.Set("init", &object.Symbol{Name: "init", Value: &STRINGINIT, IsConst: true})
	member.Set("upper", &object.Symbol{Name: "upper", Value: &STRINGUPPER, IsConst: true})
	member.Set("lower", &object.Symbol{Name: "lower", Value: &STRINGLOWER, IsConst: true})
	member.Set("startsWith", &object.Symbol{Name: "startsWith", Value: &STRINGSTARTSWITH, IsConst: true})
	member.Set("endsWith", &object.Symbol{Name: "endsWith", Value: &STRINGENDSWITH, IsConst: true})
	member.Set("split", &object.Symbol{Name: "split", Value: &STRINGSPLIT, IsConst: true})
	member.Set("trim", &object.Symbol{Name: "trim", Value: &STRINGTRIM, IsConst: true})
	member.Set("replace", &object.Symbol{Name: "replace", Value: &STRINGREPLACE, IsConst: true})
	member.Set("indexOf", &object.Symbol{Name: "indexOf", Value: &STRINGINDEXOF, IsConst: true})

	return &object.Class{
		Name:   "String",
		Parent: nil,
		Member: member,
	}
}

var (
	// STRINGINIT 表示 String 类的 init 方法
	STRINGINIT = object.BuiltinFunction{
		Name:         "init",
		Parameter:    []string{"value"},
		DefaultValue: []object.Object{&object.String{Value: ""}},
		HaveVariadic: false,
		Fn:           StringInit,
	}
	// STRINGUPPER 表示 String 类的 upper 方法
	STRINGUPPER = object.BuiltinFunction{
		Name:         "upper",
		Parameter:    []string{},
		DefaultValue: []object.Object{},
		HaveVariadic: false,
		Fn:           StringUpper,
	}
	// STRINGLOWER 表示 String 类的 lower 方法
	STRINGLOWER = object.BuiltinFunction{
		Name:         "lower",
		Parameter:    []string{},
		DefaultValue: []object.Object{},
		HaveVariadic: false,
		Fn:           StringLower,
	}
	// STRINGSTARTSWITH 表示 String 类的 startsWith 方法
	STRINGSTARTSWITH = object.BuiltinFunction{
		Name:         "startsWith",
		Parameter:    []string{"prefix"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           StringStartsWith,
	}
	// STRINGENDSWITH 表示 String 类的 endsWith 方法
	STRINGENDSWITH = object.BuiltinFunction{
		Name:         "endsWith",
		Parameter:    []string{"suffix"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           StringEndsWith,
	}
	// STRINGSPLIT 表示 String 类的 split 方法
	STRINGSPLIT = object.BuiltinFunction{
		Name:         "split",
		Parameter:    []string{"del"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           StringSplit,
	}
	// STRINGTRIM 表示 String 类的 trim 方法
	STRINGTRIM = object.BuiltinFunction{
		Name:         "trim",
		Parameter:    []string{},
		DefaultValue: []object.Object{},
		HaveVariadic: false,
		Fn:           StringTrim,
	}
	// STRINGREPLACE 表示 String 类的 replace 方法
	STRINGREPLACE = object.BuiltinFunction{
		Name:         "replace",
		Parameter:    []string{"old", "new"},
		DefaultValue: []object.Object{nil, nil},
		HaveVariadic: false,
		Fn:           StringReplace,
	}
	// STRINGINDEXOF 表示 String 类的 indexOf 方法
	STRINGINDEXOF = object.BuiltinFunction{
		Name:         "indexOf",
		Parameter:    []string{"sub"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           StringIndexOf,
	}
)

// StringUpper 实现 String 类的 upper 方法
// 将字符串转换为大写
//
// 参数:
//
//	f - 当前调用栈
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数，第一个参数是 String 实例
//
// 返回值:
//
//	object.Object - 转换为大写后的新字符串
//	error - 可能出现的错误
func StringUpper(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this := args[0]
	if this == nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "method upper() called without instance",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	strObj, ok := this.(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "upper() can only be called on String instances",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	return &object.String{Value: strings.ToUpper(strObj.Value)}, nil
}

// StringLower 实现 String 类的 lower 方法
// 将字符串转换为小写
//
// 参数:
//
//	f - 当前调用栈
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数，第一个参数是 String 实例
//
// 返回值:
//
//	object.Object - 转换为小写后的新字符串
//	error - 可能出现的错误
func StringLower(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this := args[0]
	if this == nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "method lower() called without instance",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	strObj, ok := this.(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "lower() can only be called on String instances",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	return &object.String{Value: strings.ToLower(strObj.Value)}, nil
}

// StringStartsWith 实现 String 类的 startsWith 方法
// 判断字符串是否以指定前缀开头
//
// 参数:
//
//	f - 当前调用栈
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数，第一个参数是 String 实例，第二个参数是前缀字符串
//
// 返回值:
//
//	object.Object - Bool 类型的结果，true 表示以指定前缀开头，false 表示不是
//	error - 可能出现的错误
func StringStartsWith(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this := args[0]
	if this == nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "method startsWith() called without instance",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	strObj, ok := this.(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "startsWith() can only be called on String instances",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	if len(args) < 2 {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "startsWith() expects 1 argument",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	prefix, ok := args[1].(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "startsWith() expects a String argument",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	return &object.Bool{Value: strings.HasPrefix(strObj.Value, prefix.Value)}, nil
}

// StringEndsWith 实现 String 类的 endsWith 方法
// 判断字符串是否以指定后缀结尾
//
// 参数:
//
//	f - 当前调用栈
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数，第一个参数是 String 实例，第二个参数是后缀字符串
//
// 返回值:
//
//	object.Object - Bool 类型的结果，true 表示以指定后缀结尾，false 表示不是
//	error - 可能出现的错误
func StringEndsWith(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this := args[0]
	if this == nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "method endsWith() called without instance",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	strObj, ok := this.(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "endsWith() can only be called on String instances",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	if len(args) < 2 {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "endsWith() expects 1 argument",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	suffix, ok := args[1].(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "endsWith() expects a String argument",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	return &object.Bool{Value: strings.HasSuffix(strObj.Value, suffix.Value)}, nil
}

// StringSplit 实现 String 类的 split 方法
// 按照指定分隔符将字符串分割成列表
//
// 参数:
//
//	f - 当前调用栈
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数，第一个参数是 String 实例，第二个参数是分隔符字符串
//
// 返回值:
//
//	object.Object - 分割后的列表
//	error - 可能出现的错误
func StringSplit(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this := args[0]
	if this == nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "method split() called without instance",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	strObj, ok := this.(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "split() can only be called on String instances",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	if len(args) < 2 {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "split() expects 1 argument",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	del, ok := args[1].(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "split() expects a String argument",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	parts := strings.Split(strObj.Value, del.Value)
	elements := make([]object.Object, len(parts))
	for i, part := range parts {
		elements[i] = &object.String{Value: part}
	}

	return &object.List{Elements: elements}, nil
}

// StringTrim 实现 String 类的 trim 方法
// 去除字符串首尾的空白字符
//
// 参数:
//
//	f - 当前调用栈
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数，第一个参数是 String 实例
//
// 返回值:
//
//	object.Object - 去除首尾空白字符后的新字符串
//	error - 可能出现的错误
func StringTrim(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this := args[0]
	if this == nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "method trim() called without instance",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	strObj, ok := this.(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "trim() can only be called on String instances",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	return &object.String{Value: strings.TrimSpace(strObj.Value)}, nil
}

// StringReplace 实现 String 类的 replace 方法
// 将字符串中的所有旧子字符串替换为新子字符串
//
// 参数:
//
//	f - 当前调用栈
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数，第一个参数是 String 实例，第二个是旧字符串，第三个是新字符串
//
// 返回值:
//
//	object.Object - 替换后的新字符串
//	error - 可能出现的错误
func StringReplace(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this := args[0]
	if this == nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "method replace() called without instance",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	strObj, ok := this.(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "replace() can only be called on String instances",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	if len(args) < 3 {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "replace() expects 2 arguments",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	oldStr, ok := args[1].(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "replace() expects String arguments",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	newStr, ok := args[2].(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "replace() expects String arguments",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	return &object.String{Value: strings.ReplaceAll(strObj.Value, oldStr.Value, newStr.Value)}, nil
}

// StringIndexOf 实现 String 类的 indexOf 方法
// 查找子字符串在字符串中第一次出现的位置
//
// 参数:
//
//	f - 当前调用栈
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数，第一个参数是 String 实例，第二个是要查找的子字符串
//
// 返回值:
//
//	object.Object - 子字符串第一次出现的位置，如果没有找到则返回 -1
//	error - 可能出现的错误
func StringIndexOf(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this := args[0]
	if this == nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "method indexOf() called without instance",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	strObj, ok := this.(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "indexOf() can only be called on String instances",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	if len(args) < 2 {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "indexOf() expects 1 argument",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	sub, ok := args[1].(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "indexOf() expects a String argument",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	index := strings.Index(strObj.Value, sub.Value)
	return &object.Int{Value: int64(index)}, nil
}

// StringInit 实现 String 类的 init 构造方法
// 创建一个新的 String 实例，支持从多种类型转换
//
// 参数:
//
//	f - 当前调用栈
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数，第一个参数是 String 实例，第二个参数是要转换的值
//
// 返回值:
//
//	object.Object - 新创建的 String 实例
//	error - 可能出现的错误
func StringInit(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	if len(args) < 1 {
		return &object.String{Value: ""}, nil
	}

	value := args[0]

	switch v := value.(type) {
	case *object.String:
		return &object.String{Value: v.Value}, nil
	case *object.Int:
		return &object.String{Value: fmt.Sprintf("%d", v.Value)}, nil
	case *object.Float:
		return &object.String{Value: fmt.Sprintf("%g", v.Value)}, nil
	case *object.Bool:
		return &object.String{Value: fmt.Sprintf("%t", v.Value)}, nil
	case *object.Null:
		return &object.String{Value: "null"}, nil
	default:
		return &object.String{Value: v.String()}, nil
	}
}
