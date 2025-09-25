package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/Noemiektz/mini-crm/models"
	"github.com/Noemiektz/mini-crm/store"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Mettre à jour un contact",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")
		nom, _ := cmd.Flags().GetString("nom")
		email, _ := cmd.Flags().GetString("email")

		st, err := store.GetStore()
		if err != nil {
			fmt.Println("Erreur :", err)
			return
		}

		c := models.Contact{
			ID:    uint(id),
			Nom:   nom,
			Email: email,
		}

		err = st.Update(c)
		if err != nil {
			fmt.Println("Impossible de mettre à jour :", err)
			return
		}

		fmt.Println("Contact mis à jour :", c)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().Int("id", 0, "ID du contact à mettre à jour")
	updateCmd.Flags().String("nom", "", "Nouveau nom")
	updateCmd.Flags().String("email", "", "Nouvel email")
	updateCmd.MarkFlagRequired("id")
}
