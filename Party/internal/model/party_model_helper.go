package model

import (
	"strconv"
	"strings"
)

func MapIndividualToParty(individual *Individual) (p Party) {
	party := Party{
		ID:              individual.ID,
		Status:          individual.Status,
		Title:           individual.Title,
		FirstName:       individual.GivenName,
		LastName:        individual.FamilyName,
		Gender:          individual.Gender,
		BirthDate:       individual.BirthDate,
		MaritalStatus:   individual.Status,
		Nationality:     individual.Nationality,
		Language:        "T",
		HomePhoneNumber: "",

		//IDType:       individual.IDType,
		//IDNumber:     individual.IDNumber,
		//IDExpiryDate: individual.IDExpiryDate,

		CreationDate:     individual.CreationDate,
		CreatedBy:        individual.CreatedBy,
		ModificationDate: individual.ModificationDate,
		ModifiedBy:       individual.ModifiedBy,
	}

	if len(individual.IndividualIdentification) == 1 {
		party.IDType = individual.IndividualIdentification[0].IdentificationType
		party.IDNumber = individual.IndividualIdentification[0].IdentificationId
		party.IDExpiryDate = individual.IndividualIdentification[0].ValidFor.EndDateTime
	}
	for i := range individual.LanguageAbility {
		if individual.LanguageAbility[i].IsFavouriteLanguage {
			party.Language = individual.LanguageAbility[i].LanguageCode
		}
	}

	for i := range individual.ContactMedium {
		if (individual.ContactMedium[i].Type == EntityTypeContactMediumPhone) && (individual.ContactMedium[i].MediumType == EntityTypeContactMediumTypeHomePhoneNumber) {
			party.HomePhoneNumber = strings.TrimSpace(individual.ContactMedium[i].PhoneNumber)
			break
		}
	}

	for i := range individual.ExternalReference {
		if individual.ExternalReference[i].Type == EntityTypeExternalRef {
			if individual.ExternalReference[i].IdentifierType == EntityTypeExternalRefTypeBillingExtID {
				party.BillingID = strings.TrimSpace(individual.ExternalReference[i].Name)
			}
			if individual.ExternalReference[i].IdentifierType == EntityTypeExternalRefTypeExtID {
				party.ExternalID = strings.TrimSpace(individual.ExternalReference[i].Name)
			}
		}
	}

	for i := range individual.Characteristic {
		if individual.Characteristic[i].Value != "" {
			switch individual.Characteristic[i].Name {
			case "occupationCode":
				party.OccupationCode = strings.TrimSpace(individual.Characteristic[i].Value)
			case "timeInBusiness":
				party.TimeInBusiness = strings.TrimSpace(individual.Characteristic[i].Value)
			case "salaryLevel":
				party.SalaryLevel = strings.TrimSpace(individual.Characteristic[i].Value)
			case "grade":
				party.Grade = strings.TrimSpace(individual.Characteristic[i].Value)
			case "largeCustomerIndicator":
				party.LargeCustomerInd = strings.TrimSpace(individual.Characteristic[i].Value)
			case "subscriberType":
				party.SubscriberType = strings.TrimSpace(individual.Characteristic[i].Value)
			case "sourceCustomerType":
				party.SourceCustType = strings.TrimSpace(individual.Characteristic[i].Value)
			case "parentTitle":
				party.ParentTitle = strings.TrimSpace(individual.Characteristic[i].Value)
			case "parentFirstName":
				party.ParentFirstName = strings.TrimSpace(individual.Characteristic[i].Value)
			case "parentLastName":
				party.ParentLastName = strings.TrimSpace(individual.Characteristic[i].Value)
			case "parentContact":
				party.ParentContact = strings.TrimSpace(individual.Characteristic[i].Value)
			case "parentId":
				party.ParentID = strings.TrimSpace(individual.Characteristic[i].Value)
			case "totalNoOfSubscriber":
				numbValue, err := strconv.Atoi(individual.Characteristic[i].Value)
				if err == nil {
					party.TotalSubscribers = numbValue
				}
			case "l":
				numbValue, err := strconv.Atoi(individual.Characteristic[i].Value)
				if err == nil {
					party.MaxAllowSubs = numbValue
				}
			case "dealPool":
				party.DealPool = strings.TrimSpace(individual.Characteristic[i].Value)
			}
		}
	}
	return party
}

