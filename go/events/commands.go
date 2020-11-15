// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    commands, err := UnmarshalCommands(bytes)
//    bytes, err = commands.Marshal()

package events

import "encoding/json"

func UnmarshalCommands(data []byte) (Commands, error) {
	var r Commands
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Commands) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Commands string
const (
	CommandsModifyUser Commands = "ModifyUser"
	CommandsRegisterUser Commands = "RegisterUser"
)
