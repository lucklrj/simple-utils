package system

import (
	"bufio"
	"io"
	"os"
)

func ReadAllFromFile(file io.Reader) ([]string, error) {
	result := make([]string, 0)

	line := bufio.NewReader(file)
	for {
		content, _, err := line.ReadLine()
		if err != nil && err != io.EOF {
			return make([]string, 0), err
		}
		contentStr := string(content)
		if contentStr != "" {
			result = append(result, contentStr)
		}
		if err == io.EOF {
			break
		}
	}
	return result, nil
}

func WriteFile(filename string, data []byte, perm os.FileMode) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	defer f.Close()
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		return io.ErrShortWrite
	}
	return nil
}
