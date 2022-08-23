package polis

import (
	"fmt"
)

type Package struct {
	Name   string `yaml:"Name"`
	Ensure bool   `yaml:"Ensure"`
}

func (p Package) Apply() (Status, error) {
	// Installs the package if the package is not installed
	if !p.Check() {
		fmt.Printf("Installing %s\n", p.Name)
		exitCode, err := ExecuteCommand(fmt.Sprintf(`/usr/bin/sudo apt-get -y install %s`, p.Name))

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

func (p Package) Check() bool {
	// Check if the package exists
	exitCode, err := ExecuteCommand(fmt.Sprintf(`/usr/bin/sudo dpkg-query -f '${Package}\t${db:Status-Abbrev}\t${Version}\t${Name}' -W %s`, p.Name))
	fmt.Printf("Checking for package. Exit code: %d\n", exitCode)
	if err != nil {
		fmt.Printf("Error while checking package: %s. Error: %s\n", p.Name, err)
		return false
	}

	if exitCode == Success {
		return true
	} else {
		return false
	}
}

func (p Package) TriggerExec(Trigger string) (Status, error) {
	// Execute the trigger action
	return Success, nil
}

func (p Package) UnApply() (Status, error) {
	// Installs the package if the package is not installed
	if p.Check() {
		exitCode, err := ExecuteCommand(fmt.Sprintf(`/usr/bin/sudo apt-get -y --purge remove %s`, p.Name))

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
