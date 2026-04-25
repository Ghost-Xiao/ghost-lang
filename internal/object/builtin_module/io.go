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

var IOModule = initIOModule()

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
	READ = object.BuiltinFunction{
		Name:         "read",
		Parameter:    []string{"path"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           IORead,
	}
	WRITE = object.BuiltinFunction{
		Name:         "write",
		Parameter:    []string{"path", "content"},
		DefaultValue: []object.Object{nil, nil},
		HaveVariadic: false,
		Fn:           IOWrite,
	}
	APPEND = object.BuiltinFunction{
		Name:         "append",
		Parameter:    []string{"path", "content"},
		DefaultValue: []object.Object{nil, nil},
		HaveVariadic: false,
		Fn:           IOAppend,
	}
	READLINES = object.BuiltinFunction{
		Name:         "readLines",
		Parameter:    []string{"path"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           IOReadLines,
	}
	MKDIR = object.BuiltinFunction{
		Name:         "mkdir",
		Parameter:    []string{"path"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           IOMkdir,
	}
	LISTDIR = object.BuiltinFunction{
		Name:         "listDir",
		Parameter:    []string{"path"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           IOListDir,
	}
	EXISTS = object.BuiltinFunction{
		Name:         "exists",
		Parameter:    []string{"path"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           IOExists,
	}
	ISFILE = object.BuiltinFunction{
		Name:         "isFile",
		Parameter:    []string{"path"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           IOIsFile,
	}
	ISDIR = object.BuiltinFunction{
		Name:         "isDir",
		Parameter:    []string{"path"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           IOIsDir,
	}
	FILESIZE = object.BuiltinFunction{
		Name:         "fileSize",
		Parameter:    []string{"path"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           IOFileSize,
	}
	COPY = object.BuiltinFunction{
		Name:         "copy",
		Parameter:    []string{"src", "dst"},
		DefaultValue: []object.Object{nil, nil},
		HaveVariadic: false,
		Fn:           IOCopy,
	}
	MOVE = object.BuiltinFunction{
		Name:         "move",
		Parameter:    []string{"src", "dst"},
		DefaultValue: []object.Object{nil, nil},
		HaveVariadic: false,
		Fn:           IOMove,
	}
	REMOVE = object.BuiltinFunction{
		Name:         "remove",
		Parameter:    []string{"path"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           IORemove,
	}
)

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
