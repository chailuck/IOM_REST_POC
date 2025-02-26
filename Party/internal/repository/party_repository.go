package repository

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"party/internal/model"

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
	// Prepare the data for the database
	model.PrepareIndividualForDB(party)

	err := r.db.Transaction(func(tx *gorm.DB) error {
		// Create the main party record
		if err := tx.Create(party).Error; err != nil {
			return err
		}

		// Create related addresses and address extensions
		for _, contactMedium := range party.ContactMedium {
			if contactMedium.MediumType == "address" {
				// Create address record
				address := model.PartyAddress{
					ID:           contactMedium.ID,
					PartyID:      party.ID,
					IDType:       party.IDType,
					IDNumber:     party.IDNumber,
					AddressType:  contactMedium.AddressType,
					Street1:      contactMedium.Street1,
					Street2:      contactMedium.Street2,
					Country:      contactMedium.Country,
					PostalCode:   contactMedium.PostalCode,
					PostcodeSeq:  contactMedium.PostcodeSeq,
					CreatedDate:  party.CreationDate,
					CreatedBy:    party.CreatedBy,
					ModifiedDate: party.ModificationDate,
					ModifiedBy:   party.ModifiedBy,
				}
				fmt.Printf("Address ID: %s, Object: %p", address.ID, &address)
				if err := tx.Create(&address).Error; err != nil {
					return err
				}

				// Create address extension record
				addressExt := model.PartyAddressExt{
					ID:           contactMedium.ID,
					IDType:       party.IDType,
					IDNumber:     party.IDNumber,
					AddressType:  contactMedium.AddressType,
					BuildingName: contactMedium.Building,
					HomeNo:       contactMedium.HomeNumber,
					Moo:          contactMedium.Moo,
					TumbolName:   contactMedium.Tumbol,
					CityName:     contactMedium.City,
					AccomType:    contactMedium.AccomodationType,
					TimeAtAddr:   contactMedium.TimeAtAddress,
					ZipCode:      contactMedium.PostalCode,
					CreatedDate:  party.CreationDate,
					CreatedBy:    party.CreatedBy,
					ModifiedDate: party.ModificationDate,
					ModifiedBy:   party.ModifiedBy,
				}

				if err := tx.Create(&addressExt).Error; err != nil {
					return err
				}
			}
		}

		// Create related characteristics
		for _, characteristic := range party.Characteristics {
			characteristic.PartyID = party.ID
			characteristic.IDType = party.IDType
			characteristic.IDNumber = party.IDNumber
			characteristic.CreatedDate = party.CreationDate
			characteristic.CreatedBy = party.CreatedBy
			characteristic.ModifiedDate = party.ModificationDate
			characteristic.ModifiedBy = party.ModifiedBy

			if err := tx.Create(&characteristic).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	// Prepare the data for API response
	model.PrepareIndividualForAPI(party)

	return party, nil
}

// GetByID retrieves a party by its ID with all related entities
func (r *PartyRepository) GetByID(id string) (*model.Individual, error) {
	var party model.Individual

	// Get the party record
	err := r.db.Where("paty_row_id = ?", id).First(&party).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, model.ErrPartyNotFound
		}
		return nil, err
	}

	// Get addresses
	var addresses []model.PartyAddress
	if err := r.db.Where("paty_row_id = ?", id).Find(&addresses).Error; err != nil {
		return nil, err
	}

	// Get address extensions
	for _, addr := range addresses {
		var addrExt model.PartyAddressExt
		if err := r.db.Where("addr_row_id = ?", addr.ID).First(&addrExt).Error; err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, err
			}
			// Skip if extension not found
			continue
		}

		// Create ContactMedium from address
		contactMedium := model.ContactMedium{
			ID:               addr.ID,
			MediumType:       "address",
			PartyID:          addr.PartyID,
			Preferred:        false, // Default value
			AddressType:      addr.AddressType,
			Street1:          addr.Street1,
			Street2:          addr.Street2,
			Country:          addr.Country,
			PostalCode:       addr.PostalCode,
			PostcodeSeq:      addr.PostcodeSeq,
			Building:         addrExt.BuildingName,
			HomeNumber:       addrExt.HomeNo,
			Moo:              addrExt.Moo,
			Tumbol:           addrExt.TumbolName,
			City:             addrExt.CityName,
			AccomodationType: addrExt.AccomType,
			TimeAtAddress:    addrExt.TimeAtAddr,
		}

		party.ContactMedium = append(party.ContactMedium, contactMedium)
	}

	// Add phone as ContactMedium if present
	if party.HomePhoneNumber != "" {
		phoneContact := model.ContactMedium{
			ID:          generateUniqueID("CNT"),
			MediumType:  "phone",
			PartyID:     party.ID,
			Preferred:   true, // Assume phone is preferred
			PhoneNumber: party.HomePhoneNumber,
		}
		party.ContactMedium = append(party.ContactMedium, phoneContact)
	}

	// Get characteristics
	var characteristics []model.Characteristic
	if err := r.db.Where("paty_row_id = ?", id).Find(&characteristics).Error; err != nil {
		return nil, err
	}
	party.Characteristics = characteristics

	// Prepare the data for API response
	model.PrepareIndividualForAPI(&party)

	return &party, nil
}

