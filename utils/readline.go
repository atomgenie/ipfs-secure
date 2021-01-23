package utils

import (
	"bufio"
	"os"
)

// Readline Read one line from stdin
func Readline() ([]byte, error) {
	reader := bufio.NewReader(os.Stdin)
	data, _, err := reader.ReadLine()

	return data, err
}
