/*===========================================
*   Copyright (C) 2013 All rights reserved.
*
*   company      : xiaomi
*   author       : zhangye
*   email        : zhangye@xiaomi.com
*   date         : 2013-12-12 16:21:25
*   description  : package xminfo - udp info
*
=============================================*/
package xminfo

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Udp struct {
	InDatagrams  string
	NoPorts      string
	InErrors     string
	OutDatagrams string
}

func (u *Udp) String() (udp string) {
	if u == nil {
		return
	}
	return fmt.Sprintf("InDatagrams:%s, NoPorts:%s, InErrors:%s, OutDatagrams:%s", u.InDatagrams, u.NoPorts, u.InErrors, u.OutDatagrams)
}

// 通过读取 /proc/net/snmp 获取udp信息
func (xm *Xminfo) Udp() (udp *Udp, err error) {
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
		if strings.HasPrefix(v, "Udp") {
			if 1 == i {
				u := strings.Replace(v, "Udp:", "", -1)
				u = strings.TrimSpace(u)
				uS := strings.Fields(u)
				udp = &Udp{InDatagrams: uS[0], NoPorts: uS[1], InErrors: uS[2], OutDatagrams: uS[3]}
			}
			i++
		}
	}

	return
}
