package sysinfo

import (
	"encoding/json"
	"fmt"

	"github.com/shirou/gopsutil/mem"
)

// 内存信息, virtualmemory()获取内存信息，swapmemory() 获取交换内存的信息
// type SwapMemoryStat struct {
// 	Total       uint64  `json:"total"`
// 	Used        uint64  `json:"used"`
// 	Free        uint64  `json:"free"`
// 	UsedPercent float64 `json:"usedPercent"`
// 	Sin         uint64  `json:"sin"`
// 	Sout        uint64  `json:"sout"`
// 	PgIn        uint64  `json:"pgin"` 载入页数
// 	PgOut       uint64  `json:"pgout"` 淘汰页数
// 	PgFault     uint64  `json:"pgfault"` 缺页错误数
//   }

func MemoSwap() {
	swapMemory, _ := mem.SwapMemory()
	data, _ := json.MarshalIndent(swapMemory, "", " ")
	fmt.Println(string(data))
}
