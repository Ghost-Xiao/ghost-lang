package builtinclass

import (
	"sort"

	"github.com/Ghost-Xiao/ghost-lang/internal/errors"
	"github.com/Ghost-Xiao/ghost-lang/internal/frame"
	"github.com/Ghost-Xiao/ghost-lang/internal/object"
	"github.com/Ghost-Xiao/ghost-lang/internal/util"
)

// ListClass 表示 List 内置类的类定义
var ListClass = initListClass()

// initListClass 初始化 List 类
//
// 返回值:
//
//	*object.Class - 初始化后的 List 类对象
func initListClass() *object.Class {
	member := &object.Environment{
		Name:  "List",
		Store: map[string]*object.Symbol{},
		Outer: nil,
	}

	member.Set("init", &object.Symbol{Name: "init", Value: &LISTINIT, IsConst: true})
	member.Set("append", &object.Symbol{Name: "append", Value: &LISTAPPEND, IsConst: true})
	member.Set("pop", &object.Symbol{Name: "pop", Value: &LISTPOP, IsConst: true})
	member.Set("insert", &object.Symbol{Name: "insert", Value: &LISTINSERT, IsConst: true})
	member.Set("remove", &object.Symbol{Name: "remove", Value: &LISTREMOVE, IsConst: true})
	member.Set("indexOf", &object.Symbol{Name: "indexOf", Value: &LISTINDEXOF, IsConst: true})
	member.Set("reverse", &object.Symbol{Name: "reverse", Value: &LISTREVERSE, IsConst: true})
	member.Set("sort", &object.Symbol{Name: "sort", Value: &LISTSORT, IsConst: true})
	member.Set("join", &object.Symbol{Name: "join", Value: &LISTJOIN, IsConst: true})
	member.Set("slice", &object.Symbol{Name: "slice", Value: &LISTSLICE, IsConst: true})

	return &object.Class{
		Name:   "List",
		Parent: nil,
		Member: member,
	}
}

