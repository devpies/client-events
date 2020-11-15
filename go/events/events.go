// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    events, err := UnmarshalEvents(bytes)
//    bytes, err = events.Marshal()

package events

import "encoding/json"

func UnmarshalEvents(data []byte) (Events, error) {
	var r Events
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Events) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Events string
const (
	EventsUserModified Events = "UserModified"
	EventsUserRegistered Events = "UserRegistered"
)
