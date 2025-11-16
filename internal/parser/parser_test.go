package parser

import (
	"reflect"
	"testing"

	"github.com/Ghost-Xiao/ghost-lang/internal/lexer"
	"github.com/Ghost-Xiao/ghost-lang/internal/parser/ast"
	"github.com/Ghost-Xiao/ghost-lang/internal/util"
)

func TestParser_ParseProgram(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *ast.Program
	}{
		{
			name:  "Empty Program",
			input: "",
			expected: &ast.Program{
				Statements: []ast.Statement{},
				PosStart:   util.NewPos(1, 1, 0, "<test>", ""),
				PosEnd:     util.NewPos(1, 2, 1, "<test>", ""),
			},
		},
		{
			name: "Program",
			input: `1;
var 真 = true;
"hello\n";
-1 + 1;`,
			expected: &ast.Program{
				Statements: []ast.Statement{
					&ast.ExpressionStatement{
						Expr: &ast.IntExpression{
							Value:    1,
							PosStart: util.NewPos(1, 1, 0, "<test>", "1;\nvar 真 = true;\n\"hello\\n\";\n-1 + 1;"),
							PosEnd:   util.NewPos(1, 2, 1, "<test>", "1;\nvar 真 = true;\n\"hello\\n\";\n-1 + 1;"),
						},
						PosStart: util.NewPos(1, 1, 0, "<test>", "1;\nvar 真 = true;\n\"hello\\n\";\n-1 + 1;"),
						PosEnd:   util.NewPos(1, 2, 1, "<test>", "1;\nvar 真 = true;\n\"hello\\n\";\n-1 + 1;"),
					},
					&ast.ExpressionStatement{
						Expr: &ast.VarInitializationExpression{
							IsConst: false,
							Name: &ast.IdentifierExpression{
								Name:     "真",
								PosStart: util.NewPos(2, 5, 7, "<test>", "1;\nvar 真 = true;\n\"hello\\n\";\n-1 + 1;"),
								PosEnd:   util.NewPos(2, 6, 10, "<test>", "1;\nvar 真 = true;\n\"hello\\n\";\n-1 + 1;"),
							},
							Value: &ast.BoolExpression{
								Value:    true,
								PosStart: util.NewPos(2, 9, 13, "<test>", "1;\nvar 真 = true;\n\"hello\\n\";\n-1 + 1;"),
								PosEnd:   util.NewPos(2, 13, 17, "<test>", "1;\nvar 真 = true;\n\"hello\\n\";\n-1 + 1;"),
							},
							PosStart: util.NewPos(2, 1, 3, "<test>", "1;\nvar 真 = true;\n\"hello\\n\";\n-1 + 1;"),
							PosEnd:   util.NewPos(2, 13, 17, "<test>", "1;\nvar 真 = true;\n\"hello\\n\";\n-1 + 1;"),
						},
						PosStart: util.NewPos(2, 1, 3, "<test>", "1;\nvar 真 = true;\n\"hello\\n\";\n-1 + 1;"),
						PosEnd:   util.NewPos(2, 13, 17, "<test>", "1;\nvar 真 = true;\n\"hello\\n\";\n-1 + 1;"),
					},
					&ast.ExpressionStatement{
						Expr: &ast.StringExpression{
							Value:    "hello\n",
							PosStart: util.NewPos(3, 1, 19, "<test>", "1;\nvar 真 = true;\n\"hello\\n\";\n-1 + 1;"),
							PosEnd:   util.NewPos(3, 10, 28, "<test>", "1;\nvar 真 = true;\n\"hello\\n\";\n-1 + 1;"),
						},
						PosStart: util.NewPos(3, 1, 19, "<test>", "1;\nvar 真 = true;\n\"hello\\n\";\n-1 + 1;"),
						PosEnd:   util.NewPos(3, 10, 28, "<test>", "1;\nvar 真 = true;\n\"hello\\n\";\n-1 + 1;"),
					},
					&ast.ExpressionStatement{
						Expr: &ast.InfixExpression{
							Left: &ast.PrefixExpression{
								Operator: &lexer.Token{
									Type:     lexer.MINUS,
									Literal:  "-",
									PosStart: util.NewPos(4, 1, 30, "<test>", "1;\nvar 真 = true;\n\"hello\\n\";\n-1 + 1;"),
									PosEnd:   util.NewPos(4, 2, 31, "<test>", "1;\nvar 真 = true;\n\"hello\\n\";\n-1 + 1;"),
								},
								Value: &ast.IntExpression{
									Value:    1,
									PosStart: util.NewPos(4, 2, 31, "<test>", "1;\nvar 真 = true;\n\"hello\\n\";\n-1 + 1;"),
									PosEnd:   util.NewPos(4, 3, 32, "<test>", "1;\nvar 真 = true;\n\"hello\\n\";\n-1 + 1;"),
								},
								PosStart: util.NewPos(4, 1, 30, "<test>", "1;\nvar 真 = true;\n\"hello\\n\";\n-1 + 1;"),
								PosEnd:   util.NewPos(4, 3, 32, "<test>", "1;\nvar 真 = true;\n\"hello\\n\";\n-1 + 1;"),
							},
							Operator: &lexer.Token{
								Type:     lexer.PLUS,
								Literal:  "+",
								PosStart: util.NewPos(4, 4, 33, "<test>", "1;\nvar 真 = true;\n\"hello\\n\";\n-1 + 1;"),
								PosEnd:   util.NewPos(4, 5, 34, "<test>", "1;\nvar 真 = true;\n\"hello\\n\";\n-1 + 1;"),
							},
							Right: &ast.IntExpression{
								Value:    1,
								PosStart: util.NewPos(4, 6, 35, "<test>", "1;\nvar 真 = true;\n\"hello\\n\";\n-1 + 1;"),
								PosEnd:   util.NewPos(4, 7, 36, "<test>", "1;\nvar 真 = true;\n\"hello\\n\";\n-1 + 1;"),
							},
							PosStart: util.NewPos(4, 1, 30, "<test>", "1;\nvar 真 = true;\n\"hello\\n\";\n-1 + 1;"),
							PosEnd:   util.NewPos(4, 7, 36, "<test>", "1;\nvar 真 = true;\n\"hello\\n\";\n-1 + 1;"),
						},
						PosStart: util.NewPos(4, 1, 30, "<test>", "1;\nvar 真 = true;\n\"hello\\n\";\n-1 + 1;"),
						PosEnd:   util.NewPos(4, 7, 36, "<test>", "1;\nvar 真 = true;\n\"hello\\n\";\n-1 + 1;"),
					},
				},
				PosStart: util.NewPos(1, 1, 0, "<test>", "1;\nvar 真 = true;\n\"hello\\n\";\n-1 + 1;"),
				PosEnd:   util.NewPos(4, 9, 38, "<test>", "1;\nvar 真 = true;\n\"hello\\n\";\n-1 + 1;"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer("<test>", tt.input)
			p, _ := NewParser(l)
			program := p.ParseProgram()
			if p.Err != nil {
				t.Errorf("err = %+v, expected nil", p.Err)
			}

			if !reflect.DeepEqual(program, tt.expected) {
				t.Errorf("expected %+v, got %+v", tt.expected, program)
			}
		})
	}
}

