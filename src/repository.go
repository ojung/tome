package tome

import (
	"fmt"
	"os"
)

// Repository has the methods to store shell commands.
type Repository interface {
	Store(cmd string) (bool, error)
}

// FileRepository is a basic kind of repository that simply writes to a file.
type FileRepository struct {
	path string
}

// NewFileRepository creates a new FileRepository.
func NewFileRepository(p string) Repository {
	return FileRepository{path: p}
}

// Store the given cmd in the Repository.
func (r FileRepository) Store(cmd string) (bool, error) {
	f, err := os.OpenFile(r.path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return false, err
	}
	defer f.Close()

	if _, err = f.WriteString(fmt.Sprintf("%s\n", cmd)); err != nil {
		return false, err
	}

	return true, nil
}