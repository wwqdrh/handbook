package sysinfo

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/shirou/gopsutil/host"
)

// 主机信息，获取开机时间、内核版本号、平台信息等
func HostBoot() {
	timestamp, _ := host.BootTime()
	t := time.Unix(int64(timestamp), 0)
	fmt.Println(t.Local().Format("2006-01-02 15:04:05"))
}

// 内核版本和平台信息
func HostKernelInfo() {
	version, _ := host.KernelVersion()
	fmt.Println(version)

	platform, family, version, _ := host.PlatformInformation()
	fmt.Println("platform:", platform)
	fmt.Println("family:", family)
	fmt.Println("version:", version)
}

// 中断用户
// type UserStat struct {
// 	User     string `json:"user"`
// 	Terminal string `json:"terminal"`
// 	Host     string `json:"host"`
// 	Started  int    `json:"started"`
// }
func HostUsers() {
	users, _ := host.Users()
	for _, user := range users {
		data, _ := json.MarshalIndent(user, "", " ")
		fmt.Println(string(data))
	}
}
