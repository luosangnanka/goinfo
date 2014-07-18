package goinfo

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

// Traffic is net transmit info struct.
type Traffic struct {
	Name     string `json:"name"`
	Receive  int64  `json:"receive"`
	Transmit int64  `json:"transmit"`
	Time     string `json:"time"`
}

// String format the traffic struct.
func (t *Traffic) String() (net string) {
	if t == nil {
		return
	}

	return fmt.Sprintf("%s receive:%s, transmit:%s, %v", t.Name, ByteSize(t.Receive).String(), ByteSize(t.Transmit).String(), t.Time)
}

// Net getting net info by reading linux /proc/net/dev file.
func (x *Xminfo) Net() (traffic []*Traffic, err error) {
	traffic = make([]*Traffic, 0)
	b, err := ioutil.ReadFile(gNetFile)
	if err != nil {
		return
	}
	now := time.Now().Format(gTimeFarmat)
	s := strings.SplitAfter(string(b), "\n")
	length := len(s)
	for i := 2; i < length; i++ {
		t := strings.Fields(s[i])
		if len(t) == 17 {
			name := strings.Replace(t[0], ":", "", -1)
			receive := string2Int64(t[1])
			transmit := string2Int64(t[10])
			traffic = append(traffic, &Traffic{Name: name, Receive: receive, Transmit: transmit, Time: now})
		}
	}

	return
}
