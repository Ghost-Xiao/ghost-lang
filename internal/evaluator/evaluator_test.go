package evaluator

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Ghost-Xiao/ghost-lang/internal/frame"
	"github.com/Ghost-Xiao/ghost-lang/internal/lexer"
	"github.com/Ghost-Xiao/ghost-lang/internal/object"
	"github.com/Ghost-Xiao/ghost-lang/internal/parser"
	"github.com/Ghost-Xiao/ghost-lang/internal/parser/ast"
	"github.com/Ghost-Xiao/ghost-lang/internal/util"
)

func TestEvaluator_Visit(t *testing.T) {
	env := &object.Environment{
		Store: make(map[string]*object.Symbol),
		Outer: nil,
	}
	f := &frame.Frame{
		FuncName: "<test>",
		Parent:   nil,
		PosStart: nil,
		PosEnd:   nil,
	}

	tests := []struct {
		name     string
		input    string
		excepted object.Object
	}{
		{
			name: "Program",
			input: `1;
true;`,
			excepted: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer("<test>", tt.input)
			p, _ := parser.NewParser(l)
			program := p.ParseProgram()
			e := NewEvaluator(f)
			val := e.Eval(program, env)
			if !reflect.DeepEqual(val, tt.excepted) {
				t.Errorf("excepted %+v, got %+v", tt.excepted, val)
			}
		})
	}
}

func TestEvaluator_VisitProgram(t *testing.T) {
	env := &object.Environment{
		Store: make(map[string]*object.Symbol),
		Outer: nil,
	}
	f := &frame.Frame{
		FuncName: "<test>",
		Parent:   nil,
		PosStart: nil,
		PosEnd:   nil,
	}

	tests := []struct {
		name     string
		input    string
		excepted object.Object
	}{
		{
			name: "Program",
			input: `1;
true;`,
			excepted: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer("<test>", tt.input)
			p, _ := parser.NewParser(l)
			program := p.ParseProgram()
			e := NewEvaluator(f)
			val := e.evalProgram(program, env)
			if !reflect.DeepEqual(val, tt.excepted) {
				t.Errorf("excepted %+v, got %+v", tt.excepted, val)
			}
		})
	}
}

func TestEvaluator_VisitIntExpression(t *testing.T) {
	env := &object.Environment{
		Store: make(map[string]*object.Symbol),
		Outer: nil,
	}
	f := &frame.Frame{
		FuncName: "<test>",
		Parent:   nil,
		PosStart: nil,
		PosEnd:   nil,
	}

	tests := []struct {
		name     string
		input    string
		excepted object.Object
	}{
		{
			name:  "Int Expression",
			input: `1;`,
			excepted: &object.Int{
				Value: 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer("<test>", tt.input)
			p, _ := parser.NewParser(l)
			program := p.ParseProgram()
			e := NewEvaluator(f)
			val := e.evalIntExpression(program.Statements[0].(*ast.ExpressionStatement).Expr.(*ast.IntExpression), env)
			if !reflect.DeepEqual(val, tt.excepted) {
				t.Errorf("excepted %+v, got %+v", tt.excepted, val)
			}
		})
	}
}

func TestEvaluator_VisitFloatExpression(t *testing.T) {
	env := &object.Environment{
		Store: make(map[string]*object.Symbol),
		Outer: nil,
	}
	f := &frame.Frame{
		FuncName: "<test>",
		Parent:   nil,
		PosStart: nil,
		PosEnd:   nil,
	}

	tests := []struct {
		name     string
		input    string
		excepted object.Object
	}{
		{
			name:  "Float Expression",
			input: `1.0;`,
			excepted: &object.Float{
				Value: 1.0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer("<test>", tt.input)
			p, _ := parser.NewParser(l)
			program := p.ParseProgram()
			e := NewEvaluator(f)
			val := e.evalFloatExpression(program.Statements[0].(*ast.ExpressionStatement).Expr.(*ast.FloatExpression), env)
			if !reflect.DeepEqual(val, tt.excepted) {
				t.Errorf("excepted %+v, got %+v", tt.excepted, val)
			}
		})
	}
}

