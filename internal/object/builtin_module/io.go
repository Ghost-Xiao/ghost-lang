package builtin_module

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/Ghost-Xiao/ghost-lang/internal/errors"
	"github.com/Ghost-Xiao/ghost-lang/internal/frame"
	"github.com/Ghost-Xiao/ghost-lang/internal/object"
	"github.com/Ghost-Xiao/ghost-lang/internal/util"
)

// IOModule 表示 io 内置模块
var IOModule = initIOModule()

// initIOModule 初始化 io 模块
//
// 返回值:
//
//	*object.Module - 初始化后的 io 模块
func initIOModule() *object.Module {
	env := &object.Environment{
		Name:  "io",
		Store: map[string]*object.Symbol{},
		Outer: nil,
	}

	env.Set("read", &object.Symbol{Name: "read", Value: &READ, IsConst: true})
	env.Set("write", &object.Symbol{Name: "write", Value: &WRITE, IsConst: true})
	env.Set("append", &object.Symbol{Name: "append", Value: &APPEND, IsConst: true})
	env.Set("readLines", &object.Symbol{Name: "readLines", Value: &READLINES, IsConst: true})
	env.Set("mkdir", &object.Symbol{Name: "mkdir", Value: &MKDIR, IsConst: true})
	env.Set("listDir", &object.Symbol{Name: "listDir", Value: &LISTDIR, IsConst: true})
	env.Set("exists", &object.Symbol{Name: "exists", Value: &EXISTS, IsConst: true})
	env.Set("isFile", &object.Symbol{Name: "isFile", Value: &ISFILE, IsConst: true})
	env.Set("isDir", &object.Symbol{Name: "isDir", Value: &ISDIR, IsConst: true})
	env.Set("fileSize", &object.Symbol{Name: "fileSize", Value: &FILESIZE, IsConst: true})
	env.Set("copy", &object.Symbol{Name: "copy", Value: &COPY, IsConst: true})
	env.Set("move", &object.Symbol{Name: "move", Value: &MOVE, IsConst: true})
	env.Set("remove", &object.Symbol{Name: "remove", Value: &REMOVE, IsConst: true})

	return &object.Module{
		Name: "io",
		Env:  env,
	}
}