func TestParser_ParseForStatement(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *ast.ForStatement
	}{
		{
			name:  "For Statement",
			input: "for var i = 1; i < 5; i++ 1;",
			expected: &ast.ForStatement{
				Initialization: &ast.ExpressionStatement{
					Expr: &ast.VarInitializationExpression{
						IsConst: false,
						Name: &ast.IdentifierExpression{
							Name:     "i",
							PosStart: util.NewPos(1, 9, 8, "<test>", "for var i = 1; i < 5; i++ 1;"),
							PosEnd:   util.NewPos(1, 10, 9, "<test>", "for var i = 1; i < 5; i++ 1;"),
						},
						Value: &ast.IntExpression{
							Value:    1,
							PosStart: util.NewPos(1, 13, 12, "<test>", "for var i = 1; i < 5; i++ 1;"),
							PosEnd:   util.NewPos(1, 14, 13, "<test>", "for var i = 1; i < 5; i++ 1;"),
						},
						PosStart: util.NewPos(1, 5, 4, "<test>", "for var i = 1; i < 5; i++ 1;"),
						PosEnd:   util.NewPos(1, 14, 13, "<test>", "for var i = 1; i < 5; i++ 1;"),
					},
					PosStart: util.NewPos(1, 5, 4, "<test>", "for var i = 1; i < 5; i++ 1;"),
					PosEnd:   util.NewPos(1, 14, 13, "<test>", "for var i = 1; i < 5; i++ 1;"),
				},
				Condition: &ast.InfixExpression{
					Left: &ast.IdentifierExpression{
						Name:     "i",
						PosStart: util.NewPos(1, 16, 15, "<test>", "for var i = 1; i < 5; i++ 1;"),
						PosEnd:   util.NewPos(1, 17, 16, "<test>", "for var i = 1; i < 5; i++ 1;"),
					},
					Operator: &lexer.Token{
						Type:     lexer.LT,
						Literal:  "<",
						PosStart: util.NewPos(1, 18, 17, "<test>", "for var i = 1; i < 5; i++ 1;"),
						PosEnd:   util.NewPos(1, 19, 18, "<test>", "for var i = 1; i < 5; i++ 1;"),
					},
					Right: &ast.IntExpression{
						Value:    5,
						PosStart: util.NewPos(1, 20, 19, "<test>", "for var i = 1; i < 5; i++ 1;"),
						PosEnd:   util.NewPos(1, 21, 20, "<test>", "for var i = 1; i < 5; i++ 1;"),
					},
					PosStart: util.NewPos(1, 16, 15, "<test>", "for var i = 1; i < 5; i++ 1;"),
					PosEnd:   util.NewPos(1, 21, 20, "<test>", "for var i = 1; i < 5; i++ 1;"),
				},
				Update: &ast.ExpressionStatement{
					Expr: &ast.PostfixUnaryIncDecExpression{
						Operator: &lexer.Token{
							Type:     lexer.INCREMENT,
							Literal:  "++",
							PosStart: util.NewPos(1, 24, 23, "<test>", "for var i = 1; i < 5; i++ 1;"),
							PosEnd:   util.NewPos(1, 26, 25, "<test>", "for var i = 1; i < 5; i++ 1;"),
						},
						Left: &ast.IdentifierExpression{
							Name:     "i",
							PosStart: util.NewPos(1, 23, 22, "<test>", "for var i = 1; i < 5; i++ 1;"),
							PosEnd:   util.NewPos(1, 24, 23, "<test>", "for var i = 1; i < 5; i++ 1;"),
						},
						PosStart: util.NewPos(1, 23, 22, "<test>", "for var i = 1; i < 5; i++ 1;"),
						PosEnd:   util.NewPos(1, 26, 25, "<test>", "for var i = 1; i < 5; i++ 1;"),
					},
					PosStart: util.NewPos(1, 23, 22, "<test>", "for var i = 1; i < 5; i++ 1;"),
					PosEnd:   util.NewPos(1, 26, 25, "<test>", "for var i = 1; i < 5; i++ 1;"),
				},
				Body: &ast.ExpressionStatement{
					Expr: &ast.IntExpression{
						Value:    1,
						PosStart: util.NewPos(1, 27, 26, "<test>", "for var i = 1; i < 5; i++ 1;"),
						PosEnd:   util.NewPos(1, 28, 27, "<test>", "for var i = 1; i < 5; i++ 1;"),
					},
					PosStart: util.NewPos(1, 27, 26, "<test>", "for var i = 1; i < 5; i++ 1;"),
					PosEnd:   util.NewPos(1, 28, 27, "<test>", "for var i = 1; i < 5; i++ 1;"),
				},
				PosStart: util.NewPos(1, 1, 0, "<test>", "for var i = 1; i < 5; i++ 1;"),
				PosEnd:   util.NewPos(1, 28, 27, "<test>", "for var i = 1; i < 5; i++ 1;"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer("<test>", tt.input)
			p, _ := NewParser(l)
			program := p.ParseProgram()
			expr := program.Statements[0].(*ast.ForStatement)

			if p.Err != nil {
				t.Errorf("err = %+v, expected nil", p.Err)
			}

			if !reflect.DeepEqual(expr, tt.expected) {
				t.Errorf("expected %+v, got %+v", tt.expected, expr)
			}
		})
	}
}

func TestParser_ParseFunctionDeclarationStatement(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *ast.FunctionDeclarationStatement
	}{
		{
			name:  "Function Declaration",
			input: "func f(a) 1;",
			expected: &ast.FunctionDeclarationStatement{
				Name: &ast.IdentifierExpression{
					Name:     "f",
					PosStart: util.NewPos(1, 6, 5, "<test>", "func f(a) 1;"),
					PosEnd:   util.NewPos(1, 7, 6, "<test>", "func f(a) 1;"),
				},
				Parameter: []*ast.Parameter{
					{
						Name: &ast.IdentifierExpression{
							Name:     "a",
							PosStart: util.NewPos(1, 8, 7, "<test>", "func f(a) 1;"),
							PosEnd:   util.NewPos(1, 9, 8, "<test>", "func f(a) 1;"),
						},
						DefaultValue: nil,
						PosStart:     util.NewPos(1, 8, 7, "<test>", "func f(a) 1;"),
						PosEnd:       util.NewPos(1, 9, 8, "<test>", "func f(a) 1;"),
					},
				},
				Body: &ast.ExpressionStatement{
					Expr: &ast.IntExpression{
						Value:    1,
						PosStart: util.NewPos(1, 11, 10, "<test>", "func f(a) 1;"),
						PosEnd:   util.NewPos(1, 12, 11, "<test>", "func f(a) 1;"),
					},
					PosStart: util.NewPos(1, 11, 10, "<test>", "func f(a) 1;"),
					PosEnd:   util.NewPos(1, 12, 11, "<test>", "func f(a) 1;"),
				},
				PosStart: util.NewPos(1, 1, 0, "<test>", "func f(a) 1;"),
				PosEnd:   util.NewPos(1, 12, 11, "<test>", "func f(a) 1;"),
			},
		},
		{
			name:  "Function Declaration With Default Value",
			input: "func f(a=1) 1;",
			expected: &ast.FunctionDeclarationStatement{
				Name: &ast.IdentifierExpression{
					Name:     "f",
					PosStart: util.NewPos(1, 6, 5, "<test>", "func f(a=1) 1;"),
					PosEnd:   util.NewPos(1, 7, 6, "<test>", "func f(a=1) 1;"),
				},
				Parameter: []*ast.Parameter{
					{
						Name: &ast.IdentifierExpression{
							Name:     "a",
							PosStart: util.NewPos(1, 8, 7, "<test>", "func f(a=1) 1;"),
							PosEnd:   util.NewPos(1, 9, 8, "<test>", "func f(a=1) 1;"),
						},
						DefaultValue: &ast.IntExpression{
							Value:    1,
							PosStart: util.NewPos(1, 10, 9, "<test>", "func f(a=1) 1;"),
							PosEnd:   util.NewPos(1, 11, 10, "<test>", "func f(a=1) 1;"),
						},
						PosStart: util.NewPos(1, 8, 7, "<test>", "func f(a=1) 1;"),
						PosEnd:   util.NewPos(1, 11, 10, "<test>", "func f(a=1) 1;"),
					},
				},
				Body: &ast.ExpressionStatement{
					Expr: &ast.IntExpression{
						Value:    1,
						PosStart: util.NewPos(1, 13, 12, "<test>", "func f(a=1) 1;"),
						PosEnd:   util.NewPos(1, 14, 13, "<test>", "func f(a=1) 1;"),
					},
					PosStart: util.NewPos(1, 13, 12, "<test>", "func f(a=1) 1;"),
					PosEnd:   util.NewPos(1, 14, 13, "<test>", "func f(a=1) 1;"),
				},
				PosStart: util.NewPos(1, 1, 0, "<test>", "func f(a=1) 1;"),
				PosEnd:   util.NewPos(1, 14, 13, "<test>", "func f(a=1) 1;"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer("<test>", tt.input)
			p, _ := NewParser(l)
			program := p.ParseProgram()
			expr := program.Statements[0].(*ast.FunctionDeclarationStatement)
			if p.Err != nil {
				t.Errorf("err = %+v, expected nil", p.Err)
			}

			if !reflect.DeepEqual(expr, tt.expected) {
				t.Errorf("expected %+v, got %+v", tt.expected, expr)
			}
		})
	}
}

func TestParser_ParseReturnStatement(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *ast.ReturnStatement
	}{
		{
			name:  "Return Statement",
			input: "return 1;",
			expected: &ast.ReturnStatement{
				ReturnValue: &ast.IntExpression{
					Value:    1,
					PosStart: util.NewPos(1, 8, 7, "<test>", "return 1;"),
					PosEnd:   util.NewPos(1, 9, 8, "<test>", "return 1;"),
				},
				PosStart: util.NewPos(1, 1, 0, "<test>", "return 1;"),
				PosEnd:   util.NewPos(1, 9, 8, "<test>", "return 1;"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer("<test>", tt.input)
			p, _ := NewParser(l)
			program := p.ParseProgram()
			expr := program.Statements[0].(*ast.ReturnStatement)

			if p.Err != nil {
				t.Errorf("err = %+v, expected nil", p.Err)
			}

			if !reflect.DeepEqual(expr, tt.expected) {
				t.Errorf("expected %+v, got %+v", tt.expected, expr)
			}
		})
	}
}

func TestParser_ParsePrefixExpression(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *ast.PrefixExpression
	}{
		{
			name:  "Minus Prefix Expression",
			input: "-1;",
			expected: &ast.PrefixExpression{
				Operator: &lexer.Token{
					Type:     lexer.MINUS,
					Literal:  "-",
					PosStart: util.NewPos(1, 1, 0, "<test>", "-1;"),
					PosEnd:   util.NewPos(1, 2, 1, "<test>", "-1;"),
				},
				Value: &ast.IntExpression{
					Value:    1,
					PosStart: util.NewPos(1, 2, 1, "<test>", "-1;"),
					PosEnd:   util.NewPos(1, 3, 2, "<test>", "-1;"),
				},
				PosStart: util.NewPos(1, 1, 0, "<test>", "-1;"),
				PosEnd:   util.NewPos(1, 3, 2, "<test>", "-1;"),
			},
		},
		{
			name:  "Not Prefix Expression",
			input: "!true;",
			expected: &ast.PrefixExpression{
				Operator: &lexer.Token{
					Type:     lexer.BANG,
					Literal:  "!",
					PosStart: util.NewPos(1, 1, 0, "<test>", "!true;"),
					PosEnd:   util.NewPos(1, 2, 1, "<test>", "!true;"),
				},
				Value: &ast.BoolExpression{
					Value:    true,
					PosStart: util.NewPos(1, 2, 1, "<test>", "!true;"),
					PosEnd:   util.NewPos(1, 6, 5, "<test>", "!true;"),
				},
				PosStart: util.NewPos(1, 1, 0, "<test>", "!true;"),
				PosEnd:   util.NewPos(1, 6, 5, "<test>", "!true;"),
			},
		},
		{
			name:  "BitNot Prefix Expression",
			input: "~1;",
			expected: &ast.PrefixExpression{
				Operator: &lexer.Token{
					Type:     lexer.BITWISE_NOT,
					Literal:  "~",
					PosStart: util.NewPos(1, 1, 0, "<test>", "~1;"),
					PosEnd:   util.NewPos(1, 2, 1, "<test>", "~1;"),
				},
				Value: &ast.IntExpression{
					Value:    1,
					PosStart: util.NewPos(1, 2, 1, "<test>", "~1;"),
					PosEnd:   util.NewPos(1, 3, 2, "<test>", "~1;"),
				},
				PosStart: util.NewPos(1, 1, 0, "<test>", "~1;"),
				PosEnd:   util.NewPos(1, 3, 2, "<test>", "~1;"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer("<test>", tt.input)
			p, _ := NewParser(l)
			program := p.ParseProgram()
			expr := program.Statements[0].(*ast.ExpressionStatement).Expr.(*ast.PrefixExpression)

			if p.Err != nil {
				t.Errorf("err = %+v, expected nil", p.Err)
			}

			if !reflect.DeepEqual(expr, tt.expected) {
				t.Errorf("expected %+v, got %+v", tt.expected, expr)
			}
		})
	}
}

