## goinfo
======

Getting Server system info by reading file in '/proc', such as CPU, Memory, Disk stat, IO, etc.

## Installation

go get github.com/luosangnanka/goinfo

## Features
* Fast, and little resource consumption.

	Only because it works by reading linux file but not depand on any linux command such as 'free', etc.
	
* Good compatibility.
	It works in any Linux machines.

## Reading File List.

	CPU: 		"/proc/stat"       	
	Memory: 	"/proc/meminfo"    	
	Net: 		"/proc/net/dev"    	
	Host: 		"/proc/uptime"	  	
	Disk: 		"/proc/diskstats"  	
	Loadavg: 	"/proc/loadavg"    	
	TCP & UDP: 	"/proc/net/snmp"   	
