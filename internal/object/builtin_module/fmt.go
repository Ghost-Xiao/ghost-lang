package builtin_module

import (
	"fmt"

	"github.com/Ghost-Xiao/ghost-lang/internal/errors"
	"github.com/Ghost-Xiao/ghost-lang/internal/frame"
	"github.com/Ghost-Xiao/ghost-lang/internal/object"
	"github.com/Ghost-Xiao/ghost-lang/internal/util"
)

var FmtModule = initFmtModule()

func initFmtModule() *object.Module {
	env := &object.Environment{
		Name:  "fmt",
		Store: map[string]*object.Symbol{},
		Outer: nil,
	}

	env.Set("print", &object.Symbol{Name: "print", Value: &PRINT, IsConst: true})
	env.Set("println", &object.Symbol{Name: "println", Value: &PRINTLN, IsConst: true})
	env.Set("printf", &object.Symbol{Name: "printf", Value: &PRINTF, IsConst: true})
	env.Set("sprint", &object.Symbol{Name: "sprint", Value: &SPRINT, IsConst: true})
	env.Set("sprintf", &object.Symbol{Name: "sprintf", Value: &SPRINTF, IsConst: true})

	return &object.Module{
		Name: "fmt",
		Env:  env,
	}
}

var (
	PRINT = object.BuiltinFunction{
		Name:         "print",
		Parameter:    []string{"x"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: true,
		Fn:           FmtPrint,
	}
	PRINTLN = object.BuiltinFunction{
		Name:         "println",
		Parameter:    []string{"x"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: true,
		Fn:           FmtPrintln,
	}
	PRINTF = object.BuiltinFunction{
		Name:         "printf",
		Parameter:    []string{"format", "a"},
		DefaultValue: []object.Object{nil, nil},
		HaveVariadic: true,
		Fn:           FmtPrintf,
	}
	SPRINT = object.BuiltinFunction{
		Name:         "sprint",
		Parameter:    []string{"x"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: true,
		Fn:           FmtSprint,
	}
	SPRINTF = object.BuiltinFunction{
		Name:         "sprintf",
		Parameter:    []string{"format", "a"},
		DefaultValue: []object.Object{nil, nil},
		HaveVariadic: true,
		Fn:           FmtSprintf,
	}
)

func FmtPrint(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	for _, arg := range args[0].(*object.List).Elements {
		switch v := arg.(type) {
		case *object.String:
			fmt.Print(v.Value)
		case *object.Int:
			fmt.Print(v.Value)
		case *object.Float:
			fmt.Print(v.Value)
		case *object.Bool:
			fmt.Print(v.Value)
		case *object.Null:
			fmt.Print("null")
		default:
			fmt.Print(arg.String())
		}
	}
	return &object.Null{}, nil
}

func FmtPrintln(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	for i, arg := range args[0].(*object.List).Elements {
		if i > 0 {
			fmt.Print(" ")
		}
		switch v := arg.(type) {
		case *object.String:
			fmt.Print(v.Value)
		case *object.Int:
			fmt.Print(v.Value)
		case *object.Float:
			fmt.Print(v.Value)
		case *object.Bool:
			fmt.Print(v.Value)
		case *object.Null:
			fmt.Print("null")
		default:
			fmt.Print(arg.String())
		}
	}
	fmt.Println()
	return &object.Null{}, nil
}

func FmtPrintf(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	if len(args) == 0 {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "printf requires at least one argument.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	format, ok := args[0].(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "printf() first argument must be a string.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	formatStr := format.Value
	var goArgs []interface{}
	for i := 0; i < len(args[1].(*object.List).Elements); i++ {
		arg := args[1].(*object.List).Elements[i]
		switch v := arg.(type) {
		case *object.String:
			goArgs = append(goArgs, v.Value)
		case *object.Int:
			goArgs = append(goArgs, v.Value)
		case *object.Float:
			goArgs = append(goArgs, v.Value)
		case *object.Bool:
			goArgs = append(goArgs, v.Value)
		case *object.Null:
			goArgs = append(goArgs, "null")
		default:
			goArgs = append(goArgs, arg.String())
		}
	}

	fmt.Printf(formatStr, goArgs...)
	return &object.Null{}, nil
}

func FmtSprint(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	var result string
	for _, arg := range args[0].(*object.List).Elements {
		switch v := arg.(type) {
		case *object.String:
			result += v.Value
		case *object.Int:
			result += fmt.Sprintf("%d", v.Value)
		case *object.Float:
			result += fmt.Sprintf("%g", v.Value)
		case *object.Bool:
			result += fmt.Sprintf("%t", v.Value)
		case *object.Null:
			result += "null"
		default:
			result += arg.String()
		}
	}
	return &object.String{Value: result}, nil
}

func FmtSprintf(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	if len(args) == 0 {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "sprintf requires at least one argument.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	format, ok := args[0].(*object.String)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "sprintf() first argument must be a string.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	formatStr := format.Value
	var goArgs []interface{}
	for i := 0; i < len(args[1].(*object.List).Elements); i++ {
		arg := args[1].(*object.List).Elements[i]
		switch v := arg.(type) {
		case *object.String:
			goArgs = append(goArgs, v.Value)
		case *object.Int:
			goArgs = append(goArgs, v.Value)
		case *object.Float:
			goArgs = append(goArgs, v.Value)
		case *object.Bool:
			goArgs = append(goArgs, v.Value)
		case *object.Null:
			goArgs = append(goArgs, "null")
		default:
			goArgs = append(goArgs, arg.String())
		}
	}

	result := fmt.Sprintf(formatStr, goArgs...)
	return &object.String{Value: result}, nil
}
