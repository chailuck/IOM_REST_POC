package controller

import (
	"errors"
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
	party.CreatedDate = time.Now()
	party.ModifiedDate = time.Now()

	// Set default values for addresses
	for i := range party.Address {
		if party.Address[i].ID == "" {
			party.Address[i].ID = generateUniqueID("ADR")
		}
		party.Address[i].CreatedDate = time.Now()
		party.Address[i].ModifiedDate = time.Now()
	}

	// Set default values for characteristics
	for i := range party.Characteristic {
		if party.Characteristic[i].ID == "" {
			party.Characteristic[i].ID = generateUniqueID("ATR")
		}
		party.Characteristic[i].CreatedDate = time.Now()
		party.Characteristic[i].ModifiedDate = time.Now()
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
	party.CreatedDate = existing.CreatedDate
	party.ModifiedDate = time.Now()

	// Update related entities
	for i := range party.Address {
		if party.Address[i].ID == "" {
			party.Address[i].ID = generateUniqueID("ADR")
		}
		party.Address[i].ModifiedDate = time.Now()
	}

	for i := range party.Characteristic {
		if party.Characteristic[i].ID == "" {
			party.Characteristic[i].ID = generateUniqueID("ATR")
		}
		party.Characteristic[i].ModifiedDate = time.Now()
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

// AddCharacteristic adds a new characteristic to a party
func (s *PartyService) AddCharacteristic(partyID string, characteristic *model.Characteristic) error {
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
	characteristic.CreatedDate = time.Now()
	characteristic.ModifiedDate = time.Now()

	return s.partyRepo.AddCharacteristic(characteristic)
}

// SearchParties searches for parties using various criteria
func (s *PartyService) SearchParties(params model.PartyQueryParams) ([]model.Individual, error) {
	if params.Limit == 0 {
		params.Limit = 50 // Default limit
	}

	return s.partyRepo.List(params)
}
