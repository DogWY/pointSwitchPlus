package utils

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
