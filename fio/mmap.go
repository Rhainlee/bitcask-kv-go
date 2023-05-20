package fio

import "os"
import "golang.org/x/exp/mmap"

type MMap struct {
	readerAt *mmap.ReaderAt
}

// NewMMapIOManager 初始化 MMap IO
func NewMMapIOManager(fileName string) (*MMap, error) {
	file, err := os.OpenFile(fileName, os.O_CREATE, DataFilePerm)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	readerAt, err := mmap.Open(fileName)
	if err != nil {
		return nil, err
	}
	return &MMap{readerAt: readerAt}, nil
}

func (mmap *MMap) Read(b []byte, offset int64) (int, error) {
	return mmap.readerAt.ReadAt(b, offset)
}

func (mmap *MMap) Write([]byte) (int, error) {
	panic("not implemented")
}

func (mmap *MMap) Sync() error {
	panic("not implemented")
}

func (mmap *MMap) Close() error {
	return mmap.readerAt.Close()
}

func (mmap *MMap) Size() (int64, error) {
	return int64(mmap.readerAt.Len()), nil
}
