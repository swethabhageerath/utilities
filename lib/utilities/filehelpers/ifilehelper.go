package filehelpers

import "os"

type IFileHelper interface {
	CreateFileWithCurrentDate(directoryPath string) (string, error)
	CreateFile(filePath string) (*os.File, error)
	DeleteFile(filePath string) error
	WriteFile(filePath string, data string) error
}
