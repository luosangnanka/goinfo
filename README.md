goinfo
======

通过读取 Linux 下 /proc 获得服务器获得系统信息包
the info collect for Linux, such as CPU, mem, disk, IO etc.

## Installation

go get github.com/luosangnanka/goinfo

## Features
1、快速、系统资源占用最少
2、兼容性强

## 读取的系统文件列表：

	gCpu     = "/proc/stat"       // cpu 信息
	gMem     = "/proc/meminfo"    // 内存信息
	gNet     = "/proc/net/dev"    // 流量信息
	gHost    = "/proc/uptime"	  // 主机、登陆信息
	gDisk    = "/proc/diskstats"  // 磁盘信息 
	gLoadavg = "/proc/loadavg"    // 负载信息
	gSnmp    = "/proc/net/snmp"   // tcp、udp信息