// GetByIdentification retrieves a party by ID type and number
func (r *PartyRepository) GetByIdentification(idType, idNumber string) (*model.Individual, error) {
	var party model.Individual

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
func (r *PartyRepository) Update(party *model.Individual) (*model.Individual, error) {
	// Prepare the data for the database
	model.PrepareIndividualForDB(party)

	err := r.db.Transaction(func(tx *gorm.DB) error {
		// Update main party record
		party.ModificationDate = time.Now()
		if err := tx.Save(party).Error; err != nil {
			return err
		}

		// Delete existing addresses
		if err := tx.Where("paty_row_id = ?", party.ID).Delete(&model.PartyAddress{}).Error; err != nil {
			return err
		}

		// Delete existing address extensions
		if err := tx.Where("id_type = ? AND id_numb = ?", party.IDType, party.IDNumber).Delete(&model.PartyAddressExt{}).Error; err != nil {
			return err
		}

		// Create updated addresses and address extensions
		for _, contactMedium := range party.ContactMedium {
			if contactMedium.MediumType == "address" {
				// Create address record
				address := model.PartyAddress{
					ID:           contactMedium.ID,
					PartyID:      party.ID,
					IDType:       party.IDType,
					IDNumber:     party.IDNumber,
					AddressType:  contactMedium.AddressType,
					Street1:      contactMedium.Street1,
					Street2:      contactMedium.Street2,
					Country:      contactMedium.Country,
					PostalCode:   contactMedium.PostalCode,
					PostcodeSeq:  contactMedium.PostcodeSeq,
					CreatedDate:  party.CreationDate,
					CreatedBy:    party.CreatedBy,
					ModifiedDate: party.ModificationDate,
					ModifiedBy:   party.ModifiedBy,
				}

				if err := tx.Create(&address).Error; err != nil {
					return err
				}

				// Create address extension record
				addressExt := model.PartyAddressExt{
					ID:           contactMedium.ID,
					IDType:       party.IDType,
					IDNumber:     party.IDNumber,
					AddressType:  contactMedium.AddressType,
					BuildingName: contactMedium.Building,
					HomeNo:       contactMedium.HomeNumber,
					Moo:          contactMedium.Moo,
					TumbolName:   contactMedium.Tumbol,
					CityName:     contactMedium.City,
					AccomType:    contactMedium.AccomodationType,
					TimeAtAddr:   contactMedium.TimeAtAddress,
					ZipCode:      contactMedium.PostalCode,
					CreatedDate:  party.CreationDate,
					CreatedBy:    party.CreatedBy,
					ModifiedDate: party.ModificationDate,
					ModifiedBy:   party.ModifiedBy,
				}

				if err := tx.Create(&addressExt).Error; err != nil {
					return err
				}
			}
		}

		// Delete existing characteristics
		if err := tx.Where("paty_row_id = ?", party.ID).Delete(&model.Characteristic{}).Error; err != nil {
			return err
		}

		// Create updated characteristics
		for _, characteristic := range party.Characteristics {
			characteristic.PartyID = party.ID
			characteristic.IDType = party.IDType
			characteristic.IDNumber = party.IDNumber
			characteristic.CreatedDate = party.CreationDate
			characteristic.CreatedBy = party.CreatedBy
			characteristic.ModifiedDate = party.ModificationDate
			characteristic.ModifiedBy = party.ModifiedBy

			if err := tx.Create(&characteristic).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	// Prepare the data for API response
	model.PrepareIndividualForAPI(party)

	return party, nil
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

		// Delete characteristics
		if err := tx.Where("paty_row_id = ?", id).Delete(&model.Characteristic{}).Error; err != nil {
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
func (r *PartyRepository) GetByCharacteristic(name, value string) ([]model.Individual, error) {
	var partyIDs []string

	err := r.db.Model(&model.Characteristic{}).
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
func (r *PartyRepository) AddCharacteristic(characteristic *model.Characteristic) error {
	// Get party to get identification info
	var party model.Individual
	if err := r.db.Where("paty_row_id = ?", characteristic.PartyID).First(&party).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.ErrPartyNotFound
		}
		return err
	}

	// Fill in required fields
	characteristic.IDType = party.IDType
	characteristic.IDNumber = party.IDNumber
	characteristic.CreatedDate = time.Now()
	characteristic.CreatedBy = "system" // Or pass in from context
	characteristic.ModifiedDate = time.Now()
	characteristic.ModifiedBy = "system" // Or pass in from context

	return r.db.Create(characteristic).Error
}

// Generate a unique 15-char ID
func generateUniqueID(prefix string) string {
	// Initialize random number generator
	randBytes := make([]byte, 2)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate random bytes
	for i := range randBytes {
		randBytes[i] = byte(r.Intn(256))
	}

	// Get current timestamp in format YYMMDDHHmm (10 characters)
	timestamp := time.Now().Format("0601021504")

	// Convert random bytes to readable characters
	randStr := strconv.FormatInt(int64(randBytes[0])*256+int64(randBytes[1]), 16)
	if len(randStr) > 2 {
		randStr = randStr[:2]
	} else if len(randStr) < 2 {
		randStr = "0" + randStr
	}

	// Ensure prefix is exactly 3 characters
	paddedPrefix := prefix
	if len(paddedPrefix) > 3 {
		paddedPrefix = paddedPrefix[:3]
	} else if len(paddedPrefix) < 3 {
		paddedPrefix = paddedPrefix + "XXX"[:3-len(paddedPrefix)]
	}

	// Format into 15 character ID
	id := paddedPrefix + timestamp + randStr
	if len(id) > 15 {
		id = id[:15]
	} else if len(id) < 15 {
		id = id + "XXXXXXXXXXXXX"[:15-len(id)]
	}

	return strings.ToUpper(id)
}
