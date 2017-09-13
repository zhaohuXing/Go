package basic

import (
	"errors"
	"os"
	"strings"
)

// Open is open file by name, If the file doesn't exist,create it,
// or append to the file
func Open(name string) (*os.File, error) {
	if strings.Trim(name, " ") == "" {
		return nil, errors.New("invalid args")
	}

	file, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// Write is write data to file
func Write(data []byte, file *os.File) (int, error) {
	if data == nil || file == nil {
		return -1, errors.New("invalid args")
	}

	return file.Write(data)
}

// SingleThreadCopy is copy source file to new file.
func SingleThreadCopy(sourceFile *os.File, newFile string) error {
	if strings.Trim(newFile, " ") == "" || sourceFile == nil {
		return errors.New("invalid args")
	}

	fileinfo, err := sourceFile.Stat()
	if err != nil {
		return err
	}

	size := fileinfo.Size()
	temp := make([]byte, size)
	_, err = sourceFile.Read(temp)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	nFile, err := os.Create(newFile)
	if err != nil {
		return err
	}
	_, err = nFile.Write(temp)
	if err != nil {
		return err
	}
	defer nFile.Close()
	return nil
}
