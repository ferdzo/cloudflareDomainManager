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
	Short: "The 'show' command shows the whole DNS record for the requested domain",
	Long:  `This 'show' command shows the whole DNS record for the requested domain. It shows the ID, type, name, content, TTL and proxied status of the record.`,
	Run: func(cmd *cobra.Command, args []string) {
		secrets := secrets.LoadSecrets()
		data, err := functions.Show(secrets)
		if err != nil {
			fmt.Errorf("error fetching records: %v", err)
		}
		fmt.Println(data)
	},
}

func init() {
	rootCmd.AddCommand(showCmd)

}
