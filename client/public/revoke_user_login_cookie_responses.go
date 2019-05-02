// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/OBrezhniev/hydra-sdk/models"
)

// RevokeUserLoginCookieReader is a Reader for the RevokeUserLoginCookie structure.
type RevokeUserLoginCookieReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RevokeUserLoginCookieReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 302:
		result := NewRevokeUserLoginCookieFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewRevokeUserLoginCookieNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewRevokeUserLoginCookieInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRevokeUserLoginCookieFound creates a RevokeUserLoginCookieFound with default headers values
func NewRevokeUserLoginCookieFound() *RevokeUserLoginCookieFound {
	return &RevokeUserLoginCookieFound{}
}

/*RevokeUserLoginCookieFound handles this case with default header values.

Empty responses are sent when, for example, resources are deleted. The HTTP status code for empty responses is
typically 201.
*/
type RevokeUserLoginCookieFound struct {
}

func (o *RevokeUserLoginCookieFound) Error() string {
	return fmt.Sprintf("[GET /oauth2/auth/sessions/login/revoke][%d] revokeUserLoginCookieFound ", 302)
}

func (o *RevokeUserLoginCookieFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRevokeUserLoginCookieNotFound creates a RevokeUserLoginCookieNotFound with default headers values
func NewRevokeUserLoginCookieNotFound() *RevokeUserLoginCookieNotFound {
	return &RevokeUserLoginCookieNotFound{}
}

/*RevokeUserLoginCookieNotFound handles this case with default header values.

genericError
*/
type RevokeUserLoginCookieNotFound struct {
	Payload *models.GenericError
}

func (o *RevokeUserLoginCookieNotFound) Error() string {
	return fmt.Sprintf("[GET /oauth2/auth/sessions/login/revoke][%d] revokeUserLoginCookieNotFound  %+v", 404, o.Payload)
}

func (o *RevokeUserLoginCookieNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRevokeUserLoginCookieInternalServerError creates a RevokeUserLoginCookieInternalServerError with default headers values
func NewRevokeUserLoginCookieInternalServerError() *RevokeUserLoginCookieInternalServerError {
	return &RevokeUserLoginCookieInternalServerError{}
}

/*RevokeUserLoginCookieInternalServerError handles this case with default header values.

genericError
*/
type RevokeUserLoginCookieInternalServerError struct {
	Payload *models.GenericError
}

func (o *RevokeUserLoginCookieInternalServerError) Error() string {
	return fmt.Sprintf("[GET /oauth2/auth/sessions/login/revoke][%d] revokeUserLoginCookieInternalServerError  %+v", 500, o.Payload)
}

func (o *RevokeUserLoginCookieInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
