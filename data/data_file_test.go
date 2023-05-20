package data

import (
	"fmt"
	"github.com/rhainlee/bitcask-kv-go/fio"
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

func TestOpenDataFile(t *testing.T) {
	dataFile1, err := OpenDataFile(os.TempDir(), 0, fio.StandardFIO)
	assert.Nil(t, err)
	assert.NotNil(t, dataFile1)

	dataFile2, err := OpenDataFile(os.TempDir(), 111, fio.StandardFIO)
	assert.Nil(t, err)
	assert.NotNil(t, dataFile2)

	dataFile3, err := OpenDataFile(os.TempDir(), 111, fio.StandardFIO)
	assert.Nil(t, err)
	assert.NotNil(t, dataFile3)

	// 清理测试文件
	path1 := filepath.Join(os.TempDir(), fmt.Sprintf("%09d", 0)+DataFileNameSuffix)
	fio1 := dataFile1.IoManager
	defer destoryFile(path1)

	path2 := filepath.Join(os.TempDir(), fmt.Sprintf("%09d", 111)+DataFileNameSuffix)
	fio2 := dataFile2.IoManager
	defer destoryFile(path2)

	path3 := filepath.Join(os.TempDir(), fmt.Sprintf("%09d", 111)+DataFileNameSuffix)
	fio3 := dataFile3.IoManager
	defer destoryFile(path3)

	defer fio1.Close()
	defer fio2.Close()
	defer fio3.Close()

}

func TestDataFile_Write(t *testing.T) {
	dataFile, err := OpenDataFile(os.TempDir(), 0, fio.StandardFIO)
	assert.Nil(t, err)
	assert.NotNil(t, dataFile)

	err = dataFile.Write([]byte("aaa"))
	assert.Nil(t, err)

	err = dataFile.Write([]byte("bbb"))
	assert.Nil(t, err)

	err = dataFile.Write([]byte("ccc"))
	assert.Nil(t, err)

	// 清理测试文件
	path1 := filepath.Join(os.TempDir(), fmt.Sprintf("%09d", 0)+DataFileNameSuffix)
	fio1 := dataFile.IoManager
	defer destoryFile(path1)
	defer fio1.Close()
}

func TestDataFile_Close(t *testing.T) {
	dataFile, err := OpenDataFile(os.TempDir(), 123, fio.StandardFIO)
	assert.Nil(t, err)
	assert.NotNil(t, dataFile)

	err = dataFile.Write([]byte("aaa"))
	assert.Nil(t, err)

	err = dataFile.Close()
	assert.Nil(t, err)

	// 清理测试文件
	path1 := filepath.Join(os.TempDir(), fmt.Sprintf("%09d", 123)+DataFileNameSuffix)
	fio1 := dataFile.IoManager
	defer destoryFile(path1)
	defer fio1.Close()
}

func TestDataFile_Sync(t *testing.T) {
	dataFile, err := OpenDataFile(os.TempDir(), 456, fio.StandardFIO)
	assert.Nil(t, err)
	assert.NotNil(t, dataFile)

	err = dataFile.Write([]byte("aaa"))
	assert.Nil(t, err)

	err = dataFile.Sync()
	assert.Nil(t, err)

	// 清理测试文件
	path1 := filepath.Join(os.TempDir(), fmt.Sprintf("%09d", 456)+DataFileNameSuffix)
	fio1 := dataFile.IoManager
	defer destoryFile(path1)
	defer fio1.Close()
}

func TestDataFile_ReadLogRecord(t *testing.T) {
	dataFile, err := OpenDataFile(os.TempDir(), 6666, fio.StandardFIO)
	assert.Nil(t, err)
	assert.NotNil(t, dataFile)

	// 只有一条 LogRecord
	rec1 := &LogRecord{
		Key:   []byte("name"),
		Value: []byte("bitcask kv go"),
	}
	res1, size1 := EncodeLogRecord(rec1)
	err = dataFile.Write(res1)
	assert.Nil(t, err)

	readRec1, readSize1, err := dataFile.ReadLogRecord(0)
	assert.Nil(t, err)
	assert.Equal(t, rec1, readRec1)
	assert.Equal(t, size1, readSize1)
	t.Log(readSize1)

	// 多条 LogRecord，从不同的位置读取
	rec2 := &LogRecord{
		Key:   []byte("name"),
		Value: []byte("a new value"),
	}
	res2, size2 := EncodeLogRecord(rec2)
	err = dataFile.Write(res2)
	assert.Nil(t, err)

	readRec2, readSize2, err := dataFile.ReadLogRecord(size1)
	assert.Nil(t, err)
	assert.Equal(t, rec2, readRec2)
	assert.Equal(t, size2, readSize2)

	// 被删除的数据在数据文件的末尾
	rec3 := &LogRecord{
		Key:   []byte("1"),
		Value: []byte(""),
		Type:  LogRecordDeleted,
	}
	res3, size3 := EncodeLogRecord(rec3)
	err = dataFile.Write(res3)
	assert.Nil(t, err)

	readRec3, readSize3, err := dataFile.ReadLogRecord(size1 + size2)
	assert.Nil(t, err)
	assert.Equal(t, rec3, readRec3)
	assert.Equal(t, size3, readSize3)

	// 清理测试文件
	path1 := filepath.Join(os.TempDir(), fmt.Sprintf("%09d", 6666)+DataFileNameSuffix)
	fio1 := dataFile.IoManager
	defer destoryFile(path1)
	defer fio1.Close()

}
