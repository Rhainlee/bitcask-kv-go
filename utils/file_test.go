package utils

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestDirSize(t *testing.T) {
	dir, _ := os.Getwd()
	t.Log(dir)
	size, err := DirSize(dir)
	assert.Nil(t, err)
	assert.True(t, size > 0)

}

func TestAvailableDiskSize(t *testing.T) {
	size, err := AvailableDiskSize()
	assert.Nil(t, err)
	assert.True(t, size > 0)
	//t.Log(size / 1024 / 1024 / 1024)
}
