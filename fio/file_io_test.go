package fio

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func destoryFile(name string) {
	if err := os.RemoveAll(name); err != nil {
		panic(err)
	}
}

func TestNewFileIOManager(t *testing.T) {
	path := filepath.Join("d:\\temp", "a.data") // 用于构建文件路径的函数调用，windows下绝对路径写法，要提前创建好temp目录
	fio, err := NewFileIOManager(path)
	defer destoryFile(path)
	defer fio.Close() // 只能这样才能通过测试，但是这样就不是独立测试函数了，暂时存疑

	assert.Nil(t, err)
	assert.NotNil(t, fio)
}

func TestFileIO_Write(t *testing.T) {
	path := filepath.Join("d:\\temp", "a.data")
	fio, err := NewFileIOManager(path)
	defer destoryFile(path)
	defer fio.Close()

	assert.Nil(t, err)
	assert.NotNil(t, fio)

	n, err := fio.Write([]byte(""))
	assert.Equal(t, 0, n)
	assert.Nil(t, err)

	n, err = fio.Write([]byte("bitcask kv"))
	assert.Equal(t, 10, n)
	assert.Nil(t, err)
	//t.Log(n, err)

	n, err = fio.Write([]byte("storage"))
	assert.Equal(t, 7, n)
	assert.Nil(t, err)
	//t.Log(n, err)
}

func TestFileIO_Read(t *testing.T) {
	path := filepath.Join("d:\\temp", "a.data")
	fio, err := NewFileIOManager(path)
	defer destoryFile(path)
	defer fio.Close()

	assert.Nil(t, err)
	assert.NotNil(t, fio)

	_, err = fio.Write([]byte("key-a"))
	assert.Nil(t, err)

	_, err = fio.Write([]byte("key-b"))
	assert.Nil(t, err)

	b1 := make([]byte, 5)
	n, err := fio.Read(b1, 0)
	assert.Equal(t, 5, n)
	assert.Equal(t, []byte("key-a"), b1)

	b2 := make([]byte, 5)
	n, err = fio.Read(b2, 5)
	assert.Equal(t, 5, n)
	assert.Equal(t, []byte("key-b"), b2)

}

func TestFileIO_Sync(t *testing.T) {
	path := filepath.Join("d:\\temp", "a.data")
	fio, err := NewFileIOManager(path)
	defer destoryFile(path)
	defer fio.Close()

	assert.Nil(t, err)
	assert.NotNil(t, fio)

	err = fio.Sync()
	assert.Nil(t, err)
}

func TestFileIO_Close(t *testing.T) {
	path := filepath.Join("d:\\temp", "a.data")
	fio, err := NewFileIOManager(path)
	defer destoryFile(path)
	defer fio.Close()

	assert.Nil(t, err)
	assert.NotNil(t, fio)

	err = fio.Close()
	assert.Nil(t, err)

}
