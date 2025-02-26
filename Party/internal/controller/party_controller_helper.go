package controller

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"party/internal/model"
)

var (
	ErrInvalidRequest = errors.New("invalid request")
	ErrNotAuthorized  = errors.New("not authorized")
)

// validateAddressUpdate validates address updates
func (s *PartyService) validateAddressUpdate(addresses []model.ContactMedium) error {
	for _, addr := range addresses {
		if err := addr.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// validateCharacteristicUpdate validates characteristic updates
func (s *PartyService) validateCharacteristicUpdate(characteristics []model.Characteristic) error {
	for _, char := range characteristics {
		if err := char.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// validateContactMedium validates contact medium data
func (s *PartyService) validateContactMedium(mediums []model.ContactMedium) error {
	seenPreferred := false
	for _, medium := range mediums {
		if medium.Preferred {
			if seenPreferred {
				return errors.New("only one contact medium can be preferred")
			}
			seenPreferred = true
		}

		if medium.MediumType == "" {
			return errors.New("medium type is required")
		}
	}
	return nil
}

// checkPartyExists checks if a party exists and returns it
func (s *PartyService) checkPartyExists(id string) (*model.Individual, error) {
	party, err := s.partyRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, model.ErrPartyNotFound) {
			return nil, ErrInvalidRequest
		}
		return nil, err
	}
	return party, nil
}

// validatePartyUpdate validates party update data
func (s *PartyService) validatePartyUpdate(party *model.Individual) error {
	if err := party.Validate(); err != nil {
		return err
	}

	if err := s.validateAddressUpdate(party.ContactMedium); err != nil {
		return err
	}

	if err := s.validateCharacteristicUpdate(party.Characteristics); err != nil {
		return err
	}

	if err := s.validateContactMedium(party.ContactMedium); err != nil {
		return err
	}

	return nil
}

// preparePartyResponse prepares the party response according to TMF632 format
func (s *PartyService) preparePartyResponse(party *model.Individual) *model.PartyCreateResponse {
	if party == nil {
		return &model.PartyCreateResponse{
			Error: "party not found",
		}
	}

	return &model.PartyCreateResponse{
		Data: party,
	}
}

// generateUniqueID creates a unique ID with the specified prefix
// Format: PPP + YYMMDDHHmmXXXXX (where PPP is the prefix, and XXXXX is random chars)
// This produces a 15-character ID that meets the database requirements
func generateUniqueID(prefix string) string {
	// Initialize random number generator
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Get current timestamp in format YYMMDDHHmm (10 characters)
	timestamp := time.Now().Format("0601021504")

	// Generate random suffix (2 characters)
	const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	suffix := make([]byte, 2)
	for i := range suffix {
		suffix[i] = chars[r.Intn(len(chars))]
	}

	// Ensure prefix is exactly 3 characters
	paddedPrefix := prefix
	if len(paddedPrefix) > 3 {
		paddedPrefix = paddedPrefix[:3]
	} else if len(paddedPrefix) < 3 {
		paddedPrefix = fmt.Sprintf("%-3s", paddedPrefix)
	}

	// Combine parts to create the ID
	id := fmt.Sprintf("%s%s%s", paddedPrefix, timestamp, string(suffix))

	// Ensure the ID is exactly 15 characters
	if len(id) > 15 {
		id = id[:15]
	} else if len(id) < 15 {
		id = fmt.Sprintf("%-15s", id)
	}

	return strings.ToUpper(id)
}
