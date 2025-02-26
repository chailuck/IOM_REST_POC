package model

import (
	"time"
)

// Individual represents a human party in the TMF632 model
type Individual struct {
	ID              string    `json:"id" gorm:"column:paty_row_id;primaryKey;type:char(15)"`
	IDType          string    `json:"idType" gorm:"column:id_type"`
	IDNumber        string    `json:"idNumber" gorm:"column:id_numb"`
	Language        string    `json:"language" gorm:"column:lang"`
	Title           string    `json:"title" gorm:"column:titl"`
	FirstName       string    `json:"firstName" gorm:"column:frst_name"`
	LastName        string    `json:"lastName" gorm:"column:last_name"`
	MaritalStatus   string    `json:"maritalStatus" gorm:"column:marl_stts"`
	Gender          string    `json:"gender" gorm:"column:gndr"`
	Nationality     string    `json:"nationality" gorm:"column:ntnt_code"`
	ContactLanguage string    `json:"contactLanguage" gorm:"column:cntc_lang"`
	BirthDate       time.Time `json:"birthDate" gorm:"column:date_of_brth"`
	IDExpiryDate    time.Time `json:"idExpiryDate" gorm:"column:id_exp_date"`
	Status          string    `json:"status" gorm:"column:paty_type"`
	SubType         string    `json:"subType" gorm:"column:sub_type"`
	CreatedDate     time.Time `json:"createdDate" gorm:"column:crtd_dttm"`
	CreatedBy       string    `json:"createdBy" gorm:"column:crtd_by"`
	ModifiedDate    time.Time `json:"modifiedDate" gorm:"column:last_chng_dttm"`
	ModifiedBy      string    `json:"modifiedBy" gorm:"column:last_chng_by"`

	// Related entities
	ContactMedium  []ContactMedium  `json:"contactMedium" gorm:"foreignKey:PartyID"`
	Characteristic []Characteristic `json:"characteristic" gorm:"foreignKey:PartyID"`
	Address        []PartyAddress   `json:"address" gorm:"foreignKey:PartyID"`
}

// ContactMedium represents various contact methods for a party
type ContactMedium struct {
	ID           string    `json:"id" gorm:"primaryKey"`
	PartyID      string    `json:"partyId" gorm:"column:paty_row_id"`
	MediumType   string    `json:"mediumType"`
	Preferred    bool      `json:"preferred"`
	CreatedDate  time.Time `json:"createdDate" gorm:"column:crtd_dttm"`
	ModifiedDate time.Time `json:"modifiedDate" gorm:"column:last_chng_dttm"`

	// Embedded contact details
	PhoneNumber string `json:"phoneNumber" gorm:"column:home_telp_numb"`
}

// Characteristic represents party attributes/characteristics
type Characteristic struct {
	ID           string    `json:"id" gorm:"column:attr_row_id;primaryKey;type:char(15)"`
	PartyID      string    `json:"partyId" gorm:"column:paty_row_id;type:char(15)"`
	Name         string    `json:"name" gorm:"column:attr_name"`
	Value        string    `json:"value" gorm:"column:attr_vlue"`
	CreatedDate  time.Time `json:"createdDate" gorm:"column:crtd_dttm"`
	ModifiedDate time.Time `json:"modifiedDate" gorm:"column:last_chng_dttm"`
}

// PartyAddress represents a party's address
type PartyAddress struct {
	ID           string    `json:"id" gorm:"column:addr_row_id;primaryKey;type:char(15)"`
	PartyID      string    `json:"partyId" gorm:"column:paty_row_id;type:char(15)"`
	AddressType  string    `json:"addressType" gorm:"column:addr_type"`
	Street1      string    `json:"street1" gorm:"column:adr1"`
	Street2      string    `json:"street2" gorm:"column:adr2"`
	PostalCode   string    `json:"postalCode" gorm:"column:addr_post_code"`
	Country      string    `json:"country" gorm:"column:cnty_code"`
	CreatedDate  time.Time `json:"createdDate" gorm:"column:crtd_dttm"`
	ModifiedDate time.Time `json:"modifiedDate" gorm:"column:last_chng_dttm"`

	// Extended address fields
	BuildingName string `json:"buildingName" gorm:"column:buld_name"`
	District     string `json:"district" gorm:"column:ampr_name"`
	SubDistrict  string `json:"subDistrict" gorm:"column:tmbl_name"`
	HomeNumber   string `json:"homeNumber" gorm:"column:home_no"`
	Moo          string `json:"moo" gorm:"column:moo"`
}

// TimePeriod represents a period of time with start and end dates
type TimePeriod struct {
	StartDateTime time.Time `json:"startDateTime"`
	EndDateTime   time.Time `json:"endDateTime"`
}

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

func (Individual) TableName() string {
	return "cst_paty"
}

func (Characteristic) TableName() string {
	return "cst_paty_attr"
}

func (PartyAddress) TableName() string {
	return "cst_paty_addr"
}
