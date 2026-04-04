package module

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/Ghost-Xiao/ghost-lang/internal/errors"
	"github.com/Ghost-Xiao/ghost-lang/internal/frame"
	"github.com/Ghost-Xiao/ghost-lang/internal/object"
	"github.com/Ghost-Xiao/ghost-lang/internal/util"
)

type IO struct{}

func (io *IO) Name() string {
	return "io"
}

func (io *IO) Load() *object.Environment {
	env := &object.Environment{
		Name:  io.Name(),
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

	return env
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

func IORead(f *frame.Frame, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	path, ok := args[0].(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "path must be a string.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	content, err := os.ReadFile(path.Value)
	if err != nil {
		return nil, &errors.ArgumentError{
			Frame:    f,
			Message:  fmt.Sprintf("failed to read file: %v", err),
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	return &object.String{Value: string(content)}, nil
}

func IOWrite(f *frame.Frame, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	path, ok := args[0].(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "path must be a string.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	content, ok := args[1].(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "content must be a string.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	err := os.WriteFile(path.Value, []byte(content.Value), 0644)
	if err != nil {
		return nil, &errors.ArgumentError{
			Frame:    f,
			Message:  fmt.Sprintf("failed to write file: %v", err),
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	return &object.Null{}, nil
}

func IOAppend(f *frame.Frame, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	path, ok := args[0].(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "path must be a string.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	content, ok := args[1].(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "content must be a string.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	file, err := os.OpenFile(path.Value, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, &errors.ArgumentError{
			Frame:    f,
			Message:  fmt.Sprintf("failed to open file: %v", err),
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	defer file.Close()

	_, err = file.WriteString(content.Value)
	if err != nil {
		return nil, &errors.ArgumentError{
			Frame:    f,
			Message:  fmt.Sprintf("failed to append to file: %v", err),
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	return &object.Null{}, nil
}

func IOReadLines(f *frame.Frame, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	path, ok := args[0].(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "path must be a string.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	content, err := os.ReadFile(path.Value)
	if err != nil {
		return nil, &errors.ArgumentError{
			Frame:    f,
			Message:  fmt.Sprintf("failed to read file: %v", err),
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	lines := strings.Split(string(content), "\n")
	result := make([]object.Object, 0, len(lines))
	for _, line := range lines {
		result = append(result, &object.String{Value: line})
	}

	return &object.List{Elements: result}, nil
}

func IOMkdir(f *frame.Frame, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	path, ok := args[0].(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "path must be a string.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	err := os.MkdirAll(path.Value, 0755)
	if err != nil {
		return nil, &errors.ArgumentError{
			Frame:    f,
			Message:  fmt.Sprintf("failed to create directory: %v", err),
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	return &object.Null{}, nil
}

func IOListDir(f *frame.Frame, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	path, ok := args[0].(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "path must be a string.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	entries, err := os.ReadDir(path.Value)
	if err != nil {
		return nil, &errors.ArgumentError{
			Frame:    f,
			Message:  fmt.Sprintf("failed to list directory: %v", err),
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	result := make([]object.Object, 0, len(entries))
	for _, entry := range entries {
		result = append(result, &object.String{Value: entry.Name()})
	}

	return &object.List{Elements: result}, nil
}

func IOExists(f *frame.Frame, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	path, ok := args[0].(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "path must be a string.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	_, err := os.Stat(path.Value)
	if os.IsNotExist(err) {
		return &object.Bool{Value: false}, nil
	}
	if err != nil {
		return nil, &errors.ArgumentError{
			Frame:    f,
			Message:  fmt.Sprintf("failed to check path: %v", err),
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	return &object.Bool{Value: true}, nil
}

func IOIsFile(f *frame.Frame, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	path, ok := args[0].(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "path must be a string.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	info, err := os.Stat(path.Value)
	if err != nil {
		if os.IsNotExist(err) {
			return &object.Bool{Value: false}, nil
		}
		return nil, &errors.ArgumentError{
			Frame:    f,
			Message:  fmt.Sprintf("failed to check path: %v", err),
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	return &object.Bool{Value: !info.IsDir()}, nil
}

func IOIsDir(f *frame.Frame, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	path, ok := args[0].(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "path must be a string.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	info, err := os.Stat(path.Value)
	if err != nil {
		if os.IsNotExist(err) {
			return &object.Bool{Value: false}, nil
		}
		return nil, &errors.ArgumentError{
			Frame:    f,
			Message:  fmt.Sprintf("failed to check path: %v", err),
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	return &object.Bool{Value: info.IsDir()}, nil
}

func IOFileSize(f *frame.Frame, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	path, ok := args[0].(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "path must be a string.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	info, err := os.Stat(path.Value)
	if err != nil {
		return nil, &errors.ArgumentError{
			Frame:    f,
			Message:  fmt.Sprintf("failed to get file size: %v", err),
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	return &object.Int{Value: info.Size()}, nil
}

func IOCopy(f *frame.Frame, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	src, ok := args[0].(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "src must be a string.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	dst, ok := args[1].(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "dst must be a string.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	srcFile, err := os.Open(src.Value)
	if err != nil {
		return nil, &errors.ArgumentError{
			Frame:    f,
			Message:  fmt.Sprintf("failed to open source file: %v", err),
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst.Value)
	if err != nil {
		return nil, &errors.ArgumentError{
			Frame:    f,
			Message:  fmt.Sprintf("failed to create destination file: %v", err),
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return nil, &errors.ArgumentError{
			Frame:    f,
			Message:  fmt.Sprintf("failed to copy file: %v", err),
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	return &object.Null{}, nil
}

func IOMove(f *frame.Frame, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	src, ok := args[0].(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "src must be a string.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	dst, ok := args[1].(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "dst must be a string.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	err := os.Rename(src.Value, dst.Value)
	if err != nil {
		return nil, &errors.ArgumentError{
			Frame:    f,
			Message:  fmt.Sprintf("failed to move file: %v", err),
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	return &object.Null{}, nil
}

func IORemove(f *frame.Frame, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	path, ok := args[0].(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "path must be a string.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	err := os.RemoveAll(path.Value)
	if err != nil {
		return nil, &errors.ArgumentError{
			Frame:    f,
			Message:  fmt.Sprintf("failed to remove path: %v", err),
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	return &object.Null{}, nil
}

func (io *IO) Type() string {
	return "Module"
}

func (io *IO) String() string {
	return fmt.Sprintf("module \"%s\"", io.Name())
}

func (io *IO) Negative(posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"-\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

func (io *IO) BitNot(posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"~\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

func (io *IO) Not(posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"!\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

func (io *IO) Add(_ object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"+\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

func (io *IO) Subtract(_ object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"-\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

func (io *IO) Multiply(_ object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"*\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

func (io *IO) Divide(_ object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"/\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

func (io *IO) Mod(_ object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"%\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

func (io *IO) Equal(other object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"==\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

func (io *IO) NotEqual(other object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"!=\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

func (io *IO) LessThan(_ object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"<\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

func (io *IO) GreaterThan(_ object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \">\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

func (io *IO) LessThanOrEqual(_ object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"<=\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

func (io *IO) GreaterThanOrEqual(_ object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \">=\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

func (io *IO) BitAnd(_ object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"&\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

func (io *IO) BitOr(_ object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"|\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

func (io *IO) Xor(_ object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"^\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

func (io *IO) LeftShift(_ object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"<<\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

func (io *IO) RightShift(_ object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \">>\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

func (io *IO) And(_ object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"&&\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

func (io *IO) Or(_ object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"||\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

func (io *IO) Index(other object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.TypeError{
		Frame:    frame,
		Message:  "index expression not supported for this type.",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}
