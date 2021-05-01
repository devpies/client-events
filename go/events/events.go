// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    eventTypes, err := UnmarshalEventTypes(bytes)
//    bytes, err = eventTypes.Marshal()
//
//    message, err := UnmarshalMessage(bytes)
//    bytes, err = message.Marshal()
//
//    metadata, err := UnmarshalMetadata(bytes)
//    bytes, err = metadata.Marshal()
//
//    categories, err := UnmarshalCategories(bytes)
//    bytes, err = categories.Marshal()
//
//    commands, err := UnmarshalCommands(bytes)
//    bytes, err = commands.Marshal()
//
//    enableAccountingCommand, err := UnmarshalEnableAccountingCommand(bytes)
//    bytes, err = enableAccountingCommand.Marshal()
//
//    events, err := UnmarshalEvents(bytes)
//    bytes, err = events.Marshal()
//
//    membershipCreatedEvent, err := UnmarshalMembershipCreatedEvent(bytes)
//    bytes, err = membershipCreatedEvent.Marshal()
//
//    membershipCreatedForProjectEvent, err := UnmarshalMembershipCreatedForProjectEvent(bytes)
//    bytes, err = membershipCreatedForProjectEvent.Marshal()
//
//    membershipUpdatedEvent, err := UnmarshalMembershipUpdatedEvent(bytes)
//    bytes, err = membershipUpdatedEvent.Marshal()
//
//    membershipDeletedEvent, err := UnmarshalMembershipDeletedEvent(bytes)
//    bytes, err = membershipDeletedEvent.Marshal()
//
//    projectCreatedEvent, err := UnmarshalProjectCreatedEvent(bytes)
//    bytes, err = projectCreatedEvent.Marshal()
//
//    projectUpdatedEvent, err := UnmarshalProjectUpdatedEvent(bytes)
//    bytes, err = projectUpdatedEvent.Marshal()
//
//    projectDeletedEvent, err := UnmarshalProjectDeletedEvent(bytes)
//    bytes, err = projectDeletedEvent.Marshal()

package events

import "encoding/json"

