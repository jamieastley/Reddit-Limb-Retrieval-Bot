package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
)

type FileManager interface {
	initDir(path string) error
	read(string) ([]byte, error)
	write(path string, values []string) error
}

type csvFileManager struct {
}

func (f csvFileManager) initDir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create dir: %s", err)
	}

	return nil
}

func (f csvFileManager) read(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

func (f csvFileManager) write(path string, values []string) error {
	// Create file if not exists
	if file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600); err != nil {
		return fmt.Errorf("failed to open %s, err: %s", path, err)
	} else {
		defer file.Close()

		writer := csv.NewWriter(file)
		err := writer.Write(values)
		if err != nil {
			return fmt.Errorf("failed to write to %s, err: %s", path, err)
		}
		writer.Flush()
	}

	return nil
}
