package polis

import (
	"os/exec"
	"syscall"
)

// Base polis struct which is the generic state
type Polis struct {
	ModuleType    string `yaml:"ModuleType"`
	Ensure        bool   `yaml:"Ensure"`
	Triggers      string `yaml:"Triggers"`
	TriggerAction string `yaml:"TriggerAction"`
	Module        Module `yaml:"Module"`
}

type Status int

const (
	Success = 0
	Failure = 1
)

type Module interface {
	Apply() (Status, error)
	Check() bool
	TriggerExec(string) (Status, error)
	UnApply() (Status, error)
}

func (p Polis) Execute() (string, error) {
	// Ensure the module is applied
	if p.Ensure {
		_, err := p.Module.Apply()
		if err != nil {
			return "", err
		}
	}

	// Check if trigger exists so that it can be returned
	trigger := ""
	if p.Triggers != "" {
		trigger = p.Triggers
	}

	// Check if triggerAction exists so that can be executed
	if p.TriggerAction != "" {
		_, err := p.Module.TriggerExec(p.TriggerAction)
		if err != nil {
			return "", err
		}
	}

	return trigger, nil
}

func ExecuteCommand(cmdString string) (int, error) {
	cmd := exec.Command(cmdString)
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
	return Failure, nil
}
