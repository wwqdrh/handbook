package sysinfo

import (
	"encoding/json"
	"fmt"

	"github.com/shirou/gopsutil/disk"
)

//磁盘信息
// type IOCountersStat struct {
// 	ReadCount        uint64 `json:"readCount"`
// 	MergedReadCount  uint64 `json:"mergedReadCount"`
// 	WriteCount       uint64 `json:"writeCount"`
// 	MergedWriteCount uint64 `json:"mergedWriteCount"`
// 	ReadBytes        uint64 `json:"readBytes"`
// 	WriteBytes       uint64 `json:"writeBytes"`
// 	ReadTime         uint64 `json:"readTime"`
// 	WriteTime        uint64 `json:"writeTime"`
// 	// ...
//   }

// 磁盘基本信息
func DiskBasicInfo() {
	mapStat, _ := disk.IOCounters()
	for name, stat := range mapStat {
		fmt.Println(name)
		data, _ := json.MarshalIndent(stat, "", "  ")
		fmt.Println(string(data))
	}
}

// 磁盘分区
// type PartitionStat struct {
// 	Device     string `json:"device"`
// 	Mountpoint string `json:"mountpoint"`
// 	Fstype     string `json:"fstype"`
// 	Opts       string `json:"opts"`
// }
func DiskPartionStat() {
	infos, _ := disk.Partitions(false)
	for _, info := range infos {
		data, _ := json.MarshalIndent(info, "", "  ")
		fmt.Println(string(data))
	}
}

// 磁盘使用率
// type UsageStat struct {
// 	Path              string  `json:"path"`
// 	Fstype            string  `json:"fstype"`
// 	Total             uint64  `json:"total"`
// 	Free              uint64  `json:"free"`
// 	Used              uint64  `json:"used"`
// 	UsedPercent       float64 `json:"usedPercent"`
// 	InodesTotal       uint64  `json:"inodesTotal"`
// 	InodesUsed        uint64  `json:"inodesUsed"`
// 	InodesFree        uint64  `json:"inodesFree"`
// 	InodesUsedPercent float64 `json:"inodesUsedPercent"`
// }
func DiskUsageStat() {
	info, _ := disk.Usage("~")
	data, _ := json.MarshalIndent(info, "", "  ")
	fmt.Println(string(data))
}
