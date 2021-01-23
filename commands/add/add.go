package add

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/atomgenie/ipfs-secure/utils"
	ipfsshell "github.com/ipfs/go-ipfs-api"
)

func getFile(filename string) ([]byte, error) {
	out, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	return out, nil
}

// Add Add command
func Add(argv []string) error {

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	// reccursive := addCmd.Bool("r", false, "Reccursive")
	pin := flag.Bool("pin", true, "Pin this object when adding. Default: true.")
	addCmd.Parse(argv)

	subargs := addCmd.Args()

	if len(subargs) != 1 {
		addCmd.Usage()
		return fmt.Errorf("Invalid number of args")
	}

	shell, err := utils.NewIPFS()

	if err != nil {
		return err
	}

	fmt.Println("Password:")
	reader := bufio.NewReader(os.Stdin)
	password, _, err := reader.ReadLine()

	if err != nil {
		return err
	}

	fileData, err := getFile(subargs[0])

	if err != nil {
		return err
	}

	hash := utils.HashSha256(password)
	ciphered, err := utils.EncodeAES(fileData, hash)

	if err != nil {
		return err
	}

	cid, err := shell.Add(bytes.NewReader(ciphered), ipfsshell.Pin(*pin))

	if err != nil {
		return err
	}

	println(cid)

	return nil

}
