package filehelpers

type IFileHelper interface {
	CreateFileWithCurrentDate(directoryPath string) (string, error)
	CreateFile(filePath string) error
	DeleteFile(filePath string) error
	WriteFile(filePath string, data string) error
}
