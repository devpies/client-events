// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    registerUserCommand, err := UnmarshalRegisterUserCommand(bytes)
//    bytes, err = registerUserCommand.Marshal()

package events

import "encoding/json"

func UnmarshalRegisterUserCommand(data []byte) (RegisterUserCommand, error) {
	var r RegisterUserCommand
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RegisterUserCommand) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type RegisterUserCommand struct {
	Data     RegisterUserCommandData `json:"data"`    
	ID       string                  `json:"id"`      
	Metadata Metadata                `json:"metadata"`
	Type     RegisterUserCommandType `json:"type"`    
}
