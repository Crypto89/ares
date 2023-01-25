// Code generated by go-swagger; DO NOT EDIT.

package player

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetPlayerCreatedCode is the HTTP code returned for type GetPlayerCreated
const GetPlayerCreatedCode int = 201

/*
GetPlayerCreated get player information

swagger:response getPlayerCreated
*/
type GetPlayerCreated struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewGetPlayerCreated creates GetPlayerCreated with default headers values
func NewGetPlayerCreated() *GetPlayerCreated {

	return &GetPlayerCreated{}
}

// WithPayload adds the payload to the get player created response
func (o *GetPlayerCreated) WithPayload(payload interface{}) *GetPlayerCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get player created response
func (o *GetPlayerCreated) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPlayerCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetPlayerBadRequestCode is the HTTP code returned for type GetPlayerBadRequest
const GetPlayerBadRequestCode int = 400

/*
GetPlayerBadRequest bad request

swagger:response getPlayerBadRequest
*/
type GetPlayerBadRequest struct {
}

// NewGetPlayerBadRequest creates GetPlayerBadRequest with default headers values
func NewGetPlayerBadRequest() *GetPlayerBadRequest {

	return &GetPlayerBadRequest{}
}

// WriteResponse to the client
func (o *GetPlayerBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetPlayerNotFoundCode is the HTTP code returned for type GetPlayerNotFound
const GetPlayerNotFoundCode int = 404

/*
GetPlayerNotFound report not found

swagger:response getPlayerNotFound
*/
type GetPlayerNotFound struct {
}

// NewGetPlayerNotFound creates GetPlayerNotFound with default headers values
func NewGetPlayerNotFound() *GetPlayerNotFound {

	return &GetPlayerNotFound{}
}

// WriteResponse to the client
func (o *GetPlayerNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetPlayerInternalServerErrorCode is the HTTP code returned for type GetPlayerInternalServerError
const GetPlayerInternalServerErrorCode int = 500

/*
GetPlayerInternalServerError Internal Server Error

swagger:response getPlayerInternalServerError
*/
type GetPlayerInternalServerError struct {
}

// NewGetPlayerInternalServerError creates GetPlayerInternalServerError with default headers values
func NewGetPlayerInternalServerError() *GetPlayerInternalServerError {

	return &GetPlayerInternalServerError{}
}

// WriteResponse to the client
func (o *GetPlayerInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}

// GetPlayerServiceUnavailableCode is the HTTP code returned for type GetPlayerServiceUnavailable
const GetPlayerServiceUnavailableCode int = 503

/*
GetPlayerServiceUnavailable Service Unavailable

swagger:response getPlayerServiceUnavailable
*/
type GetPlayerServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetPlayerServiceUnavailable creates GetPlayerServiceUnavailable with default headers values
func NewGetPlayerServiceUnavailable() *GetPlayerServiceUnavailable {

	return &GetPlayerServiceUnavailable{}
}

// WithPayload adds the payload to the get player service unavailable response
func (o *GetPlayerServiceUnavailable) WithPayload(payload string) *GetPlayerServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get player service unavailable response
func (o *GetPlayerServiceUnavailable) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPlayerServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