func MapIndividualToPartyAddress(individual *Individual) (addr []PartyAddress, addrExt []PartyAddressExt) {
	var addrList []PartyAddress
	var addrExtList []PartyAddressExt

	for i := range individual.ContactMedium {

		if strings.TrimSpace(individual.ContactMedium[i].Type) == EntityTypeContactMediumAddress {

			addr := PartyAddress{
				ID:           individual.ContactMedium[i].ID,
				PartyID:      individual.ID,
				IDType:       individual.IndividualIdentification[0].IdentificationType,
				IDNumber:     individual.IndividualIdentification[0].IdentificationId,
				AddressType:  individual.ContactMedium[i].AddressType,
				Street1:      individual.ContactMedium[i].Street1,
				Street2:      individual.ContactMedium[i].Street2,
				Country:      individual.ContactMedium[i].Country,
				PostalCode:   individual.ContactMedium[i].PostalCode,
				PostcodeSeq:  individual.ContactMedium[i].PostcodeSeq,
				CreatedDate:  individual.ContactMedium[i].AuditTrail.CreationDate,
				CreatedBy:    individual.ContactMedium[i].AuditTrail.CreatedBy,
				ModifiedDate: individual.ContactMedium[i].AuditTrail.ModificationDate,
				ModifiedBy:   individual.ContactMedium[i].AuditTrail.ModifiedBy,
			}
			addrExt := PartyAddressExt{
				ID:           individual.ContactMedium[i].ID,
				IDType:       individual.IndividualIdentification[0].IdentificationType,
				IDNumber:     individual.IndividualIdentification[0].IdentificationId,
				AddressType:  individual.ContactMedium[i].AddressType,
				AmphurName:   individual.ContactMedium[i].Amphur,
				BuildingName: individual.ContactMedium[i].Building,
				CityName:     individual.ContactMedium[i].City,
				HomeNo:       individual.ContactMedium[i].HomeNumber,
				Moo:          individual.ContactMedium[i].Moo,
				StreetName:   individual.ContactMedium[i].Street1 + " " + individual.ContactMedium[i].Street2,
				TimeAtAddr:   individual.ContactMedium[i].TimeAtAddress,
				TumbolName:   individual.ContactMedium[i].Tumbol,
				AccomType:    individual.ContactMedium[i].AccomodationType,
				ZipCode:      individual.ContactMedium[i].PostalCode,
				CreatedDate:  individual.ContactMedium[i].AuditTrail.CreationDate,
				CreatedBy:    individual.ContactMedium[i].AuditTrail.CreatedBy,
				ModifiedDate: individual.ContactMedium[i].AuditTrail.ModificationDate,
				ModifiedBy:   individual.ContactMedium[i].AuditTrail.ModifiedBy,
			}

			for j := range individual.ContactMedium[i].ExternalReference {
				if individual.ContactMedium[i].ExternalReference[j].Type == EntityTypeExternalRef {
					if individual.ContactMedium[i].ExternalReference[j].IdentifierType == EntityTypeExternalRefTypeBillingExtID {
						addr.BillingExtID = strings.TrimSpace(individual.ContactMedium[i].ExternalReference[j].Name)
					}
					if individual.ContactMedium[i].ExternalReference[j].IdentifierType == EntityTypeExternalRefTypeExtID {
						addr.ExtID = strings.TrimSpace(individual.ContactMedium[i].ExternalReference[j].Name)
					}
				}
			}
			addrList = append(addrList, addr)
			addrExtList = append(addrExtList, addrExt)
		}
	}
	return addrList, addrExtList
}

