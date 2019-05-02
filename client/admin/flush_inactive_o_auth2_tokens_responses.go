// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/OBrezhniev/hydra-sdk/models"
)

// FlushInactiveOAuth2TokensReader is a Reader for the FlushInactiveOAuth2Tokens structure.
type FlushInactiveOAuth2TokensReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FlushInactiveOAuth2TokensReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewFlushInactiveOAuth2TokensNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewFlushInactiveOAuth2TokensUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewFlushInactiveOAuth2TokensInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewFlushInactiveOAuth2TokensNoContent creates a FlushInactiveOAuth2TokensNoContent with default headers values
func NewFlushInactiveOAuth2TokensNoContent() *FlushInactiveOAuth2TokensNoContent {
	return &FlushInactiveOAuth2TokensNoContent{}
}

/*FlushInactiveOAuth2TokensNoContent handles this case with default header values.

Empty responses are sent when, for example, resources are deleted. The HTTP status code for empty responses is
typically 201.
*/
type FlushInactiveOAuth2TokensNoContent struct {
}

func (o *FlushInactiveOAuth2TokensNoContent) Error() string {
	return fmt.Sprintf("[POST /oauth2/flush][%d] flushInactiveOAuth2TokensNoContent ", 204)
}

func (o *FlushInactiveOAuth2TokensNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFlushInactiveOAuth2TokensUnauthorized creates a FlushInactiveOAuth2TokensUnauthorized with default headers values
func NewFlushInactiveOAuth2TokensUnauthorized() *FlushInactiveOAuth2TokensUnauthorized {
	return &FlushInactiveOAuth2TokensUnauthorized{}
}

/*FlushInactiveOAuth2TokensUnauthorized handles this case with default header values.

genericError
*/
type FlushInactiveOAuth2TokensUnauthorized struct {
	Payload *models.GenericError
}

func (o *FlushInactiveOAuth2TokensUnauthorized) Error() string {
	return fmt.Sprintf("[POST /oauth2/flush][%d] flushInactiveOAuth2TokensUnauthorized  %+v", 401, o.Payload)
}

func (o *FlushInactiveOAuth2TokensUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFlushInactiveOAuth2TokensInternalServerError creates a FlushInactiveOAuth2TokensInternalServerError with default headers values
func NewFlushInactiveOAuth2TokensInternalServerError() *FlushInactiveOAuth2TokensInternalServerError {
	return &FlushInactiveOAuth2TokensInternalServerError{}
}

/*FlushInactiveOAuth2TokensInternalServerError handles this case with default header values.

genericError
*/
type FlushInactiveOAuth2TokensInternalServerError struct {
	Payload *models.GenericError
}

func (o *FlushInactiveOAuth2TokensInternalServerError) Error() string {
	return fmt.Sprintf("[POST /oauth2/flush][%d] flushInactiveOAuth2TokensInternalServerError  %+v", 500, o.Payload)
}

func (o *FlushInactiveOAuth2TokensInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
