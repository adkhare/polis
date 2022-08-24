package polis

import (
	"fmt"
	"io/fs"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"syscall"
)

type File struct {
	Path     string      `yaml:"Path"`
	Contents string      `yaml:"Contents"`
	Owner    string      `yaml:"Owner"`
	Group    string      `yaml:"Group"`
	Perm     fs.FileMode `yaml:"Perm"`
}

/**
 * This creates a file if the file does not exist with given permissions, contents, owner and group
 */
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

	osUser, err := user.Lookup(f.Owner)
	if err != nil {
		return Failure, err
	}

	osGroup, err := user.LookupGroup(f.Group)
	if err != nil {
		return Failure, err
	}
	uid, _ := strconv.Atoi(osUser.Uid)
	gid, _ := strconv.Atoi(osGroup.Gid)
	err = os.Chown(f.Path, uid, gid)

	if err != nil {
		return Failure, err
	}

	return Changed, nil
}

/**
 * This checks if the file exists with given permissions, contents, owner and group
 */
func (f File) Check() bool {
	// Check if the file exists with given metadata and contents
	fmt.Printf("Checking file: %s\n", f.Path)
	fileInfo, err := os.Stat(f.Path)
	if err != nil {
		fmt.Printf("Error while checking file: %s. Error: %s\n", f.Path, err)
		return false
	}

	// Check if the permissions are same
	if fileInfo.Mode() != f.Perm {
		fmt.Printf("Different file permissions for file: %s\n. Source %s; Should be %s", f.Path, fileInfo.Mode(), f.Perm)
		return false
	}

	// Check if ownership is same
	stat := fileInfo.Sys().(*syscall.Stat_t)
	uidFile := stat.Uid
	gidFile := stat.Gid

	osUser, err := user.Lookup(f.Owner)
	if err != nil {
		return false
	}

	osGroup, err := user.LookupGroup(f.Group)
	if err != nil {
		return false
	}
	uid, _ := strconv.Atoi(osUser.Uid)
	gid, _ := strconv.Atoi(osGroup.Gid)

	if uid != int(uidFile) || gid != int(gidFile) {
		fmt.Printf("Different file ownership: %s\n", f.Path)
		return false
	}

	fileData, err := os.ReadFile(f.Path)
	if err != nil {
		fmt.Printf("Error while checking file: %s. Error: %s\n", f.Path, err)
		return false
	}

	if string(fileData) != f.Contents {
		fmt.Printf("Different file contents: %s\n", f.Path)
		return false
	}
	return true
}

/**
 * This is not implemented since there does not seem to be usecase to trigger a file module
 */
func (f File) TriggerExec(Trigger string) (Status, error) {
	// Return success with no error since file cannot be triggered
	return Success, nil
}

/**
 * This will delete the file
 */
func (f File) UnApply() (Status, error) {
	// Remove a file
	if f.Check() {
		err := os.RemoveAll(f.Path)
		fmt.Printf("Removing file: %s\n", f.Path)
		if err != nil {
			return Failure, err
		}
	}
	return Success, nil
}
