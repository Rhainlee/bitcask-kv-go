package fio

const DataFilePerm = 0644 // 0644 表示文件权限，它是一个八进制数，表示文件的读、写和执行权限。0644 表示文件所有者具有读写权限，而所属组和其他用户只有读权限。

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
}
