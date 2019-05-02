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

// RejectConsentRequestReader is a Reader for the RejectConsentRequest structure.
type RejectConsentRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RejectConsentRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRejectConsentRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewRejectConsentRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewRejectConsentRequestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRejectConsentRequestOK creates a RejectConsentRequestOK with default headers values
func NewRejectConsentRequestOK() *RejectConsentRequestOK {
	return &RejectConsentRequestOK{}
}

/*RejectConsentRequestOK handles this case with default header values.

completedRequest
*/
type RejectConsentRequestOK struct {
	Payload *models.RequestHandlerResponse
}

func (o *RejectConsentRequestOK) Error() string {
	return fmt.Sprintf("[PUT /oauth2/auth/requests/consent/reject][%d] rejectConsentRequestOK  %+v", 200, o.Payload)
}

func (o *RejectConsentRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestHandlerResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRejectConsentRequestNotFound creates a RejectConsentRequestNotFound with default headers values
func NewRejectConsentRequestNotFound() *RejectConsentRequestNotFound {
	return &RejectConsentRequestNotFound{}
}

/*RejectConsentRequestNotFound handles this case with default header values.

genericError
*/
type RejectConsentRequestNotFound struct {
	Payload *models.GenericError
}

func (o *RejectConsentRequestNotFound) Error() string {
	return fmt.Sprintf("[PUT /oauth2/auth/requests/consent/reject][%d] rejectConsentRequestNotFound  %+v", 404, o.Payload)
}

func (o *RejectConsentRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRejectConsentRequestInternalServerError creates a RejectConsentRequestInternalServerError with default headers values
func NewRejectConsentRequestInternalServerError() *RejectConsentRequestInternalServerError {
	return &RejectConsentRequestInternalServerError{}
}

/*RejectConsentRequestInternalServerError handles this case with default header values.

genericError
*/
type RejectConsentRequestInternalServerError struct {
	Payload *models.GenericError
}

func (o *RejectConsentRequestInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /oauth2/auth/requests/consent/reject][%d] rejectConsentRequestInternalServerError  %+v", 500, o.Payload)
}

func (o *RejectConsentRequestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
