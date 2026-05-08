package builtinclass

import (
	"github.com/Ghost-Xiao/ghost-lang/internal/errors"
	"github.com/Ghost-Xiao/ghost-lang/internal/frame"
	"github.com/Ghost-Xiao/ghost-lang/internal/object"
	"github.com/Ghost-Xiao/ghost-lang/internal/util"
)

// MapClass 表示 Map 内置类的类定义
var MapClass = initMapClass()

// initMapClass 初始化 Map 类
//
// 返回值:
//
//	*object.Class - 初始化后的 Map 类对象
func initMapClass() *object.Class {
	member := &object.Environment{
		Name:  "Map",
		Store: map[string]*object.Symbol{},
		Outer: nil,
	}

	member.Set("init", &object.Symbol{Name: "init", Value: &MAPINIT, IsConst: true})
	member.Set("keys", &object.Symbol{Name: "keys", Value: &MAPKEYS, IsConst: true})
	member.Set("values", &object.Symbol{Name: "values", Value: &MAPVALUES, IsConst: true})
	member.Set("hasKey", &object.Symbol{Name: "hasKey", Value: &MAPHASKEY, IsConst: true})
	member.Set("remove", &object.Symbol{Name: "remove", Value: &MAPREMOVE, IsConst: true})

	return &object.Class{
		Name:   "Map",
		Parent: nil,
		Member: member,
	}
}

var (
	// MAPINIT 表示 Map 类的 init 方法
	MAPINIT = object.BuiltinFunction{
		Name:         "init",
		Parameter:    []string{},
		DefaultValue: []object.Object{},
		HaveVariadic: false,
		Fn:           MapInit,
	}
	// MAPKEYS 表示 Map 类的 keys 方法
	MAPKEYS = object.BuiltinFunction{
		Name:         "keys",
		Parameter:    []string{},
		DefaultValue: []object.Object{},
		HaveVariadic: false,
		Fn:           MapKeys,
	}
	// MAPVALUES 表示 Map 类的 values 方法
	MAPVALUES = object.BuiltinFunction{
		Name:         "values",
		Parameter:    []string{},
		DefaultValue: []object.Object{},
		HaveVariadic: false,
		Fn:           MapValues,
	}
	// MAPHASKEY 表示 Map 类的 hasKey 方法
	MAPHASKEY = object.BuiltinFunction{
		Name:         "hasKey",
		Parameter:    []string{"key"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           MapHasKey,
	}
	// MAPREMOVE 表示 Map 类的 remove 方法
	MAPREMOVE = object.BuiltinFunction{
		Name:         "remove",
		Parameter:    []string{"key"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           MapRemove,
	}
)

// MapKeys 实现 Map 类的 keys 方法
// 返回 Map 中所有键的列表
//
// 参数:
//
//	f - 当前调用栈
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数，第一个参数是 Map 实例
//
// 返回值:
//
//	object.Object - 包含所有键的列表
//	error - 可能出现的错误
func MapKeys(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this := args[0]
	if this == nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "method keys() called without instance",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	mapObj, ok := this.(*object.Map)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "keys() can only be called on Map instances",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	keys := make([]object.Object, 0, len(mapObj.Pairs))
	for _, pair := range mapObj.Pairs {
		keys = append(keys, pair.Key)
	}

	return &object.List{Elements: keys}, nil
}

// MapValues 实现 Map 类的 values 方法
// 返回 Map 中所有值的列表
//
// 参数:
//
//	f - 当前调用栈
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数，第一个参数是 Map 实例
//
// 返回值:
//
//	object.Object - 包含所有值的列表
//	error - 可能出现的错误
func MapValues(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this := args[0]
	if this == nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "method values() called without instance",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	mapObj, ok := this.(*object.Map)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "values() can only be called on Map instances",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	values := make([]object.Object, 0, len(mapObj.Pairs))
	for _, pair := range mapObj.Pairs {
		values = append(values, pair.Value)
	}

	return &object.List{Elements: values}, nil
}

// MapHasKey 实现 Map 类的 hasKey 方法
// 判断 Map 中是否包含指定的键
//
// 参数:
//
//	f - 当前调用栈
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数，第一个参数是 Map 实例，第二个参数是要检查的键
//
// 返回值:
//
//	object.Object - Bool 类型的结果，true 表示包含该键，false 表示不包含
//	error - 可能出现的错误
func MapHasKey(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this := args[0]
	if this == nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "method hasKey() called without instance",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	mapObj, ok := this.(*object.Map)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "hasKey() can only be called on Map instances",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	if len(args) < 2 {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "hasKey() expects 1 argument",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	key := args[1]
	hashable, ok := key.(object.Hashable)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "hasKey() key must be hashable",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	hashKey := object.HashKey{
		Type:  hashable.Type(),
		Value: hashable.Hash(),
	}

	_, exists := mapObj.Pairs[hashKey]
	return &object.Bool{Value: exists}, nil
}

// MapRemove 实现 Map 类的 remove 方法
// 移除 Map 中指定的键并返回对应的值
//
// 参数:
//
//	f - 当前调用栈
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数，第一个参数是 Map 实例，第二个参数是要移除的键
//
// 返回值:
//
//	object.Object - 被移除的键对应的值，如果键不存在则返回 null
//	error - 可能出现的错误
func MapRemove(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this := args[0]
	if this == nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "method remove() called without instance",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	mapObj, ok := this.(*object.Map)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "remove() can only be called on Map instances",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	if len(args) < 2 {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "remove() expects 1 argument",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	key := args[1]
	hashable, ok := key.(object.Hashable)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "remove() key must be hashable",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	hashKey := object.HashKey{
		Type:  hashable.Type(),
		Value: hashable.Hash(),
	}

	pair, exists := mapObj.Pairs[hashKey]
	if !exists {
		return &object.Null{}, nil
	}

	delete(mapObj.Pairs, hashKey)
	return pair.Value, nil
}

// MapInit 实现 Map 类的 init 构造方法
// 创建一个新的空 Map 实例
//
// 参数:
//
//	f - 当前调用栈
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数，第一个参数是 Map 实例
//
// 返回值:
//
//	object.Object - 新创建的空 Map 实例
//	error - 可能出现的错误
func MapInit(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	return &object.Map{Pairs: map[object.HashKey]object.Pair{}}, nil
}
