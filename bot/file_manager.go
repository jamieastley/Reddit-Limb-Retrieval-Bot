package bot

import (
	"io/fs"
	"io/ioutil"
)

type FileManager interface {
	read(string) ([]byte, error)
	write(string, []byte, fs.FileMode) error
}

type FileManagerExecutor struct{}

func (f *FileManagerExecutor) read(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

func (f *FileManagerExecutor) write(filename string, input []byte, filemode fs.FileMode) error {
	return ioutil.WriteFile(filename, input, filemode)
}
