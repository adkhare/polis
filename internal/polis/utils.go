package polis

import (
	"os/exec"
	"syscall"
)

// Generic function that can be used to execute bash commands
func ExecuteCommand(cmdString string) (int, error) {
	cmd := exec.Command("bash", "-c", cmdString)
	if err := cmd.Start(); err != nil {
		return Failure, err
	}

	if err := cmd.Wait(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			// The program has exited with an exit code != 0

			// This works on both Unix and Windows. Although package
			// syscall is generally platform dependent, WaitStatus is
			// defined for both Unix and Windows and in both cases has
			// an ExitStatus() method with the same signature.
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				return status.ExitStatus(), nil
			}
		} else {
			return Failure, err
		}
	}
	return Success, nil
}
