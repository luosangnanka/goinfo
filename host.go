package goinfo

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// HostName host name and boot time struct.
type HostName struct {
	Name   string `json:"name"`
	Boot   string `json:"boot"`
	Uptime string `json:"uptime"`
}

// String format the host struct.
func (h *HostName) String() (host string) {
	if h == nil {
		return
	}

	return fmt.Sprintf("user:%s, boot:%s, uptime:%s", h.Name, h.Boot, h.Uptime)
}

// Host getting host info by reading linux /proc/uptime file.
func (x *Xminfo) Host() (host *HostName, err error) {
	host = new(HostName)
	b, err := ioutil.ReadFile(gHostFile)
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

	// system running time.
	d, err := time.ParseDuration(t)
	if err != nil {
		return
	}
	host.Uptime = d.String()

	// getting system booting time.
	host.Boot = time.Now().Add(-d).Format(gTimeFarmat)

	host.Name, err = os.Hostname()
	if err != nil {
		return
	}

	return
}
