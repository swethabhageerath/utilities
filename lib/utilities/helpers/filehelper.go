package helpers

import (
	"fmt"
	"os"
	"time"
)

type FileHelper struct {
	env EnvironmentHelper
}

func New(env EnvironmentHelper) FileHelper {
	return FileHelper{
		env: env,
	}
}

func (f FileHelper) CreateFileWithCurrentDate(directoryPath string) (string, error) {
	fileNameWithCurrentDate := fmt.Sprintf("%s.txt", time.Now().Format("2006-01-02"))
	fullFilePath := fmt.Sprintf("%s/%s", directoryPath, fileNameWithCurrentDate)

	if _, err := os.Stat(fullFilePath); err == nil {
		fmt.Println("first", err)
		return fullFilePath, nil
	}
	fmt.Println("path", fullFilePath)
	e := f.CreateFile(fullFilePath)
	if e != nil {
		fmt.Println("Second", e)
		return "", e
	}

	return fullFilePath, nil
}

func (f FileHelper) CreateFile(filePath string) error {
	_, err := os.Create(filePath)
	fmt.Println(err)
	if err != nil {
		return err
	}
	return nil
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
