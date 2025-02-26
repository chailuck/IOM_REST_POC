package model

import (
	"errors"
	"strconv"
	"strings"
)

var (
	ErrInvalidParty    = errors.New("invalid party data")
	ErrInvalidIDType   = errors.New("invalid ID type")
	ErrInvalidIDNumber = errors.New("invalid ID number")
	ErrInvalidName     = errors.New("given name is required")
	ErrPartyNotFound   = errors.New("party not found")
	ErrDuplicateParty  = errors.New("party already exists")
)

// Validate performs validation on Individual data
func (i *Individual) Validate() error {
	if i == nil {
		return ErrInvalidParty
	}

	// Check if we have identification information directly or in the TMF format
	if i.IDType == "" && i.IDNumber == "" {
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
	} else {
		// Check direct fields
		if strings.TrimSpace(i.IDType) == "" {
			return ErrInvalidIDType
		}

		if strings.TrimSpace(i.IDNumber) == "" {
			return ErrInvalidIDNumber
		}
	}

	return nil
}

// Validate performs validation on ContactMedium data
func (c *ContactMedium) Validate() error {
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
func (c *Characteristic) Validate() error {
	if c == nil {
		return errors.New("invalid characteristic data")
	}

	if strings.TrimSpace(c.Name) == "" {
		return errors.New("characteristic name is required")
	}

	return nil
}

// PrepareIndividualForDB prepares an Individual for database operations
// by mapping the JSON structure to database fields
func PrepareIndividualForDB(i *Individual) {
	// Set default type if not present
	if i.Type == "" {
		i.Type = "INDY"
	}

	// Extract individual identification info
	if len(i.IndividualIdentification) > 0 {
		idItem := i.IndividualIdentification[0]
		i.IDType = idItem.IdentificationType
		i.IDNumber = idItem.IdentificationId
		if idItem.ValidFor != nil {
			i.IDExpiryDate = idItem.ValidFor.EndDateTime
		}
	}

	// Extract external references
	for _, ref := range i.ExternalReference {
		if ref.Type == "PartyExtBillingID" {
			i.BillingID = ref.Name
		} else if ref.Type == "PartyExtID" {
			i.ExternalID = ref.Name
		}
	}

	// Extract characteristics
	for _, char := range i.PartyCharacteristic {
		switch char.Name {
		case "occupationCode":
			i.Characteristics = append(i.Characteristics, Characteristic{
				Name:  "occupationCode",
				Value: char.Value,
			})
		case "timeInBusiness":
			i.Characteristics = append(i.Characteristics, Characteristic{
				Name:  "timeInBusiness",
				Value: char.Value,
			})
		case "salaryLevel":
			i.Characteristics = append(i.Characteristics, Characteristic{
				Name:  "salaryLevel",
				Value: char.Value,
			})
		case "grade":
			i.Grade = char.Value
		case "largeCustomerIndicator":
			i.LargeCustomerInd = char.Value
		case "subscriberType":
			i.SubscriberType = char.Value
		case "sourceCustomerType":
			i.SourceCustType = char.Value
		case "parentTitle":
			i.ParentTitle = char.Value
		case "parentFirstName":
			i.ParentFirstName = char.Value
		case "parentLastName":
			i.ParentLastName = char.Value
		case "parentContact":
			i.ParentContact = char.Value
		case "parentId":
			i.ParentID = char.Value
		case "totalNoOfSubscriber":
			// Convert to int
			if intVal, err := strconv.Atoi(char.Value); err == nil {
				i.TotalSubscribers = intVal
			}
		case "maxAllowOnlySubscriber":
			// Convert to int
			if intVal, err := strconv.Atoi(char.Value); err == nil {
				i.MaxAllowSubs = intVal
			}
		case "dealPool":
			i.DealPool = char.Value
		default:
			// For custom attributes, add to characteristics collection
			i.Characteristics = append(i.Characteristics, Characteristic{
				Name:  char.Name,
				Value: char.Value,
			})
		}
	}

	// Process contact media for addresses
	for _, contact := range i.ContactMedium {
		if contact.MediumType == "address" {
			// Set address type if not present
			if contact.AddressType == "" {
				contact.AddressType = "HM" // Default to Home
			}
		}
	}
}

// PrepareIndividualForAPI prepares an Individual for API response
// by mapping the database fields to the JSON structure
func PrepareIndividualForAPI(i *Individual) {
	// Set TMF type
	i.Type = "INDY"

	// Prepare individual identification
	i.IndividualIdentification = []IdentificationItem{
		{
			IdentificationType: i.IDType,
			IdentificationId:   i.IDNumber,
			ValidFor: &TimePeriod{
				EndDateTime: i.IDExpiryDate,
			},
		},
	}

	// Prepare external references
	var references []ExternalReferenceItem

	if i.BillingID != "" {
		references = append(references, ExternalReferenceItem{
			Type: "PartyExtBillingID",
			Name: i.BillingID,
		})
	}

	if i.ExternalID != "" {
		references = append(references, ExternalReferenceItem{
			Type: "PartyExtID",
			Name: i.ExternalID,
		})
	}

	if len(references) > 0 {
		i.ExternalReference = references
	}

	// Prepare party characteristics
	var characteristics []CharacteristicItem

	// Core fields that map to characteristics
	addCharIfNotEmpty := func(name, value string) {
		if value != "" {
			characteristics = append(characteristics, CharacteristicItem{
				Name:  name,
				Value: value,
			})
		}
	}

	addCharIfNotEmpty("grade", i.Grade)
	addCharIfNotEmpty("largeCustomerIndicator", i.LargeCustomerInd)
	addCharIfNotEmpty("subscriberType", i.SubscriberType)
	addCharIfNotEmpty("sourceCustomerType", i.SourceCustType)
	addCharIfNotEmpty("parentTitle", i.ParentTitle)
	addCharIfNotEmpty("parentFirstName", i.ParentFirstName)
	addCharIfNotEmpty("parentLastName", i.ParentLastName)
	addCharIfNotEmpty("parentContact", i.ParentContact)
	addCharIfNotEmpty("parentId", i.ParentID)
	addCharIfNotEmpty("dealPool", i.DealPool)

	// Add totalSubscribers and maxAllowSubs if they are non-zero
	if i.TotalSubscribers > 0 {
		characteristics = append(characteristics, CharacteristicItem{
			Name:  "totalNoOfSubscriber",
			Value: strconv.Itoa(i.TotalSubscribers),
		})
	}

	if i.MaxAllowSubs > 0 {
		characteristics = append(characteristics, CharacteristicItem{
			Name:  "maxAllowOnlySubscriber",
			Value: strconv.Itoa(i.MaxAllowSubs),
		})
	}

	// Add custom characteristics from the Characteristics collection
	for _, char := range i.Characteristics {
		characteristics = append(characteristics, CharacteristicItem{
			Name:  char.Name,
			Value: char.Value,
		})
	}

	if len(characteristics) > 0 {
		i.PartyCharacteristic = characteristics
	}
}
