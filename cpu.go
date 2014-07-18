package goinfo

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// CPUInfo cpu info struct.
type CPUInfo struct {
	Total   int64 `json:"total"`
	User    int64 `json:"user"`
	Nice    int64 `json:"nice"`
	System  int64 `json:"system"`
	Idle    int64 `json:"idle"`
	Iowait  int64 `json:"iowait"`
	Irq     int64 `json:"irq"`
	Softirq int64 `json:"softirq"`
}

// String format the cpu struct.
func (c *CPUInfo) String() (cpu string) {
	if c == nil {
		return
	}

	return fmt.Sprintf("total\tuser\tnice\tsystem\tidle\tiowait\tirq\tsoftirq\n%d\t%d\t%d\t%d\t%d\t%d\t%d\t%d\t", c.Total, c.User, c.Nice, c.System, c.Idle, c.Iowait, c.Irq, c.Softirq)
}

// CPU getting cpu info by reading linux /proc/stat file.
func (x *Xminfo) CPU() (info []*CPUInfo, err error) {
	info = make([]*CPUInfo, 0)
	b, err := ioutil.ReadFile(gCPUFile)
	if err != nil {
		return
	}
	s := strings.SplitAfter(string(b), "\n")
	for _, v := range s {
		cc := strings.Fields(v)
		if len(cc) == 0 {
			continue
		}
		if strings.HasPrefix(cc[0], "cpu") {
			if len(cc) < 8 {
				err = fmt.Errorf("cpu info fields has no enough fields")
				return
			}
			user := string2Int64(cc[1])
			nice := string2Int64(cc[2])
			system := string2Int64(cc[3])
			idle := string2Int64(cc[4])
			iowait := string2Int64(cc[5])
			irq := string2Int64(cc[6])
			softirq := string2Int64(cc[7])
			total := sum(user, nice, system, idle, iowait, irq, softirq)

			info = append(info, &CPUInfo{total, user, nice, system, idle, iowait, irq, softirq})
		}
	}

	return
}
