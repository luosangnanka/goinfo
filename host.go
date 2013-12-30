package goinfo

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type HostName struct {
	Name   string
	Boot   time.Time
	Uptime string
}

func (h *HostName) String() (host string) {
	if h == nil {
		return
	}
	return fmt.Sprintf("user:%s, boot:%s, uptime:%s", h.Name, h.Boot.Format(gTimeFarmat), h.Uptime)
}

// 通过读取 /proc/uptime 获取系统运行信息
func (xm *Xminfo) Host() (host *HostName, err error) {
	host = new(HostName)
	b, err := ioutil.ReadFile(gHost)
	if err != nil {
		return
	}
	length := len(b)
	for i := 0; i < length; i++ {
		if b[i] == ' ' {
			b = b[0:i]
			break
		}
	}
	t := string(b) + "s"
	// 获得系统运行的时间
	d, err := time.ParseDuration(t)
	if err != nil {
		return
	}
	host.Uptime = d.String()
	// 获得系统登录时间
	host.Boot = time.Now().Add(-d)

	host.Name, err = os.Hostname()
	if err != nil {
		return
	}

	return
}
