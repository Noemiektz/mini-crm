package store

import "github.com/Noemiektz/mini-crm/models"

type Storer interface {
	Add(contact models.Contact) error
	List() ([]models.Contact, error)
	Delete(id uint) error
	Update(contact models.Contact) error
}
