package roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/spiegela/gorackhd/models"
)

// ModifyRoleReader is a Reader for the ModifyRole structure.
type ModifyRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModifyRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewModifyRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewModifyRoleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewModifyRoleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewModifyRoleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewModifyRoleOK creates a ModifyRoleOK with default headers values
func NewModifyRoleOK() *ModifyRoleOK {
	return &ModifyRoleOK{}
}

/*ModifyRoleOK handles this case with default header values.

Successfully modified the role
*/
type ModifyRoleOK struct {
	Payload ModifyRoleOKBody
}

func (o *ModifyRoleOK) Error() string {
	return fmt.Sprintf("[PATCH /roles/{name}][%d] modifyRoleOK  %+v", 200, o.Payload)
}

func (o *ModifyRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewModifyRoleUnauthorized creates a ModifyRoleUnauthorized with default headers values
func NewModifyRoleUnauthorized() *ModifyRoleUnauthorized {
	return &ModifyRoleUnauthorized{}
}

/*ModifyRoleUnauthorized handles this case with default header values.

Unauthorized
*/
type ModifyRoleUnauthorized struct {
	Payload ModifyRoleUnauthorizedBody
}

func (o *ModifyRoleUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /roles/{name}][%d] modifyRoleUnauthorized  %+v", 401, o.Payload)
}

func (o *ModifyRoleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewModifyRoleForbidden creates a ModifyRoleForbidden with default headers values
func NewModifyRoleForbidden() *ModifyRoleForbidden {
	return &ModifyRoleForbidden{}
}

/*ModifyRoleForbidden handles this case with default header values.

Forbidden
*/
type ModifyRoleForbidden struct {
	Payload ModifyRoleForbiddenBody
}

func (o *ModifyRoleForbidden) Error() string {
	return fmt.Sprintf("[PATCH /roles/{name}][%d] modifyRoleForbidden  %+v", 403, o.Payload)
}

func (o *ModifyRoleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewModifyRoleDefault creates a ModifyRoleDefault with default headers values
func NewModifyRoleDefault(code int) *ModifyRoleDefault {
	return &ModifyRoleDefault{
		_statusCode: code,
	}
}

/*ModifyRoleDefault handles this case with default header values.

Unexpected error
*/
type ModifyRoleDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the modify role default response
func (o *ModifyRoleDefault) Code() int {
	return o._statusCode
}

func (o *ModifyRoleDefault) Error() string {
	return fmt.Sprintf("[PATCH /roles/{name}][%d] modifyRole default  %+v", o._statusCode, o.Payload)
}

func (o *ModifyRoleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ModifyRoleForbiddenBody modify role forbidden body
swagger:model ModifyRoleForbiddenBody
*/
type ModifyRoleForbiddenBody interface{}

/*ModifyRoleOKBody modify role o k body
swagger:model ModifyRoleOKBody
*/
type ModifyRoleOKBody interface{}

/*ModifyRoleUnauthorizedBody modify role unauthorized body
swagger:model ModifyRoleUnauthorizedBody
*/
type ModifyRoleUnauthorizedBody interface{}