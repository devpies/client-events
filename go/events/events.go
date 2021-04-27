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
//    membershipAddedEvent, err := UnmarshalMembershipAddedEvent(bytes)
//    bytes, err = membershipAddedEvent.Marshal()
//
//    membershipDroppedEvent, err := UnmarshalMembershipDroppedEvent(bytes)
//    bytes, err = membershipDroppedEvent.Marshal()
//
//    projectTeamAssignedEvent, err := UnmarshalProjectTeamAssignedEvent(bytes)
//    bytes, err = projectTeamAssignedEvent.Marshal()
//
//    projectTeamUnassignedEvent, err := UnmarshalProjectTeamUnassignedEvent(bytes)
//    bytes, err = projectTeamUnassignedEvent.Marshal()

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

func UnmarshalMembershipAddedEvent(data []byte) (MembershipAddedEvent, error) {
	var r MembershipAddedEvent
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MembershipAddedEvent) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMembershipDroppedEvent(data []byte) (MembershipDroppedEvent, error) {
	var r MembershipDroppedEvent
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MembershipDroppedEvent) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalProjectTeamAssignedEvent(data []byte) (ProjectTeamAssignedEvent, error) {
	var r ProjectTeamAssignedEvent
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ProjectTeamAssignedEvent) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalProjectTeamUnassignedEvent(data []byte) (ProjectTeamUnassignedEvent, error) {
	var r ProjectTeamUnassignedEvent
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ProjectTeamUnassignedEvent) Marshal() ([]byte, error) {
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

type MembershipAddedEvent struct {
	Data     MembershipAddedEventData `json:"data"`    
	ID       string                   `json:"id"`      
	Metadata Metadata                 `json:"metadata"`
	Type     MembershipAddedEventType `json:"type"`    
}

type MembershipAddedEventData struct {
	CreatedAt string `json:"CreatedAt"`
	MemberID  string `json:"MemberId"` 
	Role      string `json:"Role"`     
	TeamID    string `json:"TeamId"`   
	UpdatedAt string `json:"UpdatedAt"`
	UserID    string `json:"UserId"`   
}

type MembershipDroppedEvent struct {
	Data     MembershipDroppedEventData `json:"data"`    
	ID       string                     `json:"id"`      
	Metadata Metadata                   `json:"metadata"`
	Type     MembershipDroppedEventType `json:"type"`    
}

type MembershipDroppedEventData struct {
	MemberID string `json:"MemberId"`
}

type ProjectTeamAssignedEvent struct {
	Data     ProjectTeamAssignedEventData `json:"data"`    
	ID       string                       `json:"id"`      
	Metadata Metadata                     `json:"metadata"`
	Type     ProjectTeamAssignedEventType `json:"type"`    
}

type ProjectTeamAssignedEventData struct {
	ProjectID string `json:"ProjectId"`
	TeamID    string `json:"TeamId"`   
}

type ProjectTeamUnassignedEvent struct {
	Data     ProjectTeamUnassignedEventData `json:"data"`    
	ID       string                         `json:"id"`      
	Metadata Metadata                       `json:"metadata"`
	Type     ProjectTeamUnassignedEventType `json:"type"`    
}

type ProjectTeamUnassignedEventData struct {
	ProjectID string `json:"ProjectId"`
}

type EventTypes string
const (
	EventTypesEnableAccounting EventTypes = "EnableAccounting"
	EventTypesMembershipAdded EventTypes = "MembershipAdded"
	EventTypesMembershipDropped EventTypes = "MembershipDropped"
	EventTypesProjectTeamAssigned EventTypes = "ProjectTeamAssigned"
	EventTypesProjectTeamUnassigned EventTypes = "ProjectTeamUnassigned"
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
	EventsMembershipAdded Events = "MembershipAdded"
	EventsMembershipDropped Events = "MembershipDropped"
	EventsProjectTeamAssigned Events = "ProjectTeamAssigned"
	EventsProjectTeamUnassigned Events = "ProjectTeamUnassigned"
)

type MembershipAddedEventType string
const (
	TypeMembershipAdded MembershipAddedEventType = "MembershipAdded"
)

type MembershipDroppedEventType string
const (
	TypeMembershipDropped MembershipDroppedEventType = "MembershipDropped"
)

type ProjectTeamAssignedEventType string
const (
	TypeProjectTeamAssigned ProjectTeamAssignedEventType = "ProjectTeamAssigned"
)

type ProjectTeamUnassignedEventType string
const (
	TypeProjectTeamUnassigned ProjectTeamUnassignedEventType = "ProjectTeamUnassigned"
)
