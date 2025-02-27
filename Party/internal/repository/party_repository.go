package repository

import (
	"errors"
	"fmt"
	"party/internal/model"
	"time"

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
func (r *PartyRepository) Create(individual *model.Individual) (*model.Individual, error) {
	// Prepare the data for the database
	//model.PrepareIndividualForDB(individual)

	err := r.db.Transaction(func(tx *gorm.DB) error {
		party := model.MapIndividualToParty(individual)
		if err := tx.Create(party).Error; err != nil {
			return err
		}
		addrList, addrExtList := model.MapIndividualToPartyAddress(individual)
		fmt.Printf("Address: %v \nExtension: %v\n", addrList, addrExtList)
		for _, addr := range addrList {
			if err := tx.Create(&addr).Error; err != nil {
				return err
			}
		}
		for _, addrExt := range addrExtList {
			if err := tx.Create(&addrExt).Error; err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	// Prepare the data for API response
	// model.PrepareIndividualForAPI(individual)

	return individual, nil
}

// GetByID retrieves a party by its ID with all related entities
func (r *PartyRepository) GetByID(id string) (*model.Individual, error) {
	var party model.Party

	// Get the party record
	err := r.db.Where("paty_row_id = ?", id).First(&party).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, model.ErrPartyNotFound
		}
		return nil, err
	}

	// Get addresses
	var addrList []model.PartyAddress
	if err := r.db.Where("paty_row_id = ?", id).Find(&addrList).Error; err != nil {
		return nil, err
	}

	// Get address extensions
	var addrExtList []model.PartyAddressExt

	for _, addr := range addrList {
		var addrExt model.PartyAddressExt
		if err := r.db.Where("addr_row_id = ?", addr.ID).First(&addrExt).Error; err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, err
			}
			// Skip if extension not found
			continue
		}
		addrExtList = append(addrExtList, addrExt)
	}
	// Get characteristics
	var attrList []model.PartyAttributes
	if err := r.db.Where("paty_row_id = ?", id).Find(&attrList).Error; err != nil {
		return nil, err
	}

	individual := model.MapPartyToIndividual(party, addrList, addrExtList, attrList)

	// Prepare the data for API response
	//model.PrepareIndividualForAPI(&party)

	return &individual, nil
}

// GetByIdentification retrieves a party by ID type and number
func (r *PartyRepository) GetByIdentification(idType, idNumber string) (*model.Individual, error) {
	var party model.Party
	fmt.Printf("ID Type: %s, ID Number: %s\n", idType, idNumber)
	err := r.db.Where("id_type = ? AND id_numb = ?", idType, idNumber).First(&party).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, model.ErrPartyNotFound
		}
		return nil, err
	}

	// Get related entities
	return r.GetByID(party.ID)
}

// Update updates an existing party and its related entities
func (r *PartyRepository) Update(individual *model.Individual) (*model.Individual, error) {
	// Prepare the data for the database
	//model.PrepareIndividualForDB(party)

	err := r.db.Transaction(func(tx *gorm.DB) error {
		party := model.MapIndividualToParty(individual)

		// Update main party record

		party.ModificationDate = time.Now()
		if err := tx.Save(party).Error; err != nil {
			return err
		}

		addrList, addrExtList := model.MapIndividualToPartyAddress(individual)
		for _, addr := range addrList {
			var queryAddr []model.PartyAddress
			err := r.db.Where("paty_row_id = ? and addr_type = ?", party.ID, addr.AddressType).Find(&queryAddr).Error
			if err != nil {
				if len(queryAddr) > 1 {
					return model.ErrMoreThanOneEntityFound
				}
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return model.ErrPartyNotFound
				}
				return err
			}
			addr_id := queryAddr[0].ID

			if err := tx.Where("paty_row_id = ? and addr_id and addr_type = ?", party.ID, addr_id, addr.AddressType).Save(addr).Error; err != nil {
				return err
			}
			for _, addrExt := range addrExtList {
				if addrExt.AddressType == addr.AddressType {
					if err := tx.Where("addr_id = ? and addr_type = ?", addrExt.ID, addrExt.AddressType).Save(addrExt).Error; err != nil {
						return err
					}
				}
			}
		}
		return nil
	}) //commit tx

	if err != nil {
		return nil, err
	}

	return individual, nil
}

