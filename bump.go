package main

import (
	"fmt"
	"github.com/coreos/go-semver/semver"
	"os"
	"strings"
)

func main() {
	ver := os.Args[1]
	if strings.HasPrefix(ver, "v") {
		ver = strings.Replace(ver, "v", "", -1)
	}

	current, err := semver.NewVersion(ver)
	if err != nil {
		fmt.Println(err.Error())
	}

	switch os.Args[2] {
	case "patch":
		current.BumpPatch()
	case "minor":
		current.BumpMinor()
	case "major":
		current.BumpMajor()
	default:
		fmt.Printf("You must provide patch, minor or major")
	}
	fmt.Printf("v%s", current)
}
