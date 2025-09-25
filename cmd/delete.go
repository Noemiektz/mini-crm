package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/Noemiektz/mini-crm/store"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Supprimer un contact",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")

		st, err := store.GetStore()
		if err != nil {
			fmt.Println("Erreur :", err)
			return
		}

		err = st.Delete(uint(id))
		if err != nil {
			fmt.Println("Impossible de supprimer :", err)
			return
		}

		fmt.Println("Contact supprimé avec succès")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().Int("id", 0, "ID du contact à supprimer")
	deleteCmd.MarkFlagRequired("id")
}
