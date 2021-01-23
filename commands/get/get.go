package get

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/atomgenie/ipfs-secure/utils"
)

func writefile(filename string, data []byte) error {
	return ioutil.WriteFile(filename, data, 0644)
}

// Get Get command
func Get(argv []string) error {
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	outfile := getCmd.String("o", "", "Out file")

	getCmd.Parse(argv)

	subargs := getCmd.Args()

	if len(subargs) != 1 {
		getCmd.Usage()
		return fmt.Errorf("Invalid usage")
	}

	ipfsHash := subargs[0]

	shell, err := utils.NewIPFS()

	if err != nil {
		return err
	}

	fmt.Println("Password:")
	password, err := utils.ReadPassword()

	if err != nil {
		return err
	}

	dataRead, err := shell.Cat(ipfsHash)

	if err != nil {
		return err
	}

	defer dataRead.Close()

	data, err := ioutil.ReadAll(dataRead)

	if err != nil {
		return err
	}

	hash := utils.HashSha256(password)

	deciphered, err := utils.DecodeAES(data, hash)

	if err != nil {
		return err
	}

	finalFile := *outfile

	if finalFile == "" {
		finalFile = ipfsHash
	}

	return writefile(finalFile, deciphered)
}