func TestParser_ParseInfixExpression(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *ast.InfixExpression
	}{
		{
			name:  "Addition Infix Expression",
			input: "1 + 1;",
			expected: &ast.InfixExpression{
				Left: &ast.IntExpression{
					Value:    1,
					PosStart: util.NewPos(1, 1, 0, "<test>", "1 + 1;"),
					PosEnd:   util.NewPos(1, 2, 1, "<test>", "1 + 1;"),
				},
				Operator: &lexer.Token{
					Type:     lexer.PLUS,
					Literal:  "+",
					PosStart: util.NewPos(1, 3, 2, "<test>", "1 + 1;"),
					PosEnd:   util.NewPos(1, 4, 3, "<test>", "1 + 1;"),
				},
				Right: &ast.IntExpression{
					Value:    1,
					PosStart: util.NewPos(1, 5, 4, "<test>", "1 + 1;"),
					PosEnd:   util.NewPos(1, 6, 5, "<test>", "1 + 1;"),
				},
				PosStart: util.NewPos(1, 1, 0, "<test>", "1 + 1;"),
				PosEnd:   util.NewPos(1, 6, 5, "<test>", "1 + 1;"),
			},
		},
		{
			name:  "Subtraction Infix Expression",
			input: "1 - 1;",
			expected: &ast.InfixExpression{
				Left: &ast.IntExpression{
					Value:    1,
					PosStart: util.NewPos(1, 1, 0, "<test>", "1 - 1;"),
					PosEnd:   util.NewPos(1, 2, 1, "<test>", "1 - 1;"),
				},
				Operator: &lexer.Token{
					Type:     lexer.MINUS,
					Literal:  "-",
					PosStart: util.NewPos(1, 3, 2, "<test>", "1 - 1;"),
					PosEnd:   util.NewPos(1, 4, 3, "<test>", "1 - 1;"),
				},
				Right: &ast.IntExpression{
					Value:    1,
					PosStart: util.NewPos(1, 5, 4, "<test>", "1 - 1;"),
					PosEnd:   util.NewPos(1, 6, 5, "<test>", "1 - 1;"),
				},
				PosStart: util.NewPos(1, 1, 0, "<test>", "1 - 1;"),
				PosEnd:   util.NewPos(1, 6, 5, "<test>", "1 - 1;"),
			},
		},
		{
			name:  "Multiplication Infix Expression",
			input: "1 * 1;",
			expected: &ast.InfixExpression{
				Left: &ast.IntExpression{
					Value:    1,
					PosStart: util.NewPos(1, 1, 0, "<test>", "1 * 1;"),
					PosEnd:   util.NewPos(1, 2, 1, "<test>", "1 * 1;"),
				},
				Operator: &lexer.Token{
					Type:     lexer.ASTERISK,
					Literal:  "*",
					PosStart: util.NewPos(1, 3, 2, "<test>", "1 * 1;"),
					PosEnd:   util.NewPos(1, 4, 3, "<test>", "1 * 1;"),
				},
				Right: &ast.IntExpression{
					Value:    1,
					PosStart: util.NewPos(1, 5, 4, "<test>", "1 * 1;"),
					PosEnd:   util.NewPos(1, 6, 5, "<test>", "1 * 1;"),
				},
				PosStart: util.NewPos(1, 1, 0, "<test>", "1 * 1;"),
				PosEnd:   util.NewPos(1, 6, 5, "<test>", "1 * 1;"),
			},
		},
		{
			name:  "Division Infix Expression",
			input: "1 / 1;",
			expected: &ast.InfixExpression{
				Left: &ast.IntExpression{
					Value:    1,
					PosStart: util.NewPos(1, 1, 0, "<test>", "1 / 1;"),
					PosEnd:   util.NewPos(1, 2, 1, "<test>", "1 / 1;"),
				},
				Operator: &lexer.Token{
					Type:     lexer.SLASH,
					Literal:  "/",
					PosStart: util.NewPos(1, 3, 2, "<test>", "1 / 1;"),
					PosEnd:   util.NewPos(1, 4, 3, "<test>", "1 / 1;"),
				},
				Right: &ast.IntExpression{
					Value:    1,
					PosStart: util.NewPos(1, 5, 4, "<test>", "1 / 1;"),
					PosEnd:   util.NewPos(1, 6, 5, "<test>", "1 / 1;"),
				},
				PosStart: util.NewPos(1, 1, 0, "<test>", "1 / 1;"),
				PosEnd:   util.NewPos(1, 6, 5, "<test>", "1 / 1;"),
			},
		},
		{
			name:  "Modulo Infix Expression",
			input: "1 % 1;",
			expected: &ast.InfixExpression{
				Left: &ast.IntExpression{
					Value:    1,
					PosStart: util.NewPos(1, 1, 0, "<test>", "1 % 1;"),
					PosEnd:   util.NewPos(1, 2, 1, "<test>", "1 % 1;"),
				},
				Operator: &lexer.Token{
					Type:     lexer.PERCENT,
					Literal:  "%",
					PosStart: util.NewPos(1, 3, 2, "<test>", "1 % 1;"),
					PosEnd:   util.NewPos(1, 4, 3, "<test>", "1 % 1;"),
				},
				Right: &ast.IntExpression{
					Value:    1,
					PosStart: util.NewPos(1, 5, 4, "<test>", "1 % 1;"),
					PosEnd:   util.NewPos(1, 6, 5, "<test>", "1 % 1;"),
				},
				PosStart: util.NewPos(1, 1, 0, "<test>", "1 % 1;"),
				PosEnd:   util.NewPos(1, 6, 5, "<test>", "1 % 1;"),
			},
		},
		{
			name:  "Equal Infix Expression",
			input: "1 == 1;",
			expected: &ast.InfixExpression{
				Left: &ast.IntExpression{
					Value:    1,
					PosStart: util.NewPos(1, 1, 0, "<test>", "1 == 1;"),
					PosEnd:   util.NewPos(1, 2, 1, "<test>", "1 == 1;"),
				},
				Operator: &lexer.Token{
					Type:     lexer.EQUALS,
					Literal:  "==",
					PosStart: util.NewPos(1, 3, 2, "<test>", "1 == 1;"),
					PosEnd:   util.NewPos(1, 5, 4, "<test>", "1 == 1;"),
				},
				Right: &ast.IntExpression{
					Value:    1,
					PosStart: util.NewPos(1, 6, 5, "<test>", "1 == 1;"),
					PosEnd:   util.NewPos(1, 7, 6, "<test>", "1 == 1;"),
				},
				PosStart: util.NewPos(1, 1, 0, "<test>", "1 == 1;"),
				PosEnd:   util.NewPos(1, 7, 6, "<test>", "1 == 1;"),
			},
		},
		{
			name:  "NotEqual Infix Expression",
			input: "1 != 1;",
			expected: &ast.InfixExpression{
				Left: &ast.IntExpression{
					Value:    1,
					PosStart: util.NewPos(1, 1, 0, "<test>", "1 != 1;"),
					PosEnd:   util.NewPos(1, 2, 1, "<test>", "1 != 1;"),
				},
				Operator: &lexer.Token{
					Type:     lexer.NOT_EQUALS,
					Literal:  "!=",
					PosStart: util.NewPos(1, 3, 2, "<test>", "1 != 1;"),
					PosEnd:   util.NewPos(1, 5, 4, "<test>", "1 != 1;"),
				},
				Right: &ast.IntExpression{
					Value:    1,
					PosStart: util.NewPos(1, 6, 5, "<test>", "1 != 1;"),
					PosEnd:   util.NewPos(1, 7, 6, "<test>", "1 != 1;"),
				},
				PosStart: util.NewPos(1, 1, 0, "<test>", "1 != 1;"),
				PosEnd:   util.NewPos(1, 7, 6, "<test>", "1 != 1;"),
			},
		},
		{
			name:  "GreaterThan Infix Expression",
			input: "1 > 1;",
			expected: &ast.InfixExpression{
				Left: &ast.IntExpression{
					Value:    1,
					PosStart: util.NewPos(1, 1, 0, "<test>", "1 > 1;"),
					PosEnd:   util.NewPos(1, 2, 1, "<test>", "1 > 1;"),
				},
				Operator: &lexer.Token{
					Type:     lexer.GT,
					Literal:  ">",
					PosStart: util.NewPos(1, 3, 2, "<test>", "1 > 1;"),
					PosEnd:   util.NewPos(1, 4, 3, "<test>", "1 > 1;"),
				},
				Right: &ast.IntExpression{
					Value:    1,
					PosStart: util.NewPos(1, 5, 4, "<test>", "1 > 1;"),
					PosEnd:   util.NewPos(1, 6, 5, "<test>", "1 > 1;"),
				},
				PosStart: util.NewPos(1, 1, 0, "<test>", "1 > 1;"),
				PosEnd:   util.NewPos(1, 6, 5, "<test>", "1 > 1;"),
			},
		},
		{
			name:  "GreaterThanEqual Infix Expression",
			input: "1 >= 1;",
			expected: &ast.InfixExpression{
				Left: &ast.IntExpression{
					Value:    1,
					PosStart: util.NewPos(1, 1, 0, "<test>", "1 >= 1;"),
					PosEnd:   util.NewPos(1, 2, 1, "<test>", "1 >= 1;"),
				},
				Operator: &lexer.Token{
					Type:     lexer.GTE,
					Literal:  ">=",
					PosStart: util.NewPos(1, 3, 2, "<test>", "1 >= 1;"),
					PosEnd:   util.NewPos(1, 5, 4, "<test>", "1 >= 1;"),
				},
				Right: &ast.IntExpression{
					Value:    1,
					PosStart: util.NewPos(1, 6, 5, "<test>", "1 >= 1;"),
					PosEnd:   util.NewPos(1, 7, 6, "<test>", "1 >= 1;"),
				},
				PosStart: util.NewPos(1, 1, 0, "<test>", "1 >= 1;"),
				PosEnd:   util.NewPos(1, 7, 6, "<test>", "1 >= 1;"),
			},
		},
		{
			name:  "LessThan Infix Expression",
			input: "1 < 1;",
			expected: &ast.InfixExpression{
				Left: &ast.IntExpression{
					Value:    1,
					PosStart: util.NewPos(1, 1, 0, "<test>", "1 < 1;"),
					PosEnd:   util.NewPos(1, 2, 1, "<test>", "1 < 1;"),
				},
				Operator: &lexer.Token{
					Type:     lexer.LT,
					Literal:  "<",
					PosStart: util.NewPos(1, 3, 2, "<test>", "1 < 1;"),
					PosEnd:   util.NewPos(1, 4, 3, "<test>", "1 < 1;"),
				},
				Right: &ast.IntExpression{
					Value:    1,
					PosStart: util.NewPos(1, 5, 4, "<test>", "1 < 1;"),
					PosEnd:   util.NewPos(1, 6, 5, "<test>", "1 < 1;"),
				},
				PosStart: util.NewPos(1, 1, 0, "<test>", "1 < 1;"),
				PosEnd:   util.NewPos(1, 6, 5, "<test>", "1 < 1;"),
			},
		},
		{
			name:  "LessThanEqual Infix Expression",
			input: "1 <= 1;",
			expected: &ast.InfixExpression{
				Left: &ast.IntExpression{
					Value:    1,
					PosStart: util.NewPos(1, 1, 0, "<test>", "1 <= 1;"),
					PosEnd:   util.NewPos(1, 2, 1, "<test>", "1 <= 1;"),
				},
				Operator: &lexer.Token{
					Type:     lexer.LTE,
					Literal:  "<=",
					PosStart: util.NewPos(1, 3, 2, "<test>", "1 <= 1;"),
					PosEnd:   util.NewPos(1, 5, 4, "<test>", "1 <= 1;"),
				},
				Right: &ast.IntExpression{
					Value:    1,
					PosStart: util.NewPos(1, 6, 5, "<test>", "1 <= 1;"),
					PosEnd:   util.NewPos(1, 7, 6, "<test>", "1 <= 1;"),
				},
				PosStart: util.NewPos(1, 1, 0, "<test>", "1 <= 1;"),
				PosEnd:   util.NewPos(1, 7, 6, "<test>", "1 <= 1;"),
			},
		},
		{
			name:  "LeftShift Infix Expression",
			input: "1 << 1;",
			expected: &ast.InfixExpression{
				Left: &ast.IntExpression{
					Value:    1,
					PosStart: util.NewPos(1, 1, 0, "<test>", "1 << 1;"),
					PosEnd:   util.NewPos(1, 2, 1, "<test>", "1 << 1;"),
				},
				Operator: &lexer.Token{
					Type:     lexer.LEFT_SHIFT,
					Literal:  "<<",
					PosStart: util.NewPos(1, 3, 2, "<test>", "1 << 1;"),
					PosEnd:   util.NewPos(1, 5, 4, "<test>", "1 << 1;"),
				},
				Right: &ast.IntExpression{
					Value:    1,
					PosStart: util.NewPos(1, 6, 5, "<test>", "1 << 1;"),
					PosEnd:   util.NewPos(1, 7, 6, "<test>", "1 << 1;"),
				},
				PosStart: util.NewPos(1, 1, 0, "<test>", "1 << 1;"),
				PosEnd:   util.NewPos(1, 7, 6, "<test>", "1 << 1;"),
			},
		},
		{
			name:  "RightShift Infix Expression",
			input: "1 >> 1;",
			expected: &ast.InfixExpression{
				Left: &ast.IntExpression{
					Value:    1,
					PosStart: util.NewPos(1, 1, 0, "<test>", "1 >> 1;"),
					PosEnd:   util.NewPos(1, 2, 1, "<test>", "1 >> 1;"),
				},
				Operator: &lexer.Token{
					Type:     lexer.RIGHT_SHIFT,
					Literal:  ">>",
					PosStart: util.NewPos(1, 3, 2, "<test>", "1 >> 1;"),
					PosEnd:   util.NewPos(1, 5, 4, "<test>", "1 >> 1;"),
				},
				Right: &ast.IntExpression{
					Value:    1,
					PosStart: util.NewPos(1, 6, 5, "<test>", "1 >> 1;"),
					PosEnd:   util.NewPos(1, 7, 6, "<test>", "1 >> 1;"),
				},
				PosStart: util.NewPos(1, 1, 0, "<test>", "1 >> 1;"),
				PosEnd:   util.NewPos(1, 7, 6, "<test>", "1 >> 1;"),
			},
		},
		{
			name:  "BitAnd Infix Expression",
			input: "1 & 1;",
			expected: &ast.InfixExpression{
				Left: &ast.IntExpression{
					Value:    1,
					PosStart: util.NewPos(1, 1, 0, "<test>", "1 & 1;"),
					PosEnd:   util.NewPos(1, 2, 1, "<test>", "1 & 1;"),
				},
				Operator: &lexer.Token{
					Type:     lexer.BITWISE_AND,
					Literal:  "&",
					PosStart: util.NewPos(1, 3, 2, "<test>", "1 & 1;"),
					PosEnd:   util.NewPos(1, 4, 3, "<test>", "1 & 1;"),
				},
				Right: &ast.IntExpression{
					Value:    1,
					PosStart: util.NewPos(1, 5, 4, "<test>", "1 & 1;"),
					PosEnd:   util.NewPos(1, 6, 5, "<test>", "1 & 1;"),
				},
				PosStart: util.NewPos(1, 1, 0, "<test>", "1 & 1;"),
				PosEnd:   util.NewPos(1, 6, 5, "<test>", "1 & 1;"),
			},
		},
		{
			name:  "BitOr Infix Expression",
			input: "1 | 1;",
			expected: &ast.InfixExpression{
				Left: &ast.IntExpression{
					Value:    1,
					PosStart: util.NewPos(1, 1, 0, "<test>", "1 | 1;"),
					PosEnd:   util.NewPos(1, 2, 1, "<test>", "1 | 1;"),
				},
				Operator: &lexer.Token{
					Type:     lexer.BITWISE_OR,
					Literal:  "|",
					PosStart: util.NewPos(1, 3, 2, "<test>", "1 | 1;"),
					PosEnd:   util.NewPos(1, 4, 3, "<test>", "1 | 1;"),
				},
				Right: &ast.IntExpression{
					Value:    1,
					PosStart: util.NewPos(1, 5, 4, "<test>", "1 | 1;"),
					PosEnd:   util.NewPos(1, 6, 5, "<test>", "1 | 1;"),
				},
				PosStart: util.NewPos(1, 1, 0, "<test>", "1 | 1;"),
				PosEnd:   util.NewPos(1, 6, 5, "<test>", "1 | 1;"),
			},
		},
		{
			name:  "Xor Infix Expression",
			input: "1 ^ 1;",
			expected: &ast.InfixExpression{
				Left: &ast.IntExpression{
					Value:    1,
					PosStart: util.NewPos(1, 1, 0, "<test>", "1 ^ 1;"),
					PosEnd:   util.NewPos(1, 2, 1, "<test>", "1 ^ 1;"),
				},
				Operator: &lexer.Token{
					Type:     lexer.BITWISE_XOR,
					Literal:  "^",
					PosStart: util.NewPos(1, 3, 2, "<test>", "1 ^ 1;"),
					PosEnd:   util.NewPos(1, 4, 3, "<test>", "1 ^ 1;"),
				},
				Right: &ast.IntExpression{
					Value:    1,
					PosStart: util.NewPos(1, 5, 4, "<test>", "1 ^ 1;"),
					PosEnd:   util.NewPos(1, 6, 5, "<test>", "1 ^ 1;"),
				},
				PosStart: util.NewPos(1, 1, 0, "<test>", "1 ^ 1;"),
				PosEnd:   util.NewPos(1, 6, 5, "<test>", "1 ^ 1;"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer("<test>", tt.input)
			p, _ := NewParser(l)
			program := p.ParseProgram()
			expr := program.Statements[0].(*ast.ExpressionStatement).Expr.(*ast.InfixExpression)

			if p.Err != nil {
				t.Errorf("err = %+v, expected nil", p.Err)
			}

			if !reflect.DeepEqual(expr, tt.expected) {
				t.Errorf("expected %+v, got %+v", tt.expected, expr)
			}
		})
	}
}

