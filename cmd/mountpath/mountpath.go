package main

import (
	"fmt"
	"os"

	"github.com/gordonmleigh/mountpath"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Fprintf(os.Stderr, "usage: mountpath <dev> [mounts file]\n")
		os.Exit(2)
	}

	var path string
	var err error

	if len(os.Args) == 3 {
		path, err = mountpath.GetMountPathEx(os.Args[1], os.Args[2])
	} else {
		path, err = mountpath.GetMountPath(os.Args[1])
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "fatal error: %v\n", err)
		os.Exit(1)
	}

	fmt.Print(path)
}
