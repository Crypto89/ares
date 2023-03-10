// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Levelable levelable
//
// swagger:model Levelable
type Levelable struct {

	// level
	Level int64 `json:"level,omitempty"`

	// technology
	Technology TechnologyInfo `json:"technology,omitempty"`
}

// Validate validates this levelable
func (m *Levelable) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this levelable based on context it is used
func (m *Levelable) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Levelable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Levelable) UnmarshalBinary(b []byte) error {
	var res Levelable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
