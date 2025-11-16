package object

// Symbol 表示一个标识符的完整信息

type Symbol struct {
	Name    string // 符号名称
	Value   Object // 值
	IsConst bool   // 是否是常量
}
