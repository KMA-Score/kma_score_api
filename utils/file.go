package utils

import (
	"errors"
	"os"
)

func CreateDirIfNotExist(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, 0755)
		if err != nil {
			return errors.New("error on create directory: " + path)
		}
	}
	return nil
}
