// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    userModifiedEvent, err := UnmarshalUserModifiedEvent(bytes)
//    bytes, err = userModifiedEvent.Marshal()

package events

import "encoding/json"

func UnmarshalUserModifiedEvent(data []byte) (UserModifiedEvent, error) {
	var r UserModifiedEvent
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UserModifiedEvent) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type UserModifiedEvent struct {
	Data     UserModifiedEventData `json:"data"`    
	ID       string                `json:"id"`      
	Metadata Metadata              `json:"metadata"`
	Type     UserModifiedEventType `json:"type"`    
}
