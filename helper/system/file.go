package system

import (
	"bufio"
	"io"
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
