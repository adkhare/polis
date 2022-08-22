package polis

import (
	"io/fs"
	"os"
	"path/filepath"
)

type File struct {
	path     string
	contents string
	owner    string
	group    string
	perm     fs.FileMode
}

func (f File) Apply() (Status, error) {
	// Check if file exists. if it does, return success
	if f.Check() {
		return Success, nil
		//TODO: check if the contents are same. if not, rewrite the file
	}

	// Create directory
	dir := filepath.Dir(f.path)
	_, err := os.Stat(dir)
	if err == nil {
		errDir := os.MkdirAll(f.path, f.perm)
		if errDir != nil {
			return Failure, err
		}
	}

	// Write a file with contents
	err = os.WriteFile(f.path, []byte(f.contents), f.perm)
	if err != nil {
		return Failure, err
	}

	return Success, nil
}

func (f File) Check() bool {
	// Check if the file exists with given metadata and contents
	_, err := os.Stat(f.path)

	if err != nil {
		return true
	} else {
		return false
	}
}
