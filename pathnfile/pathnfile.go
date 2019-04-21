// pathnfile is a simple program splits filename from paths
// It can split the output for one or multiple paths
// It will display filename and path, path or just filename based on given option

package main

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func main() {
	args := os.Args

	if len(args) < 2 || strings.HasPrefix(args[1], "-") && args[1] != "-f" && args[1] != "-d" {
		fmt.Println("invalid arguments or options")
		os.Exit(1)
	} else {
		printSplit(args)
	}
}

// printSplit prints the results of path.Split()
// -f: will print only the filenames
// -d: will print only de directories
// default: will print everything
func printSplit(a []string) {
	var df []string

	if a[1] == "-f" || a[1] == "-d" {
		df = a[2:]
	} else {
		df = a[1:]
	}

	for _, s := range df {
		d, f := path.Split(s)

		switch {
		case a[1] == "-f":
			fmt.Println(f)
		case a[1] == "-d":
			fmt.Println(d)
		default:
			fmt.Println(d, f)
		}
	}
}
