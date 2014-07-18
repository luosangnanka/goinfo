package goinfo

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Udp is udp info struct.
type Udp struct {
	InDatagrams  int64
	NoPorts      int64
	InErrors     int64
	OutDatagrams int64
}

// String format the udp struct.
func (u *Udp) String() (udp string) {
	if u == nil {
		return
	}

	return fmt.Sprintf("InDatagrams:%d, NoPorts:%d, InErrors:%d, OutDatagrams:%d", u.InDatagrams, u.NoPorts, u.InErrors, u.OutDatagrams)
}

// UDP getting udp info by reading linux /proc/net/snmp file.
func (x *Xminfo) UDP() (udp *Udp, err error) {
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
		if strings.HasPrefix(v, "Udp") {
			if 1 == i {
				u := strings.Replace(v, "Udp:", "", -1)
				u = strings.TrimSpace(u)
				uS := strings.Fields(u)
				if len(uS) < 4 {
					err = fmt.Errorf("udp info fields has no enough fields")
					return
				}
				udp = &Udp{InDatagrams: string2Int64(uS[0]), NoPorts: string2Int64(uS[1]), InErrors: string2Int64(uS[2]), OutDatagrams: string2Int64(uS[3])}
			}
			i++
		}
	}

	return
}
