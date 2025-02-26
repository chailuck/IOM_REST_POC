package repository

import (
	"party/internal/model"
)

// SearchParties performs advanced searches with flexible criteria
func (r *PartyRepository) SearchParties(criteria map[string]interface{}) ([]model.Individual, error) {
	var parties []model.Individual
	query := r.db.Model(&model.Individual{})

	// Apply all criteria as WHERE conditions
	for key, value := range criteria {
		// Map JSON field names to database columns
		switch key {
		case "givenName":
			query = query.Where("frst_name ILIKE ?", "%"+value.(string)+"%")
		case "familyName":
			query = query.Where("last_name ILIKE ?", "%"+value.(string)+"%")
		case "identificationType":
			query = query.Where("id_type = ?", value)
		case "identificationId":
			query = query.Where("id_numb = ?", value)
		case "maritalStatus":
			query = query.Where("marl_stts = ?", value)
		case "gender":
			query = query.Where("gndr = ?", value)
		case "nationality":
			query = query.Where("ntnt_code = ?", value)
		default:
			// Custom attribute search via characteristic table
			if value != nil && value.(string) != "" {
				subQuery := r.db.Table("cst_paty_attr").
					Select("paty_row_id").
					Where("attr_name = ? AND attr_vlue = ?", key, value)

				query = query.Where("paty_row_id IN (?)", subQuery)
			}
		}
	}

	// Execute query
	err := query.Find(&parties).Error
	if err != nil {
		return nil, err
	}

	// Get full party details for each result
	var result []model.Individual
	for _, party := range parties {
		fullParty, err := r.GetByID(party.ID)
		if err != nil {
			continue // Skip parties with errors
		}
		result = append(result, *fullParty)
	}

	return result, nil
}

// GetPartyByPhoneNumber finds parties by phone number
func (r *PartyRepository) GetPartyByPhoneNumber(phoneNumber string) ([]model.Individual, error) {
	var parties []model.Individual

	err := r.db.Where("home_telp_numb LIKE ?", "%"+phoneNumber+"%").Find(&parties).Error
	if err != nil {
		return nil, err
	}

	// Get full party details for each result
	var result []model.Individual
	for _, party := range parties {
		fullParty, err := r.GetByID(party.ID)
		if err != nil {
			continue // Skip parties with errors
		}
		result = append(result, *fullParty)
	}

	return result, nil
}

// GetPartyByAddress finds parties by address criteria
func (r *PartyRepository) GetPartyByAddress(postalCode, district, subDistrict string) ([]model.Individual, error) {
	// First find the address IDs that match criteria
	var addressMatches []model.PartyAddressExt
	query := r.db.Model(&model.PartyAddressExt{})

	if postalCode != "" {
		query = query.Where("zip_code = ?", postalCode)
	}

	if district != "" {
		query = query.Where("ampr_name ILIKE ?", "%"+district+"%")
	}

	if subDistrict != "" {
		query = query.Where("tmbl_name ILIKE ?", "%"+subDistrict+"%")
	}

	if err := query.Find(&addressMatches).Error; err != nil {
		return nil, err
	}

	// Extract IDs from matches
	var idNumbers []string
	for _, match := range addressMatches {
		idNumbers = append(idNumbers, match.IDNumber)
	}

	if len(idNumbers) == 0 {
		return []model.Individual{}, nil
	}

	// Find parties with these ID numbers
	var parties []model.Individual
	if err := r.db.Where("id_numb IN ?", idNumbers).Find(&parties).Error; err != nil {
		return nil, err
	}

	// Get full party details for each result
	var result []model.Individual
	for _, party := range parties {
		fullParty, err := r.GetByID(party.ID)
		if err != nil {
			continue // Skip parties with errors
		}
		result = append(result, *fullParty)
	}

	return result, nil
}

// GetPartyByExternalID finds a party by external ID
func (r *PartyRepository) GetPartyByExternalID(extIDType, extID string) (*model.Individual, error) {
	var party model.Individual

	if extIDType == "billing" {
		err := r.db.Where("bl_ext_id = ?", extID).First(&party).Error
		if err != nil {
			return nil, err
		}
	} else {
		err := r.db.Where("ext_id = ?", extID).First(&party).Error
		if err != nil {
			return nil, err
		}
	}

	return r.GetByID(party.ID)
}

// CountPartiesByType gets a count of parties by type
func (r *PartyRepository) CountPartiesByType(partyType string) (int64, error) {
	var count int64
	err := r.db.Model(&model.Individual{}).Where("paty_type = ?", partyType).Count(&count).Error
	return count, err
}
