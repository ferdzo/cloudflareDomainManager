/*
Copyright © 2024 Andrej Ferdzo Mickov <ferdzo.andrej@gmail.com>
*/
package cmd

import (
	"cloudflareDomainManager/pkg/functions"
	"cloudflareDomainManager/secrets"
	"fmt"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use: "list",
	Short: "The 'list' command allows you to list all DNS records in the zone with their ID's, types, names," +
		"contents, TTL's and proxied status.",
	Long: `This command allows you to list all DNS records in the zone with their ID's, types, names, contents, TTL's and proxied status.
		You can use this information to delete records by providing the record ID as an argument to the delete command.
		You can also use this information to update records by providing the record ID as an argument to the update command.`,
	Run: func(cmd *cobra.Command, args []string) {
		secret := secrets.LoadSecrets()
		err := functions.List(secret)
		if err != nil {
			fmt.Errorf("error fetching records: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

}
