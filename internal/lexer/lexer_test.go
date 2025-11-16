package lexer

import (
	"reflect"
	"testing"

	"github.com/Ghost-Xiao/ghost-lang/internal/util"
)

func TestLexer_Comments(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect *Token
	}{
		{
			name:  "Empty Input",
			input: "",
			expect: &Token{
				Type:     EOF,
				Literal:  "EOF",
				PosStart: util.NewPos(1, 1, 0, "<test>", ""),
				PosEnd:   util.NewPos(1, 2, 1, "<test>", ""),
			},
		},
		{
			name:  "Single Line Comment",
			input: "// This is a comment\n5",
			expect: &Token{
				Type:     INT,
				Literal:  "5",
				PosStart: util.NewPos(2, 1, 21, "<test>", "// This is a comment\n5"),
				PosEnd:   util.NewPos(2, 2, 22, "<test>", "// This is a comment\n5"),
			},
		},
		{
			name:  "Multi-Line Comment",
			input: "/* This is a multiline comment */",
			expect: &Token{
				Type:     EOF,
				Literal:  "EOF",
				PosStart: util.NewPos(1, 34, 33, "<test>", "/* This is a multiline comment */"),
				PosEnd:   util.NewPos(1, 35, 34, "<test>", "/* This is a multiline comment */"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLexer("<test>", tt.input)
			tok, err := l.NextToken()
			if err != nil {
				t.Errorf("err = %+v, expected nil", err)
			}
			if !reflect.DeepEqual(tok, tt.expect) {
				t.Errorf("tok = %+v, expected %+v", tok, tt.expect)
			}
		})
	}
}

func TestLexer_Numbers(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect *Token
	}{
		{
			name:  "Single Digit Int",
			input: "5",
			expect: &Token{
				Type:     INT,
				Literal:  "5",
				PosStart: util.NewPos(1, 1, 0, "<test>", "5"),
				PosEnd:   util.NewPos(1, 2, 1, "<test>", "5"),
			},
		},
		{
			name:  "Multi-Digit Int",
			input: "123",
			expect: &Token{
				Type:     INT,
				Literal:  "123",
				PosStart: util.NewPos(1, 1, 0, "<test>", "123"),
				PosEnd:   util.NewPos(1, 4, 3, "<test>", "123"),
			},
		},
		{
			name:  "Float",
			input: "12.34",
			expect: &Token{
				Type:     FLOAT,
				Literal:  "12.34",
				PosStart: util.NewPos(1, 1, 0, "<test>", "12.34"),
				PosEnd:   util.NewPos(1, 6, 5, "<test>", "12.34"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLexer("<test>", tt.input)
			tok, err := l.NextToken()
			if err != nil {
				t.Errorf("err = %+v, expected nil", err)
			}
			if !reflect.DeepEqual(tok, tt.expect) {
				t.Errorf("tok = %+v, expected %+v", tok, tt.expect)
			}
		})
	}
}

func TestLexer_Identifiers(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect *Token
	}{
		{
			name:  "Single Character Identifier",
			input: "a",
			expect: &Token{
				Type:     IDENT,
				Literal:  "a",
				PosStart: util.NewPos(1, 1, 0, "<test>", "a"),
				PosEnd:   util.NewPos(1, 2, 1, "<test>", "a"),
			},
		},
		{
			name:  "Multi-Character Identifier",
			input: "abc123",
			expect: &Token{
				Type:     IDENT,
				Literal:  "abc123",
				PosStart: util.NewPos(1, 1, 0, "<test>", "abc123"),
				PosEnd:   util.NewPos(1, 7, 6, "<test>", "abc123"),
			},
		},
		{
			name:  "Chinese Characters Identifier",
			input: "你好世界",
			expect: &Token{
				Type:     IDENT,
				Literal:  "你好世界",
				PosStart: util.NewPos(1, 1, 0, "<test>", "你好世界"),
				PosEnd:   util.NewPos(1, 5, 12, "<test>", "你好世界"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLexer("<test>", tt.input)
			tok, err := l.NextToken()
			if err != nil {
				t.Errorf("err = %+v, expected nil", err)
			}
			if !reflect.DeepEqual(tok, tt.expect) {
				t.Errorf("tok = %+v, expected %+v", tok, tt.expect)
			}
		})
	}
}

func TestLexer_Operators(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect *Token
	}{
		{
			name:  "Single Character Operator",
			input: "+",
			expect: &Token{
				Type:     PLUS,
				Literal:  "+",
				PosStart: util.NewPos(1, 1, 0, "<test>", "+"),
				PosEnd:   util.NewPos(1, 2, 1, "<test>", "+"),
			},
		},
		{
			name:  "Multi-Character Operator",
			input: ">=",
			expect: &Token{
				Type:     GTE,
				Literal:  ">=",
				PosStart: util.NewPos(1, 1, 0, "<test>", ">="),
				PosEnd:   util.NewPos(1, 3, 2, "<test>", ">="),
			},
		},
		{
			name:  "Multi-Character Input but Single Operator",
			input: "=>",
			expect: &Token{
				Type:     EQUAL,
				Literal:  "=",
				PosStart: util.NewPos(1, 1, 0, "<test>", "=>"),
				PosEnd:   util.NewPos(1, 2, 1, "<test>", "=>"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLexer("<test>", tt.input)
			tok, err := l.NextToken()
			if err != nil {
				t.Errorf("err = %+v, expected nil", err)
			}
			if !reflect.DeepEqual(tok, tt.expect) {
				t.Errorf("tok = %+v, expected %+v", tok, tt.expect)
			}
		})
	}
}

