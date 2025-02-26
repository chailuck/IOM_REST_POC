package model

import (
	"time"
)

// Individual represents a human party in the TMF632 model
type Individual struct {
	ID               string    `json:"id" gorm:"column:paty_row_id;primaryKey;type:char(15)"`
	Type             string    `json:"@type" gorm:"-"`
	CreationDate     time.Time `json:"@creationDate" gorm:"column:crtd_dttm"`
	CreatedBy        string    `json:"@createdBy" gorm:"column:crtd_by"`
	ModificationDate time.Time `json:"@modificationDate" gorm:"column:last_chng_dttm"`
	ModifiedBy       string    `json:"@modifiedBy" gorm:"column:last_chng_by"`

	// Core fields
	IDType          string    `json:"identificationType" gorm:"column:id_type;type:char(2)"`
	IDNumber        string    `json:"identificationId" gorm:"column:id_numb;type:char(20)"`
	Language        string    `json:"language" gorm:"column:lang;type:char(1)"`
	HomePhoneNumber string    `json:"phoneNumber" gorm:"column:home_telp_numb;type:char(21)"`
	Title           string    `json:"title" gorm:"column:titl;type:varchar(40)"`
	GivenName       string    `json:"givenName" gorm:"column:frst_name;type:varchar(80)"`
	FamilyName      string    `json:"familyName" gorm:"column:last_name;type:varchar(80)"`
	MaritalStatus   string    `json:"maritalStatus" gorm:"column:marl_stts;type:char(1)"`
	Gender          string    `json:"gender" gorm:"column:gndr;type:char(1)"`
	NameType        string    `json:"nameType" gorm:"column:name_type;type:varchar(10)"`
	ContactLanguage string    `json:"contactLanguage" gorm:"column:cntc_lang;type:char(3)"`
	Nationality     string    `json:"nationality" gorm:"column:ntnt_code;type:char(4)"`
	BirthDate       time.Time `json:"birthDate" gorm:"column:date_of_brth"`
	IDExpiryDate    time.Time `json:"validFor.endDateTime" gorm:"column:id_exp_date"`
	Grade           string    `json:"-" gorm:"column:grde;type:char(1)"`

	// External references
	BillingID  string `json:"billingId" gorm:"column:bl_ext_id;type:varchar(20)"`
	ExternalID string `json:"externalId" gorm:"column:ext_id;type:varchar(70)"`

	// Status and type fields
	LargeCustomerInd string `json:"-" gorm:"column:lrge_cust_indc;type:char(2)"`
	SubscriberType   string `json:"-" gorm:"column:sub_type;type:varchar(5)"`
	PartyType        string `json:"-" gorm:"column:paty_type;type:varchar(5)"`
	SourceCustType   string `json:"-" gorm:"column:src_cust_type;type:varchar(5)"`

	// Parent information
	ParentTitle     string `json:"-" gorm:"column:parn_titl;type:varchar(40)"`
	ParentFirstName string `json:"-" gorm:"column:parn_frst_name;type:varchar(80)"`
	ParentLastName  string `json:"-" gorm:"column:parn_last_name;type:varchar(80)"`
	ParentContact   string `json:"-" gorm:"column:parn_cntc;type:char(21)"`
	ParentID        string `json:"-" gorm:"column:parn_id;type:char(21)"`

	// Subscriber information
	TotalSubscribers int    `json:"-" gorm:"column:totl_subs"`
	MaxAllowSubs     int    `json:"-" gorm:"column:max_allw_only_subs"`
	DealPool         string `json:"-" gorm:"column:deal_pool;type:varchar(20)"`

	// Related entities
	ContactMedium   []ContactMedium  `json:"contactMedium" gorm:"-"`  // Handled manually in repository
	Address         []PartyAddress   `json:"address" gorm:"-"`        // Handled manually in repository
	Characteristics []Characteristic `json:"characteristic" gorm:"-"` // Handled manually in repository

	// Fields for individual identification in JSON
	IndividualIdentification []IdentificationItem `json:"individualIdentification,omitempty" gorm:"-"`

	// Fields for characteristics in JSON
	PartyCharacteristic []CharacteristicItem `json:"partyCharacteristic,omitempty" gorm:"-"`

	// Fields for external references in JSON
	ExternalReference []ExternalReferenceItem `json:"externalReference,omitempty" gorm:"-"`
}

