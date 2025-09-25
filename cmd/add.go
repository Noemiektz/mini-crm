package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/Noemiektz/mini-crm/models"
	"github.com/Noemiektz/mini-crm/store"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Ajouter un contact",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")
		nom, _ := cmd.Flags().GetString("nom")
		email, _ := cmd.Flags().GetString("email")

		c := models.Contact{
			ID:    uint(id),
			Nom:   nom,
			Email: email,
		}

		st, _ := store.GetStore() 
		st.Add(c)

		fmt.Println("Contact ajouté :", c)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().Int("id", 0, "ID du contact")
	addCmd.Flags().String("nom", "", "Nom du contact")
	addCmd.Flags().String("email", "", "Email du contact")
	addCmd.MarkFlagRequired("id")
	addCmd.MarkFlagRequired("nom")
	addCmd.MarkFlagRequired("email")
}
