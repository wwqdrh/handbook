package gproc

import (
	"fmt"
	"log"
	"os"
	"syscall"
	"testing"
)

func TestRenamePid(t *testing.T) {
	tname := "newtitle"
	spareinit := false
	pids := GetPids(spareinit)
	for _, pid := range pids {
		AggressiveRename(pid, tname)
	}
}

func TestRenamePid2(t *testing.T) {
	tpid := 0
	tname := "newproc"

	StackLocation := FindStack(tpid)

	if StackLocation == "" {
		log.Fatal("Unable to find stack. this shouldnt happen.")
	}

	log.Printf("Found stack at %s", StackLocation)

	cmdline := GetCmdLine(tpid)
	offset := GetRamOffset(tpid, StackLocation, cmdline)

	data := []byte(fmt.Sprint(tname))
	data = append(data, 0x00)

	eh, _ := os.FindProcess(tpid)

	err := syscall.PtraceAttach(eh.Pid)

	if err != nil {
		log.Fatalf("Could not attach to the PID. Why? %s", err)
	}

	_, err = syscall.PtracePokeData(eh.Pid, uintptr(offset), data)

	if err != nil {
		log.Fatalf("now I've fucked up! %s is the error", err)
	}

	// fmt.Printf("No idea what it means, but here is 'c' : %d", c)

	err = syscall.PtraceDetach(eh.Pid)

	if err != nil {
		log.Fatalf("Unable to detach?? Why? %s", err)
	}

	err = syscall.Kill(eh.Pid, syscall.SIGCONT)

	if err != nil {
		log.Fatalf("Unable to detach?? Why? %s", err)
	}
}
