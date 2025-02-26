package controller

import (
	"errors"
	"fmt"
	"time"

	"party/internal/model"
	"party/internal/repository"
)

type PartyService struct {
	partyRepo *repository.PartyRepository
}

func NewPartyService(partyRepo *repository.PartyRepository) *PartyService {
	return &PartyService{
		partyRepo: partyRepo,
	}
}

// CreateParty creates a new party with all related entities
func (s *PartyService) CreateParty(party *model.Individual) (*model.Individual, error) {
	// Validate party data
	if err := party.Validate(); err != nil {
		return nil, err
	}

	// Check for existing party with same identification
	existing, err := s.partyRepo.GetByIdentification(party.IDType, party.IDNumber)
	if err != nil && !errors.Is(err, model.ErrPartyNotFound) {
		return nil, err
	}
	if existing != nil {
		return nil, model.ErrDuplicateParty
	}

	// Generate a unique 15-char ID for party
	party.ID = generateUniqueID("PTY")
	party.CreationDate = time.Now()
	party.ModificationDate = time.Now()
	party.CreatedBy = "system"  // Or pass in from context
	party.ModifiedBy = "system" // Or pass in from context

	// Set default values for contact media (addresses)
	for i := range party.ContactMedium {
		if party.ContactMedium[i].ID == "" {
			party.ContactMedium[i].ID = generateUniqueID("ADR")
			fmt.Printf("ADDRESS UNIQUE:%s TYPE: %s (%s)\n", party.ContactMedium[i].ID, party.ContactMedium[i].MediumType, party.ContactMedium[i].Street1)
			//party.ContactMedium[i].ID = generateUniqueID("ADR") + "_" + party.ContactMedium[i].MediumType
		}
	}

	// Set default values for characteristics
	for i := range party.Characteristics {
		if party.Characteristics[i].ID == "" {
			party.Characteristics[i].ID = generateUniqueID("ATR")
			fmt.Printf("CHARACTERISTIC UNIQUE:%s TYPE: %s\n", party.Characteristics[i].ID, party.Characteristics[i].Name)
			//party.Characteristics[i].ID = generateUniqueID("ATR") + "_" + party.Characteristics[i].Name
		}
	}

	return s.partyRepo.Create(party)
}

// GetParty retrieves a party by ID
func (s *PartyService) GetParty(id string) (*model.Individual, error) {
	return s.partyRepo.GetByID(id)
}

// UpdateParty updates an existing party
func (s *PartyService) UpdateParty(id string, party *model.Individual) (*model.Individual, error) {
	// Validate party data
	if err := party.Validate(); err != nil {
		return nil, err
	}

	// Check if party exists
	existing, err := s.partyRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Update fields while preserving immutable data
	party.ID = existing.ID
	party.CreationDate = existing.CreationDate
	party.CreatedBy = existing.CreatedBy
	party.ModificationDate = time.Now()
	party.ModifiedBy = "system" // Or pass in from context

	// Ensure proper IDs for related entities
	for i := range party.ContactMedium {
		if party.ContactMedium[i].ID == "" {
			party.ContactMedium[i].ID = generateUniqueID("ADR")
		}
	}

	for i := range party.Characteristics {
		if party.Characteristics[i].ID == "" {
			party.Characteristics[i].ID = generateUniqueID("ATR")
		}
	}

	return s.partyRepo.Update(party)
}

// DeleteParty deletes a party by ID
func (s *PartyService) DeleteParty(id string) error {
	// Check if party exists
	_, err := s.partyRepo.GetByID(id)
	if err != nil {
		return err
	}

	return s.partyRepo.Delete(id)
}

// ListParties retrieves parties based on query parameters
func (s *PartyService) ListParties(params model.PartyQueryParams) ([]model.Individual, error) {
	return s.partyRepo.List(params)
}

// UpdatePartyStatus updates a party's status
func (s *PartyService) UpdatePartyStatus(id string, status string) error {
	// Validate status
	if status == "" {
		return errors.New("status cannot be empty")
	}

	// Check if party exists
	_, err := s.partyRepo.GetByID(id)
	if err != nil {
		return err
	}

	return s.partyRepo.UpdateStatus(id, status)
}

// AddPartyCharacteristic adds a new characteristic to a party
func (s *PartyService) AddPartyCharacteristic(partyID string, characteristic *model.Characteristic) error {
	// Validate characteristic
	if err := characteristic.Validate(); err != nil {
		return err
	}

	// Check if party exists
	_, err := s.partyRepo.GetByID(partyID)
	if err != nil {
		return err
	}

	characteristic.PartyID = partyID
	characteristic.ID = generateUniqueID("ATR")

	return s.partyRepo.AddCharacteristic(characteristic)
}

// SearchParties searches for parties using various criteria
func (s *PartyService) SearchParties(criteria map[string]interface{}) ([]model.Individual, error) {
	return s.partyRepo.SearchParties(criteria)
}

// GetPartyByPhoneNumber finds parties by phone number
func (s *PartyService) GetPartyByPhoneNumber(phoneNumber string) ([]model.Individual, error) {
	if phoneNumber == "" {
		return nil, errors.New("phone number is required")
	}

	return s.partyRepo.GetPartyByPhoneNumber(phoneNumber)
}

// GetPartyByAddress finds parties by address criteria
func (s *PartyService) GetPartyByAddress(postalCode, district, subDistrict string) ([]model.Individual, error) {
	return s.partyRepo.GetPartyByAddress(postalCode, district, subDistrict)
}

// GetPartyByExternalID finds a party by external ID
func (s *PartyService) GetPartyByExternalID(extIDType, extID string) (*model.Individual, error) {
	if extID == "" {
		return nil, errors.New("external ID is required")
	}

	return s.partyRepo.GetPartyByExternalID(extIDType, extID)
}

// CountPartiesByType gets a count of parties by type
func (s *PartyService) CountPartiesByType(partyType string) (int64, error) {
	if partyType == "" {
		return 0, errors.New("party type is required")
	}

	return s.partyRepo.CountPartiesByType(partyType)
}
