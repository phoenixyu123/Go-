package JudgeFileExist

import "os"

func PathExist(fp string) (bool, error) {
	_, err := os.Stat(fp)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
