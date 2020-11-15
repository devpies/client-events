// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    userRegisteredEvent, err := UnmarshalUserRegisteredEvent(bytes)
//    bytes, err = userRegisteredEvent.Marshal()

package events

import "encoding/json"

func UnmarshalUserRegisteredEvent(data []byte) (UserRegisteredEvent, error) {
	var r UserRegisteredEvent
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UserRegisteredEvent) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type UserRegisteredEvent struct {
	Data     UserRegisteredEventData `json:"data"`    
	ID       string                  `json:"id"`      
	Metadata Metadata                `json:"metadata"`
	Type     UserRegisteredEventType `json:"type"`    
}
