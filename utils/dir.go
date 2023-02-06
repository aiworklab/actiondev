package utils

import (
	"os"

	"github.com/aichy126/igo/log"
)

func CreatePathIsNotExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		// create directory
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return false, err
		} else {
			return true, nil
		}
	}
	return false, err
}

func DelDir(dirpath string) error {
	err := os.RemoveAll(dirpath)
	if err != nil {
		log.Error("del dir", log.Any("error", err))
		return err
	}
	return nil
}
