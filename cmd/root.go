package cmd

import (
	"fmt"
	// "github.com/spf13/cobra"
)

var PackageGlobalString="Fool"

func Print() {
	fmt.Println(PackageGlobalString)
}