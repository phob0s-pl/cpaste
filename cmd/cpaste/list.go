package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"time"

	"github.com/phob0s-pl/cpaste/pastebin"
	"github.com/spf13/cobra"
)

func list(cmd *cobra.Command, args []string) {
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

	list, err := client.List(flagListLimit)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("----------------------------------------------------------------------------------")
	fmt.Println("|           Paste URL          |  time left   | scope |  format  |  title        |")
	fmt.Println("----------------------------------------------------------------------------------")

	for i := range list {
		left := time.Unix(list[i].ExpireDate, 0).Sub(time.Unix(list[i].Date, 0)).String()
		if list[i].ExpireDate == 0 {
			left = "inf"
		}
		fmt.Printf("%-32s  %-12s %-8s %-8s  \"%s\"\n",
			list[i].URL,
			left,
			privateToString(list[i].Private),
			fmt.Sprintf("[%s]", list[i].FormatShort),
			list[i].Title,
		)
	}
}

func privateToString(private uint8) string {
	switch private {
	case pastebin.PastePublic:
		return "public"
	case pastebin.PasteUnlisted:
		return "unlisted"
	case pastebin.PastePrivate:
		return "private"
	default:
		return "unknown"
	}
}
