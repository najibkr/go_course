package common

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetUserInput(message string, allowEmpty bool) string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%v: ", message)
		input, err := reader.ReadString('\n')
		if err != nil {
			continue
		}
		if len(input) > 0 {
			return input
		}
		if len(strings.Trim(input, "\n\r\t")) == 0 && !allowEmpty {
			continue
		}

	}
}
