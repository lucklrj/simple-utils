package system

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetUserInput(name, desc string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input " + desc + "[" + name + "]:")
	result, err := reader.ReadString('\n')
	return strings.Replace(result, "\n", "", -1), err
}
