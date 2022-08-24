package polis

import "testing"

type ModuleTest struct{}

func (mt ModuleTest) Apply() (Status, error) {
	return Success, nil
}

func (mt ModuleTest) Check() bool {
	return true
}

func (mt ModuleTest) TriggerExec(string) (Status, error) {
	return Success, nil
}

func (mt ModuleTest) UnApply() (Status, error) {
	return Success, nil
}

func TestModuleInheritance(t *testing.T) {
	cases := []struct {
		m              Module // variable with type Module which is an interface
		expectedStatus Status
		expectedBool   bool
	}{
		{
			m:              ModuleTest{}, // initialize the interface variable with the inherited struct
			expectedStatus: Success,
			expectedBool:   false,
		},
	}

	for _, c := range cases {
		status, _ := c.m.Apply()
		if c.expectedStatus != status {
			t.Errorf("Expected %d but got %d", c.expectedStatus, status)
		}

		status, _ = c.m.TriggerExec("")
		if c.expectedStatus != status {
			t.Errorf("Expected %d but got %d", c.expectedStatus, status)
		}

		status, _ = c.m.UnApply()
		if c.expectedStatus != status {
			t.Errorf("Expected %d but got %d", c.expectedStatus, status)
		}

		gotBool := c.m.Check()
		if c.expectedBool != gotBool {
			t.Errorf("Expected %t but got %t", c.expectedBool, gotBool)
		}
	}
}
