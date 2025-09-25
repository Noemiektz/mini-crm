package store

import (
	"fmt"

	"github.com/Noemiektz/mini-crm/models"
	"github.com/spf13/viper"
)

type Storer interface {
	Add(contact models.Contact) error
	List() ([]models.Contact, error)
	Delete(id uint) error
	Update(contact models.Contact) error
}

func GetStore() (Storer, error) {
	storage := viper.GetString("storage")

	switch storage {
	case "gorm":
		return NewGORMStore()
	case "json":
		return NewJSONStore()
	case "memory":
		return NewMemoryStore()
	default:
		return nil, fmt.Errorf("storage inconnu : %s", storage)
	}
}

func init() {
	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Erreur lecture config :", err)
	}
}