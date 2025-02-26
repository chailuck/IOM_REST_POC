package service

import (
	"encoding/json"
	"errors"
	"strconv"

	"party/internal/model"

	iom "gitlab.com/ft25/iom/framework/rabbitmq"
)

// @Summary Create_Individual
// @Description Create a new individual party
// @Accept  json
// @Produce json
// @Param   party body model.Individual true "Individual party details"
// @Success 201 {object} model.PartyCreateResponse
// @Failure 400 {object} model.Failure
// @Failure 409 {object} model.Failure
// @Failure 500 {object} model.Failure
//
// @Router /am/api/token-service/V1/tokens/:id [put]
// func (s RestService) PUT_V1_Tokens_Groups_ByGroups_Id_ById(ctx iom.Context, serviceRequest interface{}) (serviceResponse interface{}, err error) {

// @Router /partyManagement/v1/individual [post]
func (s RestService) POST_V1_Individual(ctx iom.Context, serviceRequest interface{}) (serviceResponse interface{}, err error) {
	//func (s RestService) Create_Individual(ctx iom.Context, serviceRequest interface{}) (serviceResponse interface{}, err error) {
	var req model.Individual
	err = json.Unmarshal(serviceRequest.([]byte), &req)
	if err != nil {
		return nil, errors.New("invalid request format")
	}

	res, err := PartyService.CreateParty(&req)
	if err != nil {
		return nil, err
	}

	return &model.PartyCreateResponse{Data: res}, nil
}

// @Summary Get_Individual
// @Description Get an individual party by ID
// @Accept  json
// @Produce json
// @Param   id path string true "Individual ID"
// @Success 200 {object} model.Individual
// @Failure 404 {object} model.Failure
// @Failure 500 {object} model.Failure
// @Router /partyManagement/v1/individual/{id} [get]
// func (s RestService) Get_Individual_ById(ctx iom.Context, serviceRequest interface{}) (serviceResponse interface{}, err error) {
func (s RestService) GET_V1_Individual_Id_ById(ctx iom.Context, serviceRequest interface{}) (serviceResponse interface{}, err error) {
	id := ctx.Routing.Params("id")

	res, err := PartyService.GetParty(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// @Summary Update_Individual
// @Description Update an existing individual party
// @Accept  json
// @Produce json
// @Param   id path string true "Individual ID"
// @Param   party body model.Individual true "Updated individual details"
// @Success 200 {object} model.PartyUpdateResponse
// @Failure 400 {object} model.Failure
// @Failure 404 {object} model.Failure
// @Failure 500 {object} model.Failure
// @Router /partyManagement/v1/individual/{id} [put]
// func (s RestService) Update_Individual_ById(ctx iom.Context, serviceRequest interface{}) (serviceResponse interface{}, err error) {
func (s RestService) PATCH_V1_Individual_Id_ById(ctx iom.Context, serviceRequest interface{}) (serviceResponse interface{}, err error) {
	id := ctx.Routing.Params("id")

	var req model.Individual
	err = json.Unmarshal(serviceRequest.([]byte), &req)
	if err != nil {
		return nil, errors.New("invalid request format")
	}

	res, err := PartyService.UpdateParty(id, &req)
	if err != nil {
		return nil, err
	}

	return &model.PartyUpdateResponse{Data: res}, nil
}

// @Summary Delete_Individual
// @Description Delete an individual party
// @Accept  json
// @Produce json
// @Param   id path string true "Individual ID"
// @Success 204 "No Content"
// @Failure 404 {object} model.Failure
// @Failure 500 {object} model.Failure
// @Router /partyManagement/v1/individual/{id} [delete]
func (s RestService) DELETE_V1_Individual_Id_ById(ctx iom.Context, serviceRequest interface{}) (serviceResponse interface{}, err error) {
	id := ctx.Routing.Params("id")

	err = PartyService.DeleteParty(id)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// @Summary List_Individual
// @Description List individual parties with filtering
// @Accept  json
// @Produce json
// @Param   idType query string false "ID Type filter"
// @Param   idNumber query string false "ID Number filter"
// @Param   firstName query string false "First Name filter"
// @Param   lastName query string false "Last Name filter"
// @Param   status query string false "Status filter"
// @Param   limit query int false "Limit results" default(50)
// @Param   offset query int false "Offset results" default(0)
// @Success 200 {array} model.Individual
// @Failure 400 {object} model.Failure
// @Failure 500 {object} model.Failure
// @Router /partyManagement/v1/individual [get]
func (s RestService) List_Individual(ctx iom.Context, serviceRequest interface{}) (serviceResponse interface{}, err error) {
	params := model.PartyQueryParams{
		IDType:    ctx.Routing.Query("idType"),
		IDNumber:  ctx.Routing.Query("idNumber"),
		FirstName: ctx.Routing.Query("firstName"),
		LastName:  ctx.Routing.Query("lastName"),
		Status:    ctx.Routing.Query("status"),
	}

	// Parse pagination parameters
	if limitStr := ctx.Routing.Query("limit"); limitStr != "" {
		params.Limit, _ = strconv.Atoi(limitStr)
	}
	if offsetStr := ctx.Routing.Query("offset"); offsetStr != "" {
		params.Offset, _ = strconv.Atoi(offsetStr)
	}

	return PartyService.ListParties(params)
}

// @Summary Patch_Individual
// @Description Partially update an individual party
// @Accept  json
// @Produce json
// @Param   id path string true "Individual ID"
// @Param   party body model.Individual true "Partial individual details"
// @Success 200 {object} model.PartyUpdateResponse
// @Failure 400 {object} model.Failure
// @Failure 404 {object} model.Failure
// @Failure 500 {object} model.Failure
// @Router /partyManagement/v1/individual/{id} [patch]
func (s RestService) Patch_Individual_ById(ctx iom.Context, serviceRequest interface{}) (serviceResponse interface{}, err error) {
	id := ctx.Routing.Params("id")

	var req model.Individual
	err = json.Unmarshal(serviceRequest.([]byte), &req)
	if err != nil {
		return nil, errors.New("invalid request format")
	}

	// Get existing party
	existing, err := PartyService.GetParty(id)
	if err != nil {
		return nil, err
	}

	// Merge changes
	if req.FirstName != "" {
		existing.FirstName = req.FirstName
	}
	if req.LastName != "" {
		existing.LastName = req.LastName
	}
	// Add other field merges as needed

	res, err := PartyService.UpdateParty(id, existing)
	if err != nil {
		return nil, err
	}

	return &model.PartyUpdateResponse{Data: res}, nil
}
