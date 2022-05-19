package sysinfo

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

// cpu基本信息
func CPUBasic() {
	physicalCnt, _ := cpu.Counts(false)
	logicalCnt, _ := cpu.Counts(true)
	fmt.Printf("physical count:%d logical count:%d\n", physicalCnt, logicalCnt)

	totalPercent, _ := cpu.Percent(3*time.Second, false)
	perPercents, _ := cpu.Percent(3*time.Second, true)
	fmt.Printf("total percent:%v per percents:%v", totalPercent, perPercents)
}

// cpu详细信息
func CPUDetail() {
	infos, _ := cpu.Info()
	for _, info := range infos {
		data, _ := json.MarshalIndent(info, "", " ")
		fmt.Print(string(data))
	}
}

// cpu时间占用
// type TimesStat struct {
// 	CPU       string  `json:"cpu"`
// 	User      float64 `json:"user"`
// 	System    float64 `json:"system"`
// 	Idle      float64 `json:"idle"`
// 	Nice      float64 `json:"nice"`
// 	Iowait    float64 `json:"iowait"`
// 	Irq       float64 `json:"irq"`
// 	Softirq   float64 `json:"softirq"`
// 	Steal     float64 `json:"steal"`
// 	Guest     float64 `json:"guest"`
// 	GuestNice float64 `json:"guestNice"`
// }

func CPUTime() {
	infos, _ := cpu.Times(true)
	for _, info := range infos {
		data, _ := json.MarshalIndent(info, "", " ")
		fmt.Print(string(data))
	}
}
