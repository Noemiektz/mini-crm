package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/Noemiektz/mini-crm/store"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lister tous les contacts",
	Run: func(cmd *cobra.Command, args []string) {
		st, err := store.GetStore()
		if err != nil {
			fmt.Println("Erreur :", err)
			return
		}

		contacts, err := st.List()
		if err != nil {
			fmt.Println("Impossible de lister les contacts :", err)
			return
		}

		if len(contacts) == 0 {
			fmt.Println("📭 Aucun contact trouvé")
			return
		}

		fmt.Println("--- Liste des contacts ---")
		for _, c := range contacts {
			fmt.Printf("ID: %d | Nom: %s | Email: %s\n", c.ID, c.Nom, c.Email)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
