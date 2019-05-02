// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SwaggerFlushInactiveAccessTokens swagger flush inactive access tokens
// swagger:model swaggerFlushInactiveAccessTokens
type SwaggerFlushInactiveAccessTokens struct {

	// body
	Body *FlushInactiveOAuth2TokensRequest `json:"Body,omitempty"`
}

// Validate validates this swagger flush inactive access tokens
func (m *SwaggerFlushInactiveAccessTokens) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBody(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SwaggerFlushInactiveAccessTokens) validateBody(formats strfmt.Registry) error {

	if swag.IsZero(m.Body) { // not required
		return nil
	}

	if m.Body != nil {
		if err := m.Body.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Body")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SwaggerFlushInactiveAccessTokens) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SwaggerFlushInactiveAccessTokens) UnmarshalBinary(b []byte) error {
	var res SwaggerFlushInactiveAccessTokens
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
