package module

import (
	"github.com/Ghost-Xiao/ghost-lang/internal/object"
)

// BuiltinModules 内置模块映射
// 包含所有内置模块的名称和对应的模块实例
// 模块名称作为键，模块实例作为值
var BuiltinModules = map[string]object.BuiltinModule{
	"math": &Math{},
	"io":   &IO{},
	"fmt":  &Fmt{},
}