func TestLexer_Strings(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect *Token
	}{
		{
			name:  "Single Character String",
			input: "\"a\"",
			expect: &Token{
				Type:     STRING,
				Literal:  "a",
				PosStart: util.NewPos(1, 1, 0, "<test>", "\"a\""),
				PosEnd:   util.NewPos(1, 4, 3, "<test>", "\"a\""),
			},
		},
		{
			name:  "Multi-Character String",
			input: "\"hello world\"",
			expect: &Token{
				Type:     STRING,
				Literal:  "hello world",
				PosStart: util.NewPos(1, 1, 0, "<test>", "\"hello world\""),
				PosEnd:   util.NewPos(1, 14, 13, "<test>", "\"hello world\""),
			},
		},
		{
			name:  "String with Escaped Characters",
			input: "\"hello \\\"world\\\"\"",
			expect: &Token{
				Type:     STRING,
				Literal:  "hello \"world\"",
				PosStart: util.NewPos(1, 1, 0, "<test>", "\"hello \\\"world\\\"\""),
				PosEnd:   util.NewPos(1, 18, 17, "<test>", "\"hello \\\"world\\\"\""),
			},
		},
		{
			name:  "String with Single Chinese Character",
			input: "\"你\"",
			expect: &Token{
				Type:     STRING,
				Literal:  "你",
				PosStart: util.NewPos(1, 1, 0, "<test>", "\"你\""),
				PosEnd:   util.NewPos(1, 4, 5, "<test>", "\"你\""),
			},
		},
		{
			name:  "String with Multi-Chinese-Character",
			input: "\"你好世界\"",
			expect: &Token{
				Type:     STRING,
				Literal:  "你好世界",
				PosStart: util.NewPos(1, 1, 0, "<test>", "\"你好世界\""),
				PosEnd:   util.NewPos(1, 7, 14, "<test>", "\"你好世界\""),
			},
		},
		{
			name:  "String with Multi-Chinese-Character with Escaped Characters",
			input: "\"你好 \\\"世界\\\"\"",
			expect: &Token{
				Type:     STRING,
				Literal:  "你好 \"世界\"",
				PosStart: util.NewPos(1, 1, 0, "<test>", "\"你好 \\\"世界\\\"\""),
				PosEnd:   util.NewPos(1, 12, 19, "<test>", "\"你好 \\\"世界\\\"\""),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLexer("<test>", tt.input)
			tok, err := l.NextToken()
			if err != nil {
				t.Errorf("err = %+v, expected nil", err)
			}
			if !reflect.DeepEqual(tok, tt.expect) {
				t.Errorf("tok = %+v, expected %+v", tok, tt.expect)
			}
		})
	}
}

func TestLexer_Errors(t *testing.T) {
	tests := []struct {
		name  string
		input string
		err   error
	}{
		{
			name:  "Unclosed Multiline Comment",
			input: "/* This is an unclosed multiline comment",
			err: &SyntaxError{
				Message:  `"*/" is expected.`,
				PosStart: util.NewPos(1, 41, 40, "<test>", "/* This is an unclosed multiline comment"),
				PosEnd:   util.NewPos(1, 43, 42, "<test>", "/* This is an unclosed multiline comment"),
			},
		},
		{
			name:  "Multiple Dots in Number",
			input: "12.34.56",
			err: &IllegalTokenError{
				Message:  "illegal float literal.",
				PosStart: util.NewPos(1, 6, 5, "<test>", "12.34.56"),
				PosEnd:   util.NewPos(1, 7, 6, "<test>", "12.34.56"),
			},
		},
		{
			name:  "Unknown Escape Character",
			input: "\"hello \\zworld\"",
			err: &IllegalTokenError{
				Message:  "illegal escape character.",
				PosStart: util.NewPos(1, 8, 7, "<test>", "\"hello \\zworld\""),
				PosEnd:   util.NewPos(1, 10, 9, "<test>", "\"hello \\zworld\""),
			},
		},
		{
			name:  "Unclosed String Literal",
			input: "\"hello world",
			err: &IllegalTokenError{
				Message:  "unterminated string literal.",
				PosStart: util.NewPos(1, 1, 0, "<test>", "\"hello world"),
				PosEnd:   util.NewPos(1, 14, 13, "<test>", "\"hello world"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLexer("<test>", tt.input)
			_, err := l.NextToken()
			if err == nil {
				t.Errorf("err = %+v, expected %+v", err, tt.err)
			}
			if err.Error() != tt.err.Error() {
				t.Errorf("err = %+v, expected %+v", err, tt.err)
			}
		})
	}
}
