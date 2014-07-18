package goinfo

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Mem is mem info struct.
type Mem struct {
	Total   ByteSize `json:"total"`
	Used    ByteSize `json:"used"`
	Free    ByteSize `json:"free"`
	Buffers ByteSize `json:"buffers"`
	Cached  ByteSize `json:"cached"`
}

// Swap is swap info struct.
type Swap struct {
	Total ByteSize `json:"total"`
	Used  ByteSize `json:"used"`
	Free  ByteSize `json:"free"`
}

// Free is mem info sum.
type Free struct {
	Mem  *Mem  `json:"mem"`
	Swap *Swap `json:"swap"`
}

// String format the Mem struct.
func (m *Mem) String() (mem string) {
	if m == nil {
		return
	}

	return fmt.Sprintf("Mem:\t%s\t%s\t%s\t%s\t%s", m.Total.String(), m.Used.String(), m.Free.String(), m.Buffers.String(), m.Cached.String())
}

// String format the Swap struct.
func (s *Swap) String() (swap string) {
	if s == nil {
		return
	}

	return fmt.Sprintf("Swap:\t%s\t%s\t%s", s.Total.String(), s.Used.String(), s.Free.String())
}

// String format the free struct.
func (f *Free) String() (free string) {
	if f == nil {
		return
	}

	return fmt.Sprintf("%s\n%s", f.Mem.String(), f.Swap.String())
}

// Memory getting host memory by reading linux /proc/meminfo file.
func (x *Xminfo) Memory() (free *Free, err error) {
	free = new(Free)
	b, err := ioutil.ReadFile(gMemFile)
	if err != nil {
		return
	}
	s := strings.SplitAfter(string(b), "\n")
	var m = make([]ByteSize, 0)

	for _, v := range s {
		if v == "" {
			continue
		}
		mm := strings.Split(v, ":")
		if len(mm) < 2 {
			err = fmt.Errorf("mem info fields has no enough fields")
			return
		}
		info := strings.Replace(mm[1], "kB", "", -1)
		info = strings.TrimSpace(info)
		m = append(m, ByteSize(string2Float64(info)*1024))
	}
	if len(m) < 14 {
		err = fmt.Errorf("mem info fields has no enough fields")
		return
	}
	used := m[0] - m[1]
	swapUsed := m[13] - m[14]

	free.Mem = &Mem{Total: m[0], Used: used, Free: m[1], Buffers: m[2], Cached: m[3]}
	free.Swap = &Swap{Total: m[13], Used: swapUsed, Free: m[14]}

	return
}
