package object

// BuiltinModule 内置模块接口
// 定义了所有内置模块必须实现的方法
type BuiltinModule interface {
	// Name 获取模块名称
	//
	// 返回值:
	//
	// string - 模块名称
	Name() string

	// Load 加载模块
	//
	// 返回值:
	//
	// *Environment - 模块环境
	Load() *Environment

	// 嵌入 object.Object 接口
	Object
}
