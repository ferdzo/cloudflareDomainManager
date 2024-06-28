/*
Copyright © 2024 Andrej Ferdzo Mickov <andrej@ferdzo.xyz>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "",
	Short: "A simple command-line app to configure domain DNS records on Cloudflare.",
	Long: `A simple command-line app that can be used to create, remove,
modify, show and everything in-between for DNS records on Cloudflare.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
