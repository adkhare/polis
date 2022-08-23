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

	fmt.Printf("Creating file: %s\n", f.Path)
	// Create directory
	dir := filepath.Dir(f.Path)
	_, err := os.Stat(dir)
	if err == nil {
		errDir := os.MkdirAll(dir, f.Perm)
		if errDir != nil {
			return Failure, err
		}
	}

	// Write a file with contents
	err = os.WriteFile(f.Path, []byte(f.Contents), f.Perm)
	if err != nil {
		return Failure, err
	}

	return Changed, nil
}

func (f File) Check() bool {
	// Check if the file exists with given metadata and contents
	fmt.Printf("Checking file: %s\n", f.Path)
	fileInfo, err := os.Stat(f.Path)
	if err != nil {
		fmt.Printf("Error while checking file: %s. Error: %s\n", f.Path, err)
		return false
	}

	// Check if the permissions are same
	fmt.Printf("Checking file permissions for file: %s\n. Source %s; Should be %s", f.Path, fileInfo.Mode(), f.Perm)
	if fileInfo.Mode() != f.Perm {
		return false
	}

	fmt.Printf("Checking file contents: %s\n", f.Path)
	fileData, err := os.ReadFile(f.Path)
	if err != nil {
		fmt.Printf("Error while checking file: %s. Error: %s\n", f.Path, err)
		return false
	}

	if string(fileData) != f.Contents {
		return false
	}
	return true
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
