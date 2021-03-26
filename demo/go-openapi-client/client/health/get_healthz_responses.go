// Code generated by go-swagger; DO NOT EDIT.

package health

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetHealthzReader is a Reader for the GetHealthz structure.
type GetHealthzReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHealthzReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewGetHealthzDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewGetHealthzDefault creates a GetHealthzDefault with default headers values
func NewGetHealthzDefault(code int) *GetHealthzDefault {
	return &GetHealthzDefault{
		_statusCode: code,
	}
}

/* GetHealthzDefault describes a response with status code -1, with default header values.

OK
*/
type GetHealthzDefault struct {
	_statusCode int
}

// Code gets the status code for the get healthz default response
func (o *GetHealthzDefault) Code() int {
	return o._statusCode
}

func (o *GetHealthzDefault) Error() string {
	return fmt.Sprintf("[GET /healthz][%d] GetHealthz default ", o._statusCode)
}

func (o *GetHealthzDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}