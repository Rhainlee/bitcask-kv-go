package utils

import (
	"github.com/shirou/gopsutil/disk"
	"os"
	"path/filepath"
	"syscall"
)

// DirSize 获取一个目录的大小
func DirSize(dirPath string) (int64, error) {
	var size int64
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return nil
	})
	return size, err
}

// AvailableDiskSize 获取磁盘空间大小，字节为单位
func AvailableDiskSize() (uint64, error) {
	wd, err := syscall.Getwd()
	if err != nil {
		return 0, err
	}
	usage, err := disk.Usage(wd)
	if err != nil {
		return 0, err
	}
	return usage.Free, nil
}

// AvailableDiskSize 获取磁盘空间大小，字节为单位
// 下面的代码在windows下无法使用
//func AvailableDiskSize() (uint64, error) {
//	var stat syscall.Statfs_t
//	wd, err := syscall.Getwd()
//	if err != nil {
//		return 0, err
//	}
//
//	if err := syscall.Statfs(wd, &stat); err != nil {
//		return 0, err
//	}
//	return stat.Bavail * uint64(stat.Bsize), nil
//}
