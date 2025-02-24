package controller

import (
	"errors"

	"party/internal/model"
)

var (
	ErrInvalidRequest = errors.New("invalid request")
	ErrNotAuthorized  = errors.New("not authorized")
)

// validateAddressUpdate validates address updates
func (s *PartyService) validateAddressUpdate(addresses []model.PartyAddress) error {
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

	if err := s.validateAddressUpdate(party.Address); err != nil {
		return err
	}

	if err := s.validateCharacteristicUpdate(party.Characteristic); err != nil {
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
