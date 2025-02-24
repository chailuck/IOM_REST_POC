package repository

import (
	"internal/model"
)

// List retrieves parties based on query parameters
func (r *PartyRepository) List(params model.PartyQueryParams) ([]model.Individual, error) {
	var parties []model.Individual
	query := r.db.Model(&model.Individual{})

	// Apply filters
	if params.IDType != "" {
		query = query.Where("id_type = ?", params.IDType)
	}
	if params.IDNumber != "" {
		query = query.Where("id_numb = ?", params.IDNumber)
	}
	if params.FirstName != "" {
		query = query.Where("frst_name ILIKE ?", "%"+params.FirstName+"%")
	}
	if params.LastName != "" {
		query = query.Where("last_name ILIKE ?", "%"+params.LastName+"%")
	}
	if params.Status != "" {
		query = query.Where("paty_type = ?", params.Status)
	}

	// Apply pagination
	if params.Limit > 0 {
		query = query.Limit(params.Limit)
	}
	if params.Offset > 0 {
		query = query.Offset(params.Offset)
	}

	// Execute query with preloaded relationships
	err := query.
		Preload("Address").
		Preload("Characteristic").
		Preload("ContactMedium").
		Find(&parties).Error

	if err != nil {
		return nil, err
	}

	return parties, nil
}

// GetByCharacteristic retrieves parties by characteristic name and value
func (r *PartyRepository) GetByCharacteristic(name, value string) ([]model.Individual, error) {
	var parties []model.Individual

	err := r.db.
		Joins("JOIN characteristics c ON c.paty_row_id = individuals.paty_row_id").
		Where("c.attr_name = ? AND c.attr_vlue = ?", name, value).
		Preload("Address").
		Preload("Characteristic").
		Preload("ContactMedium").
		Find(&parties).Error

	if err != nil {
		return nil, err
	}

	return parties, nil
}

// GetByAddress retrieves parties by address criteria
func (r *PartyRepository) GetByAddress(postalCode string, addressType string) ([]model.Individual, error) {
	var parties []model.Individual

	query := r.db.
		Joins("JOIN party_addresses pa ON pa.paty_row_id = individuals.paty_row_id")

	if postalCode != "" {
		query = query.Where("pa.addr_post_code = ?", postalCode)
	}
	if addressType != "" {
		query = query.Where("pa.addr_type = ?", addressType)
	}

	err := query.
		Preload("Address").
		Preload("Characteristic").
		Preload("ContactMedium").
		Find(&parties).Error

	if err != nil {
		return nil, err
	}

	return parties, nil
}

// UpdateStatus updates a party's status
func (r *PartyRepository) UpdateStatus(id string, status string) error {
	return r.db.Model(&model.Individual{}).
		Where("paty_row_id = ?", id).
		Update("paty_type", status).Error
}

// AddCharacteristic adds a new characteristic to a party
func (r *PartyRepository) AddCharacteristic(characteristic *model.Characteristic) error {
	return r.db.Create(characteristic).Error
}

// RemoveCharacteristic removes a characteristic from a party
func (r *PartyRepository) RemoveCharacteristic(partyID string, characteristicName string) error {
	return r.db.Where("paty_row_id = ? AND attr_name = ?", partyID, characteristicName).
		Delete(&model.Characteristic{}).Error
}