func TestEvaluator_VisitBooleanExpression(t *testing.T) {
	env := &object.Environment{
		Store: make(map[string]*object.Symbol),
		Outer: nil,
	}
	f := &frame.Frame{
		FuncName: "<test>",
		Parent:   nil,
		PosStart: nil,
		PosEnd:   nil,
	}

	tests := []struct {
		name     string
		input    string
		excepted object.Object
	}{
		{
			name:  "True Expression",
			input: `true;`,
			excepted: &object.Bool{
				Value: true,
			},
		},
		{
			name:  "False Expression",
			input: `false;`,
			excepted: &object.Bool{
				Value: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer("<test>", tt.input)
			p, _ := parser.NewParser(l)
			program := p.ParseProgram()
			e := NewEvaluator(f)
			val := e.evalBooleanExpression(program.Statements[0].(*ast.ExpressionStatement).Expr.(*ast.BoolExpression), env)
			if !reflect.DeepEqual(val, tt.excepted) {
				t.Errorf("excepted %+v, got %+v", tt.excepted, val)
			}
		})
	}
}

func TestEvaluator_VisitNullExpression(t *testing.T) {
	env := &object.Environment{
		Store: make(map[string]*object.Symbol),
		Outer: nil,
	}
	f := &frame.Frame{
		FuncName: "<test>",
		Parent:   nil,
		PosStart: nil,
		PosEnd:   nil,
	}

	tests := []struct {
		name     string
		input    string
		excepted object.Object
	}{
		{
			name:     "Null Expression",
			input:    `null;`,
			excepted: &object.Null{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer("<test>", tt.input)
			p, _ := parser.NewParser(l)
			program := p.ParseProgram()
			e := NewEvaluator(f)
			val := e.evalNullExpression(program.Statements[0].(*ast.ExpressionStatement).Expr.(*ast.NullExpression), env)
			if !reflect.DeepEqual(val, tt.excepted) {
				t.Errorf("excepted %+v, got %+v", tt.excepted, val)
			}
		})
	}
}

func TestEvaluator_VisitStringExpression(t *testing.T) {
	env := &object.Environment{
		Store: make(map[string]*object.Symbol),
		Outer: nil,
	}
	f := &frame.Frame{
		FuncName: "<test>",
		Parent:   nil,
		PosStart: nil,
		PosEnd:   nil,
	}

	tests := []struct {
		name     string
		input    string
		excepted object.Object
	}{
		{
			name:  "String Expression",
			input: `"Hello World";`,
			excepted: &object.String{
				Value: "Hello World",
			},
		},
		{
			name:  "String Expression with Escape",
			input: `"Hello\nWorld";`,
			excepted: &object.String{
				Value: "Hello\nWorld",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer("<test>", tt.input)
			p, _ := parser.NewParser(l)
			program := p.ParseProgram()
			e := NewEvaluator(f)
			val := e.evalStringExpression(program.Statements[0].(*ast.ExpressionStatement).Expr.(*ast.StringExpression), env)
			if !reflect.DeepEqual(val, tt.excepted) {
				t.Errorf("excepted %+v, got %+v", tt.excepted, val)
			}
		})
	}
}