var (
	// LISTINIT 表示 List 类的 init 方法
	LISTINIT = object.BuiltinFunction{
		Name:         "init",
		Parameter:    []string{},
		DefaultValue: []object.Object{},
		HaveVariadic: false,
		Fn:           ListInit,
	}
	// LISTAPPEND 表示 List 类的 append 方法
	LISTAPPEND = object.BuiltinFunction{
		Name:         "append",
		Parameter:    []string{"item"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           ListAppend,
	}
	// LISTPOP 表示 List 类的 pop 方法
	LISTPOP = object.BuiltinFunction{
		Name:         "pop",
		Parameter:    []string{"idx"},
		DefaultValue: []object.Object{&object.Int{Value: -1}},
		HaveVariadic: false,
		Fn:           ListPop,
	}
	// LISTINSERT 表示 List 类的 insert 方法
	LISTINSERT = object.BuiltinFunction{
		Name:         "insert",
		Parameter:    []string{"idx", "item"},
		DefaultValue: []object.Object{nil, nil},
		HaveVariadic: false,
		Fn:           ListInsert,
	}
	// LISTREMOVE 表示 List 类的 remove 方法
	LISTREMOVE = object.BuiltinFunction{
		Name:         "remove",
		Parameter:    []string{"item"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           ListRemove,
	}
	// LISTINDEXOF 表示 List 类的 indexOf 方法
	LISTINDEXOF = object.BuiltinFunction{
		Name:         "indexOf",
		Parameter:    []string{"item"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           ListIndexOf,
	}
	// LISTREVERSE 表示 List 类的 reverse 方法
	LISTREVERSE = object.BuiltinFunction{
		Name:         "reverse",
		Parameter:    []string{},
		DefaultValue: []object.Object{},
		HaveVariadic: false,
		Fn:           ListReverse,
	}
	// LISTSORT 表示 List 类的 sort 方法
	LISTSORT = object.BuiltinFunction{
		Name:         "sort",
		Parameter:    []string{},
		DefaultValue: []object.Object{},
		HaveVariadic: false,
		Fn:           ListSort,
	}
	// LISTJOIN 表示 List 类的 join 方法
	LISTJOIN = object.BuiltinFunction{
		Name:         "join",
		Parameter:    []string{"sep"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           ListJoin,
	}
	// LISTSLICE 表示 List 类的 slice 方法
	LISTSLICE = object.BuiltinFunction{
		Name:         "slice",
		Parameter:    []string{"start", "end"},
		DefaultValue: []object.Object{nil, nil},
		HaveVariadic: false,
		Fn:           ListSlice,
	}
)

// ListAppend 实现 List 类的 append 方法
// 在列表末尾添加一个元素
//
// 参数:
//
//	f - 当前调用栈
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数，第一个参数是 List 实例，第二个参数是要添加的元素
//
// 返回值:
//
//	object.Object - 修改后的列表
//	error - 可能出现的错误
func ListAppend(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this := args[0]
	if this == nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "method append() called without instance",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	listObj, ok := this.(*object.List)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "append() can only be called on List instances",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	if len(args) < 2 {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "append() expects 1 argument",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	listObj.Elements = append(listObj.Elements, args[1])
	return listObj, nil
}

// ListPop 实现 List 类的 pop 方法
// 移除并返回列表中指定位置的元素，默认移除最后一个元素
//
// 参数:
//
//	f - 当前调用栈
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数，第一个参数是 List 实例，第二个参数是可选的索引位置
//
// 返回值:
//
//	object.Object - 被移除的元素
//	error - 可能出现的错误
func ListPop(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this := args[0]
	if this == nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "method pop() called without instance",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	listObj, ok := this.(*object.List)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "pop() can only be called on List instances",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	idx := int64(-1)
	if len(args) >= 2 {
		if intObj, ok := args[1].(*object.Int); ok {
			idx = intObj.Value
		}
	}

	length := int64(len(listObj.Elements))
	if idx < 0 {
		idx = length + idx
	}

	if idx < 0 || idx >= length {
		return nil, &errors.IndexError{
			Frame:    f,
			Message:  "index out of range",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	elem := listObj.Elements[idx]
	listObj.Elements = append(listObj.Elements[:idx], listObj.Elements[idx+1:]...)
	return elem, nil
}

// ListInsert 实现 List 类的 insert 方法
// 在列表的指定位置插入一个元素
//
// 参数:
//
//	f - 当前调用栈
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数，第一个参数是 List 实例，第二个是索引位置，第三个是要插入的元素
//
// 返回值:
//
//	object.Object - 修改后的列表
//	error - 可能出现的错误
func ListInsert(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this := args[0]
	if this == nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "method insert() called without instance",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	listObj, ok := this.(*object.List)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "insert() can only be called on List instances",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	if len(args) < 3 {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "insert() expects 2 arguments",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	idxObj, ok := args[1].(*object.Int)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "insert() expects an Int as first argument",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	idx := idxObj.Value
	length := int64(len(listObj.Elements))
	if idx < 0 {
		idx = length + idx
	}

	if idx < 0 || idx > length {
		return nil, &errors.IndexError{
			Frame:    f,
			Message:  "index out of range",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	item := args[2]
	listObj.Elements = append(listObj.Elements[:idx], append([]object.Object{item}, listObj.Elements[idx:]...)...)
	return listObj, nil
}

// ListRemove 实现 List 类的 remove 方法
// 移除列表中第一个匹配的元素
//
// 参数:
//
//	f - 当前调用栈
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数，第一个参数是 List 实例，第二个参数是要移除的元素
//
// 返回值:
//
//	object.Object - 修改后的列表
//	error - 可能出现的错误
func ListRemove(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this := args[0]
	if this == nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "method remove() called without instance",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	listObj, ok := this.(*object.List)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "remove() can only be called on List instances",
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

	target := args[1]
	for i, elem := range listObj.Elements {
		equal, err := elem.Equal(target, posStart, posEnd, f)
		if err != nil {
			return nil, err
		}
		if equal.(*object.Bool).Value {
			listObj.Elements = append(listObj.Elements[:i], listObj.Elements[i+1:]...)
			return listObj, nil
		}
	}

	return nil, &errors.OperationError{
		Frame:    f,
		Message:  "item not found in list",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

// ListIndexOf 实现 List 类的 indexOf 方法
// 查找元素在列表中第一次出现的位置
//
// 参数:
//
//	f - 当前调用栈
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数，第一个参数是 List 实例，第二个参数是要查找的元素
//
// 返回值:
//
//	object.Object - 元素第一次出现的位置，如果没有找到则返回 -1
//	error - 可能出现的错误
func ListIndexOf(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this := args[0]
	if this == nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "method indexOf() called without instance",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	listObj, ok := this.(*object.List)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "indexOf() can only be called on List instances",
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

	target := args[1]
	for i, elem := range listObj.Elements {
		equal, err := elem.Equal(target, posStart, posEnd, f)
		if err != nil {
			return nil, err
		}
		if equal.(*object.Bool).Value {
			return &object.Int{Value: int64(i)}, nil
		}
	}

	return &object.Int{Value: -1}, nil
}

// ListReverse 实现 List 类的 reverse 方法
// 反转列表中的元素
//
// 参数:
//
//	f - 当前调用栈
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数，第一个参数是 List 实例
//
// 返回值:
//
//	object.Object - 反转后的列表
//	error - 可能出现的错误
func ListReverse(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this := args[0]
	if this == nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "method reverse() called without instance",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	listObj, ok := this.(*object.List)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "reverse() can only be called on List instances",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	for i, j := 0, len(listObj.Elements)-1; i < j; i, j = i+1, j-1 {
		listObj.Elements[i], listObj.Elements[j] = listObj.Elements[j], listObj.Elements[i]
	}

	return listObj, nil
}

// ListSort 实现 List 类的 sort 方法
// 对列表中的元素进行排序
//
// 参数:
//
//	f - 当前调用栈
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数，第一个参数是 List 实例
//
// 返回值:
//
//	object.Object - 排序后的列表
//	error - 可能出现的错误
func ListSort(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this := args[0]
	if this == nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "method sort() called without instance",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	listObj, ok := this.(*object.List)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "sort() can only be called on List instances",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	sort.Slice(listObj.Elements, func(i, j int) bool {
		switch a := listObj.Elements[i].(type) {
		case *object.Int:
			if b, ok := listObj.Elements[j].(*object.Int); ok {
				return a.Value < b.Value
			}
		case *object.Float:
			if b, ok := listObj.Elements[j].(*object.Float); ok {
				return a.Value < b.Value
			}
		case *object.String:
			if b, ok := listObj.Elements[j].(*object.String); ok {
				return a.Value < b.Value
			}
		}
		return false
	})

	return listObj, nil
}

// ListJoin 实现 List 类的 join 方法
// 使用指定的分隔符将列表中的所有元素连接成一个字符串
//
// 参数:
//
//	f - 当前调用栈
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数，第一个参数是 List 实例，第二个参数是分隔符字符串
//
// 返回值:
//
//	object.Object - 连接后的字符串
//	error - 可能出现的错误
func ListJoin(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this := args[0]
	if this == nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "method join() called without instance",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	listObj, ok := this.(*object.List)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "join() can only be called on List instances",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	if len(args) < 2 {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "join() expects 1 argument",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	sepObj, ok := args[1].(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "join() expects a String argument",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	var parts []string
	for _, elem := range listObj.Elements {
		parts = append(parts, elem.String())
	}

	result := ""
	for i, part := range parts {
		if i > 0 {
			result += sepObj.Value
		}
		result += part
	}

	return &object.String{Value: result}, nil
}

// ListSlice 实现 List 类的 slice 方法
// 返回列表中指定范围的子列表
//
// 参数:
//
//	f - 当前调用栈
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数，第一个参数是 List 实例，第二个是起始索引，第三个是结束索引
//
// 返回值:
//
//	object.Object - 子列表
//	error - 可能出现的错误
func ListSlice(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this := args[0]
	if this == nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "method slice() called without instance",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	listObj, ok := this.(*object.List)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "slice() can only be called on List instances",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	if len(args) < 3 {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  "slice() expects 2 arguments",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	startObj, ok := args[1].(*object.Int)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "slice() expects Int arguments",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	endObj, ok := args[2].(*object.Int)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "slice() expects Int arguments",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	length := int64(len(listObj.Elements))
	start := startObj.Value
	end := endObj.Value

	if start < 0 {
		start = length + start
	}
	if end < 0 {
		end = length + end
	}

	if start < 0 {
		start = 0
	}
	if end > length {
		end = length
	}
	if start > end {
		start = end
	}

	newElements := make([]object.Object, end-start)
	copy(newElements, listObj.Elements[start:end])
	return &object.List{Elements: newElements}, nil
}

// ListInit 实现 List 类的 init 构造方法
// 创建一个新的空 List 实例
//
// 参数:
//
//	f - 当前调用栈
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数，第一个参数是 List 实例
//
// 返回值:
//
//	object.Object - 新创建的空 List 实例
//	error - 可能出现的错误
func ListInit(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	return &object.List{Elements: []object.Object{}}, nil
}
