package store

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "github.com/Noemiektz/mini-crm/models"
)

type GORMStore struct {
	db *gorm.DB
}

func NewGORMStore() (*GORMStore, error) {
	db, err := gorm.Open(sqlite.Open("contacts.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// AutoMigrate crée la table si elle n’existe pas
	db.AutoMigrate(&models.Contact{})

	return &GORMStore{db: db}, nil
}

func (s *GORMStore) Add(contact models.Contact) error {
	return s.db.Create(&contact).Error
}

func (s *GORMStore) List() ([]models.Contact, error) {
	var contacts []models.Contact
	err := s.db.Find(&contacts).Error
	return contacts, err
}

func (s *GORMStore) Delete(id uint) error {
	return s.db.Delete(&models.Contact{}, id).Error
}

func (s *GORMStore) Update(contact models.Contact) error {
	return s.db.Save(&contact).Error
}
