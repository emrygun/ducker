package main

import (
	"fmt"
	"os"

	"github.com/emreuygun/ducker/internal/container"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	var err error
	switch os.Args[1] {
	case "run":
		err = container.Run(os.Args[2:])
	case "child":
		err = container.Child(os.Args[2:])
	default:
		usage()
		os.Exit(1)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, "ducker:", err)
		os.Exit(1)
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, "usage:")
	fmt.Fprintln(os.Stderr, "  ducker run <image-rootfs> <command> [args...]")
}
