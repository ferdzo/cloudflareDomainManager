/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cloudflareDomainManager/pkg/functions"
	"cloudflareDomainManager/secrets"
	"github.com/spf13/cobra"
	"strconv"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update --rec_id [argument] --type [argument] --name [argument] --content [argument] --ttl [argument] ",
	Short: "The 'update' command lets you update a DNS record",
	Long:  `This command allows you to update a DNS record by providing the record ID, type, name, content and TTL as arguments.`,
	Run: func(cmd *cobra.Command, args []string) {
		secret := secrets.LoadSecrets()
		record := new(functions.Record)
		recId, _ := cmd.Flags().GetString("rec_id")
		record.Type, _ = cmd.Flags().GetString("type")
		record.Name, _ = cmd.Flags().GetString("name")
		record.Content, _ = cmd.Flags().GetString("content")
		TTLstr, _ := cmd.Flags().GetString("ttl")
		record.TTL, _ = strconv.Atoi(TTLstr)
		functions.Update(secret, recId, *record)

	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().String("type", "", "Type of the record to update")
	updateCmd.Flags().String("name", "", "Name of the record to update")
	updateCmd.Flags().String("content", "", "Content of the record to update")
	updateCmd.Flags().String("ttl", "", "Time to live of the record to update")
	updateCmd.Flags().String("rec_id", "", "Record ID to update")

	updateCmd.MarkFlagRequired("type")
	updateCmd.MarkFlagRequired("name")
	updateCmd.MarkFlagRequired("content")
	updateCmd.MarkFlagRequired("ttl")
	updateCmd.MarkFlagRequired("rec_id")

}
