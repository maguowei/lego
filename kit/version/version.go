package version

import (
	"fmt"
	"os"
)

var (
	GitCommit = ""
	BuildTime = ""
	GoVersion = ""
)

func ShowVersion() {
	args := os.Args
	if len(args) == 2 && (args[1] == "--version" || args[1] == "-v") {
		fmt.Printf("Git Commit: %s\n", GitCommit)
		fmt.Printf("Build Time : %s\n", BuildTime)
		fmt.Printf("Go Version : %s\n", GoVersion)
		os.Exit(0)
	}
}
