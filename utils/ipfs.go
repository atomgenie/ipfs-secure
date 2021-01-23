package utils

import (
	"fmt"

	shell "github.com/ipfs/go-ipfs-api"
)

// NewIPFS New IPFS instance
func NewIPFS() (*shell.Shell, error) {
	ipfs := shell.NewLocalShell()

	if !ipfs.IsUp() {
		fmt.Println("IPFS is not up.\nTry to run ipfs daemon forst")
		return nil, fmt.Errorf("IPFS not up")
	}

	return ipfs, nil
}
