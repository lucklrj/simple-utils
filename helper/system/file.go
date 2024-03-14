package system

import (
	"bufio"
	"io"
	"os"
	"reflect"
)

func ReadAllFromFile(file interface{}) ([][]byte, error) {
	defer func() {
		recover()
	}()

	var result [][]byte
	var line *bufio.Reader
	if reflect.TypeOf(file).String() == "string" {
		fileHandle, err := os.Open(file.(string))
		if err != nil {
			return result, err
		}
		line = bufio.NewReader(fileHandle)

	} else {
		line = bufio.NewReader(file.(io.Reader))
	}

	for {
		content, _, err := line.ReadLine()
		if err != nil && err != io.EOF {
			return result, err
		}

		result = append(result, content)
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