func TestParser_ParseIntegerExpression(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *ast.IntExpression
	}{
		{
			name:  "Single Character Integer Expression",
			input: "1;",
			expected: &ast.IntExpression{
				Value:    1,
				PosStart: util.NewPos(1, 1, 0, "<test>", "1;"),
				PosEnd:   util.NewPos(1, 2, 1, "<test>", "1;"),
			},
		},
		{
			name:  "Multi-Character Integer Expression",
			input: "123;",
			expected: &ast.IntExpression{
				Value:    123,
				PosStart: util.NewPos(1, 1, 0, "<test>", "123;"),
				PosEnd:   util.NewPos(1, 4, 3, "<test>", "123;"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer("<test>", tt.input)
			p, _ := NewParser(l)
			program := p.ParseProgram()
			expr := program.Statements[0].(*ast.ExpressionStatement).Expr.(*ast.IntExpression)

			if p.Err != nil {
				t.Errorf("err = %+v, expected nil", p.Err)
			}
			if !reflect.DeepEqual(expr, tt.expected) {
				t.Errorf("expected %+v, got %+v", tt.expected, expr)
			}
		})
	}
}

func TestParser_ParseFloatExpression(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *ast.FloatExpression
	}{
		{
			name:  "Single Character Float Expression",
			input: "1.0;",
			expected: &ast.FloatExpression{
				Value:    1.0,
				PosStart: util.NewPos(1, 1, 0, "<test>", "1.0;"),
				PosEnd:   util.NewPos(1, 4, 3, "<test>", "1.0;"),
			},
		},
		{
			name:  "Multi-Character Float Expression",
			input: "0.75;",
			expected: &ast.FloatExpression{
				Value:    0.75,
				PosStart: util.NewPos(1, 1, 0, "<test>", "0.75;"),
				PosEnd:   util.NewPos(1, 5, 4, "<test>", "0.75;"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer("<test>", tt.input)
			p, _ := NewParser(l)
			program := p.ParseProgram()
			expr := program.Statements[0].(*ast.ExpressionStatement).Expr.(*ast.FloatExpression)

			if p.Err != nil {
				t.Errorf("err = %+v, expected nil", p.Err)
			}

			if !reflect.DeepEqual(expr, tt.expected) {
				t.Errorf("expected %+v, got %+v", tt.expected, expr)
			}
		})
	}
}

