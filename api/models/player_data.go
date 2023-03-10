// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PlayerData player data
//
// swagger:model PlayerData
type PlayerData struct {

	// alliance
	Alliance *PlayerDataAlliance `json:"alliance,omitempty"`

	// buildings
	Buildings LevelableMap `json:"buildings,omitempty"`

	// class
	Class int64 `json:"class,omitempty"`

	// defence
	Defence CountableMap `json:"defence,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// planet
	Planet *Planet `json:"planet,omitempty"`

	// research
	Research LevelableMap `json:"research,omitempty"`

	// resources
	Resources *PlayerDataResources `json:"resources,omitempty"`

	// ships
	Ships CountableMap `json:"ships,omitempty"`
}

// Validate validates this player data
func (m *PlayerData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlliance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBuildings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefence(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlanet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResearch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShips(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlayerData) validateAlliance(formats strfmt.Registry) error {
	if swag.IsZero(m.Alliance) { // not required
		return nil
	}

	if m.Alliance != nil {
		if err := m.Alliance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("alliance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("alliance")
			}
			return err
		}
	}

	return nil
}

func (m *PlayerData) validateBuildings(formats strfmt.Registry) error {
	if swag.IsZero(m.Buildings) { // not required
		return nil
	}

	if m.Buildings != nil {
		if err := m.Buildings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("buildings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("buildings")
			}
			return err
		}
	}

	return nil
}

func (m *PlayerData) validateDefence(formats strfmt.Registry) error {
	if swag.IsZero(m.Defence) { // not required
		return nil
	}

	if m.Defence != nil {
		if err := m.Defence.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("defence")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("defence")
			}
			return err
		}
	}

	return nil
}

func (m *PlayerData) validatePlanet(formats strfmt.Registry) error {
	if swag.IsZero(m.Planet) { // not required
		return nil
	}

	if m.Planet != nil {
		if err := m.Planet.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("planet")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("planet")
			}
			return err
		}
	}

	return nil
}

func (m *PlayerData) validateResearch(formats strfmt.Registry) error {
	if swag.IsZero(m.Research) { // not required
		return nil
	}

	if m.Research != nil {
		if err := m.Research.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("research")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("research")
			}
			return err
		}
	}

	return nil
}

func (m *PlayerData) validateResources(formats strfmt.Registry) error {
	if swag.IsZero(m.Resources) { // not required
		return nil
	}

	if m.Resources != nil {
		if err := m.Resources.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resources")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resources")
			}
			return err
		}
	}

	return nil
}

func (m *PlayerData) validateShips(formats strfmt.Registry) error {
	if swag.IsZero(m.Ships) { // not required
		return nil
	}

	if m.Ships != nil {
		if err := m.Ships.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ships")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ships")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this player data based on the context it is used
func (m *PlayerData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAlliance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBuildings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDefence(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePlanet(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResearch(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShips(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlayerData) contextValidateAlliance(ctx context.Context, formats strfmt.Registry) error {

	if m.Alliance != nil {
		if err := m.Alliance.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("alliance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("alliance")
			}
			return err
		}
	}

	return nil
}

func (m *PlayerData) contextValidateBuildings(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Buildings.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("buildings")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("buildings")
		}
		return err
	}

	return nil
}

func (m *PlayerData) contextValidateDefence(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Defence.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("defence")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("defence")
		}
		return err
	}

	return nil
}

func (m *PlayerData) contextValidatePlanet(ctx context.Context, formats strfmt.Registry) error {

	if m.Planet != nil {
		if err := m.Planet.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("planet")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("planet")
			}
			return err
		}
	}

	return nil
}

func (m *PlayerData) contextValidateResearch(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Research.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("research")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("research")
		}
		return err
	}

	return nil
}

func (m *PlayerData) contextValidateResources(ctx context.Context, formats strfmt.Registry) error {

	if m.Resources != nil {
		if err := m.Resources.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resources")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resources")
			}
			return err
		}
	}

	return nil
}

func (m *PlayerData) contextValidateShips(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Ships.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ships")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ships")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PlayerData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlayerData) UnmarshalBinary(b []byte) error {
	var res PlayerData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PlayerDataAlliance player data alliance
//
// swagger:model PlayerDataAlliance
type PlayerDataAlliance struct {

	// name
	// Required: true
	Name *string `json:"name"`

	// tag
	// Required: true
	Tag *string `json:"tag"`
}

// Validate validates this player data alliance
func (m *PlayerDataAlliance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTag(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlayerDataAlliance) validateName(formats strfmt.Registry) error {

	if err := validate.Required("alliance"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *PlayerDataAlliance) validateTag(formats strfmt.Registry) error {

	if err := validate.Required("alliance"+"."+"tag", "body", m.Tag); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this player data alliance based on context it is used
func (m *PlayerDataAlliance) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PlayerDataAlliance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlayerDataAlliance) UnmarshalBinary(b []byte) error {
	var res PlayerDataAlliance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PlayerDataResources player data resources
//
// swagger:model PlayerDataResources
type PlayerDataResources struct {

	// crystal
	Crystal int64 `json:"crystal,omitempty"`

	// deuterium
	Deuterium int64 `json:"deuterium,omitempty"`

	// energy
	Energy int64 `json:"energy,omitempty"`

	// metal
	Metal int64 `json:"metal,omitempty"`
}

// Validate validates this player data resources
func (m *PlayerDataResources) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this player data resources based on context it is used
func (m *PlayerDataResources) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PlayerDataResources) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlayerDataResources) UnmarshalBinary(b []byte) error {
	var res PlayerDataResources
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
