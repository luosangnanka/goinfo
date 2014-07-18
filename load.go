package goinfo

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Loadavg is load info struct.
type Loadavg struct {
	La1       float64 `json:"la1"`
	La5       float64 `json:"la2"`
	La15      float64 `json:"la3"`
	Processes string  `json:"processes"`
	MaxPid    int64   `json:"max_pid"`
}

// String format the loadavg struct.
func (l *Loadavg) String() (load string) {
	if l == nil {
		return
	}

	return fmt.Sprintf("La1:%.2f, La5:%.2f, La15:%.2f, Processes:%s, MaxPid:%d", l.La1, l.La5, l.La15, l.Processes, l.MaxPid)
}

// Load getting load info by reading linux /proc/loadavg file.
func (x *Xminfo) Load() (load *Loadavg, err error) {
	load = new(Loadavg)
	b, err := ioutil.ReadFile(gLoadavgFile)
	if err != nil {
		return
	}
	s := strings.Fields(string(b))
	if len(s) < 5 {
		err = fmt.Errorf("loadavg info fields has no enough fields")
		return
	}

	load.La1 = string2Float64(s[0])
	load.La5 = string2Float64(s[1])
	load.La15 = string2Float64(s[2])
	load.Processes = s[3]
	load.MaxPid = string2Int64(s[4])

	return
}
