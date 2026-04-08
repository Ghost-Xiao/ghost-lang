package module

import (
	"fmt"
	"strings"

	"github.com/Ghost-Xiao/ghost-lang/internal/errors"
	"github.com/Ghost-Xiao/ghost-lang/internal/frame"
	"github.com/Ghost-Xiao/ghost-lang/internal/object"
	"github.com/Ghost-Xiao/ghost-lang/internal/util"
)

type Fmt struct{}

func (f *Fmt) Name() string {
	return "fmt"
}

func (f *Fmt) Load() *object.Environment {
	env := &object.Environment{
		Name:  f.Name(),
		Store: map[string]*object.Symbol{},
		Outer: nil,
	}

	env.Set("print", &object.Symbol{Name: "print", Value: &PRINT, IsConst: true})
	env.Set("println", &object.Symbol{Name: "println", Value: &PRINTLN, IsConst: true})
	env.Set("sprintf", &object.Symbol{Name: "sprintf", Value: &SPRINTF, IsConst: true})
	env.Set("printf", &object.Symbol{Name: "printf", Value: &PRINTF, IsConst: true})

	return env
}

var (
	PRINT = object.BuiltinFunction{
		Name:         "print",
		Parameter:    []string{"a"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: true,
		Fn:           FmtPrint,
	}
	PRINTLN = object.BuiltinFunction{
		Name:         "println",
		Parameter:    []string{"a"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: true,
		Fn:           FmtPrintln,
	}
	SPRINTF = object.BuiltinFunction{
		Name:         "sprintf",
		Parameter:    []string{"format", "a"},
		DefaultValue: []object.Object{nil, nil},
		HaveVariadic: true,
		Fn:           FmtSprintf,
	}
	PRINTF = object.BuiltinFunction{
		Name:         "printf",
		Parameter:    []string{"format", "a"},
		DefaultValue: []object.Object{nil, nil},
		HaveVariadic: true,
		Fn:           FmtPrintf,
	}
)

func formatValue(obj object.Object) string {
	switch v := obj.(type) {
	case *object.Int:
		return fmt.Sprintf("%d", v.Value)
	case *object.Float:
		return fmt.Sprintf("%g", v.Value)
	case *object.String:
		return v.Value
	case *object.Bool:
		if v.Value {
			return "true"
		}
		return "false"
	case *object.Null:
		return "null"
	case *object.List:
		var sb strings.Builder
		sb.WriteString("[")
		for i, elem := range v.Elements {
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(formatValue(elem))
		}
		sb.WriteString("]")
		return sb.String()
	case *object.Map:
		var sb strings.Builder
		sb.WriteString("{")
		first := true
		for _, pair := range v.Pairs {
			if !first {
				sb.WriteString(", ")
			}
			first = false
			sb.WriteString(formatValue(pair.Key))
			sb.WriteString(": ")
			sb.WriteString(formatValue(pair.Value))
		}
		sb.WriteString("}")
		return sb.String()
	default:
		return obj.String()
	}
}

func formatValueWithVerb(obj object.Object, verb rune) (string, error) {
	switch verb {
	case 'v':
		return formatValue(obj), nil
	case 'd':
		switch v := obj.(type) {
		case *object.Int:
			return fmt.Sprintf("%d", v.Value), nil
		case *object.Float:
			return fmt.Sprintf("%.0f", v.Value), nil
		default:
			return "", fmt.Errorf("expected number for %%d")
		}
	case 'f':
		switch v := obj.(type) {
		case *object.Int:
			return fmt.Sprintf("%f", float64(v.Value)), nil
		case *object.Float:
			return fmt.Sprintf("%f", v.Value), nil
		default:
			return "", fmt.Errorf("expected number for %%f")
		}
	case 's':
		switch v := obj.(type) {
		case *object.String:
			return v.Value, nil
		default:
			return "", fmt.Errorf("expected string for %%s")
		}
	case 't':
		switch v := obj.(type) {
		case *object.Bool:
			if v.Value {
				return "true", nil
			}
			return "false", nil
		default:
			return "", fmt.Errorf("expected bool for %%t")
		}
	case 'b':
		switch v := obj.(type) {
		case *object.Int:
			return fmt.Sprintf("%b", v.Value), nil
		case *object.Float:
			return "", fmt.Errorf("expected integer for %%b")
		default:
			return "", fmt.Errorf("expected integer for %%b")
		}
	case 'x':
		switch v := obj.(type) {
		case *object.Int:
			return fmt.Sprintf("%x", v.Value), nil
		case *object.Float:
			return "", fmt.Errorf("expected integer for %%x")
		default:
			return "", fmt.Errorf("expected integer for %%x")
		}
	case '%':
		return "%", nil
	default:
		return "", fmt.Errorf("unknown verb '%%%c'", verb)
	}
}

func FmtSprintf(f *frame.Frame, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	if len(args) == 0 {
		return nil, &errors.ArgumentError{
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
			Message:  "format must be a string.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	sb := strings.Builder{}
	argIndex := 1
	i := 0
	formatStr := format.Value

	for i < len(formatStr) {
		if formatStr[i] == '%' {
			if i+1 >= len(formatStr) {
				return nil, &errors.ArgumentError{
					Frame:    f,
					Message:  "invalid format: trailing %",
					PosStart: posStart,
					PosEnd:   posEnd,
				}
			}
			verb := rune(formatStr[i+1])
			if verb == '%' {
				sb.WriteRune('%')
				i += 2
				continue
			}

			if argIndex >= len(args) {
				return nil, &errors.ArgumentError{
					Frame:    f,
					Message:  "not enough arguments for format string",
					PosStart: posStart,
					PosEnd:   posEnd,
				}
			}

			formatted, err := formatValueWithVerb(args[argIndex], verb)
			if err != nil {
				return nil, &errors.ArgumentError{
					Frame:    f,
					Message:  err.Error(),
					PosStart: posStart,
					PosEnd:   posEnd,
				}
			}
			sb.WriteString(formatted)
			argIndex++
			i += 2
		} else {
			sb.WriteByte(formatStr[i])
			i++
		}
	}

	return &object.String{Value: sb.String()}, nil
}

func FmtPrint(f *frame.Frame, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	for i, arg := range args {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(formatValue(arg))
	}
	return &object.Null{}, nil
}

func FmtPrintln(f *frame.Frame, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	for i, arg := range args {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(formatValue(arg))
	}
	fmt.Println()
	return &object.Null{}, nil
}

func FmtPrintf(f *frame.Frame, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	result, err := FmtSprintf(f, posStart, posEnd, args...)
	if err != nil {
		return nil, err
	}
	str, ok := result.(*object.String)
	if ok {
		fmt.Print(str.Value)
	}
	return &object.Null{}, nil
}

func (f *Fmt) Type() string {
	return "Module"
}

func (f *Fmt) String() string {
	return fmt.Sprintf("module \"%s\"", f.Name())
}

func (f *Fmt) Negative(posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"-\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

func (f *Fmt) BitNot(posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"~\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

func (f *Fmt) Not(posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"!\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

func (f *Fmt) Add(_ object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"+\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

func (f *Fmt) Subtract(_ object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"-\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

func (f *Fmt) Multiply(_ object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"*\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

func (f *Fmt) Divide(_ object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"/\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

func (f *Fmt) Mod(_ object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"%\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

func (f *Fmt) Equal(other object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"==\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

func (f *Fmt) NotEqual(other object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"!=\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

func (f *Fmt) LessThan(_ object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"<\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

func (f *Fmt) GreaterThan(_ object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \">\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

func (f *Fmt) LessThanOrEqual(_ object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"<=\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

func (f *Fmt) GreaterThanOrEqual(_ object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \">=\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

func (f *Fmt) BitAnd(_ object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"&\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

func (f *Fmt) BitOr(_ object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"|\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

func (f *Fmt) Xor(_ object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"^\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

func (f *Fmt) LeftShift(_ object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"<<\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

func (f *Fmt) RightShift(_ object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \">>\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

func (f *Fmt) And(_ object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"&&\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

func (f *Fmt) Or(_ object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.OperationError{
		Frame:    frame,
		Message:  "invalid operation \"||\".",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}

func (f *Fmt) Index(other object.Object, posStart, posEnd *util.Pos, frame *frame.Frame) (object.Object, error) {
	return nil, &errors.TypeError{
		Frame:    frame,
		Message:  "index expression not supported for this type.",
		PosStart: posStart,
		PosEnd:   posEnd,
	}
}