// ContactMedium represents contact methods in the TMF model
type ContactMedium struct {
	ID           string    `json:"id" gorm:"column:addr_row_id;primaryKey;type:char(15)"`
	MediumType   string    `json:"type" gorm:"-"`
	PartyID      string    `json:"-" gorm:"column:paty_row_id;type:char(15)"`
	Preferred    bool      `json:"preferred" gorm:"-"`
	CreatedDate  time.Time `json:"-" gorm:"column:crtd_dttm"`
	CreatedBy    string    `json:"-" gorm:"column:crtd_by"`
	ModifiedDate time.Time `json:"-" gorm:"column:last_chng_dttm"`
	ModifiedBy   string    `json:"-" gorm:"column:last_chng_by"`

	// Address fields from cst_paty_addr
	PhoneNumber string `json:"phoneNumber,omitempty" gorm:"-"`
	AddressType string `json:"addressType" gorm:"column:addr_type;type:char(2)"`
	Street1     string `json:"street1,omitempty" gorm:"column:adr1;type:varchar(80)"`
	Street2     string `json:"street2,omitempty" gorm:"column:adr2;type:varchar(80)"`
	Country     string `json:"country,omitempty" gorm:"column:cnty_code;type:char(4)"`
	PostalCode  string `json:"postCode,omitempty" gorm:"column:addr_post_code;type:char(8)"`
	PostcodeSeq int16  `json:"postcodeSequence,omitempty" gorm:"column:pscd_seqn_numb"`

	// Extended address fields from cst_paty_addr_ext
	Building         string `json:"building,omitempty" gorm:"-"`
	HomeNumber       string `json:"homenumber,omitempty" gorm:"-"`
	Moo              string `json:"moo,omitempty" gorm:"-"`
	Tumbol           string `json:"tumbol,omitempty" gorm:"-"`
	City             string `json:"city,omitempty" gorm:"-"`
	AccomodationType string `json:"accomdationType,omitempty" gorm:"-"`
	TimeAtAddress    string `json:"time_at_addr,omitempty" gorm:"-"`
}

// PartyAddress represents a party's address in the database model
type PartyAddress struct {
	ID           string    `json:"id" gorm:"column:addr_row_id;primaryKey;type:char(15)"`
	PartyID      string    `json:"-" gorm:"column:paty_row_id;type:char(15)"`
	BillingExtID string    `json:"-" gorm:"column:bl_ext_id;type:varchar(20)"`
	ExtID        string    `json:"-" gorm:"column:ext_id;type:varchar(70)"`
	IDType       string    `json:"-" gorm:"column:id_type;type:char(2)"`
	IDNumber     string    `json:"-" gorm:"column:id_numb;type:char(20)"`
	AddressType  string    `json:"-" gorm:"column:addr_type;type:char(2)"`
	Street1      string    `json:"-" gorm:"column:adr1;type:varchar(80)"`
	Street2      string    `json:"-" gorm:"column:adr2;type:varchar(80)"`
	Country      string    `json:"-" gorm:"column:cnty_code;type:char(4)"`
	PostalCode   string    `json:"-" gorm:"column:addr_post_code;type:char(8)"`
	PostcodeSeq  int16     `json:"-" gorm:"column:pscd_seqn_numb"`
	AddressPCXT  string    `json:"-" gorm:"column:addr_pcxt;type:char(5)"`
	CreatedDate  time.Time `json:"-" gorm:"column:crtd_dttm"`
	CreatedBy    string    `json:"-" gorm:"column:crtd_by"`
	ModifiedDate time.Time `json:"-" gorm:"column:last_chng_dttm"`
	ModifiedBy   string    `json:"-" gorm:"column:last_chng_by"`
}

