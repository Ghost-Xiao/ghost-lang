package object

// Number 数值类型接口，定义所有数值类型共有的操作
// 实现此接口的类型包括Int和Float
// 接口方法包括算术运算、位运算等数值操作
type Number interface {
	Object
}
