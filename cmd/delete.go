/*
Copyright © 2024 Andrej Ferdzo Mickov <andrej@ferdzo.xyz>
*/
package cmd

import (
	"cloudflareDomainManager/pkg/functions"
	"cloudflareDomainManager/secrets"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [record_id]",
	Short: "This command allows you to delete DNS records",
	Long:  `This command allows you to delete DNS records by providing the record ID as an argument.`,
	Run: func(cmd *cobra.Command, args []string) {
		rec_id := args[0]
		secret := secrets.LoadSecrets()
		functions.Delete(secret, rec_id)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

}
