package model

import (
	"time"
)

// DATABASE MODEL

func (Party) TableName() string {
	return "cst_paty"
}

type Party struct {
	ID            string    `gorm:"column:paty_row_id;primaryKey;type:char(15)"`
	Status        string    `gorm:"status"`
	Title         string    `gorm:"column:titl;type:varchar(40)"`
	FirstName     string    `gorm:"column:frst_name;type:varchar(80)"`
	LastName      string    `gorm:"column:last_name;type:varchar(80)"`
	Gender        string    `gorm:"column:gndr;type:char(1)"`
	BirthDate     time.Time `gorm:"column:date_of_brth"`
	MaritalStatus string    `gorm:"column:marl_stts;type:char(1)"`
	Nationality   string    `gorm:"column:ntnt_code;type:char(4)"`
	Language      string    `gorm:"column:lang;type:char(1)"`

	// Core fields - IdentificationItem
	IDType       string    `gorm:"column:id_type;type:char(2)"`
	IDNumber     string    `gorm:"column:id_numb;type:char(20)"`
	IDExpiryDate time.Time `gorm:"column:id_exp_date"`

	HomePhoneNumber string `gorm:"column:home_telp_numb;type:char(21)"`

	// Parent information
	ParentTitle     string `gorm:"column:parn_titl;type:varchar(40)"`
	ParentFirstName string `gorm:"column:parn_frst_name;type:varchar(80)"`
	ParentLastName  string `gorm:"column:parn_last_name;type:varchar(80)"`
	ParentContact   string `gorm:"column:parn_cntc;type:char(21)"`
	ParentID        string `gorm:"column:parn_id;type:char(21)"`

	BillingID  string `gorm:"column:bl_ext_id;type:varchar(20)"`
	ExternalID string `gorm:"column:ext_id;type:varchar(70)"`

	// Party Charateristic
	OccupationCode       string `gorm:"column:occp_code;type:char(3)"`
	TimeInBusiness       string `gorm:"column:time_in_buss;type:varchar(10)"`
	SalaryLevel          string `gorm:"column:salr_levl;type:char(1)"`
	InitialTimeInAddress string `gorm:"column:init_time_in_addr;type:varchar(10)"`
	Grade                string `gorm:"column:grde;type:char(1)"`
	LargeCustomerInd     string `gorm:"column:lrge_cust_indc;type:char(2)"`
	SubscriberType       string `gorm:"column:sub_type;type:varchar(5)"`
	SourceCustType       string `gorm:"column:src_cust_type;type:varchar(5)"`
	TotalSubscribers     int    `gorm:"column:totl_subs"`
	MaxAllowSubs         int    `gorm:"column:max_allw_only_subs"`
	DealPool             string `gorm:"column:deal_pool;type:varchar(20)"`

	CreationDate     time.Time `gorm:"column:crtd_dttm"`
	CreatedBy        string    `gorm:"column:crtd_by"`
	ModificationDate time.Time `gorm:"column:last_chng_dttm"`
	ModifiedBy       string    `gorm:"column:last_chng_by"`
}

// PartyAddress represents a party's address in the database model
func (PartyAddress) TableName() string {
	return "cst_paty_addr"
}

type PartyAddress struct {
	ID           string    `json:"-" gorm:"column:addr_row_id;primaryKey;type:char(15)"`
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

func (PartyAddressExt) TableName() string {
	return "cst_paty_addr_ext"
}

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

func (PartyAttributes) TableName() string {
	return "cst_paty_attr"
}

type PartyAttributes struct {
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
