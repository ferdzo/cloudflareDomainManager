/*
Copyright © 2024 Andrej Ferdzo Mickov <andrej@ferdzo.xyz>
*/
package cmd

import (
	"cloudflareDomainManager/pkg/functions"
	"cloudflareDomainManager/secrets"
	"fmt"
	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "This command shows the whole DNS record for the requested domain",
	Long:  `This command shows the whole DNS record for the requested domain. It shows the ID, type, name, content, TTL and proxied status of the record.`,
	Run: func(cmd *cobra.Command, args []string) {
		secrets := secrets.LoadSecrets()
		data, err := functions.Show(secrets)
		if err != nil {
			panic(err)
		}
		fmt.Println(data)
	},
}

func init() {
	rootCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