func TestParser_ParseIdentifierExpression(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *ast.IdentifierExpression
	}{
		{
			name:  "Identifier",
			input: "x;",
			expected: &ast.IdentifierExpression{
				Name:     "x",
				PosStart: util.NewPos(1, 1, 0, "<test>", "x;"),
				PosEnd:   util.NewPos(1, 2, 1, "<test>", "x;"),
			},
		},
		{
			name:  "Identifier with Chinese Character",
			input: "真;",
			expected: &ast.IdentifierExpression{
				Name:     "真",
				PosStart: util.NewPos(1, 1, 0, "<test>", "真;"),
				PosEnd:   util.NewPos(1, 2, 3, "<test>", "真;"),
			},
		},
		{
			name:  "Identifier with Multi-Chinese-Character",
			input: "你好;",
			expected: &ast.IdentifierExpression{
				Name:     "你好",
				PosStart: util.NewPos(1, 1, 0, "<test>", "你好;"),
				PosEnd:   util.NewPos(1, 3, 6, "<test>", "你好;"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer("<test>", tt.input)
			p, _ := NewParser(l)
			program := p.ParseProgram()
			expr := program.Statements[0].(*ast.ExpressionStatement).Expr.(*ast.IdentifierExpression)

			if p.Err != nil {
				t.Errorf("err = %+v, expected nil", p.Err)
			}

			if !reflect.DeepEqual(expr, tt.expected) {
				t.Errorf("expected %+v, got %+v", tt.expected, expr)
			}
		})
	}
}

func TestParser_ParseBooleanExpression(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *ast.BoolExpression
	}{
		{
			name:  "True Boolean Expression",
			input: "true;",
			expected: &ast.BoolExpression{
				Value:    true,
				PosStart: util.NewPos(1, 1, 0, "<test>", "true;"),
				PosEnd:   util.NewPos(1, 5, 4, "<test>", "true;"),
			},
		},
		{
			name:  "False Boolean Expression",
			input: "false;",
			expected: &ast.BoolExpression{
				Value:    false,
				PosStart: util.NewPos(1, 1, 0, "<test>", "false;"),
				PosEnd:   util.NewPos(1, 6, 5, "<test>", "false;"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer("<test>", tt.input)
			p, _ := NewParser(l)
			program := p.ParseProgram()
			expr := program.Statements[0].(*ast.ExpressionStatement).Expr.(*ast.BoolExpression)

			if p.Err != nil {
				t.Errorf("err = %+v, expected nil", p.Err)
			}

			if !reflect.DeepEqual(expr, tt.expected) {
				t.Errorf("expected %+v, got %+v", tt.expected, expr)
			}
		})
	}
}

func TestParser_ParseNullExpression(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *ast.NullExpression
	}{
		{
			name:  "Null Expression",
			input: "null;",
			expected: &ast.NullExpression{
				PosStart: util.NewPos(1, 1, 0, "<test>", "null;"),
				PosEnd:   util.NewPos(1, 5, 4, "<test>", "null;"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer("<test>", tt.input)
			p, _ := NewParser(l)
			program := p.ParseProgram()
			expr := program.Statements[0].(*ast.ExpressionStatement).Expr.(*ast.NullExpression)

			if p.Err != nil {
				t.Errorf("err = %+v, expected nil", p.Err)
			}

			if !reflect.DeepEqual(expr, tt.expected) {
				t.Errorf("expected %+v, got %+v", tt.expected, expr)
			}
		})
	}
}

func TestParser_ParseStringExpression(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *ast.StringExpression
	}{
		{
			name:  "Single-Line String Expression",
			input: "\"hello\";",
			expected: &ast.StringExpression{
				Value:    "hello",
				PosStart: util.NewPos(1, 1, 0, "<test>", "\"hello\";"),
				PosEnd:   util.NewPos(1, 8, 7, "<test>", "\"hello\";"),
			},
		},
		{
			name:  "Multi-Line String Expression",
			input: "\"hello\nworld\";",
			expected: &ast.StringExpression{
				Value:    "hello\nworld",
				PosStart: util.NewPos(1, 1, 0, "<test>", "\"hello\nworld\";"),
				PosEnd:   util.NewPos(2, 7, 13, "<test>", "\"hello\nworld\";"),
			},
		},
		{
			name:  "String Expression with Escape Character",
			input: "\"hello\\nworld\\t\";",
			expected: &ast.StringExpression{
				Value:    "hello\nworld\t",
				PosStart: util.NewPos(1, 1, 0, "<test>", "\"hello\\nworld\\t\";"),
				PosEnd:   util.NewPos(1, 17, 16, "<test>", "\"hello\\nworld\\t\";"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer("<test>", tt.input)
			p, _ := NewParser(l)
			program := p.ParseProgram()
			expr := program.Statements[0].(*ast.ExpressionStatement).Expr.(*ast.StringExpression)

			if p.Err != nil {
				t.Errorf("err = %+v, expected nil", p.Err)
			}

			if !reflect.DeepEqual(expr, tt.expected) {
				t.Errorf("expected %+v, got %+v", tt.expected, expr)
			}
		})
	}
}

