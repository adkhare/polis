package polis

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

type File struct {
	Path     string      `yaml:"Path"`
	Contents string      `yaml:"Contents"`
	Owner    string      `yaml:"Owner"`
	Group    string      `yaml:"Group"`
	Perm     fs.FileMode `yaml:"Perm"`
}

func (f File) Apply() (Status, error) {
	// Check if file exists. if it does, return success
	if f.Check() {
		return Success, nil
		//TODO: check if the contents are same. if not, rewrite the file
	}

	fmt.Printf("Creating file: %s", f.Path)
	// Create directory
	dir := filepath.Dir(f.Path)
	_, err := os.Stat(dir)
	if err == nil {
		errDir := os.MkdirAll(f.Path, f.Perm)
		if errDir != nil {
			return Failure, err
		}
	}

	// Write a file with contents
	err = os.WriteFile(f.Path, []byte(f.Contents), f.Perm)
	if err != nil {
		return Failure, err
	}

	return Success, nil
}

func (f File) Check() bool {
	// Check if the file exists with given metadata and contents
	fileInfo, err := os.Stat(f.Path)
	fmt.Printf("Checking for file. Exit code: %v", fileInfo)
	if err != nil {
		return false
	} else {
		return true
	}
}

func (f File) TriggerExec(Trigger string) (Status, error) {
	// Execute the trigger action
	return Success, nil
}

func (f File) UnApply() (Status, error) {
	// Check if file exists. if it does not, return success
	if !f.Check() {
		return Success, nil
		//TODO: check if the contents are same. if not, rewrite the file
	}

	// Write a file with contents
	err := os.RemoveAll(f.Path)
	if err != nil {
		return Failure, err
	}

	return Success, nil
}
