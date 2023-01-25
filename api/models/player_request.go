// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PlayerRequest player request
//
// swagger:model PlayerRequest
type PlayerRequest struct {

	// fleet
	// Required: true
	Fleet *int64 `json:"fleet"`

	// key
	// Required: true
	// Pattern: sr-.*
	Key *string `json:"key"`

	// party
	// Required: true
	// Enum: [attackers defenders]
	Party *string `json:"party"`
}

// Validate validates this player request
func (m *PlayerRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFleet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParty(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlayerRequest) validateFleet(formats strfmt.Registry) error {

	if err := validate.Required("fleet", "body", m.Fleet); err != nil {
		return err
	}

	return nil
}

func (m *PlayerRequest) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", m.Key); err != nil {
		return err
	}

	if err := validate.Pattern("key", "body", *m.Key, `sr-.*`); err != nil {
		return err
	}

	return nil
}

var playerRequestTypePartyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["attackers","defenders"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		playerRequestTypePartyPropEnum = append(playerRequestTypePartyPropEnum, v)
	}
}

const (

	// PlayerRequestPartyAttackers captures enum value "attackers"
	PlayerRequestPartyAttackers string = "attackers"

	// PlayerRequestPartyDefenders captures enum value "defenders"
	PlayerRequestPartyDefenders string = "defenders"
)

// prop value enum
func (m *PlayerRequest) validatePartyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, playerRequestTypePartyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PlayerRequest) validateParty(formats strfmt.Registry) error {

	if err := validate.Required("party", "body", m.Party); err != nil {
		return err
	}

	// value enum
	if err := m.validatePartyEnum("party", "body", *m.Party); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this player request based on context it is used
func (m *PlayerRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PlayerRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlayerRequest) UnmarshalBinary(b []byte) error {
	var res PlayerRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}