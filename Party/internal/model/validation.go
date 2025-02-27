package model

import (
	"errors"
	"strings"
)

var (
	ErrInvalidParty           = errors.New("invalid party data")
	ErrInvalidIDType          = errors.New("invalid ID type")
	ErrInvalidIDNumber        = errors.New("invalid ID number")
	ErrInvalidGender          = errors.New("invalid gender")
	ErrInvalidName            = errors.New("given name is required")
	ErrPartyNotFound          = errors.New("party not found")
	ErrDuplicateParty         = errors.New("party already exists")
	ErrMoreThanOneEntityFound = errors.New("more than 1 entitity found")
)

// Validate performs validation on Individual data
func (i *Individual) Validate() error {
	if i == nil {
		return ErrInvalidParty
	}

	// Check in TMF-style individualIdentification
	if len(i.IndividualIdentification) == 0 {
		return ErrInvalidIDType
	}
	idItem := i.IndividualIdentification[0]
	if strings.TrimSpace(idItem.IdentificationType) == "" {
		return ErrInvalidIDType
	}

	if strings.TrimSpace(idItem.IdentificationId) == "" {
		return ErrInvalidIDNumber
	}
	gd := strings.TrimSpace(i.Gender)
	if gd != "M" && gd != "F" && gd != "U" {
		return ErrInvalidGender
	}

	return nil
}

// Validate performs validation on ContactMedium data
func (c *ContactMediumItem) Validate() error {
	if c == nil {
		return errors.New("invalid contact medium data")
	}

	// Different validation based on medium type
	switch c.MediumType {
	case "email":
		if strings.TrimSpace(c.PhoneNumber) == "" {
			return errors.New("email address is required for email contact type")
		}
	case "phone":
		if strings.TrimSpace(c.PhoneNumber) == "" {
			return errors.New("phone number is required for phone contact type")
		}
	case "address":
		if strings.TrimSpace(c.Street1) == "" && strings.TrimSpace(c.HomeNumber) == "" {
			return errors.New("either street or home number is required for address")
		}

		if strings.TrimSpace(c.PostalCode) == "" {
			return errors.New("postal code is required for address")
		}
	}

	return nil
}

// Validate performs validation on Characteristic data
func (c *CharacteristicItem) Validate() error {
	if c == nil {
		return errors.New("invalid characteristic data")
	}

	if strings.TrimSpace(c.Name) == "" {
		return errors.New("characteristic name is required")
	}

	return nil
}
