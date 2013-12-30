package goinfo

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Cpu struct {
	Total  int64
	User   int64
	Nice   int64
	System int64
	Idle   int64
	Iowait int64
	Irq    int64
	Sftirq int64
}

func (c *Cpu) String() (cpu string) {
	if c == nil {
		return
	}
	return fmt.Sprintf("total\tuser\tnice\tsystem\tidle\tiowait\tirq\tsoftirq\n%d\t%d\t%d\t%d\t%d\t%d\t%d\t%d\t", c.Total, c.User, c.Nice, c.System, c.Idle, c.Iowait, c.Irq, c.Sftirq)
}

// 通过读取 /proc/stat 获取cpu状态
func (xm *Xminfo) Cpu() (cpu []*Cpu, err error) {
	cpu = make([]*Cpu, 0)
	b, err := ioutil.ReadFile(gCpu)
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
			user := string2int64(cc[1])
			nice := string2int64(cc[2])
			system := string2int64(cc[3])
			idle := string2int64(cc[4])
			iowait := string2int64(cc[5])
			irq := string2int64(cc[6])
			softirq := string2int64(cc[7])
			total := sum(user, nice, system, idle, iowait, irq, softirq)

			cpu = append(cpu, &Cpu{total, user, nice, system, idle, iowait, irq, softirq})
		}
	}

	return
}
