package utils

import "os"

func Stat(path string) (fileInfo os.FileInfo, err error) {
	if fileInfo, err = os.Stat(path); os.IsNotExist(err) {
		fileInfo, err = os.Stat(path + ".json")
	}
	return
}
