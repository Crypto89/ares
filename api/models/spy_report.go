// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SpyReport spy report
//
// swagger:model SpyReport
type SpyReport struct {

	// activity
	Activity int64 `json:"activity,omitempty"`

	// attacker
	Attacker *PlayerData `json:"attacker,omitempty"`

	// defender
	Defender *PlayerData `json:"defender,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// loot percentage
	LootPercentage int64 `json:"loot_percentage,omitempty"`

	// spy fail chance
	SpyFailChance int64 `json:"spy_fail_chance,omitempty"`

	// time
	Time string `json:"time,omitempty"`

	// timestamp
	Timestamp int64 `json:"timestamp,omitempty"`

	// total defence count
	TotalDefenceCount int64 `json:"total_defence_count,omitempty"`

	// total ship count
	TotalShipCount int64 `json:"total_ship_count,omitempty"`
}

// Validate validates this spy report
func (m *SpyReport) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttacker(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefender(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SpyReport) validateAttacker(formats strfmt.Registry) error {
	if swag.IsZero(m.Attacker) { // not required
		return nil
	}

	if m.Attacker != nil {
		if err := m.Attacker.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attacker")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attacker")
			}
			return err
		}
	}

	return nil
}

func (m *SpyReport) validateDefender(formats strfmt.Registry) error {
	if swag.IsZero(m.Defender) { // not required
		return nil
	}

	if m.Defender != nil {
		if err := m.Defender.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("defender")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("defender")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this spy report based on the context it is used
func (m *SpyReport) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttacker(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDefender(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SpyReport) contextValidateAttacker(ctx context.Context, formats strfmt.Registry) error {

	if m.Attacker != nil {
		if err := m.Attacker.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attacker")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attacker")
			}
			return err
		}
	}

	return nil
}

func (m *SpyReport) contextValidateDefender(ctx context.Context, formats strfmt.Registry) error {

	if m.Defender != nil {
		if err := m.Defender.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("defender")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("defender")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SpyReport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SpyReport) UnmarshalBinary(b []byte) error {
	var res SpyReport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}