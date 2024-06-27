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

// createCmd represents the create command

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "This command is to create a new DNS record",
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
		//fmt.Printf("Creating item of type '%s' with name '%s' and  content '%s'\n", rec.rtype, rec.name, rec.content)

	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().String("type", "", "Type of the record to create")
	createCmd.Flags().String("name", "", "Name of the record to create")
	createCmd.Flags().String("content", "", "Content of the record to create")
	createCmd.Flags().String("ttl", "", "Time to live of the record to create")
	createCmd.Flags().String("proxied", "", "Is the record proxied")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