func MapIndividualToPartyAttribute(individual *Individual) (attr []PartyAttributes) {
	var attrList []PartyAttributes

	for i := range individual.Characteristic {
		attr := PartyAttributes{
			ID:       individual.Characteristic[i].ID,
			PartyID:  individual.ID,
			IDType:   individual.IndividualIdentification[0].IdentificationType,
			IDNumber: individual.IndividualIdentification[0].IdentificationId,
			//BillingExtID
			//ExtID
			Name:         individual.Characteristic[i].Name,
			Value:        individual.Characteristic[i].Value,
			CreatedDate:  individual.Characteristic[i].AuditTrail.CreationDate,
			CreatedBy:    individual.Characteristic[i].AuditTrail.CreatedBy,
			ModifiedDate: individual.Characteristic[i].AuditTrail.ModificationDate,
			ModifiedBy:   individual.Characteristic[i].AuditTrail.ModifiedBy,
		}

		for i := range individual.Characteristic[i].ExternalReference {
			if individual.Characteristic[i].ExternalReference[i].Type == EntityTypeExternalRef {
				if individual.Characteristic[i].ExternalReference[i].IdentifierType == EntityTypeExternalRefTypeBillingExtID {
					attr.BillingExtID = strings.TrimSpace(individual.ContactMedium[i].ExternalReference[i].Name)
				}
				if individual.ExternalReference[i].IdentifierType == EntityTypeExternalRefTypeExtID {
					attr.ExtID = strings.TrimSpace(individual.ContactMedium[i].ExternalReference[i].Name)
				}
			}
		}
		attrList = append(attrList, attr)
	}
	return attrList
}

func MapCharacteristicToPartyAttribute(partyId string, idType string, idNumber string, characteristic *CharacteristicItem) (attr PartyAttributes) {

	attr = PartyAttributes{
		ID:           characteristic.ID,
		PartyID:      partyId,
		IDType:       idType,
		IDNumber:     idNumber,
		Name:         characteristic.Name,
		Value:        characteristic.Value,
		CreatedDate:  characteristic.AuditTrail.CreationDate,
		CreatedBy:    characteristic.AuditTrail.CreatedBy,
		ModifiedDate: characteristic.AuditTrail.ModificationDate,
		ModifiedBy:   characteristic.AuditTrail.ModifiedBy,
	}

	for i := range characteristic.ExternalReference {
		if characteristic.ExternalReference[i].Type == EntityTypeExternalRef {
			if characteristic.ExternalReference[i].IdentifierType == EntityTypeExternalRefTypeBillingExtID {
				attr.BillingExtID = strings.TrimSpace(characteristic.ExternalReference[i].Name)
			}
			if characteristic.ExternalReference[i].IdentifierType == EntityTypeExternalRefTypeExtID {
				attr.ExtID = strings.TrimSpace(characteristic.ExternalReference[i].Name)
			}
		}
	}
	return attr
}

