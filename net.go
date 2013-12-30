package goinfo

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

type Traffic struct {
	Name     string
	Receive  int64
	Transmit int64
	Time     time.Time
}

func (t *Traffic) String() (net string) {
	if t == nil {
		return
	}
	return fmt.Sprintf("%s receive:%s, transmit:%s, %v", t.Name, ByteSize(t.Receive).String(), ByteSize(t.Transmit).String(), t.Time.Format(gTimeFarmat))
}

// 通过读取 /proc/net/dev 获取流量信息
func (xm *Xminfo) Net() (traffic []*Traffic, err error) {
	traffic = make([]*Traffic, 0)
	b, err := ioutil.ReadFile(gNet)
	if err != nil {
		return
	}
	now := time.Now()
	s := strings.SplitAfter(string(b), "\n")
	length := len(s)
	for i := 2; i < length; i++ {
		t := strings.Fields(s[i])
		if len(t) == 17 {
			name := strings.Replace(t[0], ":", "", -1)
			receive := string2int64(t[1])
			transmit := string2int64(t[10])
			traffic = append(traffic, &Traffic{Name: name, Receive: receive, Transmit: transmit, Time: now})
		}
	}

	return
}
