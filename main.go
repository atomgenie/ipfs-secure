package main

import (
	"fmt"
	"os"

	"github.com/atomgenie/ipfs-secure/commands/add"
	"github.com/atomgenie/ipfs-secure/commands/get"
)

func main() {
	argv := os.Args[1:]

	if len(argv) == 0 {
		return
	}

	command := argv[0]

	var err error = nil
	switch command {
	case "add":
		err = add.Add(argv[1:])
	case "get":
		err = get.Get(argv[1:])
	default:
		fmt.Println("Not a valid command")
	}

	if err != nil {
		fmt.Print("Error: ")
		fmt.Println(err)
	}

}
