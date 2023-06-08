package index

import (
	"github.com/rhainlee/bitcask-kv-go/data"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestBPlusTree_Put(t *testing.T) {
	path := filepath.Join("D:\\work\\tmp") // path := filepath.Join("/tmp")
	_ = os.MkdirAll(path, os.ModePerm)
	defer func() {
		if err := os.RemoveAll(path); err != nil {
			log.Println(err)
		}
	}()
	tree := NewBPlusTree(path, false)
	defer tree.Close()

	res1 := tree.Put([]byte("aac"), &data.LogRecordPos{Fid: 123, Offset: 999})
	assert.Nil(t, res1)
	tree.Put([]byte("abc"), &data.LogRecordPos{Fid: 123, Offset: 999})
	tree.Put([]byte("acc"), &data.LogRecordPos{Fid: 123, Offset: 999})

	res2 := tree.Put([]byte("acc"), &data.LogRecordPos{Fid: 123, Offset: 111})
	assert.Equal(t, uint32(123), res2.Fid)
	assert.Equal(t, int64(999), res2.Offset)

}

func TestBPlusTree_Get(t *testing.T) {
	path := filepath.Join("D:\\work\\tmp")
	_ = os.MkdirAll(path, os.ModePerm)
	defer func() {
		os.RemoveAll(path)
	}()
	tree := NewBPlusTree(path, false)
	defer func() {
		tree.Close()
	}()

	pos := tree.Get([]byte("not exist"))
	assert.Nil(t, pos)

	tree.Put([]byte("aac"), &data.LogRecordPos{Fid: 123, Offset: 999})
	pos1 := tree.Get([]byte("aac"))
	assert.Equal(t, pos1.Fid, uint32(123))
	assert.Equal(t, pos1.Offset, int64(999))

	tree.Put([]byte("aac"), &data.LogRecordPos{Fid: 9884, Offset: 1232})
	pos2 := tree.Get([]byte("aac"))
	assert.Equal(t, pos2.Fid, uint32(9884))
	assert.Equal(t, pos2.Offset, int64(1232))
}

func TestBPlusTree_Delete(t *testing.T) {
	path := filepath.Join(os.TempDir(), "bptree-delete")
	_ = os.MkdirAll(path, os.ModePerm)
	defer func() {
		os.RemoveAll(path)
	}()
	tree := NewBPlusTree(path, false)
	defer tree.Close()

	res1, ok1 := tree.Delete([]byte("not exist"))
	assert.Nil(t, res1)
	assert.False(t, ok1)

	tree.Put([]byte("aac"), &data.LogRecordPos{Fid: 123, Offset: 999})
	res2, ok2 := tree.Delete([]byte("aac"))
	assert.Equal(t, uint32(123), res2.Fid)
	assert.Equal(t, int64(999), res2.Offset)
	assert.True(t, ok2)

	pos1 := tree.Get([]byte("aac"))
	assert.Nil(t, pos1)
}

func TestBPlusTree_Size(t *testing.T) {
	path := filepath.Join(os.TempDir(), "bptree-size")
	_ = os.MkdirAll(path, os.ModePerm)
	defer func() {
		os.RemoveAll(path)
	}()
	tree := NewBPlusTree(path, false)
	defer func() {
		tree.Close()
	}()
	assert.Equal(t, 0, tree.Size())

	tree.Put([]byte("aac"), &data.LogRecordPos{Fid: 123, Offset: 999})
	tree.Put([]byte("abc"), &data.LogRecordPos{Fid: 123, Offset: 999})
	tree.Put([]byte("acc"), &data.LogRecordPos{Fid: 123, Offset: 999})

	assert.Equal(t, 3, tree.Size())
}

func TestBPlusTree_Iterator(t *testing.T) {
	path := filepath.Join(os.TempDir(), "bptree-iter")
	_ = os.MkdirAll(path, os.ModePerm)
	defer func() {
		os.RemoveAll(path)
	}()
	tree := NewBPlusTree(path, false)

	tree.Put([]byte("caac"), &data.LogRecordPos{Fid: 123, Offset: 999})
	tree.Put([]byte("bbca"), &data.LogRecordPos{Fid: 123, Offset: 999})
	tree.Put([]byte("acce"), &data.LogRecordPos{Fid: 123, Offset: 999})
	tree.Put([]byte("ccec"), &data.LogRecordPos{Fid: 123, Offset: 999})
	tree.Put([]byte("bbba"), &data.LogRecordPos{Fid: 123, Offset: 999})

	iter := tree.Iterator(true)
	for iter.Rewind(); iter.Valid(); iter.Next() {
		//t.Log(string(iter.Key()))
		assert.NotNil(t, iter.Key())
		assert.NotNil(t, iter.Value())
	}
}
