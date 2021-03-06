// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// CreateUserBadRequestCode is the HTTP code returned for type CreateUserBadRequest
const CreateUserBadRequestCode int = 400

/*CreateUserBadRequest Error creating user

swagger:response createUserBadRequest
*/
type CreateUserBadRequest struct {
}

// NewCreateUserBadRequest creates CreateUserBadRequest with default headers values
func NewCreateUserBadRequest() *CreateUserBadRequest {

	return &CreateUserBadRequest{}
}

// WriteResponse to the client
func (o *CreateUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// CreateUserConflictCode is the HTTP code returned for type CreateUserConflict
const CreateUserConflictCode int = 409

/*CreateUserConflict User exists

swagger:response createUserConflict
*/
type CreateUserConflict struct {
}

// NewCreateUserConflict creates CreateUserConflict with default headers values
func NewCreateUserConflict() *CreateUserConflict {

	return &CreateUserConflict{}
}

// WriteResponse to the client
func (o *CreateUserConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(409)
}

/*CreateUserDefault successful operation

swagger:response createUserDefault
*/
type CreateUserDefault struct {
	_statusCode int
}

// NewCreateUserDefault creates CreateUserDefault with default headers values
func NewCreateUserDefault(code int) *CreateUserDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateUserDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create user default response
func (o *CreateUserDefault) WithStatusCode(code int) *CreateUserDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create user default response
func (o *CreateUserDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WriteResponse to the client
func (o *CreateUserDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(o._statusCode)
}
