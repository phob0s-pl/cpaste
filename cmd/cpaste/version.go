package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func printVersion(cmd *cobra.Command, args []string) {
	fmt.Println(version)
}
