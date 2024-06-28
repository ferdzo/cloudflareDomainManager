/*
Copyright © 2024 Andrej Ferdzo Mickov <andrej@ferdzo.xyz>
*/
package cmd

import (
	"cloudflareDomainManager/pkg/functions"
	"cloudflareDomainManager/secrets"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

var createCmd = &cobra.Command{
	Use:   "create --type [argument] --name [argument] --content [argument] --ttl [argument] --proxied [argument]",
	Short: "The 'create' command is to create a new DNS record",
	Long:  `This 'create' command is to create a new DNS record by providing the type, name, content, ttl and proxied status of the record.`,

	Run: func(cmd *cobra.Command, args []string) {
		secrets := secrets.LoadSecrets()
		rec := new(functions.Record)
		rec.Type, _ = cmd.Flags().GetString("type")
		rec.Name, _ = cmd.Flags().GetString("name")
		rec.Content, _ = cmd.Flags().GetString("content")
		TTLstr, _ := cmd.Flags().GetString("ttl")
		rec.TTL, _ = strconv.Atoi(TTLstr)
		proxied, _ := cmd.Flags().GetString("proxied")
		rec.Proxied = proxied == "true"
		if proxied != "true" && proxied != "false" {
			log.Fatal("Wrong Proxied statement")
		}

		functions.Create(secrets, rec)

	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().String("type", "", "Type of the record to create")
	createCmd.Flags().String("name", "", "Name of the record to create")
	createCmd.Flags().String("content", "", "Content of the record to create")
	createCmd.Flags().String("ttl", "", "Time to live of the record to create")
	createCmd.Flags().String("proxied", "", "Is the record proxied")

	createCmd.MarkFlagRequired("type")
	createCmd.MarkFlagRequired("name")
	createCmd.MarkFlagRequired("content")
	createCmd.MarkFlagRequired("ttl")
	createCmd.MarkFlagRequired("proxied")

}