func TestParser_ParseGroupedExpression(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *ast.ExpressionStatement
	}{
		{
			name:  "Grouped Expression",
			input: "(1 + 2);",
			expected: &ast.ExpressionStatement{
				Expr: &ast.GroupedExpression{
					Expr: &ast.InfixExpression{
						Left: &ast.IntExpression{
							Value:    1,
							PosStart: util.NewPos(1, 2, 1, "<test>", "(1 + 2);"),
							PosEnd:   util.NewPos(1, 3, 2, "<test>", "(1 + 2);"),
						},
						Operator: &lexer.Token{
							Type:     lexer.PLUS,
							Literal:  "+",
							PosStart: util.NewPos(1, 4, 3, "<test>", "(1 + 2);"),
							PosEnd:   util.NewPos(1, 5, 4, "<test>", "(1 + 2);"),
						},
						Right: &ast.IntExpression{
							Value:    2,
							PosStart: util.NewPos(1, 6, 5, "<test>", "(1 + 2);"),
							PosEnd:   util.NewPos(1, 7, 6, "<test>", "(1 + 2);"),
						},
						PosStart: util.NewPos(1, 2, 1, "<test>", "(1 + 2);"),
						PosEnd:   util.NewPos(1, 7, 6, "<test>", "(1 + 2);"),
					},
					PosStart: util.NewPos(1, 1, 0, "<test>", "(1 + 2);"),
					PosEnd:   util.NewPos(1, 8, 7, "<test>", "(1 + 2);"),
				},
				PosStart: util.NewPos(1, 1, 0, "<test>", "(1 + 2);"),
				PosEnd:   util.NewPos(1, 8, 7, "<test>", "(1 + 2);"),
			},
		},
		{
			name:  "Grouped Expression to Change Priority",
			input: "(1 + 2) * 3;",
			expected: &ast.ExpressionStatement{
				Expr: &ast.InfixExpression{
					Left: &ast.GroupedExpression{
						Expr: &ast.InfixExpression{
							Left: &ast.IntExpression{
								Value:    1,
								PosStart: util.NewPos(1, 2, 1, "<test>", "(1 + 2) * 3;"),
								PosEnd:   util.NewPos(1, 3, 2, "<test>", "(1 + 2) * 3;"),
							},
							Operator: &lexer.Token{
								Type:     lexer.PLUS,
								Literal:  "+",
								PosStart: util.NewPos(1, 4, 3, "<test>", "(1 + 2) * 3;"),
								PosEnd:   util.NewPos(1, 5, 4, "<test>", "(1 + 2) * 3;"),
							},
							Right: &ast.IntExpression{
								Value:    2,
								PosStart: util.NewPos(1, 6, 5, "<test>", "(1 + 2) * 3;"),
								PosEnd:   util.NewPos(1, 7, 6, "<test>", "(1 + 2) * 3;"),
							},
							PosStart: util.NewPos(1, 2, 1, "<test>", "(1 + 2) * 3;"),
							PosEnd:   util.NewPos(1, 7, 6, "<test>", "(1 + 2) * 3;"),
						},
						PosStart: util.NewPos(1, 1, 0, "<test>", "(1 + 2) * 3;"),
						PosEnd:   util.NewPos(1, 8, 7, "<test>", "(1 + 2) * 3;"),
					},
					Operator: &lexer.Token{
						Type:     lexer.ASTERISK,
						Literal:  "*",
						PosStart: util.NewPos(1, 9, 8, "<test>", "(1 + 2) * 3;"),
						PosEnd:   util.NewPos(1, 10, 9, "<test>", "(1 + 2) * 3;"),
					},
					Right: &ast.IntExpression{
						Value:    3,
						PosStart: util.NewPos(1, 11, 10, "<test>", "(1 + 2) * 3;"),
						PosEnd:   util.NewPos(1, 12, 11, "<test>", "(1 + 2) * 3;"),
					},
					PosStart: util.NewPos(1, 1, 0, "<test>", "(1 + 2) * 3;"),
					PosEnd:   util.NewPos(1, 12, 11, "<test>", "(1 + 2) * 3;"),
				},
				PosStart: util.NewPos(1, 1, 0, "<test>", "(1 + 2) * 3;"),
				PosEnd:   util.NewPos(1, 12, 11, "<test>", "(1 + 2) * 3;"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer("<test>", tt.input)
			p, _ := NewParser(l)
			program := p.ParseProgram()
			expr := program.Statements[0].(*ast.ExpressionStatement)

			if p.Err != nil {
				t.Errorf("err = %+v, expected nil", p.Err)
			}

			if !reflect.DeepEqual(expr, tt.expected) {
				t.Errorf("expected %+v, got %+v", tt.expected, expr)
			}
		})
	}
}

func TestParser_ParseVarInitializationExpression(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *ast.VarInitializationExpression
	}{
		{
			name:  "Simple Var Expression",
			input: "var a = 1;",
			expected: &ast.VarInitializationExpression{
				IsConst: false,
				Name: &ast.IdentifierExpression{
					Name:     "a",
					PosStart: util.NewPos(1, 5, 4, "<test>", "var a = 1;"),
					PosEnd:   util.NewPos(1, 6, 5, "<test>", "var a = 1;"),
				},
				Value: &ast.IntExpression{
					Value:    1,
					PosStart: util.NewPos(1, 9, 8, "<test>", "var a = 1;"),
					PosEnd:   util.NewPos(1, 10, 9, "<test>", "var a = 1;"),
				},
				PosStart: util.NewPos(1, 1, 0, "<test>", "var a = 1;"),
				PosEnd:   util.NewPos(1, 10, 9, "<test>", "var a = 1;"),
			},
		},
		{
			name:  "Simple Const Expression",
			input: "const a = 1;",
			expected: &ast.VarInitializationExpression{
				IsConst: true,
				Name: &ast.IdentifierExpression{
					Name:     "a",
					PosStart: util.NewPos(1, 7, 6, "<test>", "const a = 1;"),
					PosEnd:   util.NewPos(1, 8, 7, "<test>", "const a = 1;"),
				},
				Value: &ast.IntExpression{
					Value:    1,
					PosStart: util.NewPos(1, 11, 10, "<test>", "const a = 1;"),
					PosEnd:   util.NewPos(1, 12, 11, "<test>", "const a = 1;"),
				},
				PosStart: util.NewPos(1, 1, 0, "<test>", "const a = 1;"),
				PosEnd:   util.NewPos(1, 12, 11, "<test>", "const a = 1;"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer("<test>", tt.input)
			p, _ := NewParser(l)
			program := p.ParseProgram()
			expr := program.Statements[0].(*ast.ExpressionStatement).Expr.(*ast.VarInitializationExpression)

			if p.Err != nil {
				t.Errorf("err = %+v, expected nil", p.Err)
			}

			if !reflect.DeepEqual(expr, tt.expected) {
				t.Errorf("expected %+v, got %+v", tt.expected, expr)
			}
		})
	}
}

func TestParser_ParseCompoundAssignmentExpression(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *ast.CompoundAssignmentExpression
	}{
		{
			name:  "Simple Compound Assignment Expression",
			input: "a += 1;",
			expected: &ast.CompoundAssignmentExpression{
				Name: &ast.IdentifierExpression{
					Name:     "a",
					PosStart: util.NewPos(1, 1, 0, "<test>", "a += 1;"),
					PosEnd:   util.NewPos(1, 2, 1, "<test>", "a += 1;"),
				},
				Operator: &lexer.Token{
					Type:     lexer.PLUS_EQUAL,
					Literal:  "+=",
					PosStart: util.NewPos(1, 3, 2, "<test>", "a += 1;"),
					PosEnd:   util.NewPos(1, 5, 4, "<test>", "a += 1;"),
				},
				Right: &ast.IntExpression{
					Value:    1,
					PosStart: util.NewPos(1, 6, 5, "<test>", "a += 1;"),
					PosEnd:   util.NewPos(1, 7, 6, "<test>", "a += 1;"),
				},
				PosStart: util.NewPos(1, 1, 0, "<test>", "a += 1;"),
				PosEnd:   util.NewPos(1, 7, 6, "<test>", "a += 1;"),
			},
		},
		{
			name:  "Simple Compound Assignment Expression",
			input: "a -= 1;",
			expected: &ast.CompoundAssignmentExpression{
				Name: &ast.IdentifierExpression{
					Name:     "a",
					PosStart: util.NewPos(1, 1, 0, "<test>", "a -= 1;"),
					PosEnd:   util.NewPos(1, 2, 1, "<test>", "a -= 1;"),
				},
				Operator: &lexer.Token{
					Type:     lexer.MINUS_EQUAL,
					Literal:  "-=",
					PosStart: util.NewPos(1, 3, 2, "<test>", "a -= 1;"),
					PosEnd:   util.NewPos(1, 5, 4, "<test>", "a -= 1;"),
				},
				Right: &ast.IntExpression{
					Value:    1,
					PosStart: util.NewPos(1, 6, 5, "<test>", "a -= 1;"),
					PosEnd:   util.NewPos(1, 7, 6, "<test>", "a -= 1;"),
				},
				PosStart: util.NewPos(1, 1, 0, "<test>", "a -= 1;"),
				PosEnd:   util.NewPos(1, 7, 6, "<test>", "a -= 1;"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer("<test>", tt.input)
			p, _ := NewParser(l)
			program := p.ParseProgram()
			expr := program.Statements[0].(*ast.ExpressionStatement).Expr.(*ast.CompoundAssignmentExpression)

			if p.Err != nil {
				t.Errorf("err = %+v, expected nil", p.Err)
			}

			if !reflect.DeepEqual(expr, tt.expected) {
				t.Errorf("expected %+v, got %+v", tt.expected, expr)
			}
		})
	}
}

func TestParser_ParsePrefixUnaryIncDecExpression(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *ast.PrefixUnaryIncDecExpression
	}{
		{
			name:  "Prefix Unary Increment Expression",
			input: "++a;",
			expected: &ast.PrefixUnaryIncDecExpression{
				Operator: &lexer.Token{
					Type:     lexer.INCREMENT,
					Literal:  "++",
					PosStart: util.NewPos(1, 1, 0, "<test>", "++a;"),
					PosEnd:   util.NewPos(1, 3, 2, "<test>", "++a;"),
				},
				Right: &ast.IdentifierExpression{
					Name:     "a",
					PosStart: util.NewPos(1, 3, 2, "<test>", "++a;"),
					PosEnd:   util.NewPos(1, 4, 3, "<test>", "++a;"),
				},
				PosStart: util.NewPos(1, 1, 0, "<test>", "++a;"),
				PosEnd:   util.NewPos(1, 4, 3, "<test>", "++a;"),
			},
		},
		{
			name:  "Prefix Unary Decrement Expression",
			input: "--a;",
			expected: &ast.PrefixUnaryIncDecExpression{
				Operator: &lexer.Token{
					Type:     lexer.DECREMENT,
					Literal:  "--",
					PosStart: util.NewPos(1, 1, 0, "<test>", "--a;"),
					PosEnd:   util.NewPos(1, 3, 2, "<test>", "--a;"),
				},
				Right: &ast.IdentifierExpression{
					Name:     "a",
					PosStart: util.NewPos(1, 3, 2, "<test>", "--a;"),
					PosEnd:   util.NewPos(1, 4, 3, "<test>", "--a;"),
				},
				PosStart: util.NewPos(1, 1, 0, "<test>", "--a;"),
				PosEnd:   util.NewPos(1, 4, 3, "<test>", "--a;"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer("<test>", tt.input)
			p, _ := NewParser(l)
			program := p.ParseProgram()
			expr := program.Statements[0].(*ast.ExpressionStatement).Expr.(*ast.PrefixUnaryIncDecExpression)

			if p.Err != nil {
				t.Errorf("err = %+v, expected nil", p.Err)
			}

			if !reflect.DeepEqual(expr, tt.expected) {
				t.Errorf("expected %+v, got %+v", tt.expected, expr)
			}
		})
	}
}

