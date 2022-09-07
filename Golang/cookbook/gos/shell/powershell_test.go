package shell

import "testing"

func TestLocalShell(t *testing.T) {
	// choose a backend
	back := &Local{}

	// start a local powershell process
	shell, err := ps.New(back)
	if err != nil {
		panic(err)
	}
	defer shell.Exit()

	// ... and interact with it
	stdout, stderr, err := shell.Execute("Get-WmiObject -Class Win32_Processor")
	if err != nil {
		panic(err)
	}

	fmt.Println(stdout)
}

func TestRemoteShell(t *testing.T) {
	// choose a backend
	back := &Local{}

	// start a local powershell process
	shell, err := ps.New(back)
	if err != nil {
		panic(err)
	}

	// prepare remote session configuration
	config := NewSessionConfig()
	config.ComputerName = "remote-pc-1"

	// create a new shell by wrapping the existing one in the session middleware
	session, err := NewSession(shell, config)
	if err != nil {
		panic(err)
	}
	defer session.Exit() // will also close the underlying ps shell!

	// everything run via the session is run on the remote machine
	stdout, stderr, err = session.Execute("Get-WmiObject -Class Win32_Processor")
	if err != nil {
		panic(err)
	}

	fmt.Println(stdout)
}
