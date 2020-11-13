// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    eventTypes, err := UnmarshalEventTypes(bytes)
//    bytes, err = eventTypes.Marshal()
//
//    metadata, err := UnmarshalMetadata(bytes)
//    bytes, err = metadata.Marshal()
//
//    registeredUserEvent, err := UnmarshalRegisteredUserEvent(bytes)
//    bytes, err = registeredUserEvent.Marshal()
//
//    changeUserEvent, err := UnmarshalChangeUserEvent(bytes)
//    bytes, err = changeUserEvent.Marshal()
//
//    userChangedEvent, err := UnmarshalUserChangedEvent(bytes)
//    bytes, err = userChangedEvent.Marshal()

package main

import "encoding/json"

func UnmarshalEventTypes(data []byte) (EventTypes, error) {
	var r EventTypes
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *EventTypes) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMetadata(data []byte) (Metadata, error) {
	var r Metadata
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Metadata) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalRegisteredUserEvent(data []byte) (RegisteredUserEvent, error) {
	var r RegisteredUserEvent
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RegisteredUserEvent) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalChangeUserEvent(data []byte) (ChangeUserEvent, error) {
	var r ChangeUserEvent
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ChangeUserEvent) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUserChangedEvent(data []byte) (UserChangedEvent, error) {
	var r UserChangedEvent
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UserChangedEvent) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type RegisteredUserEvent struct {
	Data     RegisteredUserEventData `json:"data"`    
	ID       string                  `json:"id"`      
	Metadata Metadata                `json:"metadata"`
	Type     RegisteredUserEventType `json:"type"`    
}

type RegisteredUserEventData struct {
	Auth0ID       string `json:"auth0Id"`      
	Email         string `json:"email"`        
	EmailVerified bool   `json:"emailVerified"`
	FirstName     string `json:"firstName"`    
	ID            string `json:"id"`           
	LastName      string `json:"lastName"`     
	Locale        string `json:"locale"`       
	Picture       string `json:"picture"`      
}

type Metadata struct {
	TraceID string `json:"traceId"`
	UserID  string `json:"userId"` 
}

type ChangeUserEvent struct {
	Data     ChangeUserEventData `json:"data"`    
	ID       string              `json:"id"`      
	Metadata Metadata            `json:"metadata"`
	Type     ChangeUserEventType `json:"type"`    
}

type ChangeUserEventData struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"` 
	Locale    string `json:"locale"`   
	Picture   string `json:"picture"`  
}

type UserChangedEvent struct {
	Data     UserChangedEventData `json:"data"`    
	ID       string               `json:"id"`      
	Metadata Metadata             `json:"metadata"`
	Type     UserChangedEventType `json:"type"`    
}

type UserChangedEventData struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"` 
	Locale    string `json:"locale"`   
	Picture   string `json:"picture"`  
}

type EventTypes string
const (
	EventTypesChangeUser EventTypes = "ChangeUser"
	EventTypesRegisteredUser EventTypes = "RegisteredUser"
	EventTypesUserChanged EventTypes = "UserChanged"
)

type RegisteredUserEventType string
const (
	TypeRegisteredUser RegisteredUserEventType = "RegisteredUser"
)

type ChangeUserEventType string
const (
	TypeChangeUser ChangeUserEventType = "ChangeUser"
)

type UserChangedEventType string
const (
	TypeUserChanged UserChangedEventType = "UserChanged"
)
