// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SourceBlock source block
//
// swagger:model SourceBlock
type SourceBlock struct {

	// chain Id
	ChainID int64 `json:"chainId,omitempty"`

	// hash
	Hash string `json:"hash,omitempty"`

	// number
	Number int64 `json:"number,omitempty"`
}

// Validate validates this source block
func (m *SourceBlock) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this source block based on context it is used
func (m *SourceBlock) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SourceBlock) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SourceBlock) UnmarshalBinary(b []byte) error {
	var res SourceBlock
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}