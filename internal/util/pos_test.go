package util

import (
	"reflect"
	"testing"
)

func TestPos_Advance(t *testing.T) {
	tests := []struct {
		name        string
		pos         *Pos
		expectedPos *Pos
	}{
		{
			name:        "Advance",
			pos:         NewPos(1, 1, 0, "<text>", "Hello, World!"),
			expectedPos: NewPos(1, 2, 1, "<text>", "Hello, World!"),
		},
		{
			name:        "New line",
			pos:         NewPos(1, 5, 4, "<text>", "Hello\nWorld!"),
			expectedPos: NewPos(1, 6, 5, "<text>", "Hello\nWorld!"),
		},
		{
			name:        "After new line",
			pos:         NewPos(1, 6, 5, "<text>", "Hello\nWorld!"),
			expectedPos: NewPos(2, 1, 6, "<text>", "Hello\nWorld!"),
		},
		{
			name:        "End of text",
			pos:         NewPos(2, 5, 11, "<text>", "Hello\nWorld!"),
			expectedPos: NewPos(2, 6, 12, "<text>", "Hello\nWorld!"),
		},
		{
			name:        "Empty text",
			pos:         NewPos(1, 1, 0, "<text>", ""),
			expectedPos: NewPos(1, 2, 1, "<text>", ""),
		},
		{
			name:        "Single char",
			pos:         NewPos(1, 1, 0, "<text>", "a"),
			expectedPos: NewPos(1, 2, 1, "<text>", "a"),
		},
		{
			name:        "Single new line",
			pos:         NewPos(1, 1, 0, "<text>", "\n"),
			expectedPos: NewPos(2, 1, 1, "<text>", "\n"),
		},
		{
			name:        "Chinese words",
			pos:         NewPos(1, 1, 0, "<text>", "你好，世界！"),
			expectedPos: NewPos(1, 2, 3, "<text>", "你好，世界！"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.pos.Advance()
			if !reflect.DeepEqual(tt.pos, tt.expectedPos) {
				t.Errorf("tt.pos = %+v, expected %+v", tt.pos, tt.expectedPos)
			}
		})
	}
}

func TestPos_Backup(t *testing.T) {
	tests := []struct {
		name        string
		pos         *Pos
		expectedPos *Pos
	}{
		{
			name:        "Backup",
			pos:         NewPos(1, 2, 1, "<text>", "Hello, World!"),
			expectedPos: NewPos(1, 1, 0, "<text>", "Hello, World!"),
		},
		{
			name:        "New Line",
			pos:         NewPos(1, 6, 5, "<text>", "Hello\nWorld!"),
			expectedPos: NewPos(1, 5, 4, "<text>", "Hello\nWorld!"),
		},
		{
			name:        "After New Line",
			pos:         NewPos(2, 1, 6, "<text>", "Hello\nWorld!"),
			expectedPos: NewPos(1, 6, 5, "<text>", "Hello\nWorld!"),
		},
		{
			name:        "End of Text",
			pos:         NewPos(2, 6, 12, "<text>", "Hello\nWorld!"),
			expectedPos: NewPos(2, 5, 11, "<text>", "Hello\nWorld!"),
		},
		{
			name:        "Empty Text",
			pos:         NewPos(1, 1, 0, "<text>", ""),
			expectedPos: NewPos(1, 0, -1, "<text>", ""),
		},
		{
			name:        "Single Char",
			pos:         NewPos(1, 2, 1, "<text>", "a"),
			expectedPos: NewPos(1, 1, 0, "<text>", "a"),
		},
		{
			name:        "Single New Line",
			pos:         NewPos(2, 1, 1, "<text>", "\n"),
			expectedPos: NewPos(1, 1, 0, "<text>", "\n"),
		},
		{
			name:        "Chinese Words",
			pos:         NewPos(1, 2, 3, "<text>", "你好，世界！"),
			expectedPos: NewPos(1, 1, 0, "<text>", "你好，世界！"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.pos.Backup()
			if !reflect.DeepEqual(tt.pos, tt.expectedPos) {
				t.Errorf("tt.pos = %+v, expected %+v", tt.pos, tt.expectedPos)
			}
		})
	}
}
