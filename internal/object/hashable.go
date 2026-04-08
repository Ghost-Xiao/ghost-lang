package object

// Hashable 表示可哈希接口
type Hashable interface {
	// Hash 返回可哈希对象的哈希值
	//
	// 返回值:
	//
	//	uint64 - 可哈希对象的哈希值
	Hash() uint64

	// 嵌入 Object 接口
	Object
}
