package benchmark

import (
	"fmt"
	bitcask "github.com/rhainlee/bitcask-kv-go"
	"github.com/rhainlee/bitcask-kv-go/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

var db *bitcask.DB

func init() {
	//	初始化用于基准测试的 DB 实例
	var err error
	options := bitcask.DefaultOptions
	options.DirPath = "/tmp/bitcask-go-bench"
	db, err = bitcask.Open(options)
	if err != nil {
		panic(fmt.Sprintf("failed to open db: %v", err))
	}
}

func Benchmark_Put(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		err := db.Put(utils.GetTestKey(i), utils.RandomValue(1024))
		assert.Nil(b, err)
	}
}

func Benchmark_Get(b *testing.B) {
	for i := 0; i < 100000; i++ {
		err := db.Put(utils.GetTestKey(i), utils.RandomValue(1024))
		assert.Nil(b, err)
	}

	rand.Seed(time.Now().Unix())
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := db.Get(utils.GetTestKey(rand.Int()))
		if err != nil && err != bitcask.ErrKeyNotFound {
			b.Fatal(err)
		}
	}
}

func Benchmark_Delete(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		err := db.Delete(utils.GetTestKey(rand.Int()))
		assert.Nil(b, err)
	}
}
