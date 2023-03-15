package bitcask_kv_go

import "errors"

var (
	// ErrKeyIsEmpty 定义一个名为 ErrkeyIsEmpty 的全部变量，其类型为error
	ErrKeyIsEmpty = errors.New("the key is empty")

	ErrIndexUpdateFailed = errors.New("failed to update index")

	ErrKeyNotFound = errors.New("key not found in database")

	ErrDataFileNotFound = errors.New("data file is not found")
)
