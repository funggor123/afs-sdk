package  afs_sdk
import (
	"fmt"
	"os"
)

func DeleteFile(filePath string, fileName string) error {
	err := os.Remove(filePath + fileName)
	if err != nil {
		return err
	}
	return nil
}

func DeleteDirectory(fileFullPath string) error {
	err := os.RemoveAll(fileFullPath)
	if err != nil {
		return err
	}
	return nil
}

func DeleteFileFullPath(fileFullPath string) error {
	err := os.Remove(fileFullPath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func OpenFile(filePath string, fileName string) (*os.File, error) {
	file, err := os.Open(filePath + fileName)
	if err != nil {
		return file, err
	}
	return file, nil
}

func OpenFileFullPath(fileFullPath string) (*os.File, error) {
	file, err := os.Open(fileFullPath)
	if err != nil {
		return file, err
	}
	return file, nil
}

func IsFileExists(filePath string, fileName string) bool {
	_, err := OpenFile(filePath, fileName)
	if err != nil {
		return false
	}
	return true
}

func IsDirExists(dirPath string, isCreate bool) bool {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		if isCreate {
			os.Mkdir(dirPath, 0755)
			return true
		}
		return false
	}
	return true
}

func Terminate() {
	os.Exit(3)
}

func GetFileSize(f *os.File) (int64, error) {
	fi, err := f.Stat()
	if err != nil {
		return 0, err
	}
	return fi.Size(), nil
}

func MoveFile(oldLocation string, newLocation string) error {
	err := os.Rename(oldLocation, newLocation)
	if err != nil {
		return err
	}
	return nil
}