func TestParser_ParsePostfixUnaryIncDecExpression(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *ast.PostfixUnaryIncDecExpression
	}{
		{
			name:  "Postfix Unary Increment Expression",
			input: "a++;",
			expected: &ast.PostfixUnaryIncDecExpression{
				Left: &ast.IdentifierExpression{
					Name:     "a",
					PosStart: util.NewPos(1, 1, 0, "<test>", "a++;"),
					PosEnd:   util.NewPos(1, 2, 1, "<test>", "a++;"),
				},
				Operator: &lexer.Token{
					Type:     lexer.INCREMENT,
					Literal:  "++",
					PosStart: util.NewPos(1, 2, 1, "<test>", "a++;"),
					PosEnd:   util.NewPos(1, 4, 3, "<test>", "a++;"),
				},
				PosStart: util.NewPos(1, 1, 0, "<test>", "a++;"),
				PosEnd:   util.NewPos(1, 4, 3, "<test>", "a++;"),
			},
		},
		{
			name:  "Postfix Unary Decrement Expression",
			input: "a--;",
			expected: &ast.PostfixUnaryIncDecExpression{
				Left: &ast.IdentifierExpression{
					Name:     "a",
					PosStart: util.NewPos(1, 1, 0, "<test>", "a--;"),
					PosEnd:   util.NewPos(1, 2, 1, "<test>", "a--;"),
				},
				Operator: &lexer.Token{
					Type:     lexer.DECREMENT,
					Literal:  "--",
					PosStart: util.NewPos(1, 2, 1, "<test>", "a--;"),
					PosEnd:   util.NewPos(1, 4, 3, "<test>", "a--;"),
				},
				PosStart: util.NewPos(1, 1, 0, "<test>", "a--;"),
				PosEnd:   util.NewPos(1, 4, 3, "<test>", "a--;"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer("<test>", tt.input)
			p, _ := NewParser(l)
			program := p.ParseProgram()
			expr := program.Statements[0].(*ast.ExpressionStatement).Expr.(*ast.PostfixUnaryIncDecExpression)

			if p.Err != nil {
				t.Errorf("err = %+v, expected nil", p.Err)
			}

			if !reflect.DeepEqual(expr, tt.expected) {
				t.Errorf("expected %+v, got %+v", tt.expected, expr)
			}
		})
	}
}

func TestParser_ParseBlockExpression(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *ast.BlockExpression
	}{
		{
			name:  "Block Expression",
			input: "{ 1 };",
			expected: &ast.BlockExpression{
				Statements: []ast.Statement{
					&ast.ExpressionStatement{
						Expr: &ast.IntExpression{
							Value:    1,
							PosStart: util.NewPos(1, 3, 2, "<test>", "{ 1 };"),
							PosEnd:   util.NewPos(1, 4, 3, "<test>", "{ 1 };"),
						},
						PosStart: util.NewPos(1, 3, 2, "<test>", "{ 1 };"),
						PosEnd:   util.NewPos(1, 4, 3, "<test>", "{ 1 };"),
					},
				},
				PosStart: util.NewPos(1, 1, 0, "<test>", "{ 1 };"),
				PosEnd:   util.NewPos(1, 6, 5, "<test>", "{ 1 };"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer("<test>", tt.input)
			p, _ := NewParser(l)
			program := p.ParseProgram()
			expr := program.Statements[0].(*ast.ExpressionStatement).Expr.(*ast.BlockExpression)

			if p.Err != nil {
				t.Errorf("err = %+v, expected nil", p.Err)
			}

			if !reflect.DeepEqual(expr, tt.expected) {
				t.Errorf("expected %+v, got %+v", tt.expected, expr)
			}
		})
	}
}

func TestParser_ParseIfExpression(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *ast.IfExpression
	}{
		{
			name:  "Simple If Expression",
			input: "if true 1;",
			expected: &ast.IfExpression{
				Condition: &ast.BoolExpression{
					Value:    true,
					PosStart: util.NewPos(1, 4, 3, "<test>", "if true 1;"),
					PosEnd:   util.NewPos(1, 8, 7, "<test>", "if true 1;"),
				},
				Consequence: &ast.ExpressionStatement{
					Expr: &ast.IntExpression{
						Value:    1,
						PosStart: util.NewPos(1, 9, 8, "<test>", "if true 1;"),
						PosEnd:   util.NewPos(1, 10, 9, "<test>", "if true 1;"),
					},
					PosStart: util.NewPos(1, 9, 8, "<test>", "if true 1;"),
					PosEnd:   util.NewPos(1, 10, 9, "<test>", "if true 1;"),
				},
				Alternative: nil,
				PosStart:    util.NewPos(1, 1, 0, "<test>", "if true 1;"),
				PosEnd:      util.NewPos(1, 10, 9, "<test>", "if true 1;"),
			},
		},
		{
			name:  "If-Else Expression",
			input: "if false 1 else 2;",
			expected: &ast.IfExpression{
				Condition: &ast.BoolExpression{
					Value:    false,
					PosStart: util.NewPos(1, 4, 3, "<test>", "if false 1 else 2;"),
					PosEnd:   util.NewPos(1, 9, 8, "<test>", "if false 1 else 2;"),
				},
				Consequence: &ast.ExpressionStatement{
					Expr: &ast.IntExpression{
						Value:    1,
						PosStart: util.NewPos(1, 10, 9, "<test>", "if false 1 else 2;"),
						PosEnd:   util.NewPos(1, 11, 10, "<test>", "if false 1 else 2;"),
					},
					PosStart: util.NewPos(1, 10, 9, "<test>", "if false 1 else 2;"),
					PosEnd:   util.NewPos(1, 11, 10, "<test>", "if false 1 else 2;"),
				},
				Alternative: &ast.ExpressionStatement{
					Expr: &ast.IntExpression{
						Value:    2,
						PosStart: util.NewPos(1, 17, 16, "<test>", "if false 1 else 2;"),
						PosEnd:   util.NewPos(1, 18, 17, "<test>", "if false 1 else 2;"),
					},
					PosStart: util.NewPos(1, 17, 16, "<test>", "if false 1 else 2;"),
					PosEnd:   util.NewPos(1, 18, 17, "<test>", "if false 1 else 2;"),
				},
				PosStart: util.NewPos(1, 1, 0, "<test>", "if false 1 else 2;"),
				PosEnd:   util.NewPos(1, 18, 17, "<test>", "if false 1 else 2;"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer("<test>", tt.input)
			p, _ := NewParser(l)
			program := p.ParseProgram()
			expr := program.Statements[0].(*ast.ExpressionStatement).Expr.(*ast.IfExpression)

			if p.Err != nil {
				t.Errorf("err = %+v, expected nil", p.Err)
			}

			if !reflect.DeepEqual(expr, tt.expected) {
				t.Errorf("expected %+v, got %+v", tt.expected, expr)
			}
		})
	}
}

func TestParser_ParseCallExpression(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *ast.CallExpression
	}{
		{
			name:  "Call Expression",
			input: "f();",
			expected: &ast.CallExpression{
				Function: &ast.IdentifierExpression{
					Name:     "f",
					PosStart: util.NewPos(1, 1, 0, "<test>", "f();"),
					PosEnd:   util.NewPos(1, 2, 1, "<test>", "f();"),
				},
				Argument: make([]ast.Expression, 0),
				PosStart: util.NewPos(1, 1, 0, "<test>", "f();"),
				PosEnd:   util.NewPos(1, 4, 3, "<test>", "f();"),
			},
		},
		{
			name:  "Call Expression",
			input: "f(1, , 2);",
			expected: &ast.CallExpression{
				Function: &ast.IdentifierExpression{
					Name:     "f",
					PosStart: util.NewPos(1, 1, 0, "<test>", "f(1, , 2);"),
					PosEnd:   util.NewPos(1, 2, 1, "<test>", "f(1, , 2);"),
				},
				Argument: []ast.Expression{
					&ast.IntExpression{
						Value:    1,
						PosStart: util.NewPos(1, 3, 2, "<test>", "f(1, , 2);"),
						PosEnd:   util.NewPos(1, 4, 3, "<test>", "f(1, , 2);"),
					},
					nil,
					&ast.IntExpression{
						Value:    2,
						PosStart: util.NewPos(1, 8, 7, "<test>", "f(1, , 2);"),
						PosEnd:   util.NewPos(1, 9, 8, "<test>", "f(1, , 2);"),
					},
				},
				PosStart: util.NewPos(1, 1, 0, "<test>", "f(1, , 2);"),
				PosEnd:   util.NewPos(1, 10, 9, "<test>", "f(1, , 2);"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer("<test>", tt.input)
			p, _ := NewParser(l)
			program := p.ParseProgram()
			expr := program.Statements[0].(*ast.ExpressionStatement).Expr.(*ast.CallExpression)

			if p.Err != nil {
				t.Errorf("err = %+v, expected nil", p.Err)
			}

			if !reflect.DeepEqual(expr, tt.expected) {
				t.Errorf("expected %+v, got %+v", tt.expected, expr)
			}
		})
	}
}

