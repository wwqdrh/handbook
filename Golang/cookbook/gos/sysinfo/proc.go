package sysinfo

import (
	"fmt"

	"github.com/shirou/gopsutil/process"
)

// 进程信息, 可以获取当前运行的进程信，对进程进行操作
// 并不能创建新进程 只是通过传入pid管理这个实例
func BasicProcess() {
	var rootProcess *process.Process
	processes, _ := process.Processes()
	for _, p := range processes {
		if p.Pid == 0 {
			rootProcess = p
			break
		}
	}

	fmt.Println(rootProcess)

	fmt.Println("children:")
	children, _ := rootProcess.Children()
	for _, p := range children {
		fmt.Println(p)
	}
}
