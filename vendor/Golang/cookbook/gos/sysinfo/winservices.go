//go:build windows

package sysinfo

import (
	"fmt"
	"syscall"
	"unsafe"
)

// winservices子包可以获取 Windows 系统中的服务信息，内部使用了golang.org/x/sys包。在winservices中，一个服务对应一个Service结构：
// type Service struct {
// 	Name   string
// 	Config mgr.Config
// 	Status ServiceStatus
// 	// contains filtered or unexported fields
// }

// mgr.Config为包golang.org/x/sys中的结构，该结构详细记录了服务类型、启动类型（自动/手动）、二进制文件路径等信息：

// ServiceStatus结构记录了服务的状态：
// State：为服务状态，有已停止、运行、暂停等；
// Accepts：表示服务接收哪些操作，有暂停、继续、会话切换等；
// Pid：进程 ID；
// Win32ExitCode：应用程序退出状态码。

func WinServiceList() {
	services, _ := winservices.ListServices()

	for _, service := range services {
		newservice, _ := winservices.NewService(service.Name)
		newservice.GetServiceDetail()
		fmt.Println("Name:", newservice.Name, "Binary Path:", newservice.Config.BinaryPathName, "State: ", newservice.Status.State)
	}
}

func GetFreeDisk(path string) (int64, error) {
	kernel32, err := syscall.LoadLibrary("Kernel32.dll")
	if err != nil {
		return 0, err
	}
	defer syscall.FreeLibrary(kernel32)
	GetDiskFreeSpaceEx, err := syscall.GetProcAddress(syscall.Handle(kernel32), "GetDiskFreeSpaceExW")

	if err != nil {
		return 0, err
	}

	lpFreeBytesAvailable := int64(0)
	lpTotalNumberOfBytes := int64(0)
	lpTotalNumberOfFreeBytes := int64(0)
	syscall.Syscall6(uintptr(GetDiskFreeSpaceEx), 4,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("C:"))),
		uintptr(unsafe.Pointer(&lpFreeBytesAvailable)),
		uintptr(unsafe.Pointer(&lpTotalNumberOfBytes)),
		uintptr(unsafe.Pointer(&lpTotalNumberOfFreeBytes)), 0, 0)

	return lpTotalNumberOfFreeBytes / 1024 / 1024.0, nil
}
