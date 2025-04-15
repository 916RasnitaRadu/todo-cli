package file

import (
	"fmt"
	"log"
	"os"
	"syscall"
)

func LoadFile(path string) (*os.File, error) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Println("[ERROR]: error at opening the file")
		return nil, fmt.Errorf("error at opening the file")
	}

	// exclusive lock obtained on a file descriptor
	if err := syscall.Flock(int(file.Fd()), syscall.LOCK_EX); err != nil {
		_ = file.Close()
		log.Printf("[ERROR]: %v", err)
		return nil, err
	}

	return file, nil
}

func CloseFile(f *os.File) error {
	syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
	return f.Close()
}
