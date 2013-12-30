package goinfo

import (
	"fmt"
	"strconv"
)

// 格式化单位为 KB MB GB TB等格式
type ByteSize float64

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
)

func (b ByteSize) String() (rs string) {
	switch {
	case b >= TB:
		return fmt.Sprintf("%.2fTB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%2.fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%2.fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%2.fKB", b/KB)
	}
	return fmt.Sprintf("%.2fB", b)
}

// string转换成int64类型
func string2int64(src string) (rs int64) {
	rs, _ = strconv.ParseInt(src, 10, 64)
	return
}

// 计算不定项 int64 的和
func sum(src ...int64) (rs int64) {
	rs = 0
	for _, v := range src {
		rs += v
	}

	return
}
