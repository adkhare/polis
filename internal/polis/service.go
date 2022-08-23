package polis

import (
	"fmt"
)

type Service struct {
	Name string `yaml:"Name"`
}

func (s Service) Apply() (Status, error) {
	// Starts or restarts the service
	if !s.Check() {
		fmt.Printf("Starting service %s\n", s.Name)
		exitCode, err := ExecuteCommand(fmt.Sprintf(`/usr/bin/sudo service %s start`, s.Name))
		if err != nil {
			return Failure, err
		}

		if exitCode == Success {
			fmt.Printf("%s installed successfully\n", s.Name)
			return Success, nil
		} else {
			return Failure, nil
		}
	}

	return Success, nil
}

func (s Service) Check() bool {
	// Check if the service is running
	exitCode, err := ExecuteCommand(fmt.Sprintf(`/usr/bin/sudo systemctl status %s`, s.Name))
	fmt.Printf("Checking for service - %s. Exit code: %d\n", s.Name, exitCode)
	if err != nil {
		fmt.Printf("Error while checking service: %s. Error: %s\n", s.Name, err)
		return false
	}

	if exitCode == Success {
		return true
	} else {
		return false
	}
}

func (s Service) TriggerExec(Trigger string) (Status, error) {
	// Execute the trigger action
	exitCode, err := ExecuteCommand(fmt.Sprintf(`/usr/bin/sudo systemctl %s %s`, Trigger, s.Name))
	if err != nil {
		return Failure, err
	}

	if exitCode == Success {
		return Success, nil
	} else {
		return Failure, nil
	}
}

func (s Service) UnApply() (Status, error) {
	// stops the service
	if s.Check() {
		exitCode, err := ExecuteCommand(fmt.Sprintf(`/usr/bin/sudo systemctl stop %s`, s.Name))

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
