package generator

import (
	"fmt"
	"os"
)

type File struct {
	Name    string
	Path    string
	Content []byte

	osFile *os.File
}

func NewFile(name string, path string, content []byte) *File {
	return &File{
		Name:    name,
		Path:    path,
		Content: content,
		osFile:  nil,
	}
}

func (f *File) Save() (*os.File, error) {
	nf, err := os.Create(fmt.Sprintf("%s/%s", f.Path, f.Name))
	if err != nil {
		return nil, err
	}

	defer nf.Close()

	_, err = nf.Write(f.Content)
	if err != nil {
		return nil, err
	}

	return nf, nil
}

func (f *File) Write(p []byte) (n int, err error) {
	f.Content = append(f.Content, p...)
	return len(p), nil
}
