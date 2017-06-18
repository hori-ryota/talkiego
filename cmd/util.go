package cmd

import "os"

func IsExistFile(filepath string) bool {
	_, err := os.Stat(filepath)
	return !os.IsNotExist(err)
}