func MapPartyToIndividual(party Party, addrList []PartyAddress, addrExtList []PartyAddressExt, attr []PartyAttributes) (individual Individual) {
	individual = Individual{
		ID:            party.ID,
		Type:          EntityTypePartyIndividual,
		BaseType:      EntityBaseTypeParty,
		Status:        party.Status,
		Href:          party.ID,
		Title:         party.Title,
		GivenName:     party.FirstName,
		FamilyName:    party.LastName,
		Gender:        party.Gender,
		BirthDate:     party.BirthDate,
		MaritalStatus: party.MaritalStatus,
		Nationality:   party.Nationality,
		// Core field
		CreationDate:     party.CreationDate,
		CreatedBy:        party.CreatedBy,
		ModificationDate: party.ModificationDate,
		ModifiedBy:       party.ModifiedBy,

		/*		ExternalReference
				IndividualIdentif
				LanguageAbility
				ContactMedium
				Characteristic   */
	}
	var IDValidity TimePeriod
	IDValidity = TimePeriod{
		EndDateTime: party.IDExpiryDate,
	}

	var identification IdentificationItem
	identification = IdentificationItem{
		IdentificationType: party.IDType,
		IdentificationId:   party.IDNumber,
		ValidFor:           IDValidity,
	}
	individual.IndividualIdentification = append(individual.IndividualIdentification, identification)

	var extRef ExternalReferenceItem
	if strings.TrimSpace(party.BillingID) != "" {
		extRef = ExternalReferenceItem{
			Type:           party.BillingID,
			Name:           EntityTypeExternalRefTypeBillingExtID,
			IdentifierType: EntityTypeExternalRef,
		}
		individual.ExternalReference = append(individual.ExternalReference, extRef)
	}

	if strings.TrimSpace(party.ExternalID) != "" {
		extRef = ExternalReferenceItem{
			Type:           party.ExternalID,
			Name:           EntityTypeExternalRefTypeExtID,
			IdentifierType: EntityTypeExternalRef,
		}
		individual.ExternalReference = append(individual.ExternalReference, extRef)
	}

	//home phone number
	if strings.TrimSpace(party.HomePhoneNumber) != "" {
		individual.ContactMedium = append(individual.ContactMedium, ContactMediumItem{
			ID:          "HOMEPHONE_" + party.ID,
			Type:        EntityTypeContactMediumPhone,
			BaseType:    EntityBaseTypeContactMedium,
			MediumType:  EntityTypeContactMediumTypeHomePhoneNumber,
			Preferred:   true,
			PhoneNumber: party.HomePhoneNumber,
		})
	}

	// address
	addressContactList := MapPartyAddressToContactMediums(addrList, addrExtList)
	if len(addressContactList) > 0 {
		individual.ContactMedium = append(individual.ContactMedium, addressContactList...)
	}

	assignCharStringIfNotEmpty := func(name string, value string) {
		if value != "" {
			individual.Characteristic = append(individual.Characteristic, CharacteristicItem{
				Name:      name,
				Value:     value,
				Type:      "StringCharacteristic",
				ValueType: "String",
			})
		}
	}

	assignCharIntegerIfNotEmpty := func(name string, value int) {

		if value != -99 {
			strValue := strconv.Itoa(value)

			individual.Characteristic = append(individual.Characteristic, CharacteristicItem{
				Name:      name,
				Value:     strValue,
				Type:      "IntegerCharacteristic",
				ValueType: "Number",
			})
		}
	}
	assignCharStringIfNotEmpty("occupationCode", party.OccupationCode)
	assignCharStringIfNotEmpty("timeInBusiness", party.TimeInBusiness)
	assignCharStringIfNotEmpty("salaryLevel", party.SalaryLevel)
	assignCharStringIfNotEmpty("grade", party.Grade)
	assignCharStringIfNotEmpty("largeCustomerIndicator", party.LargeCustomerInd)
	assignCharStringIfNotEmpty("subscriberType", party.SubscriberType)
	assignCharStringIfNotEmpty("sourceCustomerType", party.SourceCustType)
	assignCharStringIfNotEmpty("parentTitle", party.ParentTitle)
	assignCharStringIfNotEmpty("parentFirstName", party.ParentFirstName)
	assignCharStringIfNotEmpty("parentLastName", party.ParentLastName)
	assignCharStringIfNotEmpty("parentContact", party.ParentContact)
	assignCharStringIfNotEmpty("parentId", party.ParentID)
	assignCharIntegerIfNotEmpty("totalNoOfSubscriber", party.TotalSubscribers)
	assignCharIntegerIfNotEmpty("maxAllowOnlySubscriber", party.MaxAllowSubs)
	assignCharStringIfNotEmpty("dealPool", party.DealPool)

	return individual
}

func MapPartyAddressToContactMediums(addrList []PartyAddress, addrExtList []PartyAddressExt) (cm []ContactMediumItem) {
	var contactMediumList []ContactMediumItem
	for i := range addrList {
		addr := addrList[i]
		contactItem := ContactMediumItem{
			Type:        EntityTypeContactMediumAddress,
			BaseType:    EntityBaseTypeContactMedium,
			ID:          addr.ID,
			Preferred:   true,
			MediumType:  addr.AddressType,
			Street1:     addr.Street1,
			Street2:     addr.Street2,
			Country:     addr.Country,
			PostalCode:  addr.PostalCode,
			PostcodeSeq: addr.PostcodeSeq,
		}
		if strings.TrimSpace(addr.Street1+addr.Street2) == "" {
			for j := range addrExtList {
				addrExt := addrExtList[j]
				if addrExt.ID == addr.ID {
					contactItem.Amphur = addrExt.AmphurName
					contactItem.Building = addrExt.BuildingName
					contactItem.City = addrExt.CityName
					contactItem.HomeNumber = addrExt.HomeNo
					contactItem.Moo = addrExt.Moo
					contactItem.Tumbol = addrExt.TumbolName
					contactItem.TimeAtAddress = addrExt.TimeAtAddr
					contactItem.Street1 = addrExt.StreetName
					contactItem.AccomodationType = addrExt.AccomType
					contactItem.PostalCode = addrExt.ZipCode
				}
			}
		}
		contactMediumList = append(contactMediumList, contactItem)
	}
	return contactMediumList
}
