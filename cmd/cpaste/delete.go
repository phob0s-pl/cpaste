package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/phob0s-pl/cpaste/pastebin"
	"github.com/spf13/cobra"
)

func deletePaste(cmd *cobra.Command, args []string) {
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	path := filepath.Join(currentUser.HomeDir, configPath)
	client, err := pastebin.NewClient(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	output, err := client.Delete(flagPasteKey)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(output)
}
