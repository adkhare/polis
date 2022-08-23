package polis

// Base polis struct which is the generic state
type Polis struct {
	ModuleType    string `yaml:"ModuleType"`    // Type of module: File, Service, Package
	Ensure        bool   `yaml:"Ensure"`        // Ensure to be present or absent
	Triggers      string `yaml:"Triggers"`      // ID of the module to be triggered
	TriggerAction string `yaml:"TriggerAction"` // If this module is triggered, what should be the action to be taken
	Module        Module `yaml:"Module"`        // Module spec File, Service or Package
}

type Status int

const (
	Success = 0
	Failure = 1
	Changed = 2
)

type Module interface {
	Apply() (Status, error)
	Check() bool
	TriggerExec(string) (Status, error)
	UnApply() (Status, error)
}

// Execute a specific module
func (p Polis) Execute() (string, error) {
	// Ensure the module is applied
	isChanged := false //Mark this flag if any changes are applied
	if p.Ensure {
		ensureStatus, err := p.Module.Apply()
		if err != nil {
			return "", err
		}
		if ensureStatus == Changed {
			isChanged = true
		}
	} else {
		ensureStatus, err := p.Module.UnApply()
		if err != nil {
			return "", err
		}
		if ensureStatus == Changed {
			isChanged = true
		}
	}

	// Check if trigger exists so that it can be returned
	trigger := ""
	if p.Triggers != "" && isChanged {
		trigger = p.Triggers
	}

	return trigger, nil
}
