package model

import (
	"errors"
	"strings"
)

var (
	ErrInvalidParty    = errors.New("invalid party data")
	ErrInvalidIDType   = errors.New("invalid ID type")
	ErrInvalidIDNumber = errors.New("invalid ID number")
	ErrInvalidName     = errors.New("first name is required")
	ErrPartyNotFound   = errors.New("party not found")
	ErrDuplicateParty  = errors.New("party already exists")
)

// Validate performs validation on Individual data
func (i *Individual) Validate() error {
	if i == nil {
		return ErrInvalidParty
	}

	if strings.TrimSpace(i.IDType) == "" {
		return ErrInvalidIDType
	}

	if strings.TrimSpace(i.IDNumber) == "" {
		return ErrInvalidIDNumber
	}

	if strings.TrimSpace(i.FirstName) == "" {
		return ErrInvalidName
	}

	return nil
}

// Validate performs validation on PartyAddress data
func (a *PartyAddress) Validate() error {
	if a == nil {
		return errors.New("invalid address data")
	}

	if strings.TrimSpace(a.AddressType) == "" {
		return errors.New("address type is required")
	}

	if strings.TrimSpace(a.Street1) == "" && strings.TrimSpace(a.HomeNumber) == "" {
		return errors.New("either street or home number is required")
	}

	if strings.TrimSpace(a.PostalCode) == "" {
		return errors.New("postal code is required")
	}

	return nil
}

// Validate performs validation on Characteristic data
func (c *Characteristic) Validate() error {
	if c == nil {
		return errors.New("invalid characteristic data")
	}

	if strings.TrimSpace(c.Name) == "" {
		return errors.New("characteristic name is required")
	}

	if strings.TrimSpace(c.Value) == "" {
		return errors.New("characteristic value is required")
	}

	return nil
}
