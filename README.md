# bitcask-kv-go
## 简介
基于高效`bitcask`模型的高性能键值（KV）存储引擎。提供快速可靠的数据检索和存储功能。其对速度和简单性的关注使其成为在平衡存储成本的同时优先考虑性能的应用程序的有价值的替代方案。

## 使用
1. 使用Go命令行工具安装，将该项目作为嵌入式的库在`Go`的其他应用中使用。

  ```
  go get github.com/rhainlee/bitcask-kv-go
  ```

  下面是一个简单的使用示例（Windows环境）：

  ```go
  package main
  
  import (
  	"fmt"
  	bitcask "github.com/rhainlee/bitcask-kv-go"
  )
  
  func main() {
  	opts := bitcask.DefaultOptions
  	opts.DirPath = "D:\\temp"
  	db, err := bitcask.Open(opts)
  	if err != nil {
  		panic(err)
  	}
  
  	db.Put([]byte("name"), []byte("bitcask-example"))
  	if err != nil {
  		panic(err)
  	}
  
  	val, err := db.Get([]byte("name"))
  	if err != nil {
  		panic(err)
  	}
  	fmt.Println("val = ", string(val))
  
  	err = db.Delete([]byte("name"))
  	if err != nil {
  		panic(err)
  	}
  
  }
  ```

  

2. 对外提供http接口，详情请参阅`bitcask-kv-go/http/main.go`

3. 兼容`RESP`协议，支持`String`、`Hash`、`Set`、`List`、`Sorted Set`等数据结构，详情请参阅`bitcask-kv-go/redis/`

## 基准测试

针对当前版本做了一个简单的benchmark基准测试。本次测试主要针对大规模数据的读写，测试结果如下：

### BTree索引

```
cpu: AMD Ryzen 7 4800H with Radeon Graphics
Benchmark_Put-16           132208              9350 ns/op            4624 B/op          9 allocs/op
Benchmark_Get-16         3962122               313.8 ns/op           135 B/op          4 allocs/op
Benchmark_Delete-16      3932242               312.3 ns/op           135 B/op          4 allocs/op
```

### ARTree索引

```
cpu: AMD Ryzen 7 4800H with Radeon Graphics
Benchmark_Put-16            133409             8917 ns/op            4592 B/op          8 allocs/op
Benchmark_Get-16         4025162               305.5 ns/op           103 B/op          3 allocs/op
Benchmark_Delete-16      4003288               303.3 ns/op           103 B/op          3 allocs/op
```




## TODO List

1. 集成日志、配置系统
2. 支持HTTP/RESP/RPC等多种常见通信方式，对外提供更友好的API
3. 支持事务并实现多种隔离级别比如RR等
4. 支持分布式集群部署

