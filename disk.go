package goinfo

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// DiskStat disk stat info struct.
type DiskStat struct {
	ID     int64  `json:"id"`     // 编号
	SdName string `json:"sdname"` // 设备名称

	RCompleNum      int64 `json:"r_comple_num"`       // 读完成次数
	RCompleMergeNum int64 `json:"r_comple_merge_num"` // 合并完成就
	RSectorsNum     int64 `json:"r_sectors_num"`      // 读扇区次数
	RSpentMill      int64 `json:"r_spent_mill"`       // 读操作花费毫秒数

	WConpleNum      int64 `json:"w_conple_num"`       // 写完成次数
	WCompleMergeNum int64 `json:"w_comple_merge_num"` // 合并写完成次数
	WSectirsNum     int64 `json:"w_sectirs_num"`      // 写扇区次数
	WSpentMill      int64 `json:"w_spent_mill"`       // 写操作花费毫秒数

	RWResquestNum int64 `json:"rw_resquest_num"` // 正在处理的输入/输出请求书
	RWSpentMill   int64 `json:"rw_spent_mill"`   // 输入/输出操作花费的毫秒数
	RWSpentMillW  int64 `json:"rw_spent_mill_w"` // 输入/输出操作花费的加权毫秒数
}

// String format the disk stat struct.
func (d *DiskStat) String() (disk string) {
	if d == nil {
		return
	}

	return fmt.Sprintf("ID: %d, SdName: %s, RCompleNum: %d, RCompleMergeNum: %d, RSectorsNum: %d, RSpentMill: %d, WConpleNum: %d, WCompleMergeNum: %d, WSectirsNum: %d, WSpentMill: %d, RWResquestNum: %d, RWSpentMill: %d, RWSpentMillW: %d", d.ID, d.SdName, d.RCompleNum, d.RCompleMergeNum, d.RSectorsNum, d.RSpentMill, d.WConpleNum, d.WCompleMergeNum, d.WSectirsNum, d.WSpentMill, d.RWResquestNum, d.RWResquestNum, d.RWSpentMillW)
}

// Disk getting disk stat info by reading linux /proc/diskstats file.
func (x *Xminfo) Disk() (disk []*DiskStat, err error) {
	disk = make([]*DiskStat, 0)
	b, err := ioutil.ReadFile(gDiskFile)
	if err != nil {
		return
	}
	s := strings.SplitAfter(string(b), "\n")
	for _, v := range s {
		dT := strings.Fields(v)
		if len(dT) < 3 {
			continue
		}
		if strings.HasPrefix(dT[2], "sd") {
			if len(dT) < 14 {
				err = fmt.Errorf("disk stat info fields has no enough fields")
				return
			}
			disk = append(disk, &DiskStat{string2Int64(dT[1]), dT[2], string2Int64(dT[3]), string2Int64(dT[4]), string2Int64(dT[5]), string2Int64(dT[6]), string2Int64(dT[7]), string2Int64(dT[8]), string2Int64(dT[9]), string2Int64(dT[10]), string2Int64(dT[11]), string2Int64(dT[12]), string2Int64(dT[13])})
		}
	}

	return
}
