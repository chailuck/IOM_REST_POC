package repository

import (
	"errors"
	"time"

	"internal/model"

	"gorm.io/gorm"
)

type PartyRepository struct {
	db *gorm.DB
}

func NewPartyRepository(db *gorm.DB) *PartyRepository {
	return &PartyRepository{
		db: db,
	}
}

// Create creates a new party record with related entities
func (r *PartyRepository) Create(party *model.Individual) (*model.Individual, error) {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		// Create the main party record
		if err := tx.Create(party).Error; err != nil {
			return err
		}

		// Create related addresses
		for i := range party.Address {
			party.Address[i].PartyID = party.ID
			if err := tx.Create(&party.Address[i]).Error; err != nil {
				return err
			}
		}

		// Create related characteristics
		for i := range party.Characteristic {
			party.Characteristic[i].PartyID = party.ID
			if err := tx.Create(&party.Characteristic[i]).Error; err != nil {
				return err
			}
		}

		// Create related contact mediums
		for i := range party.ContactMedium {
			party.ContactMedium[i].PartyID = party.ID
			if err := tx.Create(&party.ContactMedium[i]).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return party, nil
}

// GetByID retrieves a party by its ID with all related entities
func (r *PartyRepository) GetByID(id string) (*model.Individual, error) {
	var party model.Individual

	err := r.db.
		Preload("Address").
		Preload("Characteristic").
		Preload("ContactMedium").
		Where("paty_row_id = ?", id).
		First(&party).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, model.ErrPartyNotFound
		}
		return nil, err
	}

	return &party, nil
}

// GetByIdentification retrieves a party by ID type and number
func (r *PartyRepository) GetByIdentification(idType, idNumber string) (*model.Individual, error) {
	var party model.Individual

	err := r.db.
		Preload("Address").
		Preload("Characteristic").
		Preload("ContactMedium").
		Where("id_type = ? AND id_numb = ?", idType, idNumber).
		First(&party).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, model.ErrPartyNotFound
		}
		return nil, err
	}

	return &party, nil
}

// Update updates an existing party and its related entities
func (r *PartyRepository) Update(party *model.Individual) (*model.Individual, error) {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		// Update main party record
		party.ModifiedDate = time.Now()
		if err := tx.Save(party).Error; err != nil {
			return err
		}

		// Update addresses
		if err := tx.Where("paty_row_id = ?", party.ID).Delete(&model.PartyAddress{}).Error; err != nil {
			return err
		}
		for i := range party.Address {
			party.Address[i].PartyID = party.ID
			if err := tx.Create(&party.Address[i]).Error; err != nil {
				return err
			}
		}

		// Update characteristics
		if err := tx.Where("paty_row_id = ?", party.ID).Delete(&model.Characteristic{}).Error; err != nil {
			return err
		}
		for i := range party.Characteristic {
			party.Characteristic[i].PartyID = party.ID
			if err := tx.Create(&party.Characteristic[i]).Error; err != nil {
				return err
			}
		}

		// Update contact mediums
		if err := tx.Where("paty_row_id = ?", party.ID).Delete(&model.ContactMedium{}).Error; err != nil {
			return err
		}
		for i := range party.ContactMedium {
			party.ContactMedium[i].PartyID = party.ID
			if err := tx.Create(&party.ContactMedium[i]).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return party, nil
}

// Delete deletes a party and all related entities
func (r *PartyRepository) Delete(id string) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("paty_row_id = ?", id).Delete(&model.PartyAddress{}).Error; err != nil {
			return err
		}
		if err := tx.Where("paty_row_id = ?", id).Delete(&model.Characteristic{}).Error; err != nil {
			return err
		}
		if err := tx.Where("paty_row_id = ?", id).Delete(&model.ContactMedium{}).Error; err != nil {
			return err
		}
		if err := tx.Delete(&model.Individual{}, "paty_row_id = ?", id).Error; err != nil {
			return err
		}
		return nil
	})
}