func UnmarshalEventTypes(data []byte) (EventTypes, error) {
	var r EventTypes
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *EventTypes) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMessage(data []byte) (Message, error) {
	var r Message
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Message) Marshal() ([]byte, error) {
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

func UnmarshalCategories(data []byte) (Categories, error) {
	var r Categories
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Categories) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCommands(data []byte) (Commands, error) {
	var r Commands
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Commands) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalEnableAccountingCommand(data []byte) (EnableAccountingCommand, error) {
	var r EnableAccountingCommand
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *EnableAccountingCommand) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalEvents(data []byte) (Events, error) {
	var r Events
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Events) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMembershipCreatedEvent(data []byte) (MembershipCreatedEvent, error) {
	var r MembershipCreatedEvent
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MembershipCreatedEvent) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMembershipCreatedForProjectEvent(data []byte) (MembershipCreatedForProjectEvent, error) {
	var r MembershipCreatedForProjectEvent
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MembershipCreatedForProjectEvent) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMembershipUpdatedEvent(data []byte) (MembershipUpdatedEvent, error) {
	var r MembershipUpdatedEvent
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MembershipUpdatedEvent) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMembershipDeletedEvent(data []byte) (MembershipDeletedEvent, error) {
	var r MembershipDeletedEvent
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MembershipDeletedEvent) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalProjectCreatedEvent(data []byte) (ProjectCreatedEvent, error) {
	var r ProjectCreatedEvent
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ProjectCreatedEvent) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalProjectUpdatedEvent(data []byte) (ProjectUpdatedEvent, error) {
	var r ProjectUpdatedEvent
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ProjectUpdatedEvent) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalProjectDeletedEvent(data []byte) (ProjectDeletedEvent, error) {
	var r ProjectDeletedEvent
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ProjectDeletedEvent) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Message struct {
	Data     interface{} `json:"data"`    
	ID       string      `json:"id"`      
	Metadata Metadata    `json:"metadata"`
	Type     EventTypes  `json:"type"`    
}

type Metadata struct {
	TraceID string `json:"traceId"`
	UserID  string `json:"userId"` 
}

type EnableAccountingCommand struct {
	Data     EnableAccountingCommandData `json:"data"`    
	ID       string                      `json:"id"`      
	Metadata Metadata                    `json:"metadata"`
	Type     Commands                    `json:"type"`    
}

type EnableAccountingCommandData struct {
	Auth0ID string `json:"auth0Id"`
	Token   string `json:"token"`  
}

type MembershipCreatedEvent struct {
	Data     MembershipCreatedEventData `json:"data"`    
	ID       string                     `json:"id"`      
	Metadata Metadata                   `json:"metadata"`
	Type     MembershipCreatedEventType `json:"type"`    
}

type MembershipCreatedEventData struct {
	CreatedAt    string `json:"createdAt"`   
	MembershipID string `json:"membershipId"`
	Role         string `json:"role"`        
	TeamID       string `json:"teamId"`      
	UpdatedAt    string `json:"updatedAt"`   
	UserID       string `json:"userId"`      
}

type MembershipCreatedForProjectEvent struct {
	Data     MembershipCreatedForProjectEventData `json:"data"`    
	ID       string                               `json:"id"`      
	Metadata Metadata                             `json:"metadata"`
	Type     MembershipCreatedForProjectEventType `json:"type"`    
}

type MembershipCreatedForProjectEventData struct {
	CreatedAt    string `json:"createdAt"`   
	MembershipID string `json:"membershipId"`
	ProjectID    string `json:"projectId"`   
	Role         string `json:"role"`        
	TeamID       string `json:"teamId"`      
	UpdatedAt    string `json:"updatedAt"`   
	UserID       string `json:"userId"`      
}

type MembershipUpdatedEvent struct {
	Data     MembershipUpdatedEventData `json:"data"`    
	ID       string                     `json:"id"`      
	Metadata Metadata                   `json:"metadata"`
	Type     MembershipUpdatedEventType `json:"type"`    
}

type MembershipUpdatedEventData struct {
	MembershipID string `json:"membershipId"`
	Role         string `json:"role"`        
	UpdatedAt    string `json:"updatedAt"`   
}

type MembershipDeletedEvent struct {
	Data     MembershipDeletedEventData `json:"data"`    
	ID       string                     `json:"id"`      
	Metadata Metadata                   `json:"metadata"`
	Type     MembershipDeletedEventType `json:"type"`    
}

type MembershipDeletedEventData struct {
	MembershipID string `json:"membershipId"`
}

type ProjectCreatedEvent struct {
	Data     ProjectCreatedEventData `json:"data"`    
	ID       string                  `json:"id"`      
	Metadata Metadata                `json:"metadata"`
	Type     ProjectCreatedEventType `json:"type"`    
}

type ProjectCreatedEventData struct {
	Active      bool     `json:"active"`     
	ColumnOrder []string `json:"columnOrder"`
	CreatedAt   string   `json:"createdAt"`  
	Name        string   `json:"name"`       
	ProjectID   string   `json:"projectId"`  
	Public      bool     `json:"public"`     
	TeamID      string   `json:"teamId"`     
	UpdatedAt   string   `json:"updatedAt"`  
	UserID      string   `json:"userId"`     
}

type ProjectUpdatedEvent struct {
	Data     ProjectUpdatedEventData `json:"data"`    
	ID       string                  `json:"id"`      
	Metadata Metadata                `json:"metadata"`
	Type     ProjectUpdatedEventType `json:"type"`    
}

type ProjectUpdatedEventData struct {
	Active      *bool    `json:"active,omitempty"`     
	ColumnOrder []string `json:"columnOrder,omitempty"`
	Name        *string  `json:"name,omitempty"`       
	ProjectID   string   `json:"projectId"`            
	Public      *bool    `json:"public,omitempty"`     
	TeamID      *string  `json:"teamId,omitempty"`     
	UpdatedAt   string   `json:"updatedAt"`            
}

type ProjectDeletedEvent struct {
	Data     ProjectDeletedEventData `json:"data"`    
	ID       string                  `json:"id"`      
	Metadata Metadata                `json:"metadata"`
	Type     ProjectDeletedEventType `json:"type"`    
}

type ProjectDeletedEventData struct {
	ProjectID string `json:"projectId"`
}

type EventTypes string
const (
	EventTypesEnableAccounting EventTypes = "EnableAccounting"
	EventTypesMembershipCreated EventTypes = "MembershipCreated"
	EventTypesMembershipCreatedForProject EventTypes = "MembershipCreatedForProject"
	EventTypesMembershipDeleted EventTypes = "MembershipDeleted"
	EventTypesMembershipUpdated EventTypes = "MembershipUpdated"
	EventTypesProjectCreated EventTypes = "ProjectCreated"
	EventTypesProjectDeleted EventTypes = "ProjectDeleted"
	EventTypesProjectUpdated EventTypes = "ProjectUpdated"
)

type Categories string
const (
	Accounting Categories = "Accounting"
	Estimation Categories = "Estimation"
	Project Categories = "Project"
	Users Categories = "Users"
)

type Commands string
const (
	CommandsEnableAccounting Commands = "EnableAccounting"
)

type Events string
const (
	EventsMembershipCreated Events = "MembershipCreated"
	EventsMembershipCreatedForProject Events = "MembershipCreatedForProject"
	EventsMembershipDeleted Events = "MembershipDeleted"
	EventsMembershipUpdated Events = "MembershipUpdated"
	EventsProjectCreated Events = "ProjectCreated"
	EventsProjectDeleted Events = "ProjectDeleted"
	EventsProjectUpdated Events = "ProjectUpdated"
)

type MembershipCreatedEventType string
const (
	TypeMembershipCreated MembershipCreatedEventType = "MembershipCreated"
)

type MembershipCreatedForProjectEventType string
const (
	TypeMembershipCreatedForProject MembershipCreatedForProjectEventType = "MembershipCreatedForProject"
)

type MembershipUpdatedEventType string
const (
	TypeMembershipUpdated MembershipUpdatedEventType = "MembershipUpdated"
)

type MembershipDeletedEventType string
const (
	TypeMembershipDeleted MembershipDeletedEventType = "MembershipDeleted"
)

type ProjectCreatedEventType string
const (
	TypeProjectCreated ProjectCreatedEventType = "ProjectCreated"
)

type ProjectUpdatedEventType string
const (
	TypeProjectUpdated ProjectUpdatedEventType = "ProjectUpdated"
)

type ProjectDeletedEventType string
const (
	TypeProjectDeleted ProjectDeletedEventType = "ProjectDeleted"
)
