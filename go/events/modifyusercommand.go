// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    modifyUserCommand, err := UnmarshalModifyUserCommand(bytes)
//    bytes, err = modifyUserCommand.Marshal()

package events

import "encoding/json"

func UnmarshalModifyUserCommand(data []byte) (ModifyUserCommand, error) {
	var r ModifyUserCommand
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ModifyUserCommand) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ModifyUserCommand struct {
	Data     ModifyUserCommandData `json:"data"`    
	ID       string                `json:"id"`      
	Metadata Metadata              `json:"metadata"`
	Type     ModifyUserCommandType `json:"type"`    
}
