package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mini-crm",
	Short: "Mini-CRM CLI : gérer vos contacts facilement",
	Long:  "Mini-CRM CLI est un gestionnaire de contacts simple avec plusieurs backends de stockage",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
