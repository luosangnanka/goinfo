/*===========================================
*   Copyright (C) 2013 All rights reserved.
*
*   company      : xiaomi
*   author       : zhangye
*   email        : zhangye@xiaomi.com
*   date         : 2013-12-12 11:49:25
*   description  : package info
*
=============================================*/
package xminfo

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
