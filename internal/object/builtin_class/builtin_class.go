package builtinclass

import (
	"github.com/Ghost-Xiao/ghost-lang/internal/object"
)

var BuiltinClasses = map[string]*object.Class{
	"Int":       IntClass,
	"Float":     FloatClass,
	"Bool":      BoolClass,
	"String":    StringClass,
	"List":      ListClass,
	"Map":       MapClass,
	"Namespace": NamespaceClass,
	"Module":    ModuleClass,
	"Class":     ClassClass,
	"Instance":  InstanceClass,
	"Function":  FunctionClass,
	"Method":    MethodClass,
}