// Delete deletes a party and all related entities
func (r *PartyRepository) Delete(id string) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Get party to get identification info
		var party model.Individual
		if err := tx.Where("paty_row_id = ?", id).First(&party).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return model.ErrPartyNotFound
			}
			return err
		}

		// Get address IDs to delete address extensions
		var addressIDs []string
		if err := tx.Model(&model.PartyAddress{}).Where("paty_row_id = ?", id).Pluck("addr_row_id", &addressIDs).Error; err != nil {
			return err
		}

		// Delete address extensions
		for _, addrID := range addressIDs {
			if err := tx.Where("addr_row_id = ?", addrID).Delete(&model.PartyAddressExt{}).Error; err != nil {
				return err
			}
		}

		// Delete addresses
		if err := tx.Where("paty_row_id = ?", id).Delete(&model.PartyAddress{}).Error; err != nil {
			return err
		}

		// Delete attributes
		if err := tx.Where("paty_row_id = ?", id).Delete(&model.PartyAttributes{}).Error; err != nil {
			return err
		}

		// Delete party
		if err := tx.Delete(&model.Individual{}, "paty_row_id = ?", id).Error; err != nil {
			return err
		}

		return nil
	})
}

// List retrieves parties based on query parameters
func (r *PartyRepository) List(params model.PartyQueryParams) ([]model.Individual, error) {
	var parties []model.Individual
	query := r.db.Model(&model.Individual{})

	// Apply filters
	if params.IDType != "" {
		query = query.Where("id_type = ?", params.IDType)
	}
	if params.IDNumber != "" {
		query = query.Where("id_numb = ?", params.IDNumber)
	}
	if params.FirstName != "" {
		query = query.Where("frst_name ILIKE ?", "%"+params.FirstName+"%")
	}
	if params.LastName != "" {
		query = query.Where("last_name ILIKE ?", "%"+params.LastName+"%")
	}
	if params.Status != "" {
		query = query.Where("paty_type = ?", params.Status)
	}

	// Apply pagination
	if params.Limit > 0 {
		query = query.Limit(params.Limit)
	}
	if params.Offset > 0 {
		query = query.Offset(params.Offset)
	}

	// Execute query
	err := query.Find(&parties).Error
	if err != nil {
		return nil, err
	}

	// Get related entities for each party
	var result []model.Individual
	for _, party := range parties {
		fullParty, err := r.GetByID(party.ID)
		if err != nil {
			continue // Skip parties with errors
		}
		result = append(result, *fullParty)
	}

	return result, nil
}

// GetByCharacteristic retrieves parties by characteristic name and value
func (r *PartyRepository) GetByAttributes(name, value string) ([]model.Individual, error) {
	var partyIDs []string

	err := r.db.Model(&model.PartyAttributes{}).
		Where("attr_name = ? AND attr_vlue = ?", name, value).
		Pluck("paty_row_id", &partyIDs).Error
	if err != nil {
		return nil, err
	}

	var result []model.Individual
	for _, id := range partyIDs {
		party, err := r.GetByID(id)
		if err != nil {
			continue // Skip parties with errors
		}
		result = append(result, *party)
	}

	return result, nil
}

// UpdateStatus updates a party's status
func (r *PartyRepository) UpdateStatus(id string, status string) error {
	return r.db.Model(&model.Individual{}).
		Where("paty_row_id = ?", id).
		Update("paty_type", status).Error
}

// AddCharacteristic adds a new characteristic to a party
func (r *PartyRepository) AddAttributes(PartyID string, characteristic *model.CharacteristicItem) error {
	// Get party to get identification info
	var party model.Party
	if err := r.db.Where("paty_row_id = ?", PartyID).First(&party).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.ErrPartyNotFound
		}
		return err
	}

	attr := model.MapCharacteristicToPartyAttribute(party.ID, party.IDType, party.IDNumber, characteristic)

	return r.db.Create(attr).Error
}
