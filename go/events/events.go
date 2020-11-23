// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    subjects, err := UnmarshalSubjects(bytes)
//    bytes, err = subjects.Marshal()
//
//    message, err := UnmarshalMessage(bytes)
//    bytes, err = message.Marshal()
//
//    metadata, err := UnmarshalMetadata(bytes)
//    bytes, err = metadata.Marshal()
//
//    commands, err := UnmarshalCommands(bytes)
//    bytes, err = commands.Marshal()
//
//    addUserCommand, err := UnmarshalAddUserCommand(bytes)
//    bytes, err = addUserCommand.Marshal()
//
//    modifyUserCommand, err := UnmarshalModifyUserCommand(bytes)
//    bytes, err = modifyUserCommand.Marshal()
//
//    events, err := UnmarshalEvents(bytes)
//    bytes, err = events.Marshal()
//
//    userAddedEvent, err := UnmarshalUserAddedEvent(bytes)
//    bytes, err = userAddedEvent.Marshal()
//
//    userModifiedEvent, err := UnmarshalUserModifiedEvent(bytes)
//    bytes, err = userModifiedEvent.Marshal()

package events

import "encoding/json"

func UnmarshalSubjects(data []byte) (Subjects, error) {
	var r Subjects
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Subjects) Marshal() ([]byte, error) {
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

func UnmarshalCommands(data []byte) (Commands, error) {
	var r Commands
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Commands) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalAddUserCommand(data []byte) (AddUserCommand, error) {
	var r AddUserCommand
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AddUserCommand) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalModifyUserCommand(data []byte) (ModifyUserCommand, error) {
	var r ModifyUserCommand
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ModifyUserCommand) Marshal() ([]byte, error) {
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

func UnmarshalUserAddedEvent(data []byte) (UserAddedEvent, error) {
	var r UserAddedEvent
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UserAddedEvent) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUserModifiedEvent(data []byte) (UserModifiedEvent, error) {
	var r UserModifiedEvent
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UserModifiedEvent) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Message struct {
	Data     interface{} `json:"data"`    
	ID       string      `json:"id"`      
	Metadata Metadata    `json:"metadata"`
	Subject  Subjects    `json:"subject"` 
}

type Metadata struct {
	TraceID string `json:"traceId"`
	UserID  string `json:"userId"` 
}

type AddUserCommand struct {
	Data     AddUserCommandData    `json:"data"`    
	ID       string                `json:"id"`      
	Metadata Metadata              `json:"metadata"`
	Subject  AddUserCommandSubject `json:"subject"` 
}

type AddUserCommandData struct {
	Auth0ID       string `json:"auth0Id"`      
	Email         string `json:"email"`        
	EmailVerified bool   `json:"emailVerified"`
	FirstName     string `json:"firstName"`    
	ID            string `json:"id"`           
	LastName      string `json:"lastName"`     
	Locale        string `json:"locale"`       
	Picture       string `json:"picture"`      
}

type ModifyUserCommand struct {
	Data     ModifyUserCommandData    `json:"data"`    
	ID       string                   `json:"id"`      
	Metadata Metadata                 `json:"metadata"`
	Subject  ModifyUserCommandSubject `json:"subject"` 
}

type ModifyUserCommandData struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"` 
	Locale    string `json:"locale"`   
	Picture   string `json:"picture"`  
}

type UserAddedEvent struct {
	Data     UserAddedEventData    `json:"data"`    
	ID       string                `json:"id"`      
	Metadata Metadata              `json:"metadata"`
	Subject  UserAddedEventSubject `json:"subject"` 
}

type UserAddedEventData struct {
	Auth0ID       string `json:"auth0Id"`      
	Email         string `json:"email"`        
	EmailVerified bool   `json:"emailVerified"`
	FirstName     string `json:"firstName"`    
	ID            string `json:"id"`           
	LastName      string `json:"lastName"`     
	Locale        string `json:"locale"`       
	Picture       string `json:"picture"`      
}

type UserModifiedEvent struct {
	Data     UserModifiedEventData    `json:"data"`    
	ID       string                   `json:"id"`      
	Metadata Metadata                 `json:"metadata"`
	Subject  UserModifiedEventSubject `json:"subject"` 
}

type UserModifiedEventData struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"` 
	Locale    string `json:"locale"`   
	Picture   string `json:"picture"`  
}

type Subjects string
const (
	SubjectsAddUser Subjects = "AddUser"
	SubjectsModifyUser Subjects = "ModifyUser"
	SubjectsUserAdded Subjects = "UserAdded"
	SubjectsUserModified Subjects = "UserModified"
)

type Commands string
const (
	CommandsAddUser Commands = "AddUser"
	CommandsModifyUser Commands = "ModifyUser"
)

type AddUserCommandSubject string
const (
	SubjectAddUser AddUserCommandSubject = "AddUser"
)

type ModifyUserCommandSubject string
const (
	SubjectModifyUser ModifyUserCommandSubject = "ModifyUser"
)

type Events string
const (
	EventsUserAdded Events = "UserAdded"
	EventsUserModified Events = "UserModified"
)

type UserAddedEventSubject string
const (
	SubjectUserAdded UserAddedEventSubject = "UserAdded"
)

type UserModifiedEventSubject string
const (
	SubjectUserModified UserModifiedEventSubject = "UserModified"
)
