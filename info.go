package goinfo

import (
// "fmt"
)

type Xminfo struct {
}

var gTimeFarmat = "2006-01-02 15:04:05"

var (
	gCpu     = "/proc/stat"
	gMem     = "/proc/meminfo"
	gNet     = "/proc/net/dev"
	gHost    = "/proc/uptime"
	gDisk    = "/proc/diskstats"
	gLoadavg = "/proc/loadavg"
	gSnmp    = "/proc/net/snmp"
)

func NewXminfo() (xminfo *Xminfo) {
	return &Xminfo{}
}