var (
	// READ 读取文件内容函数
	READ = object.BuiltinFunction{
		Name:         "read",
		Parameter:    []string{"path"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           IORead,
	}
	// WRITE 写入文件内容函数
	WRITE = object.BuiltinFunction{
		Name:         "write",
		Parameter:    []string{"path", "content"},
		DefaultValue: []object.Object{nil, nil},
		HaveVariadic: false,
		Fn:           IOWrite,
	}
	// APPEND 追加内容到文件函数
	APPEND = object.BuiltinFunction{
		Name:         "append",
		Parameter:    []string{"path", "content"},
		DefaultValue: []object.Object{nil, nil},
		HaveVariadic: false,
		Fn:           IOAppend,
	}
	// READLINES 按行读取文件函数
	READLINES = object.BuiltinFunction{
		Name:         "readLines",
		Parameter:    []string{"path"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           IOReadLines,
	}
	// MKDIR 创建目录函数
	MKDIR = object.BuiltinFunction{
		Name:         "mkdir",
		Parameter:    []string{"path"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           IOMkdir,
	}
	// LISTDIR 列出目录内容函数
	LISTDIR = object.BuiltinFunction{
		Name:         "listDir",
		Parameter:    []string{"path"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           IOListDir,
	}
	// EXISTS 检查文件或目录是否存在函数
	EXISTS = object.BuiltinFunction{
		Name:         "exists",
		Parameter:    []string{"path"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           IOExists,
	}
	// ISFILE 检查是否为文件函数
	ISFILE = object.BuiltinFunction{
		Name:         "isFile",
		Parameter:    []string{"path"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           IOIsFile,
	}
	// ISDIR 检查是否为目录函数
	ISDIR = object.BuiltinFunction{
		Name:         "isDir",
		Parameter:    []string{"path"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           IOIsDir,
	}
	// FILESIZE 获取文件大小函数
	FILESIZE = object.BuiltinFunction{
		Name:         "fileSize",
		Parameter:    []string{"path"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           IOFileSize,
	}
	// COPY 复制文件函数
	COPY = object.BuiltinFunction{
		Name:         "copy",
		Parameter:    []string{"src", "dst"},
		DefaultValue: []object.Object{nil, nil},
		HaveVariadic: false,
		Fn:           IOCopy,
	}
	// MOVE 移动文件函数
	MOVE = object.BuiltinFunction{
		Name:         "move",
		Parameter:    []string{"src", "dst"},
		DefaultValue: []object.Object{nil, nil},
		HaveVariadic: false,
		Fn:           IOMove,
	}
	// REMOVE 删除文件或目录函数
	REMOVE = object.BuiltinFunction{
		Name:         "remove",
		Parameter:    []string{"path"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           IORemove,
	}
)

// IORead 实现 read 函数
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
//	object.Object - 函数返回值
//	error - 可能出现的错误
func IORead(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	pathArg, ok := args[0].(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "path must be a string.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	content, err := os.ReadFile(pathArg.Value)
	if err != nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  fmt.Sprintf("failed to read file: %s", err.Error()),
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	return &object.String{Value: string(content)}, nil
}

// IOWrite 实现 write 函数
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
//	object.Object - 空结果
//	error - 可能出现的错误
func IOWrite(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	pathArg, ok := args[0].(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "path must be a string.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	contentArg, ok := args[1].(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "content must be a string.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	err := os.WriteFile(pathArg.Value, []byte(contentArg.Value), 0644)
	if err != nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  fmt.Sprintf("failed to write file: %s", err.Error()),
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	return &object.Null{}, nil
}

// IOAppend 实现 append 函数
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
//	object.Object - 空结果
//	error - 可能出现的错误
func IOAppend(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	pathArg, ok := args[0].(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "path must be a string.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	contentArg, ok := args[1].(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "content must be a string.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	file, err := os.OpenFile(pathArg.Value, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  fmt.Sprintf("failed to open file: %s", err.Error()),
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	defer file.Close()

	_, err = file.WriteString(contentArg.Value)
	if err != nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  fmt.Sprintf("failed to append to file: %s", err.Error()),
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	return &object.Null{}, nil
}

// IOReadLines 实现 readLines 函数
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
//	object.Object - 字符串列表
//	error - 可能出现的错误
func IOReadLines(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	pathArg, ok := args[0].(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "path must be a string.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	content, err := os.ReadFile(pathArg.Value)
	if err != nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  fmt.Sprintf("failed to read file: %s", err.Error()),
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	lines := strings.Split(string(content), "\n")
	elements := make([]object.Object, 0, len(lines))
	for _, line := range lines {
		elements = append(elements, &object.String{Value: line})
	}

	return &object.List{Elements: elements}, nil
}

// IOMkdir 实现 mkdir 函数
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
//	object.Object - 空结果
//	error - 可能出现的错误
func IOMkdir(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	pathArg, ok := args[0].(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "path must be a string.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	err := os.MkdirAll(pathArg.Value, 0755)
	if err != nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  fmt.Sprintf("failed to create directory: %s", err.Error()),
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	return &object.Null{}, nil
}

// IOListDir 实现 listDir 函数
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
//	object.Object - 文件列表
//	error - 可能出现的错误
func IOListDir(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	pathArg, ok := args[0].(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "path must be a string.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	entries, err := os.ReadDir(pathArg.Value)
	if err != nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  fmt.Sprintf("failed to list directory: %s", err.Error()),
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	elements := make([]object.Object, 0, len(entries))
	for _, entry := range entries {
		elements = append(elements, &object.String{Value: entry.Name()})
	}

	return &object.List{Elements: elements}, nil
}

// IOExists 实现 exists 函数
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
//	object.Object - 布尔值
//	error - 可能出现的错误
func IOExists(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	pathArg, ok := args[0].(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "path must be a string.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	_, err := os.Stat(pathArg.Value)
	if err != nil {
		if os.IsNotExist(err) {
			return &object.Bool{Value: false}, nil
		}
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  fmt.Sprintf("failed to check path: %s", err.Error()),
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	return &object.Bool{Value: true}, nil
}

// IOIsFile 实现 isFile 函数
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
//	object.Object - 布尔值
//	error - 可能出现的错误
func IOIsFile(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	pathArg, ok := args[0].(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "path must be a string.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	info, err := os.Stat(pathArg.Value)
	if err != nil {
		if os.IsNotExist(err) {
			return &object.Bool{Value: false}, nil
		}
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  fmt.Sprintf("failed to check path: %s", err.Error()),
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	return &object.Bool{Value: !info.IsDir()}, nil
}

// IOIsDir 实现 isDir 函数
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
//	object.Object - 布尔值
//	error - 可能出现的错误
func IOIsDir(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	pathArg, ok := args[0].(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "path must be a string.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	info, err := os.Stat(pathArg.Value)
	if err != nil {
		if os.IsNotExist(err) {
			return &object.Bool{Value: false}, nil
		}
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  fmt.Sprintf("failed to check path: %s", err.Error()),
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	return &object.Bool{Value: info.IsDir()}, nil
}

// IOFileSize 实现 fileSize 函数
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
//	object.Object - 文件大小
//	error - 可能出现的错误
func IOFileSize(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	pathArg, ok := args[0].(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "path must be a string.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	info, err := os.Stat(pathArg.Value)
	if err != nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  fmt.Sprintf("failed to get file size: %s", err.Error()),
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	return &object.Int{Value: info.Size()}, nil
}

// IOCopy 实现 copy 函数
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
//	object.Object - 空结果
//	error - 可能出现的错误
func IOCopy(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	srcArg, ok := args[0].(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "src must be a string.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	dstArg, ok := args[1].(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "dst must be a string.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	srcFile, err := os.Open(srcArg.Value)
	if err != nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  fmt.Sprintf("failed to open source file: %s", err.Error()),
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dstArg.Value)
	if err != nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  fmt.Sprintf("failed to create destination file: %s", err.Error()),
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  fmt.Sprintf("failed to copy file: %s", err.Error()),
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	return &object.Null{}, nil
}

// IOMove 实现 move 函数
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
//	object.Object - 空结果
//	error - 可能出现的错误
func IOMove(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	srcArg, ok := args[0].(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "src must be a string.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	dstArg, ok := args[1].(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "dst must be a string.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	err := os.Rename(srcArg.Value, dstArg.Value)
	if err != nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  fmt.Sprintf("failed to move file: %s", err.Error()),
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	return &object.Null{}, nil
}

// removeAll 递归删除文件或目录
//
// 参数:
//
//	path - 文件或目录路径
//
// 返回值:
//
//	error - 可能出现的错误
func removeAll(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	if !info.IsDir() {
		return os.Remove(path)
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		entryPath := filepath.Join(path, entry.Name())
		if err := removeAll(entryPath); err != nil {
			return err
		}
	}

	return os.Remove(path)
}

// IORemove 实现 remove 函数
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
//	object.Object - 空结果
//	error - 可能出现的错误
func IORemove(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	pathArg, ok := args[0].(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "path must be a string.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	err := removeAll(pathArg.Value)
	if err != nil {
		return nil, &errors.OperationError{
			Frame:    f,
			Message:  fmt.Sprintf("failed to remove: %s", err.Error()),
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	return &object.Null{}, nil
}
