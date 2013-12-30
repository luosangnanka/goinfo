package goinfo

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Tcp struct {
	ActiveOpens  string
	PassiveOpens string
	InSegs       string
	OutSegs      string
	RetransSegs  string
}

func (t *Tcp) String() (tcp string) {
	if t == nil {
		return
	}
	return fmt.Sprintf("ActiveOpens:%s, PassiveOpens:%s, InSegs:%s, OutSegs:%s, RetransSegs:%s", t.ActiveOpens, t.PassiveOpens, t.InSegs, t.OutSegs, t.RetransSegs)
}

// 通过读取 /proc/net/snmp 获取tcp信息
func (xm *Xminfo) Tcp() (tcp *Tcp, err error) {
	var i int = 0
	b, err := ioutil.ReadFile(gSnmp)
	if err != nil {
		return
	}
	s := strings.SplitAfter(string(b), "\n")
	for _, v := range s {
		if v == "" {
			continue
		}
		if strings.HasPrefix(v, "Tcp") {
			if 1 == i {
				t := strings.Replace(v, "Tcp:", "", -1)
				t = strings.TrimSpace(t)
				tS := strings.Fields(t)
				tcp = &Tcp{ActiveOpens: tS[4], PassiveOpens: tS[5], InSegs: tS[9], OutSegs: tS[10], RetransSegs: tS[11]}
			}
			i++
		}
	}

	return
}
