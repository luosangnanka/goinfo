/*===========================================
*   Copyright (C) 2013 All rights reserved.
*
*   company      : xiaomi
*   author       : zhangye
*   email        : zhangye@xiaomi.com
*   date         : 2013-12-12 16:42:25
*   description  : package xminfo - disk info
*
=============================================*/
package xminfo

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type DiskStat struct {
	ID     string // 编号
	SdName string // 设备名称

	RCompleNum      int64 // 读完成次数
	RCompleMergeNum int64 // 合并完成就
	RSectorsNum     int64 // 读扇区次数
	RSpentMill      int64 // 读操作花费毫秒数

	WConpleNum      int64 // 写完成次数
	WCompleMergeNum int64 // 合并写完成次数
	WSectirsNum     int64 // 写扇区次数
	WSpentMill      int64 // 写操作花费毫秒数

	RWResquestNum int64 // 正在处理的输入/输出请求书
	RWSpentMill   int64 // 输入/输出操作花费的毫秒数
	RWSpentMillW  int64 // 输入/输出操作花费的加权毫秒数
}

func (d *DiskStat) String() (disk string) {
	if d == nil {
		return
	}
	return fmt.Sprintf("ID: %s, SdName: %s, RCompleNum: %d, RCompleMergeNum: %d, RSectorsNum: %d, RSpentMill: %d, WConpleNum: %d, WCompleMergeNum: %d, WSectirsNum: %d, WSpentMill: %d, RWResquestNum: %d, RWSpentMill: %d, RWSpentMillW: %d", d.ID, d.SdName, d.RCompleNum, d.RCompleMergeNum, d.RSectorsNum, d.RSpentMill, d.WConpleNum, d.WCompleMergeNum, d.WSectirsNum, d.WSpentMill, d.RWResquestNum, d.RWResquestNum, d.RWSpentMillW)
}

// 通过读取 /proc/diskstats 获取磁盘使用情况
func (xm *Xminfo) Disk() (disk []*DiskStat, err error) {
	disk = make([]*DiskStat, 0)
	b, err := ioutil.ReadFile(gDisk)
	if err != nil {
		return
	}
	s := strings.SplitAfter(string(b), "\n")
	for _, v := range s {
		dT := strings.Fields(v)
		if len(dT) == 0 {
			continue
		}
		if strings.HasPrefix(dT[2], "sd") {
			disk = append(disk, &DiskStat{dT[1], dT[2], string2int64(dT[3]), string2int64(dT[4]), string2int64(dT[5]), string2int64(dT[6]), string2int64(dT[7]), string2int64(dT[8]), string2int64(dT[9]), string2int64(dT[10]), string2int64(dT[11]), string2int64(dT[12]), string2int64(dT[13])})
		}
	}

	return
}
