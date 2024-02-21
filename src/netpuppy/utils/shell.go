package utils

import (
	"fmt"
	"io"
	"os/exec"
)

func StartHelperShell(thisPeer *Peer) error {
	// Which peer are we? -- changes shell execution
	fmt.Printf("this peer connection type: %v", thisPeer.ConnectionType)
	if thisPeer.ConnectionType == "connect-back" {
		fmt.Println("made it into cb if statement\n")
		// If bash exists on the system, find it, save the path:
		bashCopPath, err := exec.LookPath(`/bin/bash`) // bashPath @0xfaraday
		if err != nil {
			return err
		}
		bCmd := exec.Command(bashCopPath)
		bashIn, eRr := bCmd.StdinPipe()
		if eRr != nil {
			return eRr
		}

		// If bash exists, attach the exec.Cmd struct to the peer:
		thisPeer.ShellProcess = *bCmd

		//	// Start the shell:
		//	var erR error = thisPeer.ShellProcess.Start()
		//	if erR != nil {
		//		return erR
		//	}

		// Test pipe into stdin
		go func() {
			defer bashIn.Close()
			io.WriteString(bashIn, "/usr/bin/date")
		}()

		out, eRR := bCmd.CombinedOutput()
		if eRR != nil {
			return eRR
		}

		fmt.Printf("output: %v\n", string(out))
	}

	fmt.Printf("Shell mod: process field: %v\n", thisPeer.ShellProcess.Process)
	return nil
}
