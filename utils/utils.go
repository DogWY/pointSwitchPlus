package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

// RemoveDuplicates 删除重复元素
func RemoveDuplicates(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	return nums
}

// ByteToAddr 字节转地址
func ByteToAddr(b []byte) int {
	return int(b[0])<<8 + int(b[1])
}

// AddrToByte 地址转字节
func AddrToByte(a int) []byte {
	return []byte{byte(a >> 8), byte(a & 255)}
}

// 定位项目根目录
func findProjectRoot(currentDir string) (string, error) {
	for {
		if _, err := os.Stat(filepath.Join(currentDir, "go.mod")); err == nil {
			return currentDir, nil
		}

		newDir := filepath.Dir(currentDir)
		if newDir == currentDir {
			break
		}

		currentDir = newDir
	}

	return "", fmt.Errorf("项目根目录未找到")
}
