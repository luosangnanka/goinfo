package goinfo

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Mem struct {
	Total   int64
	Used    int64
	Free    int64
	Buffers int64
	Cached  int64
}

type Swap struct {
	Total int64
	Used  int64
	Free  int64
}

type Free struct {
	Mem  *Mem
	Swap *Swap
}

// 通过读取 /proc/meminfo 获得内存信息
func (f *Free) String() (memory string) {
	if f == nil {
		return
	}
	mem, swap := f.Mem, f.Swap
	memory = fmt.Sprintf("Mem:\t%s\t%s\t%s\t%s\t%s\nSwap:\t%s\t%s\t%s",
		ByteSize(mem.Total).String(), ByteSize(mem.Used).String(), ByteSize(mem.Free).String(), ByteSize(mem.Buffers).String(), ByteSize(mem.Cached).String(),
		ByteSize(swap.Total).String(), ByteSize(swap.Used).String(), ByteSize(swap.Free).String())

	return
}

func (xm *Xminfo) Memory() (free *Free, err error) {
	free = new(Free)
	b, err := ioutil.ReadFile(gMem)
	if err != nil {
		return
	}
	s := strings.SplitAfter(string(b), "\n")
	m := make([]int64, 0)
	for _, v := range s {
		if v == "" {
			continue
		}
		mm := strings.Split(v, ":")
		info := strings.Replace(mm[1], "kB", "", -1)
		info = strings.TrimSpace(info)
		m = append(m, string2int64(info)*1024)
	}
	used := m[0] - m[1]
	swapUsed := m[13] - m[14]
	free.Mem = &Mem{Total: m[0], Used: used, Free: m[1], Buffers: m[2], Cached: m[3]}
	free.Swap = &Swap{Total: m[13], Used: swapUsed, Free: m[14]}

	return
}