func TestParser_ParseListExpression(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *ast.ListExpression
	}{
		{
			name:  "Empty List Expression",
			input: "[];",
			expected: &ast.ListExpression{
				Value:    make([]ast.Expression, 0),
				PosStart: util.NewPos(1, 1, 0, "<test>", "[];"),
				PosEnd:   util.NewPos(1, 3, 2, "<test>", "[];"),
			},
		},
		{
			name:  "Single Integer Element List",
			input: "[1];",
			expected: &ast.ListExpression{
				Value: []ast.Expression{
					&ast.IntExpression{
						Value:    1,
						PosStart: util.NewPos(1, 2, 1, "<test>", "[1];"),
						PosEnd:   util.NewPos(1, 3, 2, "<test>", "[1];"),
					},
				},
				PosStart: util.NewPos(1, 1, 0, "<test>", "[1];"),
				PosEnd:   util.NewPos(1, 4, 3, "<test>", "[1];"),
			},
		},
		{
			name:  "Multiple Integer Elements List",
			input: "[1, 2, 3];",
			expected: &ast.ListExpression{
				Value: []ast.Expression{
					&ast.IntExpression{
						Value:    1,
						PosStart: util.NewPos(1, 2, 1, "<test>", "[1, 2, 3];"),
						PosEnd:   util.NewPos(1, 3, 2, "<test>", "[1, 2, 3];"),
					},
					&ast.IntExpression{
						Value:    2,
						PosStart: util.NewPos(1, 5, 4, "<test>", "[1, 2, 3];"),
						PosEnd:   util.NewPos(1, 6, 5, "<test>", "[1, 2, 3];"),
					},
					&ast.IntExpression{
						Value:    3,
						PosStart: util.NewPos(1, 8, 7, "<test>", "[1, 2, 3];"),
						PosEnd:   util.NewPos(1, 9, 8, "<test>", "[1, 2, 3];"),
					},
				},
				PosStart: util.NewPos(1, 1, 0, "<test>", "[1, 2, 3];"),
				PosEnd:   util.NewPos(1, 10, 9, "<test>", "[1, 2, 3];"),
			},
		},
		{
			name:  "Multiple String Elements List",
			input: "[\"hello\", \"world\"];",
			expected: &ast.ListExpression{
				Value: []ast.Expression{
					&ast.StringExpression{
						Value:    "hello",
						PosStart: util.NewPos(1, 2, 1, "<test>", "[\"hello\", \"world\"];"),
						PosEnd:   util.NewPos(1, 9, 8, "<test>", "[\"hello\", \"world\"];"),
					},
					&ast.StringExpression{
						Value:    "world",
						PosStart: util.NewPos(1, 11, 10, "<test>", "[\"hello\", \"world\"];"),
						PosEnd:   util.NewPos(1, 18, 17, "<test>", "[\"hello\", \"world\"];"),
					},
				},
				PosStart: util.NewPos(1, 1, 0, "<test>", "[\"hello\", \"world\"];"),
				PosEnd:   util.NewPos(1, 19, 18, "<test>", "[\"hello\", \"world\"];"),
			},
		},
		{
			name:  "Boolean Elements List",
			input: "[true, false, true];",
			expected: &ast.ListExpression{
				Value: []ast.Expression{
					&ast.BoolExpression{
						Value:    true,
						PosStart: util.NewPos(1, 2, 1, "<test>", "[true, false, true];"),
						PosEnd:   util.NewPos(1, 6, 5, "<test>", "[true, false, true];"),
					},
					&ast.BoolExpression{
						Value:    false,
						PosStart: util.NewPos(1, 8, 7, "<test>", "[true, false, true];"),
						PosEnd:   util.NewPos(1, 13, 12, "<test>", "[true, false, true];"),
					},
					&ast.BoolExpression{
						Value:    true,
						PosStart: util.NewPos(1, 15, 14, "<test>", "[true, false, true];"),
						PosEnd:   util.NewPos(1, 19, 18, "<test>", "[true, false, true];"),
					},
				},
				PosStart: util.NewPos(1, 1, 0, "<test>", "[true, false, true];"),
				PosEnd:   util.NewPos(1, 20, 19, "<test>", "[true, false, true];"),
			},
		},
		{
			name:  "Expression Elements List",
			input: "[1 + 2, 3 + 4];",
			expected: &ast.ListExpression{
				Value: []ast.Expression{
					&ast.InfixExpression{
						Left: &ast.IntExpression{
							Value:    1,
							PosStart: util.NewPos(1, 2, 1, "<test>", "[1 + 2, 3 + 4];"),
							PosEnd:   util.NewPos(1, 3, 2, "<test>", "[1 + 2, 3 + 4];"),
						},
						Operator: &lexer.Token{
							Type:     lexer.PLUS,
							Literal:  "+",
							PosStart: util.NewPos(1, 4, 3, "<test>", "[1 + 2, 3 + 4];"),
							PosEnd:   util.NewPos(1, 5, 4, "<test>", "[1 + 2, 3 + 4];"),
						},
						Right: &ast.IntExpression{
							Value:    2,
							PosStart: util.NewPos(1, 6, 5, "<test>", "[1 + 2, 3 + 4];"),
							PosEnd:   util.NewPos(1, 7, 6, "<test>", "[1 + 2, 3 + 4];"),
						},
						PosStart: util.NewPos(1, 2, 1, "<test>", "[1 + 2, 3 + 4];"),
						PosEnd:   util.NewPos(1, 7, 6, "<test>", "[1 + 2, 3 + 4];"),
					},
					&ast.InfixExpression{
						Left: &ast.IntExpression{
							Value:    3,
							PosStart: util.NewPos(1, 9, 8, "<test>", "[1 + 2, 3 + 4];"),
							PosEnd:   util.NewPos(1, 10, 9, "<test>", "[1 + 2, 3 + 4];"),
						},
						Operator: &lexer.Token{
							Type:     lexer.PLUS,
							Literal:  "+",
							PosStart: util.NewPos(1, 11, 10, "<test>", "[1 + 2, 3 + 4];"),
							PosEnd:   util.NewPos(1, 12, 11, "<test>", "[1 + 2, 3 + 4];"),
						},
						Right: &ast.IntExpression{
							Value:    4,
							PosStart: util.NewPos(1, 13, 12, "<test>", "[1 + 2, 3 + 4];"),
							PosEnd:   util.NewPos(1, 14, 13, "<test>", "[1 + 2, 3 + 4];"),
						},
						PosStart: util.NewPos(1, 9, 8, "<test>", "[1 + 2, 3 + 4];"),
						PosEnd:   util.NewPos(1, 14, 13, "<test>", "[1 + 2, 3 + 4];"),
					},
				},
				PosStart: util.NewPos(1, 1, 0, "<test>", "[1 + 2, 3 + 4];"),
				PosEnd:   util.NewPos(1, 15, 14, "<test>", "[1 + 2, 3 + 4];"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer("<test>", tt.input)
			p, _ := NewParser(l)
			program := p.ParseProgram()
			expr := program.Statements[0].(*ast.ExpressionStatement).Expr.(*ast.ListExpression)

			if p.Err != nil {
				t.Errorf("err = %+v, expected nil", p.Err)
			}

			if !reflect.DeepEqual(expr, tt.expected) {
				t.Errorf("expected %+v, got %+v", tt.expected, expr)
			}
		})
	}
}

func TestParser_ParseIndexExpression_Structure(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *ast.IndexExpression
	}{
		{
			name:  "List Index Simple",
			input: "[1][0];",
			expected: &ast.IndexExpression{
				Target: &ast.ListExpression{
					Value: []ast.Expression{
						&ast.IntExpression{
							Value:    1,
							PosStart: util.NewPos(1, 2, 1, "<test>", "[1][0];"),
							PosEnd:   util.NewPos(1, 3, 2, "<test>", "[1][0];"),
						},
					},
					PosStart: util.NewPos(1, 1, 0, "<test>", "[1][0];"),
					PosEnd:   util.NewPos(1, 4, 3, "<test>", "[1][0];"),
				},
				Index: &ast.IntExpression{
					Value:    0,
					PosStart: util.NewPos(1, 5, 4, "<test>", "[1][0];"),
					PosEnd:   util.NewPos(1, 6, 5, "<test>", "[1][0];"),
				},
				PosStart: util.NewPos(1, 1, 0, "<test>", "[1][0];"),
				PosEnd:   util.NewPos(1, 7, 6, "<test>", "[1][0];"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer("<test>", tt.input)
			p, _ := NewParser(l)
			program := p.ParseProgram()
			expr := program.Statements[0].(*ast.ExpressionStatement).Expr.(*ast.IndexExpression)

			if p.Err != nil {
				t.Errorf("err = %+v, expected nil", p.Err)
			}

			if !reflect.DeepEqual(expr, tt.expected) {
				t.Errorf("expected %+v, got %+v", tt.expected, expr)
			}
		})
	}
}

func TestParser_Errors(t *testing.T) {
	tests := []struct {
		name  string
		input string
		err   error
	}{
		{
			name:  "Invalid Prefix Expression",
			input: "*1;",
			err: &SyntaxError{
				Message:  "unexpected \"ASTERISK\".",
				PosStart: util.NewPos(1, 1, 0, "<test>", "*1;"),
				PosEnd:   util.NewPos(1, 2, 1, "<test>", "*1;"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer("<test>", tt.input)
			p, _ := NewParser(l)
			p.ParseProgram()
			if p.Err.Error() != tt.err.Error() {
				t.Errorf("err = %+v, expected %+v", p.Err, tt.err)
			}
		})
	}
}
