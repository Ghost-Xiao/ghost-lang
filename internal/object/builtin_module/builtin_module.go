package builtin_module

import (
	"github.com/Ghost-Xiao/ghost-lang/internal/object"
)

// Modules 内置模块映射
// 包含所有内置模块的名称和对应的模块实例
// 模块名称作为键，模块实例作为值
var Modules = map[string]*object.Module{
	"math": MathModule,
	"io":   IOModule,
	"fmt":  FmtModule,
	"time": TimeModule,
}