// PartyAddressExt represents extended address details in the database model
type PartyAddressExt struct {
	ID           string    `json:"-" gorm:"column:addr_row_id;primaryKey;type:char(15)"`
	IDType       string    `json:"-" gorm:"column:id_type;type:char(2)"`
	IDNumber     string    `json:"-" gorm:"column:id_numb;type:char(20)"`
	AddressType  string    `json:"-" gorm:"column:addr_type;type:char(2)"`
	AmphurName   string    `json:"-" gorm:"column:ampr_name;type:char(40)"`
	BuildingName string    `json:"-" gorm:"column:buld_name;type:char(40)"`
	CityName     string    `json:"-" gorm:"column:city_name;type:char(40)"`
	HomeNo       string    `json:"-" gorm:"column:home_no;type:char(20)"`
	Moo          string    `json:"-" gorm:"column:moo;type:char(10)"`
	StreetName   string    `json:"-" gorm:"column:strt_name;type:char(40)"`
	TimeAtAddr   string    `json:"-" gorm:"column:time_at_addr;type:char(40)"`
	TumbolName   string    `json:"-" gorm:"column:tmbl_name;type:char(10)"`
	AccomType    string    `json:"-" gorm:"column:accm_type;type:char(10)"`
	ZipCode      string    `json:"-" gorm:"column:zip_code;type:char(10)"`
	CreatedDate  time.Time `json:"-" gorm:"column:crtd_dttm"`
	CreatedBy    string    `json:"-" gorm:"column:crtd_by"`
	ModifiedDate time.Time `json:"-" gorm:"column:last_chng_dttm"`
	ModifiedBy   string    `json:"-" gorm:"column:last_chng_by"`
}

// Characteristic represents party attributes/characteristics in the database model
type Characteristic struct {
	ID           string    `json:"id" gorm:"column:attr_row_id;primaryKey;type:char(15)"`
	PartyID      string    `json:"-" gorm:"column:paty_row_id;type:char(15)"`
	IDType       string    `json:"-" gorm:"column:id_type;type:char(2)"`
	IDNumber     string    `json:"-" gorm:"column:id_numb;type:char(20)"`
	BillingExtID string    `json:"-" gorm:"column:bl_ext_id;type:varchar(20)"`
	ExtID        string    `json:"-" gorm:"column:ext_id;type:varchar(70)"`
	Name         string    `json:"-" gorm:"column:attr_name;type:varchar(20)"`
	Value        string    `json:"-" gorm:"column:attr_vlue;type:varchar(40)"`
	CreatedDate  time.Time `json:"-" gorm:"column:crtd_dttm"`
	CreatedBy    string    `json:"-" gorm:"column:crtd_by"`
	ModifiedDate time.Time `json:"-" gorm:"column:last_chng_dttm"`
	ModifiedBy   string    `json:"-" gorm:"column:last_chng_by"`
}

// Helper structs for JSON marshaling/unmarshaling

// IdentificationItem represents individual identification in TMF632
type IdentificationItem struct {
	IdentificationType string      `json:"identificationType"`
	IdentificationId   string      `json:"identificationId"`
	ValidFor           *TimePeriod `json:"validFor,omitempty"`
}

// ExternalReferenceItem represents external references in TMF632
type ExternalReferenceItem struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

// CharacteristicItem represents characteristics in TMF632
type CharacteristicItem struct {
	Name  string `json:"name"`
	Value string `json:"value"`
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

func (Individual) TableName() string {
	return "cst_paty"
}

func (Characteristic) TableName() string {
	return "cst_paty_attr"
}

func (PartyAddress) TableName() string {
	return "cst_paty_addr"
}

func (PartyAddressExt) TableName() string {
	return "cst_paty_addr_ext"
}
