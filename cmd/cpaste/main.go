package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	configPath = ".config/cpaste.json"
	version    = "0.2.0"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "cpaste",
		Short: "cpaste is command line pastebin.com client",
		Long:  "cpaste is command line pastebin.com client",
		Run:   nil,
	}

	var publishCmd = &cobra.Command{
		Use:   "publish",
		Short: "publish paste on the server",
		Long:  "publish paste on the server",
		Run:   publish,
	}

	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "list published pastes",
		Long:  "list published pastes",
		Run:   list,
	}

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "print version of the application",
		Long:  "print version of the application",
		Run:   printVersion,
	}

	var deleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "delete paste with give paste key from server",
		Long:  "delete paste with give paste key from server",
		Run:   deletePaste,
	}
	deleteCmd.Flags().StringVarP(&flagPasteKey, "paste-key", "", "",
		"paste key to delete")

	var sessionCmd = &cobra.Command{
		Use:   "session",
		Short: "acquire session key from pastebin.com",
		Long:  "acquire session key from pastebin.com",
		Run:   session,
	}
	sessionCmd.Flags().StringVarP(&flagDevKey, "devkey", "", "",
		"dev key value")
	sessionCmd.Flags().StringVarP(&flagUser, "user", "", "",
		"user name")
	sessionCmd.Flags().StringVarP(&flagPass, "pass", "", "",
		"password")
	sessionCmd.MarkFlagRequired("devkey") // nolint: gosec
	sessionCmd.MarkFlagRequired("user")   // nolint: gosec
	sessionCmd.MarkFlagRequired("pass")   // nolint: gosec

	var fileCmd = &cobra.Command{
		Use:   "file",
		Short: "upload file to pastebin.com",
		Long:  "upload file to pastebin.com",
		Run:   file,
	}
	fileCmd.Flags().StringVarP(&flagFormat, "format", "f", "",
		"format of the paste, \"\" to check by file extension")
	fileCmd.Flags().BoolVarP(&flagPrivate, "private", "", false,
		"make paste private")
	fileCmd.Flags().BoolVarP(&flagPublic, "public", "", false,
		"make paste public")
	fileCmd.Flags().StringVarP(&flagName, "name", "n", "",
		"title of the paste, \"\" to use file name")
	fileCmd.Flags().StringVarP(&flagExpiry, "expiration", "e", "1M",
		"expiration date")
	fileCmd.Flags().StringVarP(&flagFilePath, "path", "p", "",
		"path of the file to upload")
	fileCmd.MarkFlagRequired("path") // nolint: gosec

	publishCmd.Flags().StringVarP(&flagFormat, "format", "f", "text",
		"format of the paste")
	publishCmd.Flags().BoolVarP(&flagPrivate, "private", "", false,
		"make paste private")
	publishCmd.Flags().BoolVarP(&flagPublic, "public", "", false,
		"make paste public")
	publishCmd.Flags().StringVarP(&flagName, "name", "n", "",
		"title of the paste")
	publishCmd.Flags().StringVarP(&flagExpiry, "expiration", "e", "1M",
		"expiration date")

	listCmd.Flags().IntVarP(&flagListLimit, "limit", "l", 50,
		"limit of results")

	rootCmd.AddCommand(publishCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(fileCmd)
	rootCmd.AddCommand(sessionCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(deleteCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
