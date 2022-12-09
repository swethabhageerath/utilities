package filehelpers

import (
	"fmt"
	"github.com/pkg/errors"
	"libraries/utilities/lib/utilities/environmenthelper"
	"os"
	"time"
)

type FileHelper struct {
	env environmenthelper.IEnvironmentHelper
}

func New(env environmenthelper.IEnvironmentHelper) FileHelper {
	return FileHelper{
		env: env,
	}
}

func (f FileHelper) CreateFileWithCurrentDate(directoryPath string) (string, error) {
	fileNameWithCurrentDate := fmt.Sprintf("%s.txt", time.Now().Format("2006-01-02"))
	fullFilePath := fmt.Sprintf("%s/%s", directoryPath, fileNameWithCurrentDate)

	if _, err := os.Stat(fullFilePath); errors.As(err, os.ErrNotExist) {
		e := f.CreateFile(fullFilePath)
		if e != nil {
			return "", e
		}
	}

	return fullFilePath, nil
}

func (f FileHelper) CreateFile(filePath string) error {
	_, err := os.Create(filePath)
	return err
}

func (f FileHelper) DeleteFile(filePath string) error {
	return os.Remove(filePath)
}

func (f FileHelper) WriteFile(filePath string, data string) error {
	file, err := f.OpenFileForAppendWrite(filePath)
	if err != nil {
		return err
	}
	_, e := file.WriteString(data)
	return e
}

func (f FileHelper) OpenFileForAppendWrite(filePath string) (*os.File, error) {
	return os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0644)
}
