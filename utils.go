package goinfo

import (
	"fmt"
	"strconv"
)

// ByteSize type for the format of KB MB GB TB.
type ByteSize float64

const (
	_ = iota
	// KB size.
	KB ByteSize = 1 << (10 * iota)
	// MB size.
	MB
	// GB size.
	GB
	// TB size.
	TB
)

// String format the size into the KB MB GB TB format.
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

func string2Int64(src string) (rs int64) {
	rs, _ = strconv.ParseInt(src, 10, 64)
	return
}

func string2Float64(src string) (rs float64) {
	rs, _ = strconv.ParseFloat(src, 64)
	return
}

func sum(src ...int64) (rs int64) {
	rs = 0
	for _, v := range src {
		rs += v
	}

	return
}
