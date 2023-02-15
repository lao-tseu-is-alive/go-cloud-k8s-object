// Package objects provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package objects

import (
	"time"
)

const (
	JWTAuthScopes = "JWTAuth.Scopes"
)

// Error defines model for Error.
type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// Object defines model for Object.
type Object struct {
	Comment              *string    `json:"comment,omitempty"`
	CreateTime           time.Time  `json:"create_time"`
	Creator              int32      `json:"creator"`
	Description          *string    `json:"description,omitempty"`
	ExternalId           *int32     `json:"external_id,omitempty"`
	Id                   int32      `json:"id"`
	IdValidator          *int32     `json:"id_validator,omitempty"`
	InactivationReason   *string    `json:"inactivation_reason,omitempty"`
	InactivationTime     *time.Time `json:"inactivation_time,omitempty"`
	IsActive             bool       `json:"is_active"`
	IsValidated          *bool      `json:"is_validated,omitempty"`
	LastModificationTime *time.Time `json:"last_modification_time,omitempty"`
	LastModificationUser *int32     `json:"last_modification_user,omitempty"`
	Name                 string     `json:"name"`
	TypeObjectId         int32      `json:"type_object_id"`
}

// ObjectList defines model for ObjectList.
type ObjectList struct {
	CreateTime   time.Time `json:"create_time"`
	Creator      int32     `json:"creator"`
	Description  *string   `json:"description,omitempty"`
	ExternalId   *int32    `json:"external_id,omitempty"`
	Id           int32     `json:"id"`
	IsActive     bool      `json:"is_active"`
	Name         string    `json:"name"`
	TypeObjectId int32     `json:"type_object_id"`
}

// TypeObject defines model for TypeObject.
type TypeObject struct {
	Comment              *string    `json:"comment,omitempty"`
	CreateTime           time.Time  `json:"create_time"`
	Creator              int32      `json:"creator"`
	Id                   int32      `json:"id"`
	InactivationReason   *string    `json:"inactivation_reason,omitempty"`
	InactivationTime     *time.Time `json:"inactivation_time,omitempty"`
	IsActive             bool       `json:"is_active"`
	LastModificationTime *time.Time `json:"last_modification_time,omitempty"`
	LastModificationUser *int32     `json:"last_modification_user,omitempty"`
	Name                 string     `json:"name"`
}

// TypeObjectList defines model for TypeObjectList.
type TypeObjectList struct {
	Id       int32  `json:"id"`
	IsActive bool   `json:"is_active"`
	Name     string `json:"name"`
}

// ListParams defines parameters for List.
type ListParams struct {
	// maximum number of results to return
	Limit *int32 `form:"limit,omitempty" json:"limit,omitempty"`
}

// CreateJSONBody defines parameters for Create.
type CreateJSONBody = Object

// ListByTypeParams defines parameters for ListByType.
type ListByTypeParams struct {
	// maximum number of results to return
	Limit *int32 `form:"limit,omitempty" json:"limit,omitempty"`
}

// TypeObjectListParams defines parameters for TypeObjectList.
type TypeObjectListParams struct {
	// maximum number of results to return
	Limit *int32 `form:"limit,omitempty" json:"limit,omitempty"`
}

// TypeObjectCreateJSONBody defines parameters for TypeObjectCreate.
type TypeObjectCreateJSONBody = TypeObject

// CreateJSONRequestBody defines body for Create for application/json ContentType.
type CreateJSONRequestBody = CreateJSONBody

// TypeObjectCreateJSONRequestBody defines body for TypeObjectCreate for application/json ContentType.
type TypeObjectCreateJSONRequestBody = TypeObjectCreateJSONBody
