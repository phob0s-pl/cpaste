package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"

	"github.com/phob0s-pl/cpaste/pastebin"
	"github.com/spf13/cobra"
)

func publish(cmd *cobra.Command, args []string) {
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

	content, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	paste := &pastebin.Paste{
		Title:       flagName,
		Expire:      flagExpiry,
		FormatShort: flagFormat,
	}

	if flagPublic {
		paste.Private = pastebin.PastePublic
	}

	if flagPrivate {
		paste.Private = pastebin.PastePrivate
	}

	pasteURL, err := client.Publish(paste, content)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(pasteURL)
}
