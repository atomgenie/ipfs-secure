package utils

import (
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

// ReadPassword Get password from terminal
func ReadPassword() ([]byte, error) {
	return terminal.ReadPassword(syscall.Stdin)
}
