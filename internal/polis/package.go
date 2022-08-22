package polis

import (
	"fmt"
)

type Package struct {
	name string
}

func (p Package) Apply() (Status, error) {
	// Installs the package if the package is not installed
	if !p.Check() {
		exitCode, err := ExecuteCommand(fmt.Sprintf(`sudo apt-get update && sudo apt-get install %s`, p.name))

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
	exitCode, err := ExecuteCommand(fmt.Sprintf(`sudo dpkg-query -f '${Package}\t${db:Status-Abbrev}\t${Version}\t${Name}' -W %s`, p.name))
	if err != nil {
		return false
	}

	if exitCode == Success {
		return true
	} else {
		return false
	}
}
