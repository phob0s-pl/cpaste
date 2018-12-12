package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/phob0s-pl/cpaste/pastebin"

	"github.com/spf13/cobra"
)

func session(cmd *cobra.Command, args []string) {
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	path := filepath.Join(currentUser.HomeDir, configPath)
	credentials := &pastebin.Credentials{
		DevKey:   flagDevKey,
		Password: flagPass,
		User:     flagUser,
	}

	if err := pastebin.RequestUserKey(credentials, path); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("user key obtained successfully")
}
