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
	Use:   "update",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		secret := secrets.LoadSecrets()
		record := new(functions.Record)
		record.Type, _ = cmd.Flags().GetString("type")
		record.Name, _ = cmd.Flags().GetString("name")
		record.Content, _ = cmd.Flags().GetString("content")
		TTLstr, _ := cmd.Flags().GetString("ttl")
		record.TTL, _ = strconv.Atoi(TTLstr)
		rec_id, _ := cmd.Flags().GetString("rec_id")
		functions.Update(secret, rec_id, *record)

	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
