package goinfo

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Loadavg struct {
	La1, La5, La15 string
	Processes      string
	MaxPid         string
}

func (l *Loadavg) String() (load string) {
	if l == nil {
		return
	}
	return fmt.Sprintf("La1:%s, La5:%s, La15:%s, Processes:%s, MaxPid:%s", l.La1, l.La5, l.La15, l.Processes, l.MaxPid)
}

// 通过读取 /proc/loadavg 获得负载信息
func (xm *Xminfo) Load() (load *Loadavg, err error) {
	b, err := ioutil.ReadFile(gLoadavg)
	if err != nil {
		return
	}
	s := strings.Fields(string(b))

	load = &Loadavg{s[0], s[1], s[2], s[3], s[4]}

	return
}
