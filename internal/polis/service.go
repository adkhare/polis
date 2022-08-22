package polis

import (
	"fmt"
)

type Service struct {
	name string
}

func (s Service) Apply() (Status, error) {
	// Starts or restarts the service
	if !s.Check() {
		exitCode, err := ExecuteCommand(fmt.Sprintf(`sudo service %s start`, s.name))

		if err != nil {
			return Failure, err
		}

		if exitCode == Success {
			return Success, nil
		} else {
			return Failure, nil
		}
	}

	return Success, nil
}

func (s Service) Check() bool {
	// Check if the service is running
	exitCode, err := ExecuteCommand(fmt.Sprintf(`sudo service %s status`, s.name))
	if err != nil {
		return false
	}

	if exitCode == Success {
		return true
	} else {
		return false
	}
}
