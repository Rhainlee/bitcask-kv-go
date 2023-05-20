package fio

const DataFilePerm = 0644 // 0644 表示文件权限，它是一个八进制数，表示文件的读、写和执行权限。0644 表示文件所有者具有读写权限，而所属组和其他用户只有读权限。

type FileIOType = byte

const (
	// StandardFIO 标准文件 IO
	StandardFIO FileIOType = iota

	// MemoryMap 内存文件映射
	MemoryMap
)

// IOManager 抽象IO 管理接口，可以接入不同的 IO 类型，目前支持标准文件 IO
type IOManager interface {
	// Read 从文件的给定位置读取对应的数据
	Read([]byte, int64) (int, error)

	// Write 写入字节数组到文件中
	Write([]byte) (int, error)

	// Sync 持久化数据
	Sync() error

	// Close 关闭文件
	Close() error

	// Size 获取到文件大小
	Size() (int64, error)
}

// NewIOManager 初始化 IOManager
func NewIOManager(fileName string, ioType FileIOType) (IOManager, error) {
	switch ioType {
	case StandardFIO:
		return NewFileIOManager(fileName)
	case MemoryMap:
		return NewMMapIOManager(fileName)
	default:
		panic("unsupported io type")
	}
}
