package model

import (
	"time"
)

const (
	EntityBaseTypeParty       = "Party"
	EntityTypePartyIndividual = "INDY"
	PartyNameTypeIndividual   = "INDY"

	EntityBaseTypeContactMedium                = "ContacMedium"
	EntityTypeContactMediumPhone               = "PhoneContactMedium"
	EntityTypeContactMediumAddress             = "GeographicAddressContactMedium"
	EntityTypeContactMediumTypeHomePhoneNumber = "homePhoneNumber"

	EntityTypeExternalRef                 = "ExternalIdentifier"
	EntityTypeExternalRefTypeBillingExtID = "PartyExtBillingID"
	EntityTypeExternalRefTypeExtID        = "PartyExtID"
)

// Individual represents a human party in the TMF632 model
type Individual struct {
	ID            string    `json:"id"`
	Type          string    `json:"@type"`
	BaseType      string    `json:"@baseType"`
	Status        string    `json:"-"`
	Href          string    `json:"href"`
	Title         string    `json:"title"`
	GivenName     string    `json:"givenName"`
	FamilyName    string    `json:"familyName"`
	Gender        string    `json:"gender"`
	BirthDate     time.Time `json:"birthDate"`
	MaritalStatus string    `json:"maritalStatus"`
	Nationality   string    `json:"nationality"`

	// Core fields - IdentificationItem
	//IDType       string    `json:"individualIdentification.identificationType"`
	//IDNumber     string    `json:"individualIdentification.identificationId"`
	IDExpiryDate time.Time `json:"individualIdentification.validFor.endDateTime"`

	CreationDate     time.Time `json:"@creationDate"`
	CreatedBy        string    `json:"@createdBy"`
	ModificationDate time.Time `json:"@modificationDate"`
	ModifiedBy       string    `json:"@modifiedBy"`

	AuditTrail DBAuditTrail

	ExternalReference        []ExternalReferenceItem `json:"externalReference,omitempty"`
	IndividualIdentification []IdentificationItem    `json:"individualIdentification,omitempty"`
	LanguageAbility          []languageAbilityItem   `json:"languageAbility,omitempty"`
	ContactMedium            []ContactMediumItem     `json:"contactMedium"`
	Characteristic           []CharacteristicItem    `json:"characteristic,omitempty"`
}

// ContactMedium represents contact methods in the TMF model
type ContactMediumItem struct {
	ID       string `json:"id"`
	Type     string `json:"@type" `
	BaseType string `json:"@baseType" `

	MediumType string `json:"mediumType"`
	Preferred  bool   `json:"preferred"`

	// Fields for external references in JSON
	ExternalReference []ExternalReferenceItem `json:"externalReference,omitempty"`

	// Type PhoneContact Medium
	PhoneNumber string `json:"phoneNumber,omitempty"`

	// Type GeographicAddress fields from cst_paty_addr
	AddressType string `json:"addressType"`
	Street1     string `json:"street1,omitempty"`
	Street2     string `json:"street2,omitempty"`
	Country     string `json:"country,omitempty"`
	PostalCode  string `json:"postCode,omitempty"`
	PostcodeSeq int16  `json:"postcodeSequence,omitempty"`
	// Type GeographicAddress fields - Extended address fields from cst_paty_addr_ext
	Building         string       `json:"building,omitempty"`
	HomeNumber       string       `json:"homenumber,omitempty"`
	Moo              string       `json:"moo,omitempty"`
	Tumbol           string       `json:"tumbol,omitempty"`
	Amphur           string       `json:"amphur,omitempty"`
	City             string       `json:"city,omitempty"`
	AccomodationType string       `json:"accomdationType,omitempty"`
	TimeAtAddress    string       `json:"time_at_addr,omitempty"`
	AuditTrail       DBAuditTrail `json:"-"`
}

// IdentificationItem represents individual identification in TMF632
type IdentificationItem struct {
	IdentificationType string `json:"identificationType"`
	IdentificationId   string `json:"identificationId"`
	//ValidFor           time.Time `json:"validFor,omitempty"`
	ValidFor TimePeriod `json:"validFor,omitempty"`
}

// ExternalReferenceItem represents external references in TMF632
type ExternalReferenceItem struct {
	Type           string `json:"@type"`
	Name           string `json:"name"`
	IdentifierType string `json:"externalIdentifierType"`
}

// CharacteristicItem represents characteristics in TMF632
type CharacteristicItem struct {
	ID                string                  `json:"id"`
	Type              string                  `json:"@type"`
	Name              string                  `json:"name"`
	Value             string                  `json:"value"`
	ValueType         string                  `json:"valueType"`
	AuditTrail        DBAuditTrail            `json:"-"`
	ExternalReference []ExternalReferenceItem `json:"externalReference,omitempty"`
}

// LanguageAbilityItem represents characteristics in TMF632
type languageAbilityItem struct {
	LanguageCode        string `json:"languageCode"`
	IsFavouriteLanguage bool   `json:"isFavouriteLanguage"`
}

// TimePeriod represents a period of time with start and end dates
type TimePeriod struct {
	StartDateTime time.Time `json:"startDateTime,omitempty"`
	EndDateTime   time.Time `json:"endDateTime,omitempty"`
}

// Response structures

// PartyCreateResponse represents the response for party creation
type PartyCreateResponse struct {
	Data  *Individual `json:"data"`
	Error string      `json:"error,omitempty"`
}

// PartyUpdateResponse represents the response for party updates
type PartyUpdateResponse struct {
	Data  *Individual `json:"data"`
	Error string      `json:"error,omitempty"`
}

// PartyQueryParams represents query parameters for party search
type PartyQueryParams struct {
	IDType    string `json:"idType"`
	IDNumber  string `json:"idNumber"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Status    string `json:"status"`
	Limit     int    `json:"limit"`
	Offset    int    `json:"offset"`
}

type DBAuditTrail struct {
	CreationDate     time.Time `json:"-"`
	CreatedBy        string    `json:"-"`
	ModificationDate time.Time `json:"-"`
	ModifiedBy       string    `json:"-"`
}
