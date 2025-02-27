package controller

import (
	"errors"
	"fmt"
	"strings"
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
func (s *PartyService) CreateIndividual(individual *model.Individual) (*model.Individual, error) {
	// Validate party data
	if err := individual.Validate(); err != nil {
		return nil, err
	}

	// Check for existing party with same identification
	existing, err := s.partyRepo.GetByIdentification(individual.IDType, individual.IDNumber)
	if err != nil && !errors.Is(err, model.ErrPartyNotFound) {
		return nil, err
	}
	if existing != nil {
		return nil, model.ErrDuplicateParty
	}

	// Generate a unique 15-char ID for party
	individual.ID = generateUniqueID("PTY")
	individual.CreationDate = time.Now()
	individual.ModificationDate = time.Now()
	createdBy := "system" // Or pass in from context
	if strings.TrimSpace(individual.CreatedBy) != "" {
		createdBy = strings.TrimSpace(individual.CreatedBy)
	}
	modifiedBy := "system" // Or pass in from context
	if strings.TrimSpace(individual.ModifiedBy) != "" {
		individual.ModifiedBy = strings.TrimSpace(individual.ModifiedBy)
	}
	individual.Status = "P"            //Prospect
	individual.CreatedBy = createdBy   // Or pass in from context
	individual.ModifiedBy = modifiedBy // Or pass in from context

	// Set default values for contact media (addresses)
	for i := range individual.ContactMedium {
		if individual.ContactMedium[i].ID == "" {
			if (individual.ContactMedium[i].Type == model.EntityTypeContactMediumPhone) && (individual.ContactMedium[i].MediumType == model.EntityTypeContactMediumTypeHomePhoneNumber) {
				individual.ContactMedium[i].ID = "HOMEPHONE_" + individual.ID
			} else {
				individual.ContactMedium[i].ID = generateUniqueID("ADR")
			}
			fmt.Printf("ADDRESS UNIQUE:%s TYPE: %s (%s)\n", individual.ContactMedium[i].ID, individual.ContactMedium[i].MediumType, individual.ContactMedium[i].Street1)
		}
		individual.ContactMedium[i].AuditTrail.CreationDate = time.Now()
		individual.ContactMedium[i].AuditTrail.CreatedBy = createdBy
		individual.ContactMedium[i].AuditTrail.ModificationDate = time.Now()
		individual.ContactMedium[i].AuditTrail.ModifiedBy = modifiedBy

	}

	// Set default values for characteristics
	for i := range individual.Characteristic {
		if individual.Characteristic[i].ID == "" {
			individual.Characteristic[i].ID = generateUniqueID("ATR")
			individual.Characteristic[i].AuditTrail.CreationDate = time.Now()
			individual.Characteristic[i].AuditTrail.CreatedBy = createdBy
			individual.Characteristic[i].AuditTrail.ModificationDate = time.Now()
			individual.Characteristic[i].AuditTrail.ModifiedBy = modifiedBy
			fmt.Printf("CHARACTERISTIC UNIQUE:%s TYPE: %s\n", individual.Characteristic[i].ID, individual.Characteristic[i].Name)
		}
	}

	return s.partyRepo.Create(individual)
}

// GetParty retrieves a party by ID
func (s *PartyService) GetParty(id string) (*model.Individual, error) {
	return s.partyRepo.GetByID(id)
}

// UpdateParty updates an existing party
func (s *PartyService) UpdateIndividual(id string, individual *model.Individual) (*model.Individual, error) {
	// Validate party data
	if err := individual.Validate(); err != nil {
		return nil, err
	}

	// Check if party exists
	existing, err := s.partyRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Update fields while preserving immutable data
	individual.ID = existing.ID
	individual.CreationDate = existing.CreationDate
	individual.CreatedBy = existing.CreatedBy
	individual.ModificationDate = time.Now()
	individual.ModifiedBy = "system" // Or pass in from context

	// Ensure proper IDs for related entities
	for i := range individual.ContactMedium {
		if individual.ContactMedium[i].ID == "" {
			individual.ContactMedium[i].ID = generateUniqueID("ADR")
		}
	}

	for i := range individual.Characteristic {
		if individual.Characteristic[i].ID == "" {
			individual.Characteristic[i].ID = generateUniqueID("ATR")
		}
	}

	return s.partyRepo.Update(individual)
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
func (s *PartyService) AddPartyCharacteristic(partyID string, characteristic *model.CharacteristicItem) error {
	// Validate characteristic
	if err := characteristic.Validate(); err != nil {
		return err
	}

	// Check if party exists
	_, err := s.partyRepo.GetByID(partyID)
	if err != nil {
		return err
	}

	characteristic.ID = generateUniqueID("ATR")

	return s.partyRepo.AddAttributes(partyID, characteristic)
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
