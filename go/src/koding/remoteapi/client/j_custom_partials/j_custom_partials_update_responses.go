package j_custom_partials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// JCustomPartialsUpdateReader is a Reader for the JCustomPartialsUpdate structure.
type JCustomPartialsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JCustomPartialsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewJCustomPartialsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJCustomPartialsUpdateOK creates a JCustomPartialsUpdateOK with default headers values
func NewJCustomPartialsUpdateOK() *JCustomPartialsUpdateOK {
	return &JCustomPartialsUpdateOK{}
}

/*JCustomPartialsUpdateOK handles this case with default header values.

OK
*/
type JCustomPartialsUpdateOK struct {
	Payload JCustomPartialsUpdateOKBody
}

func (o *JCustomPartialsUpdateOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JCustomPartials.update/{id}][%d] jCustomPartialsUpdateOK  %+v", 200, o.Payload)
}

func (o *JCustomPartialsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*JCustomPartialsUpdateOKBody j custom partials update o k body
swagger:model JCustomPartialsUpdateOKBody
*/
type JCustomPartialsUpdateOKBody struct {
	models.JCustomPartials

	models.DefaultResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *JCustomPartialsUpdateOKBody) UnmarshalJSON(raw []byte) error {

	var jCustomPartialsUpdateOKBodyAO0 models.JCustomPartials
	if err := swag.ReadJSON(raw, &jCustomPartialsUpdateOKBodyAO0); err != nil {
		return err
	}
	o.JCustomPartials = jCustomPartialsUpdateOKBodyAO0

	var jCustomPartialsUpdateOKBodyAO1 models.DefaultResponse
	if err := swag.ReadJSON(raw, &jCustomPartialsUpdateOKBodyAO1); err != nil {
		return err
	}
	o.DefaultResponse = jCustomPartialsUpdateOKBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o JCustomPartialsUpdateOKBody) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	jCustomPartialsUpdateOKBodyAO0, err := swag.WriteJSON(o.JCustomPartials)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jCustomPartialsUpdateOKBodyAO0)

	jCustomPartialsUpdateOKBodyAO1, err := swag.WriteJSON(o.DefaultResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jCustomPartialsUpdateOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this j custom partials update o k body
func (o *JCustomPartialsUpdateOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.JCustomPartials.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.DefaultResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
