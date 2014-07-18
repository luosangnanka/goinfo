package goinfo

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Tcp is tcp info struct.
type Tcp struct {
	ActiveOpens  int64
	PassiveOpens int64
	InSegs       int64
	OutSegs      int64
	RetransSegs  int64
}

// String format the tcp struct.
func (t *Tcp) String() (tcp string) {
	if t == nil {
		return
	}

	return fmt.Sprintf("ActiveOpens:%d, PassiveOpens:%d, InSegs:%d, OutSegs:%d, RetransSegs:%d", t.ActiveOpens, t.PassiveOpens, t.InSegs, t.OutSegs, t.RetransSegs)
}

// TCP getting tcp info by reading linux /proc/net/snmp file.
func (x *Xminfo) TCP() (tcp *Tcp, err error) {
	var i = 0
	b, err := ioutil.ReadFile(gSnmpFile)
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
				if len(tS) < 11 {
					err = fmt.Errorf("tcp info fields has no enough fields")
					return
				}

				tcp = &Tcp{ActiveOpens: string2Int64(tS[4]), PassiveOpens: string2Int64(tS[5]), InSegs: string2Int64(tS[9]), OutSegs: string2Int64(tS[10]), RetransSegs: string2Int64(tS[11])}
			}
		}
	}

	return
}