func TestEvaluator_VisitListExpression(t *testing.T) {
	env := &object.Environment{
		Store: make(map[string]*object.Symbol),
		Outer: nil,
	}
	f := &frame.Frame{
		FuncName: "<test>",
		Parent:   nil,
		PosStart: nil,
		PosEnd:   nil,
	}

	tests := []struct {
		name     string
		input    string
		excepted object.Object
	}{
		{
			name:  "Empty List Expression",
			input: `[];`,
			excepted: &object.List{
				Elements: make([]object.Object, 0),
			},
		},
		{
			name:  "Single Integer Element List",
			input: `[1];`,
			excepted: &object.List{
				Elements: []object.Object{
					&object.Int{Value: 1},
				},
			},
		},
		{
			name:  "Multiple Integer Elements List",
			input: `[1, 2, 3];`,
			excepted: &object.List{
				Elements: []object.Object{
					&object.Int{Value: 1},
					&object.Int{Value: 2},
					&object.Int{Value: 3},
				},
			},
		},
		{
			name:  "Multiple String Elements List",
			input: `["hello", "world"];`,
			excepted: &object.List{
				Elements: []object.Object{
					&object.String{Value: "hello"},
					&object.String{Value: "world"},
				},
			},
		},
		{
			name:  "Boolean Elements List",
			input: `[true, false, true];`,
			excepted: &object.List{
				Elements: []object.Object{
					&object.Bool{Value: true},
					&object.Bool{Value: false},
					&object.Bool{Value: true},
				},
			},
		},
		{
			name:  "Float Elements List",
			input: `[1.5, 2.7, 3.14];`,
			excepted: &object.List{
				Elements: []object.Object{
					&object.Float{Value: 1.5},
					&object.Float{Value: 2.7},
					&object.Float{Value: 3.14},
				},
			},
		},
		{
			name:  "Expression Elements List",
			input: `[1 + 2, 3 * 4];`,
			excepted: &object.List{
				Elements: []object.Object{
					&object.Int{Value: 3},
					&object.Int{Value: 12},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer("<test>", tt.input)
			p, _ := parser.NewParser(l)
			program := p.ParseProgram()
			e := NewEvaluator(f)
			val := e.evalListExpression(program.Statements[0].(*ast.ExpressionStatement).Expr.(*ast.ListExpression), env)
			if !reflect.DeepEqual(val, tt.excepted) {
				t.Errorf("excepted %+v, got %+v", tt.excepted, val)
			}
		})
	}
}

func TestEvaluator_VisitIdentifierExpression(t *testing.T) {
	env := &object.Environment{
		Store: map[string]*object.Symbol{
			"a": {
				Name: "a",
				Value: &object.Int{
					Value: 1,
				},
			},
		},
		Outer: &object.Environment{
			Store: map[string]*object.Symbol{
				"b": {
					Name: "b",
					Value: &object.Bool{
						Value: true,
					},
				},
			},
			Outer: nil,
		},
	}
	f := &frame.Frame{
		FuncName: "<test>",
		Parent:   nil,
		PosStart: nil,
		PosEnd:   nil,
	}

	tests := []struct {
		name     string
		input    string
		excepted object.Object
	}{
		{
			name:  "Identifier Expression",
			input: "a;",
			excepted: &object.Int{
				Value: 1,
			},
		},
		{
			name:  "Eval Identifier of Outer Context",
			input: "b;",
			excepted: &object.Bool{
				Value: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer("<test>", tt.input)
			p, _ := parser.NewParser(l)
			program := p.ParseProgram()
			e := NewEvaluator(f)
			val := e.evalIdentifierExpression(program.Statements[0].(*ast.ExpressionStatement).Expr.(*ast.IdentifierExpression), env)
			if !reflect.DeepEqual(val, tt.excepted) {
				t.Errorf("excepted %+v, got %+v", tt.excepted, val)
			}
		})
	}
}

func TestEvaluator_VisitVarInitializationExpression(t *testing.T) {
	env := &object.Environment{
		Store: make(map[string]*object.Symbol),
		Outer: nil,
	}
	f := &frame.Frame{
		FuncName: "<test>",
		Parent:   nil,
		PosStart: nil,
		PosEnd:   nil,
	}

	tests := []struct {
		name     string
		input    string
		excepted object.Object
	}{
		{
			name:  "Var Initialization",
			input: "var a = 1;",
			excepted: &object.Int{
				Value: 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer("<test>", tt.input)
			p, _ := parser.NewParser(l)
			program := p.ParseProgram()
			e := NewEvaluator(f)
			val := e.evalVarInitializationExpression(program.Statements[0].(*ast.ExpressionStatement).Expr.(*ast.VarInitializationExpression), env)
			if !reflect.DeepEqual(val, tt.excepted) {
				t.Errorf("excepted %+v, got %+v", tt.excepted, val)
			}
		})
	}
}

func TestEvaluator_VisitVarAssignmentExpression(t *testing.T) {
	env := &object.Environment{
		Store: map[string]*object.Symbol{
			"a": {
				Name: "a",
				Value: &object.Int{
					Value: 1,
				},
				IsConst: false,
			},
			"l": {
				Name: "l",
				Value: &object.List{
					Elements: []object.Object{
						&object.Int{Value: 1},
						&object.Int{Value: 2},
						&object.Int{Value: 3},
					},
				},
				IsConst: false,
			},
		},
		Outer: nil,
	}
	f := &frame.Frame{
		FuncName: "<test>",
		Parent:   nil,
		PosStart: nil,
		PosEnd:   nil,
	}

	tests := []struct {
		name     string
		input    string
		excepted object.Object
	}{
		{
			name:  "Var Assignment",
			input: "a = 2;",
			excepted: &object.Int{
				Value: 2,
			},
		},
		{
			name:     "List Index Assignment",
			input:    "l[0] = 4;",
			excepted: &object.Int{Value: 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer("<test>", tt.input)
			p, _ := parser.NewParser(l)
			program := p.ParseProgram()
			e := NewEvaluator(f)
			val := e.evalVarAssignmentExpression(program.Statements[0].(*ast.ExpressionStatement).Expr.(*ast.VarAssignmentExpression), env)
			if !reflect.DeepEqual(val, tt.excepted) {
				t.Errorf("excepted %+v, got %+v", tt.excepted, val)
			}
		})
	}
}

func TestEvaluator_VisitCompoundAssignmentExpression(t *testing.T) {
	env := &object.Environment{
		Store: map[string]*object.Symbol{
			"a": {
				Name: "a",
				Value: &object.Int{
					Value: 1,
				},
			},
			"l": {
				Name: "l",
				Value: &object.List{
					Elements: []object.Object{
						&object.Int{Value: 1},
						&object.Int{Value: 2},
						&object.Int{Value: 3},
					},
				},
			},
		},
		Outer: nil,
	}
	f := &frame.Frame{
		FuncName: "<test>",
		Parent:   nil,
		PosStart: nil,
		PosEnd:   nil,
	}

	tests := []struct {
		name     string
		input    string
		excepted object.Object
	}{
		{
			name:  "Compound Assignment",
			input: "a += 1;",
			excepted: &object.Int{
				Value: 2,
			},
		},
		{
			name:     "List Index Compound Assignment",
			input:    "l[0] += 1;",
			excepted: &object.Int{Value: 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer("<test>", tt.input)
			p, _ := parser.NewParser(l)
			program := p.ParseProgram()
			e := NewEvaluator(f)
			val := e.evalCompoundAssignmentExpression(program.Statements[0].(*ast.ExpressionStatement).Expr.(*ast.CompoundAssignmentExpression), env)
			if !reflect.DeepEqual(val, tt.excepted) {
				t.Errorf("excepted %+v, got %+v", tt.excepted, val)
			}
		})
	}
}

func TestEvaluator_VisitPrefixUnaryIncDecExpression(t *testing.T) {
	env := &object.Environment{
		Store: map[string]*object.Symbol{
			"a": {
				Name: "a",
				Value: &object.Int{
					Value: 1,
				},
			},
			"l": {
				Name: "l",
				Value: &object.List{
					Elements: []object.Object{
						&object.Int{Value: 1},
						&object.Int{Value: 2},
						&object.Int{Value: 3},
					},
				},
			},
		},
		Outer: nil,
	}
	f := &frame.Frame{
		FuncName: "<test>",
		Parent:   nil,
		PosStart: nil,
		PosEnd:   nil,
	}

	tests := []struct {
		name     string
		input    string
		excepted object.Object
	}{
		{
			name:  "Prefix Unary Increment Expression",
			input: "++a;",
			excepted: &object.Int{
				Value: 2,
			},
		},
		{
			name:  "Prefix Unary Decrement Expression",
			input: "--l[0];",
			excepted: &object.Int{
				Value: 0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer("<test>", tt.input)
			p, _ := parser.NewParser(l)
			program := p.ParseProgram()
			e := NewEvaluator(f)
			val := e.evalPrefixUnaryIncDecExpression(program.Statements[0].(*ast.ExpressionStatement).Expr.(*ast.PrefixUnaryIncDecExpression), env)
			if !reflect.DeepEqual(val, tt.excepted) {
				t.Errorf("excepted %+v, got %+v", tt.excepted, val)
			}
		})
	}
}

func TestEvaluator_VisitPostfixUnaryIncDecExpression(t *testing.T) {
	env := &object.Environment{
		Store: map[string]*object.Symbol{
			"a": {
				Name: "a",
				Value: &object.Int{
					Value: 1,
				},
			},
			"l": {
				Name: "l",
				Value: &object.List{
					Elements: []object.Object{
						&object.Int{Value: 1},
						&object.Int{Value: 2},
						&object.Int{Value: 3},
					},
				},
			},
		},
		Outer: nil,
	}
	f := &frame.Frame{
		FuncName: "<test>",
		Parent:   nil,
		PosStart: nil,
		PosEnd:   nil,
	}

	tests := []struct {
		name     string
		input    string
		excepted object.Object
	}{
		{
			name:  "Prefix Unary Increment Expression",
			input: "a++;",
			excepted: &object.Int{
				Value: 1,
			},
		},
		{
			name:  "Prefix Unary Decrement Expression",
			input: "l[0]--;",
			excepted: &object.Int{
				Value: 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer("<test>", tt.input)
			p, _ := parser.NewParser(l)
			program := p.ParseProgram()
			e := NewEvaluator(f)
			val := e.evalPostfixUnaryIncDecExpression(program.Statements[0].(*ast.ExpressionStatement).Expr.(*ast.PostfixUnaryIncDecExpression), env)
			if !reflect.DeepEqual(val, tt.excepted) {
				t.Errorf("excepted %+v, got %+v", tt.excepted, val)
			}
		})
	}
}

func TestEvaluator_VisitPrefixExpression(t *testing.T) {
	env := &object.Environment{
		Store: make(map[string]*object.Symbol),
		Outer: nil,
	}
	f := &frame.Frame{
		FuncName: "<test>",
		Parent:   nil,
		PosStart: nil,
		PosEnd:   nil,
	}

	tests := []struct {
		name     string
		input    string
		excepted object.Object
	}{
		{
			name:  "Negative Int",
			input: `-1;`,
			excepted: &object.Int{
				Value: -1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer("<test>", tt.input)
			p, _ := parser.NewParser(l)
			program := p.ParseProgram()
			e := NewEvaluator(f)
			val := e.evalPrefixExpression(program.Statements[0].(*ast.ExpressionStatement).Expr.(*ast.PrefixExpression), env)
			if !reflect.DeepEqual(val, tt.excepted) {
				t.Errorf("excepted %+v, got %+v", tt.excepted, val)
			}
		})
	}
}

func TestEvaluator_VisitBlockExpression(t *testing.T) {
	env := &object.Environment{
		Store: make(map[string]*object.Symbol),
		Outer: nil,
	}
	f := &frame.Frame{
		FuncName: "<test>",
		Parent:   nil,
		PosStart: nil,
		PosEnd:   nil,
	}

	tests := []struct {
		name     string
		input    string
		excepted object.Object
	}{
		{
			name:  "Block Expression",
			input: `{ 1 };`,
			excepted: &object.Int{
				Value: 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer("<test>", tt.input)
			p, _ := parser.NewParser(l)
			program := p.ParseProgram()
			e := NewEvaluator(f)
			val := e.evalBlockExpression(program.Statements[0].(*ast.ExpressionStatement).Expr.(*ast.BlockExpression), env)
			if !reflect.DeepEqual(val, tt.excepted) {
				t.Errorf("excepted %+v, got %+v", tt.excepted, val)
			}
		})
	}
}

func TestEvaluator_VisitIfExpression(t *testing.T) {
	env := &object.Environment{
		Store: make(map[string]*object.Symbol),
		Outer: nil,
	}
	f := &frame.Frame{
		FuncName: "<test>",
		Parent:   nil,
		PosStart: nil,
		PosEnd:   nil,
	}

	tests := []struct {
		name     string
		input    string
		excepted object.Object
	}{
		{
			name:  "If Expression",
			input: `if true 1;`,
			excepted: &object.Int{
				Value: 1,
			},
		},
		{
			name:  "If-Else Expression",
			input: `if false 1 else 2;`,
			excepted: &object.Int{
				Value: 2,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer("<test>", tt.input)
			p, _ := parser.NewParser(l)
			program := p.ParseProgram()
			e := NewEvaluator(f)
			val := e.evalIfExpression(program.Statements[0].(*ast.ExpressionStatement).Expr.(*ast.IfExpression), env)
			if !reflect.DeepEqual(val, tt.excepted) {
				t.Errorf("excepted %+v, got %+v", tt.excepted, val)
			}
		})
	}
}

func TestEvaluator_VisitCallExpression(t *testing.T) {
	f := &frame.Frame{
		FuncName: "<test>",
		Parent:   nil,
		PosStart: nil,
		PosEnd:   nil,
	}

	env := &object.Environment{
		Store: map[string]*object.Symbol{
			"f": {
				Name: "f",
				Value: &object.Function{
					Name:      "f",
					Parameter: make([]*ast.Parameter, 0),
					Body: &ast.ExpressionStatement{
						Expr: &ast.IntExpression{
							Value:    1,
							PosStart: util.NewPos(1, 10, 9, "<test>", "func f() 1;"),
							PosEnd:   util.NewPos(1, 11, 10, "<test>", "func f() 1;"),
						},
					},
					Env: &object.Environment{
						Store: make(map[string]*object.Symbol),
						Outer: nil,
					},
				},
				IsConst: true,
			},
			"g": {
				Name: "g",
				Value: &object.Function{
					Name: "g",
					Parameter: []*ast.Parameter{
						{
							Name: &ast.IdentifierExpression{
								Name:     "a",
								PosStart: util.NewPos(1, 8, 7, "<test>", "func g(a=1) 1;"),
								PosEnd:   util.NewPos(1, 9, 8, "<test>", "func g(a=1) 1;"),
							},
							DefaultValue: &ast.IntExpression{
								Value:    1,
								PosStart: util.NewPos(1, 10, 9, "<test>", "func g(a=1) 1;"),
								PosEnd:   util.NewPos(1, 11, 10, "<test>", "func g(a=1) 1;"),
							},
							PosStart: util.NewPos(1, 8, 7, "<test>", "func g(a=1) 1;"),
							PosEnd:   util.NewPos(1, 11, 10, "<test>", "func g(a=1) 1;"),
						},
					},
					Body: &ast.ExpressionStatement{
						Expr: &ast.IntExpression{
							Value:    1,
							PosStart: util.NewPos(1, 13, 12, "<test>", "func g(a=1) 1;"),
							PosEnd:   util.NewPos(1, 14, 13, "<test>", "func g(a=1) 1;"),
						},
					},
					Env: &object.Environment{
						Store: make(map[string]*object.Symbol),
						Outer: nil,
					},
				},
				IsConst: true,
			},
			"len": {
				Name: "len",
				Value: &object.BuiltinFunction{
					Name:      "len",
					Parameter: []string{"a"},
					Fn: func(f *frame.Frame, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
						switch args[0].(type) {
						case *object.String:
							return &object.Int{
								Value: int64(len(args[0].(*object.String).Value)),
							}, nil
						default:
							return nil, &TypeError{
								Frame: &frame.Frame{
									FuncName: "<builtin \"len\">",
									Parent:   f,
									PosStart: posStart,
									PosEnd:   posEnd,
								},
								Message:  fmt.Sprintf("invalid argument type %s for len()", args[0].Type()),
								PosStart: posStart,
								PosEnd:   posEnd,
							}
						}
					},
				},
				IsConst: true,
			},
		},

		Outer: nil,
	}

	tests := []struct {
		name     string
		input    string
		excepted object.Object
	}{
		{
			name:  "Call Expression",
			input: `f();`,
			excepted: &object.Int{
				Value: 1,
			},
		},
		{
			name:  "Call Function with Default Argument",
			input: `g();`,
			excepted: &object.Int{
				Value: 1,
			},
		},
		{
			name:  "Call Builtin Function",
			input: `len("hello");`,
			excepted: &object.Int{
				Value: 5,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer("<test>", tt.input)
			p, _ := parser.NewParser(l)
			program := p.ParseProgram()
			e := NewEvaluator(f)
			val := e.evalCallExpression(program.Statements[0].(*ast.ExpressionStatement).Expr.(*ast.CallExpression), env)
			if !reflect.DeepEqual(val, tt.excepted) {
				t.Errorf("excepted %+v, got %+v", tt.excepted, val)
			}
		})
	}
}

func TestEvaluator_VisitIndexExpression(t *testing.T) {
	f := &frame.Frame{
		FuncName: "<test>",
		Parent:   nil,
		PosStart: nil,
		PosEnd:   nil,
	}

	env := &object.Environment{
		Store: map[string]*object.Symbol{
			"lst": {
				Name: "lst",
				Value: &object.List{
					Elements: []object.Object{
						&object.Int{
							Value: 1,
						},
						&object.Int{
							Value: 2,
						},
						&object.Int{
							Value: 3,
						},
					},
				},
				IsConst: true,
			},
		},
		Outer: nil,
	}

	tests := []struct {
		name     string
		input    string
		excepted object.Object
	}{
		{
			name:  "Index Expression",
			input: `lst[1];`,
			excepted: &object.Int{
				Value: 2,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer("<test>", tt.input)
			p, _ := parser.NewParser(l)
			program := p.ParseProgram()
			e := NewEvaluator(f)
			val := e.evalIndexExpression(program.Statements[0].(*ast.ExpressionStatement).Expr.(*ast.IndexExpression), env)
			if !reflect.DeepEqual(val, tt.excepted) {
				t.Errorf("excepted %+v, got %+v", tt.excepted, val)
			}
		})
	}
}
